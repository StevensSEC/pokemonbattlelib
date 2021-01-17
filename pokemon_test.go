package pokemonbattlelib

import "testing"

func TestPokemonName(t *testing.T) {
	p := Pokemon{
		NatDex: 1,
	}
	if p.GetName() != "Bulbasaur" {
		t.Fail()
	}
}
