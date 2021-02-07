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

func TestStringNoNickname(t *testing.T) {
    p := Pokemon{
        NatDex: 1,
        Gender: Female,
        Level: 5,
        CurrentHP: 11,
        Stats: [6]uint{11, 6, 5, 6, 5, 5},
    }

    if (p.String() != "Bulbasaur\u2640\tLv5\nHP: 11/11\n") {
        t.Fail()
    }
}

