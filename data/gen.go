package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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
	Power    int
	PP       int
	Accuracy int
	Priority int
	// see: move_targets.csv
	Targets     int
	DamageClass int
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
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panicln(err)
	}
	return n
}

func getCsvReader(path string) *csv.Reader {
	log.Printf("Reading csv: %s\n", path)
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

func createCodeOutput() *os.File {
	file, err := os.Create("pokedex_GEN.go")
	if err != nil {
		log.Panicln(err)
	}
	_, err = file.WriteString("// Code generated - DO NOT EDIT.\n" +
		"// Regenerate with `go generate`.\n\n" +
		"package pokemonbattlelib\n\n")
	return file
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("current directory: %s\n", path)

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
	log.Println("finding pokemon names")
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
	log.Println("finding pokemon dex numbers")
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

	// print out pokemon
	log.Println("generating code for pokemon")
	output := createCodeOutput()
	_, err = output.WriteString("var ALL_POKEMON = []Pokemon{\n")
	if err != nil {
		log.Panicln(err)
	}
	for _, p := range pokemon {
		// fmt.Printf("%v\n", p)
		// TODO: add more fields
		_, err = output.WriteString(fmt.Sprintf("\t{ NatDex: %d },\n", p.NatDex))
		if err != nil {
			log.Panicln(err)
		}
	}
	_, err = output.WriteString("}\n\n" +
		"// A map of national pokedex numbers to pokemon names.\n" +
		"var PokemonNames = map[uint16]string{\n")
	if err != nil {
		log.Panicln(err)
	}
	for _, p := range pokemon {
		// fmt.Printf("%v\n", p)
		if p.NatDex == 0 {
			continue
		}
		_, err = output.WriteString(fmt.Sprintf("\t%d: \"%s\",\n", p.NatDex, p.Name))
		if err != nil {
			log.Panicln(err)
		}
	}
	_, err = output.WriteString("}\n\n")
	if err != nil {
		log.Panicln(err)
	}
	// find all moves
	moves := []data_move{}
	log.Println("finding availble moves")
	moves_csv := getCsvReader("data/moves.csv")
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
			Power:       power,
			PP:          pp,
			Accuracy:    accuracy,
			Priority:    priority,
			Targets:     targets,
			DamageClass: damageClass,
			Effect:      effect,
		})
	}

	// find all the move names
	log.Println("finding move names")
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
	log.Println("finding move flags")
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

	log.Println("TODO: generate code for moves")

	// Generate hold item data
	items := make([]data_item, 0)
	items_csv := getCsvReader("data/items.csv")
	records, err := items_csv.ReadAll()
	if err != nil {
		log.Panicln(err)
	}
	for _, r := range records[1:] {
		item := new(data_item)
		item.ID = r[0]
		item.Identifier = r[1]

	}
}
