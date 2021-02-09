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
	Type     int
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

func getCsvReader(path string) *csv.Reader {
	log.Printf("Reading csv: %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}
	reader := csv.NewReader(file)
	reader.ReuseRecord = true
	reader.Read() // skip header line
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

func createCodeOutput(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		log.Panicln(err)
	}
	file.WriteString("// Code generated - DO NOT EDIT.\n")
	file.WriteString("// Regenerate with `go generate`.\n\n")
	file.WriteString("package pokemonbattlelib\n\n")
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
		gen_id, err := strconv.Atoi(record[2])
		if gen_id <= HIGHEST_GEN {
			group_id, _ := strconv.Atoi(record[0])
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
		group_id, err := strconv.Atoi(record[1])
		if contains(valid_version_groups, group_id) {
			vid, _ := strconv.Atoi(record[0])
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
		vid, err := strconv.Atoi(record[1])
		if contains(valid_versions, vid) {
			pid, _ := strconv.Atoi(record[0])
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
		id, err := strconv.Atoi(record[0])
		if !contains(valid_pkmn_ids, id) {
			continue
		}
		species_id, err := strconv.Atoi(record[2])
		height, err := strconv.Atoi(record[3])
		weight, err := strconv.Atoi(record[4])
		baseexp, err := strconv.Atoi(record[5])
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
		lang, err := strconv.Atoi(record[1])
		if lang != ENGLISH_LANGUAGE_ID {
			continue
		}
		sid, err := strconv.Atoi(record[0])

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
		dexid, err := strconv.Atoi(record[1])
		if dexid != NATIONAL_DEX_ID {
			continue
		}
		sid, err := strconv.Atoi(record[0])

		for i, p := range pokemon {
			if p.SpeciesId != sid {
				continue
			}
			dexnum, _ := strconv.Atoi(record[2])
			(&pokemon[i]).NatDex = uint16(dexnum)
			break
		}
	}

	// print out pokemon
	log.Println("generating code for pokemon")
	output := createCodeOutput("pokedex_GEN.go")
	output.WriteString("var ALL_POKEMON = []Pokemon{\n")
	for _, p := range pokemon {
		// fmt.Printf("%v\n", p)
		// TODO: add more fields
		output.WriteString(fmt.Sprintf("\t{ NatDex: %d },\n", p.NatDex))
	}
	output.WriteString("}\n\n")
	output.WriteString("// A map of national pokedex numbers to pokemon names.\n")
	output.WriteString("var PokemonNames = map[uint16]string{\n")
	for _, p := range pokemon {
		// fmt.Printf("%v\n", p)
		if p.NatDex == 0 {
			continue
		}
		output.WriteString(fmt.Sprintf("\t%d: \"%s\",\n", p.NatDex, p.Name))
	}
	output.WriteString("}\n\n")

	// find all moves
	moves := []data_move{}
	log.Println("finding available moves")
	moves_csv := getCsvReader("data/moves.csv")
	for {
		record, err := moves_csv.Read()
		if err == io.EOF {
			break
		}
		gid, err := strconv.Atoi(record[2])
		if gid > HIGHEST_GEN {
			continue
		}
		mid, err := strconv.Atoi(record[0])
		moveType, err := strconv.Atoi(record[3])
		power, err := strconv.Atoi(record[4])
		pp, err := strconv.Atoi(record[5])
		accuracy, err := strconv.Atoi(record[6])
		priority, err := strconv.Atoi(record[7])
		targets, err := strconv.Atoi(record[8])
		damageClass, err := strconv.Atoi(record[9])
		effect, err := strconv.Atoi(record[10])
		moves = append(moves, data_move{
			Id:          mid,
			Identifier:  record[1],
			Type:        moveType,
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
		lang, err := strconv.Atoi(record[1])
		if lang != ENGLISH_LANGUAGE_ID {
			continue
		}
		mid, err := strconv.Atoi(record[0])

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
		mid, err := strconv.Atoi(record[0])
		flag, err := strconv.Atoi(record[1])
		for i, m := range moves {
			if m.Id != mid {
				continue
			}
			(&moves[i]).Flags |= data_move_flags(1 << flag)
			break
		}
	}

	log.Println("generating code for moves")
	output = createCodeOutput("moves_GEN.go")
	output.WriteString("var ALL_MOVES = []Move{\n")
	for _, p := range moves {

		// fmt.Printf("%v\n", p)
		// TODO: add more fields
		output.WriteString(fmt.Sprintf("\t{ID: %d, Name: %q, Type: %d, Category: %d, Max_PP: %d,"+
			" Priority: %d, Power: %d, Accuracy: %d},\n", p.Id, p.Name, p.Type, p.DamageClass, p.PP, p.Priority, p.Power, p.Accuracy))
	}
	output.WriteString("}")
}
