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
