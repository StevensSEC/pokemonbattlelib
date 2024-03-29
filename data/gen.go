package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const EnglishLanguageID = 9
const NationalDexID = 1
const HighestGen = 4
const HighestDexNum = 493

var statNames = map[int]string{
	1: "StatHP",
	2: "StatAtk",
	3: "StatDef",
	4: "StatSpAtk",
	5: "StatSpDef",
	6: "StatSpeed",
	7: "StatAccuracy",
	8: "StatEvasion",
}

var ailmentNames = map[int]string{
	-1: "StatusNone",
	0:  "StatusNone",
	1:  "StatusParalyze",
	2:  "StatusSleep",
	3:  "StatusFreeze",
	4:  "StatusBurn",
	5:  "StatusPoison",
	6:  "StatusConfusion",
	7:  "StatusInfatuation",
	8:  "StatusCantEscape",
	9:  "StatusNightmare",
	12: "StatusTorment",
	13: "StatusNone", // FIXME: not implemented
	14: "StatusNone", // FIXME: not implemented
	15: "StatusHealBlock",
	17: "StatusNone", // FIXME: not implemented
	18: "StatusLeechSeed",
	19: "StatusEmbargo",
	20: "StatusPerishSong",
	21: "StatusNone", // FIXME: not implemented
	24: "StatusNone", // FIXME: not implemented
}

var moveMetaCategoryNames = map[int]string{
	0:  "MoveMetaCategoryDamage",
	1:  "MoveMetaCategoryAilment",
	2:  "MoveMetaCategoryNetGoodStats",
	3:  "MoveMetaCategoryHeal",
	4:  "MoveMetaCategoryDamageAilment",
	5:  "MoveMetaCategorySwagger",
	6:  "MoveMetaCategoryDamageLower",
	7:  "MoveMetaCategoryDamageRaise",
	8:  "MoveMetaCategoryDamageHeal",
	9:  "MoveMetaCategoryOhko",
	10: "MoveMetaCategoryWholeFieldEffect",
	11: "MoveMetaCategoryFieldEffect",
	12: "MoveMetaCategoryForceSwitch",
	13: "MoveMetaCategoryUnique",
}

type data_pokemon struct {
	Identifier        string
	SpeciesId         int
	Height            int
	Weight            int
	BaseExperience    int
	IsBiGender        bool
	GenderRate        int
	IsLegendary       bool
	IsMythical        bool
	CaptureRate       int
	EvolvesFrom       int
	HasAlternateForms bool

	Name       string
	NatDex     uint16
	Type       int
	Ability    string
	Stats      [6]int
	Evs        [6]int
	GrowthRate int
}

type data_move struct {
	Id             int
	Identifier     string
	SpeciesId      int
	Height         int
	Weight         int
	BaseExperience int

	Name     string
	Type     int
	Power    int
	PP       int
	Accuracy int
	Priority int
	// see: move_targets.csv
	Targets     int
	DamageClass string
	Effect      int
	// Metadata
	MinHits       int
	MaxHits       int
	MinTurns      int
	MaxTurns      int
	Drain         int
	Healing       int
	CritRate      int
	AilmentChance int
	FlinchChance  int
	StatChance    int
	Flags         []string
	AffectedStat  string
	StatChange    int
	Ailment       string
	MetaCategory  string
}

type data_item struct {
	ID            string
	Identifier    string
	CategoryID    int
	FlingPower    int
	FlingEffectID int
	Flags         []string
}

type data_nature struct {
	id       int
	name     string
	statup   int
	statdown int
}

func parseInt(s string) (n int) {
	if s == "" {
		return 0
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panicln(err)
	}
	return n
}

func parseBool(s string) bool {
	return s == "1"
}

func cleanName(s string) string {
	name := strings.ReplaceAll(s, "♀", "F")
	name = strings.ReplaceAll(name, "♂", "M")
	name = strings.ReplaceAll(name, "É", "E")
	name = strings.ReplaceAll(name, " ", "")
	re, err := regexp.Compile(`[^a-zA-Z_0-9]`)
	if err != nil {
		log.Panicln(err)
	}
	return re.ReplaceAllString(name, "")
}

