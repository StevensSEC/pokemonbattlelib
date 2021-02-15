package pokemonbattlelib

import "testing"

func TestNewPokemon(t *testing.T) {
	p := NewPokemon(8)
	if p.NatDex != 8 {
		t.Errorf("expected Pokemon national dex to be 8, received %v", p.NatDex)
	}
}

func TestPokemonName(t *testing.T) {
	p := NewPokemon(1)
	if p.GetName() != "Bulbasaur" {
		t.Errorf("expected Pokemon name to be Bulbasaur, received %v", p.GetName())
	}
}

func TestStringNoNickname(t *testing.T) {
	p := Pokemon{
		NatDex:    1,
		Gender:    Female,
		Level:     5,
		CurrentHP: 11,
		Stats:     [6]uint{11, 6, 5, 6, 5, 5},
	}

	if p.String() != "Bulbasaur\u2640\tLv5\nHP: 11/11\n" {
		t.Fail()
	}
}
