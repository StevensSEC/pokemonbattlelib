package pokemonbattlelib

import (
	"fmt"
	"testing"
)

func TestGetPokemon(t *testing.T) {
	p := GetPokemon(8)
	if p.NatDex != 8 {
		t.Errorf("expected Pokemon national dex to be 8, received %v", p.NatDex)
	}
}

func TestPokemonName(t *testing.T) {
	p := GetPokemon(1)
	if p.GetName() != "Bulbasaur" {
		t.Errorf("expected Pokemon name to be Bulbasaur, received %v", p.GetName())
	}
}

func TestPokemonStringer(t *testing.T) {
	tests := []struct {
		pkmn Pokemon
		want string
	}{
		{
			pkmn: Pokemon{
				NatDex:    1,
				Gender:    Female,
				Level:     5,
				CurrentHP: 11,
				Stats:     [6]uint{11, 6, 5, 6, 5, 5},
			},
			want: "Bulbasaur♀\tLv5\nHP: 11/11\n",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Pokemon Stringer: %d", tt.pkmn.NatDex), func(t *testing.T) {
			got := fmt.Sprintf("%s", tt.pkmn)
			if got != tt.want {
				t.Errorf("Pokemon Stringer %d got %v, want %v", tt.pkmn.NatDex, got, tt.want)
			}
		})
	}
}
