package pokemonbattlelib

import "testing"

func TestNewPokemon(t *testing.T) {
	p := NewPokemon(8)
	if p.NatDex != 8 {
		t.Fail()
	}
}
