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