func getCsvReader(path string) *csv.Reader {
	log.Printf("Reading CSV: %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}
	reader := csv.NewReader(file)
	reader.ReuseRecord = true
	_, err = reader.Read() // skip header line
	if err != nil {
		log.Panicln(err)
	}
	return reader
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func createLevelTableStringFromArray(growth_rate_name string, level_array []int) string {
	output := growth_rate_name + ": {\n"
	for level, experience := range level_array {
		output += fmt.Sprintf("%d: %d,\n", level, experience)
	}
	output += "}"
	return output
}

func createCodeOutput(s string) {
	file, err := os.Create("pokedex_GEN.go")
	if err != nil {
		log.Panicln(err)
	}
	_, err = file.WriteString("// Code generated - DO NOT EDIT.\n" +
		"// Regenerate with `go generate`.\n\n" +
		"package pokemonbattlelib\n\n" + s)
	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	var records [][]string
	var err error
	path, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("Current directory: %s\n", path)
	output := ""

	growth_rate_strings := map[int]string{
		1: "GrowthSlow",
		2: "GrowthMediumFast",
		3: "GrowthFast",
		4: "GrowthMediumSlow",
		5: "GrowthErratic",
		6: "GrowthFluctuating",
	}

	// get all valid version ids
	valid_version_groups := []int{}
	version_groups_csv := getCsvReader("data/version_groups.csv")
	for {
		record, err := version_groups_csv.Read()
		if err == io.EOF {
			break
		}
		gen_id := parseInt(record[2])
		if gen_id <= HighestGen {
			group_id := parseInt(record[0])
			valid_version_groups = append(valid_version_groups, group_id)
		}
	}
	valid_versions := []int{}
	versions_csv := getCsvReader("data/versions.csv")
	for {
		record, err := versions_csv.Read()
		if err == io.EOF {
			break
		}
		group_id := parseInt(record[1])
		if contains(valid_version_groups, group_id) {
			vid := parseInt(record[0])
			valid_versions = append(valid_versions, vid)
		}
	}

	// Abilities
	log.Println("Generating code for Abilities")
	// read abilities data
	abilities_csv := getCsvReader("data/abilities.csv")
	abilities := []int{}
	for {
		record, err := abilities_csv.Read()
		if err == io.EOF {
			break
		}

		if parseInt(record[2]) > HighestGen {
			continue
		}

		abilities = append(abilities, parseInt(record[0]))
	}
	ability_names_csv := getCsvReader("data/ability_names.csv")
	ability_names := map[int]string{}
	for {
		record, err := ability_names_csv.Read()
		if err == io.EOF {
			break
		}

		if parseInt(record[1]) != EnglishLanguageID {
			continue
		}

		ability_names[parseInt(record[0])] = record[2]
	}
	sort.Slice(abilities, func(i, j int) bool {
		return ability_names[abilities[i]] < ability_names[abilities[j]]
	})
	// generate Ability Constants
	log.Println("Generating Ability constants")
	output += "const (\n"
	for i, n := range abilities {
		if i == 0 {
			output += fmt.Sprintf("Ability%s Ability = iota + 1\n", cleanName(ability_names[n]))
		} else {
			output += fmt.Sprintf("Ability%s\n", cleanName(ability_names[n]))
		}
	}
	output += ")\n\n"
	// generate Ability String()
	log.Println("Generating Ability String()")
	output += "// Get the string name of this Ability.\n" +
		"func (n Ability) String() string {\n" +
		"switch n {\n"
	for _, n := range abilities {
		output += fmt.Sprintf("case Ability%s:\n", cleanName(ability_names[n]))
		output += fmt.Sprintf("return \"%s\"\n", ability_names[n])
	}
	output += "}\n" +
		"panic(\"Unknown ability\")" +
		"}\n\n"

	// get valid pokemon based on generation
	valid_pkmn_ids := []int{}
	pkmn_game_csv := getCsvReader("data/pokemon_game_indices.csv")
	for {
		record, err := pkmn_game_csv.Read()
		if err == io.EOF {
			break
		}
		vid := parseInt(record[1])
		if contains(valid_versions, vid) {
			pid := parseInt(record[0])
			valid_pkmn_ids = append(valid_pkmn_ids, pid)
		}
	}

	// get all the pokemon
	pokemon := []data_pokemon{}
	pkmn_csv := getCsvReader("data/pokemon.csv")
	for {
		record, err := pkmn_csv.Read()
		if err == io.EOF {
			break
		}
		// fmt.Printf("%v\n", record)
		id := parseInt(record[0])
		if !contains(valid_pkmn_ids, id) {
			continue
		}
		species_id := parseInt(record[2])
		height := parseInt(record[3])
		weight := parseInt(record[4])
		baseexp := parseInt(record[5])
		dp := data_pokemon{
			Identifier:     record[1],
			SpeciesId:      species_id,
			Height:         height,
			Weight:         weight,
			BaseExperience: baseexp,
		}
		if species_id == 94 { // gengar
			// HACK: see #267
			dp.Ability = "AbilityLevitate"
		}
		pokemon = append(pokemon, dp)
	}

	log.Println("Getting species metadata")
	pkmn_species_csv := getCsvReader("data/pokemon_species.csv")
	for {
		record, err := pkmn_species_csv.Read()
		if err == io.EOF {
			break
		}
		// fmt.Printf("%v\n", record)
		sid := parseInt(record[0])
		for i, p := range pokemon {
			if p.SpeciesId != sid {
				continue
			}
			(&pokemon[i]).EvolvesFrom = parseInt(record[3])
			gender_rate := parseInt(record[8])
			if gender_rate >= 0 {
				(&pokemon[i]).IsBiGender = true
				(&pokemon[i]).GenderRate = gender_rate
			}
			(&pokemon[i]).HasAlternateForms = parseBool(record[15])
			(&pokemon[i]).IsLegendary = parseBool(record[16])
			(&pokemon[i]).IsMythical = parseBool(record[17])
			break
		}
	}

	// find all the pokemon names
	log.Println("Getting Pokemon names")
	pkmn_names_csv := getCsvReader("data/pokemon_species_names.csv")
	for {
		record, err := pkmn_names_csv.Read()
		if err == io.EOF {
			break
		}
		lang := parseInt(record[1])
		if lang != EnglishLanguageID {
			continue
		}
		sid := parseInt(record[0])
		for i, p := range pokemon {
			if p.SpeciesId != sid {
				continue
			}
			(&pokemon[i]).Name = record[2]
			break
		}
	}
	// get all Pokemon types
	log.Println("Getting Pokemon types")
	pkmn_types_csv := getCsvReader("data/pokemon_types.csv")
	records, err = pkmn_types_csv.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, r := range records {
		id := parseInt(r[0])
		pType := parseInt(r[1])
		for i, p := range pokemon {
			if p.SpeciesId != id {
				continue
			}
			pokemon[i].Type |= (1 << (pType - 1))
		}
	}
	for {
		record, err := pkmn_names_csv.Read()
		if err == io.EOF {
			break
		}
		lang := parseInt(record[1])
		if lang != EnglishLanguageID {
			continue
		}
		sid := parseInt(record[0])
		for i, p := range pokemon {
			if p.SpeciesId != sid {
				continue
			}
			(&pokemon[i]).Name = record[2]
			break
		}
	}

	// get all Pokemon abilities
	log.Println("Getting Pokemon abilities")
	pkmn_abilities_csv := getCsvReader("data/pokemon_abilities.csv")
	for {
		record, err := pkmn_abilities_csv.Read()
		if err == io.EOF {
			break
		}
		sid := parseInt(record[0])
		aid := parseInt(record[1])
		is_hidden := record[2] == "1"
		if is_hidden {
			// Hidden abilities were introduced in gen 5
			continue
		}
		if !contains(abilities, aid) {
			// skip abilities that aren't available in gen 4
			continue
		}
		for i, p := range pokemon {
			if p.SpeciesId != sid {
				continue
			}
			(&pokemon[i]).Ability = fmt.Sprintf("Ability%s", cleanName(ability_names[aid]))
			break
		}
	}

	// find national dex numbers
	log.Println("Getting Pokemon dex numbers")
	pkmn_dex_nums_csv := getCsvReader("data/pokemon_dex_numbers.csv")
	for {
		record, err := pkmn_dex_nums_csv.Read()
		if err == io.EOF {
			break
		}
		dexid := parseInt(record[1])
		if dexid != NationalDexID {
			continue
		}
		sid := parseInt(record[0])
		for i, p := range pokemon {
			if p.SpeciesId != sid {
				continue
			}
			dexnum := parseInt(record[2])
			(&pokemon[i]).NatDex = uint16(dexnum)
			break
		}
	}

	// pokemon base stat + EV yield
	log.Println("Creating base stat table + EV yield table")
	pokemon_stats_csv := getCsvReader("data/pokemon_stats.csv")
	for {
		record, err := pokemon_stats_csv.Read()

		if err == io.EOF {
			break
		}

		dex_num := parseInt(record[0])

		if dex_num > HighestDexNum {
			break
		}

		stat_id := parseInt(record[1]) - 1
		stat_value := parseInt(record[2])
		ev := parseInt(record[3])
		pokemon[dex_num-1].Stats[stat_id] = stat_value
		pokemon[dex_num-1].Evs[stat_id] = ev
	}

	log.Println("Mapping growth rates to dex numbers")
	// map growth rates to pokemon national dex numbers
	pokemon_species_csv := getCsvReader("data/pokemon_species.csv")
	for {
		record, err := pokemon_species_csv.Read()

		if err == io.EOF {
			break
		}

		growth_rate_id := parseInt(record[14])
		dex_num := parseInt(record[0])

		if dex_num == 0 {
			continue
		}

		if dex_num > HighestDexNum {
			break
		}

		pokemon[dex_num-1].GrowthRate = growth_rate_id
	}

	output += "\n\n" +
		"// An array of all registered pokemon data ordered by national pokedex numbers.\n" +
		"var AllPokemonData = []PokemonData{\n"
	for _, p := range pokemon {
		if p.NatDex == 0 {
			continue
		}
		output += fmt.Sprintf("{NatDex: %d, Name: \"%s\", Type: %v, Ability: %s, BaseStats: %#v, EvYield: %#v, GrowthRate: %s, IsBiGender: %v, GenderRate: %d, HasAlternateForms: %v, IsLegendary: %v, IsMythical: %v, EvolvesFrom: %d},\n", p.NatDex, p.Name, p.Type, p.Ability, p.Stats, p.Evs, growth_rate_strings[p.GrowthRate], p.IsBiGender, p.GenderRate, p.HasAlternateForms, p.IsLegendary, p.IsMythical, p.EvolvesFrom)
	}
	output += "}\n\n"
	output += "// Pokemon const enum for quick lookup\nconst (\n"
	for _, p := range pokemon {
		if p.NatDex == 0 {
			continue
		}
		name := cleanName(p.Name)
		output += fmt.Sprintf("\tPkmn%s = %v\n", name, p.NatDex)
	}
	output += ")\n\n"
	// find all moves
	moves := []data_move{}
	log.Println("Getting all available moves")
	moves_csv := getCsvReader("data/moves.csv")
	moveMap := map[int]string{
		1: "MoveCategoryStatus",
		2: "MoveCategoryPhysical",
		3: "MoveCategorySpecial",
	}
	for {
		record, err := moves_csv.Read()
		if err == io.EOF {
			break
		}
		gid := parseInt(record[2])
		if gid > HighestGen {
			continue
		}
		mid := parseInt(record[0])
		if mid >= 10000 {
			continue // skip these, because they are the shaow moves from pokemon XD, see #355
		}
		moveType := 1 << (parseInt(record[3]) - 1)
		power := parseInt(record[4])
		pp := parseInt(record[5])
		accuracy := parseInt(record[6])
		priority := parseInt(record[7])
		targets := parseInt(record[8])
		damageClass := parseInt(record[9])
		effect := parseInt(record[10])
		moves = append(moves, data_move{
			Id:          mid,
			Identifier:  record[1],
			Type:        moveType,
			Power:       power,
			PP:          pp,
			Accuracy:    accuracy,
			Priority:    priority,
			Targets:     targets,
			DamageClass: moveMap[damageClass],
			Effect:      effect,
		})
	}
	// find all the move names
	log.Println("Getting move names")
	move_names_csv := getCsvReader("data/move_names.csv")
	for {
		record, err := move_names_csv.Read()
		if err == io.EOF {
			break
		}
		lang := parseInt(record[1])
		if lang != EnglishLanguageID {
			continue
		}
		mid := parseInt(record[0])
		for i, m := range moves {
			if m.Id != mid {
				continue
			}
			(&moves[i]).Name = record[2]
			break
		}
	}
	// find all the move flags
	log.Println("Getting move flags")
	moveFlagMap := map[int]string{
		1:  "FlagContact",
		2:  "FlagCharge",
		3:  "FlagRecharge",
		4:  "FlagProtect",
		5:  "FlagReflectable",
		6:  "FlagSnatch",
		7:  "FlagMirror",
		8:  "FlagPunch",
		9:  "FlagSound",
		10: "FlagGravity",
		11: "FlagDefrost",
		12: "FlagDistance",
		13: "FlagHeal",
		14: "FlagAuthentic",
		15: "FlagPowder",
		16: "FlagBite",
		17: "FlagPulse",
		18: "FlagBallistics",
		19: "FlagMental",
		20: "FlagNonSkyBattle",
		21: "FlagDance",
	}
	move_flag_map_csv := getCsvReader("data/move_flag_map.csv")
	for {
		record, err := move_flag_map_csv.Read()
		if err == io.EOF {
			break
		}
		mid := parseInt(record[0])
		flag := parseInt(record[1])
		for i, m := range moves {
			if m.Id != mid {
				continue
			}
			(&moves[i]).Flags = append((&moves[i]).Flags, moveFlagMap[flag])
			break
		}
	}
	// Getting move metadata
	log.Println("Getting move metadata")
	move_meta_csv := getCsvReader("data/move_meta.csv")
	records, err = move_meta_csv.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, v := range records {
		mid := parseInt(v[0])
		for i, m := range moves {
			if m.Id != mid {
				continue
			}
			moves[i].MetaCategory = moveMetaCategoryNames[parseInt(v[1])]
			moves[i].Ailment = ailmentNames[parseInt(v[2])]
			moves[i].MinHits = parseInt(v[3])
			moves[i].MaxHits = parseInt(v[4])
			moves[i].MinTurns = parseInt(v[5])
			moves[i].MaxTurns = parseInt(v[6])
			moves[i].Drain = parseInt(v[7])
			moves[i].Healing = parseInt(v[8])
			moves[i].CritRate = parseInt(v[9])
			moves[i].AilmentChance = parseInt(v[10])
			moves[i].FlinchChance = parseInt(v[11])
			moves[i].StatChance = parseInt(v[12])
		}
	}
	// get all stat changes that moves do
	move_stat_change := getCsvReader("data/move_meta_stat_changes.csv")
	for {
		record, err := move_stat_change.Read()
		if err == io.EOF {
			break
		}
		mid := parseInt(record[0])
		stat := parseInt(record[1])
		stages := parseInt(record[2])
		for i, m := range moves {
			if m.Id != mid {
				continue
			}
			(&moves[i]).AffectedStat = statNames[stat]
			(&moves[i]).StatChange = stages
			break
		}
	}
	output += "var AllMoves = []MoveData{\n"
	for _, m := range moves {
		flags := strings.Join(m.Flags, " | ")
		if len(m.Flags) == 0 {
			flags = "0"
		}
		affectedStat := m.AffectedStat
		if len(m.AffectedStat) == 0 {
			affectedStat = "0"
		}
		output += fmt.Sprintf("\t{%q, %d, %s, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %d, %s, %s, %d, %s, %s},\n",
			m.Name, m.Type, m.DamageClass, m.Targets, m.Priority, m.Power, m.Accuracy, m.PP,
			m.MinHits, m.MaxHits, m.MinTurns, m.MaxTurns, m.Drain, m.Healing, m.CritRate,
			m.AilmentChance, m.FlinchChance, m.StatChance, flags, affectedStat, m.StatChange, m.Ailment, m.MetaCategory)
	}
	output += "}\n\n"
	// Add move constants
	output += "// Move constants for quick and easy reference\nconst (\n" +
		"MoveNone MoveId = iota\n"
	for _, m := range moves {
		name := cleanName(m.Name)
		output += fmt.Sprintf("\tMove%s\n", name)
	}
	output += ")\n\n"
	// Generate hold item data
	/* id,identifier (we only care about items with flag 4, 5, or 7)
	1,countable
	2,consumable
	3,usable-overworld
	4,usable-in-battle
	5,holdable
	6,holdable-passive
	7,holdable-active
	8,underground */
	log.Println("Getting all items/item flags")
	item_flags_csv := getCsvReader("data/item_flag_map.csv")
	records, err = item_flags_csv.ReadAll()
	if err != nil {
		log.Panicln(err)
	}
	itemFlagMap := map[int]string{
		2: "FlagConsumable",
		4: "FlagUsableInBattle",
		5: "FlagHoldable",
		6: "FlagHoldablePassive",
		7: "FlagHoldableActive",
	}
	item_flags := make(map[string][]string)
	for _, r := range records {
		v := parseInt(r[1])
		if v == 2 || v >= 4 && v <= 7 {
			item_flags[r[0]] = append(item_flags[r[0]], itemFlagMap[v])
		}
	}
	output += "// Item constants for quick and easy reference\n"
	item_names_csv := getCsvReader("data/item_names.csv")
	item_vars := make(map[string]string) // data IDs to constants
	item_names := make(map[string]string)
	records, err = item_names_csv.ReadAll()
	if err != nil {
		log.Panicln(err)
	}
	// Add item names
	output += "const (\nItemNone Item = iota\n"
	for _, r := range records {
		if _, ok := item_flags[r[0]]; !ok {
			continue
		}
		if parseInt(r[1]) != EnglishLanguageID {
			continue
		}
		name := r[2]
		item_names[r[0]] = name
		varname := fmt.Sprintf("Item%s", cleanName(name))
		item_vars[r[0]] = varname
		output += fmt.Sprintf("%s\n", varname)
	}
	output += ")\n"

	// HACK: data is missing these flags on items in these categories for some reason
	itemCategoryImplicitFlags := map[int][]string{
		3:  {"FlagUsableInBattle", "FlagConsumable", "FlagHoldable"},
		5:  {"FlagUsableInBattle", "FlagConsumable", "FlagHoldable"},
		6:  {"FlagUsableInBattle", "FlagConsumable", "FlagHoldable"},
		7:  {"FlagUsableInBattle", "FlagConsumable", "FlagHoldable"},
		17: {"FlagHoldable", "FlagHoldablePassive"},
		19: {"FlagHoldablePassive"},
	}

	items := make([]data_item, 0)
	items_csv := getCsvReader("data/items.csv")
	records, err = items_csv.ReadAll()
	if err != nil {
		log.Panicln(err)
	}
	for _, r := range records {
		if _, ok := item_flags[r[0]]; !ok {
			continue
		}
		item := data_item{
			ID:            r[0],
			Identifier:    r[1],
			CategoryID:    parseInt(r[2]),
			FlingPower:    parseInt(r[4]),
			FlingEffectID: parseInt(r[5]),
			Flags:         item_flags[r[0]],
		}
		if impliedFlags, ok := itemCategoryImplicitFlags[item.CategoryID]; ok {
			item.Flags = append(item.Flags, impliedFlags...)
		}
		items = append(items, item)
	}
	// Add item data to generated output
	output += "// A collection of all items in the game\n" + "var AllItems = []ItemData{\n"
	for _, item := range items {
		output += fmt.Sprintf("{Name: \"%s\", Category: %d, FlingPower: %d, FlingEffect: %d, Flags: %s},\n",
			item_names[item.ID], item.CategoryID, item.FlingPower, item.FlingEffectID, strings.Join(item.Flags, " | "))
	}
	output += "}\n\n"

	// create table of levels to the exp at that level

	log.Println("Getting experience table")
	experience_csv := getCsvReader("data/experience.csv")

	slow_leveling := make([]int, 101)
	med_fast_leveling := make([]int, 101)
	fast_leveling := make([]int, 101)
	med_slow_leveling := make([]int, 101)
	erratic_leveling := make([]int, 101)
	fluctuating_leveling := make([]int, 101)

	output += "//A table of levels mapped to the total experience at that level for each growth rate\n" +
		"var ExpTable = map[int]map[int]int{\n"

	for {
		record, err := experience_csv.Read()

		if err == io.EOF {
			break
		}

		growth_rate_id := parseInt(record[0])
		level := parseInt(record[1])
		experience := parseInt(record[2])
		switch growth_rate_id {
		case 1:
			slow_leveling[level] = experience
		case 2:
			med_fast_leveling[level] = experience
		case 3:
			fast_leveling[level] = experience
		case 4:
			med_slow_leveling[level] = experience
		case 5:
			erratic_leveling[level] = experience
		case 6:
			fluctuating_leveling[level] = experience
		}
	}

	output += createLevelTableStringFromArray("GrowthSlow", slow_leveling) + ","
	output += createLevelTableStringFromArray("GrowthMediumFast", med_fast_leveling) + ","
	output += createLevelTableStringFromArray("GrowthFast", fast_leveling) + ","
	output += createLevelTableStringFromArray("GrowthMediumSlow", med_slow_leveling) + ","
	output += createLevelTableStringFromArray("GrowthErratic", erratic_leveling) + ","
	output += createLevelTableStringFromArray("GrowthFluctuating", fluctuating_leveling) + ","
	output += "}\n\n"

	// Natures
	log.Println("Generating code for Natures")
	// read natures data
	natures_csv := getCsvReader("data/natures.csv")
	natures := []data_nature{}
	for {
		record, err := natures_csv.Read()
		if err == io.EOF {
			break
		}

		natures = append(natures, data_nature{
			id:       parseInt(record[0]),
			name:     strings.Title(record[1]),
			statdown: parseInt(record[2]),
			statup:   parseInt(record[3]),
		})
	}
	sort.Slice(natures, func(i, j int) bool {
		return natures[i].name < natures[j].name
	})
	// generate Nature Constants
	log.Println("Generating Nature constants")
	output += "const (\n"
	for i, n := range natures {
		if i == 0 {
			output += fmt.Sprintf("Nature%s Nature = iota\n", n.name)
		} else {
			output += fmt.Sprintf("Nature%s\n", n.name)
		}
	}
	output += ")\n\n"
	// generate Nature GetStatModifiers()
	log.Println("Generating Nature GetStatModifiers()")
	output += "// Get the stat modifiers that this nature gives.\n" +
		"func (n Nature) GetStatModifiers() (statUp, statDown int) {\n" +
		"switch n {\n"
	for _, n := range natures {
		output += fmt.Sprintf("case Nature%s:\n", n.name)
		output += fmt.Sprintf("return %s, %s\n", statNames[n.statup], statNames[n.statdown])
	}
	output += "}\n" +
		"panic(\"Unknown nature\")" +
		"}\n\n"
	// generate Nature String()
	log.Println("Generating Nature String()")
	output += "// Get the string name of this Nature.\n" +
		"func (n Nature) String() string {\n" +
		"switch n {\n"
	for _, n := range natures {
		output += fmt.Sprintf("case Nature%s:\n", n.name)
		output += fmt.Sprintf("return \"%s\"\n", n.name)
	}
	output += "}\n" +
		"panic(\"Unknown nature\")" +
		"}\n\n"

	createCodeOutput(output)

	// run gofmt on generated code
	log.Println("Formatting generated code...")
	cmd := exec.Command("gofmt", "-w", "pokedex_GEN.go")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to format generated code: %v", err)
	}
}
