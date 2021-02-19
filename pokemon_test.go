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

func TestHealPokemon(t *testing.T) {
	p := GetPokemon(1)
	p.Stats[0] = 100
	p.CurrentHP = 0
	p.heal(25)
	if p.CurrentHP != 25 {
		t.Errorf("expected Pokemon to have 25 HP, received %v", p.CurrentHP)
	}
	p.heal(100)
	if p.CurrentHP != 100 {
		t.Errorf("expected Pokemon to have 100 HP, received %v", p.CurrentHP)
	}
}

func TestDamagePokemon(t *testing.T) {
	p := GetPokemon(1)
	p.Stats[0] = 100
	p.CurrentHP = 100
	p.damage(25)
	if p.CurrentHP != 75 {
		t.Errorf("expected Pokemon to have 75 HP, received %v", p.CurrentHP)
	}
	p.damage(100)
	if p.CurrentHP != 0 {
		t.Errorf("expected Pokemon to have 0 HP, received %v", p.CurrentHP)
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
