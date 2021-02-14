package pokemonbattlelib

//go:generate go run data/gen.go

// Get a new pokemon by nation dex number.
func GetPokemon(natdex uint16) Pokemon {
	for _, p := range ALL_POKEMON {
		if p.NatDex == natdex {
			return p
		}
	}
	// Not exactly the best way to handle this
	return Pokemon{}
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
