package pokemonbattlelib

type Pokemon struct {
	// National Pokedex Number
	NatDex uint16
	// TODO
}

func (p *Pokemon) GetName() string {
	return PokemonNames[p.NatDex]
}
