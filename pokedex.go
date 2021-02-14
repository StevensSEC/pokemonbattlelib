package pokemonbattlelib

import "log"

//go:generate go run data/gen.go

// Creates a new Pokemon given its national dex number
func NewPokemon(natdex uint16) *Pokemon {
	for _, p := range ALL_POKEMON {
		if p.NatDex == natdex {
			return &p
		}
	}
	// Not exactly the best way to handle this
	log.Panicf("unknown Pokedex number %v\n", natdex)
	return nil
}

// Get a Pokemon Move by ID.
func GetMove(id int) Move {
	for _, m := range ALL_MOVES {
		if m.ID == id {
			return m
		}
	}
	panic("move not found")
}
