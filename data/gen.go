package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const ENGLISH_LANGUAGE_ID = 9
const NATIONAL_DEX_ID = 1
const HIGHEST_GEN = 4

type data_pokemon struct {
	Identifier     string
	SpeciesId      int
	Height         int
	Weight         int
	BaseExperience int

	Name   string
	NatDex uint16
}

type data_move_flags uint32

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
	Flags       data_move_flags
}

type data_item struct {
	ID            string
	Identifier    string
	CategoryID    int
	FlingPower    int
	FlingEffectID int
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
	var err error
	path, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("Current directory: %s\n", path)
	output := ""

	// get all valid version ids
	valid_version_groups := []int{}
	version_groups_csv := getCsvReader("data/version_groups.csv")
	for {
		record, err := version_groups_csv.Read()
		if err == io.EOF {
			break
		}
		gen_id := parseInt(record[2])
		if gen_id <= HIGHEST_GEN {
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
		pokemon = append(pokemon, data_pokemon{
			Identifier:     record[1],
			SpeciesId:      species_id,
			Height:         height,
			Weight:         weight,
			BaseExperience: baseexp,
		})
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
		if lang != ENGLISH_LANGUAGE_ID {
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

	// find national dex numbers
	log.Println("Getting Pokemon dex numbers")
	pkmn_dex_nums_csv := getCsvReader("data/pokemon_dex_numbers.csv")
	for {
		record, err := pkmn_dex_nums_csv.Read()
		if err == io.EOF {
			break
		}
		dexid := parseInt(record[1])
		if dexid != NATIONAL_DEX_ID {
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
	output += "var ALL_POKEMON = []Pokemon{\n"
	for _, p := range pokemon {
		// fmt.Printf("%v\n", p)
		output += "{\n" +
			fmt.Sprintf("NatDex:%d,\n", p.NatDex) +
			fmt.Sprintf("Level:%d,\n", uint8(1)) +
			// fmt.Sprintf("Ability:%s,\n", "new(Ability)") +
			// fmt.Sprintf("TotalExperience:%d,\n", uint(0)) +
			fmt.Sprintf("Gender:%s,\n", "Genderless") +
			fmt.Sprintf("IVs:%s,\n", "[6]uint8{0, 0, 0, 0, 0, 0}") +
			fmt.Sprintf("EVs:%s,\n", "[6]uint8{0, 0, 0, 0, 0, 0}") +
			// fmt.Sprintf("Nature:%s,\n", "new(Nature)") +
			fmt.Sprintf("Stats:%s,\n", "[6]uint{0, 0, 0, 0, 0, 0}") +
			// fmt.Sprintf("CurrentHP:%d,\n", uint(0)) +
			// fmt.Sprintf("HeldItem:%v,\n", "new(Item)") +
			fmt.Sprintf("Moves:%s,\n", "[4]*Move{}") +
			// fmt.Sprintf("Friendship:%d,\n", uint8(0)) +
			// fmt.Sprintf("OriginalTrainerID: %d,\n", uint16(0)) +
			"},\n"
	}
	output += "}\n\n" +
		"// A map of national pokedex numbers to pokemon names.\n" +
		"var PokemonNames = map[uint16]string{\n"
	for _, p := range pokemon {
		if p.NatDex == 0 {
			continue
		}
		output += fmt.Sprintf("\t%d: \"%s\",\n", p.NatDex, p.Name)
	}
	output += "}\n\n"
	// find all moves
	moves := []data_move{}
	log.Println("Getting all available moves")
	moves_csv := getCsvReader("data/moves.csv")
	moveMap := map[int]string{
		1: "Status",
		2: "Physical",
		3: "Special",
	}
	for {
		record, err := moves_csv.Read()
		if err == io.EOF {
			break
		}
		gid := parseInt(record[2])
		if gid > HIGHEST_GEN {
			continue
		}
		mid := parseInt(record[0])
		moveType := parseInt(record[3])
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
		if lang != ENGLISH_LANGUAGE_ID {
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
			(&moves[i]).Flags |= data_move_flags(1 << flag)
			break
		}
	}
	output += "var ALL_MOVES = []Move{\n"
	for _, p := range moves {
		output += fmt.Sprintf("\t{ID: %d, Name: %q, Type: %d, Category: %s, CurrentPP: %d, MaxPP: %d,"+
			" Priority: %d, Power: %d, Accuracy: %d},\n", p.Id, p.Name, p.Type, p.DamageClass, p.PP, p.PP, p.Priority, p.Power, p.Accuracy)
	}
	output += "}\n\n"
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
	records, err := item_flags_csv.ReadAll()
	if err != nil {
		log.Panicln(err)
	}
	valid_items := make(map[string]bool)
	for _, r := range records {
		if r[1] == "4" || r[1] == "5" || r[1] == "7" {
			valid_items[r[0]] = true
		}
	}
	output += "// Create item constant enum for quick reference\nconst (\n"
	item_names_csv := getCsvReader("data/item_names.csv")
	item_names := make(map[string]string)
	records, err = item_names_csv.ReadAll()
	if err != nil {
		log.Panicln(err)
	}
	// Add item names
	for _, r := range records {
		if _, ok := valid_items[r[0]]; !ok {
			continue
		}
		if parseInt(r[1]) != ENGLISH_LANGUAGE_ID {
			continue
		}
		item_names[r[0]] = r[2]
		name := strings.ToUpper(r[2])
		name = strings.ReplaceAll(name, "É", "E")
		name = strings.ReplaceAll(name, " ", "_")
		re, err := regexp.Compile(`[^a-zA-Z_0-9]`)
		if err != nil {
			log.Panicln(err)
		}
		name = re.ReplaceAllString(name, "")
		output += fmt.Sprintf("\tITEM_%s = %v\n", name, r[0])
	}
	output += ")\n"
	items := make([]data_item, 0)
	items_csv := getCsvReader("data/items.csv")
	records, err = items_csv.ReadAll()
	if err != nil {
		log.Panicln(err)
	}
	for _, r := range records {
		if _, ok := valid_items[r[0]]; !ok {
			continue
		}
		item := data_item{
			ID:            r[0],
			Identifier:    r[1],
			CategoryID:    parseInt(r[2]),
			FlingPower:    parseInt(r[4]),
			FlingEffectID: parseInt(r[5]),
		}
		items = append(items, item)
	}
	// Add item data to generated output
	output += "// A collection of all items in the game\n" + "var ALL_ITEMS = []Item{\n"
	for _, item := range items {
		output += fmt.Sprintf("\t{ID: %s, Name: \"%s\", Category: %d, FlingPower: %d, FlingEffect: %d},\n",
			item.ID, item_names[item.ID], item.CategoryID, item.FlingPower, item.FlingEffectID)
	}
	output += "}\n\n"
	createCodeOutput(output)
	// run gofmt on generated code
	log.Println("Formatting generated code...")
	cmd := exec.Command("gofmt", "-w", "pokedex_GEN.go")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to format generated code: %v", err)
	}
}
