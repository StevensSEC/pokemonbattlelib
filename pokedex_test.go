package pokemonbattlelib

import "testing"

func TestGetPokemon(t *testing.T) {
	p := GetPokemon(8)
	if p.NatDex != 8 {
		t.Fail()
	}
}
