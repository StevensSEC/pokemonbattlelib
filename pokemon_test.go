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

func TestPokemonConstructor(t *testing.T) {
	tests := []struct {
		pkmn Pokemon
		want Pokemon
	}{
		{
			pkmn: GeneratePokemon(uint16(393), uint8(5)), // constructor w/ dex number and level
			want: Pokemon{
				NatDex:          393, // piplup if you're curious
				Level:           5,
				TotalExperience: 135,
				CurrentHP:       20,
				IVs:             [6]uint8{0, 0, 0, 0, 0, 0},
				EVs:             [6]uint8{0, 0, 0, 0, 0, 0},
				Nature:          GetNatureTable()["hardy"],
				Stats:           [6]uint{20, 10, 10, 10, 10, 10},
			},
		},
		{
			pkmn: GeneratePokemon(uint16(393), uint(135)), // constructor w/ dex number and total exp
			want: Pokemon{
				NatDex:          393,
				Level:           5,
				TotalExperience: 135,
				CurrentHP:       20,
				IVs:             [6]uint8{0, 0, 0, 0, 0, 0},
				EVs:             [6]uint8{0, 0, 0, 0, 0, 0},
				Nature:          GetNatureTable()["hardy"],
				Stats:           [6]uint{20, 10, 10, 10, 10, 10},
			},
		},
		{
			pkmn: GeneratePokemon(uint16(393), uint8(5), [6]uint8{31, 31, 31, 31, 31, 31}), // constructor w/ dex number, level, ivs
			want: Pokemon{
				NatDex:          393,
				Level:           5,
				TotalExperience: 135,
				CurrentHP:       20,
				IVs:             [6]uint8{31, 31, 31, 31, 31, 31},
				EVs:             [6]uint8{0, 0, 0, 0, 0, 0},
				Nature:          GetNatureTable()["hardy"],
				Stats:           [6]uint{21, 11, 11, 11, 11, 11},
			},
		},
		{
			pkmn: GeneratePokemon(uint16(393), uint8(5), [6]uint8{0, 0, 0, 0, 0, 0}, [6]uint8{0, 252, 6, 0, 0, 252}), // constructor w/ dex number, level, ivs, evs
			want: Pokemon{
				NatDex:          393,
				Level:           5,
				TotalExperience: 135,
				CurrentHP:       20,
				IVs:             [6]uint8{0, 0, 0, 0, 0, 0},
				EVs:             [6]uint8{0, 252, 6, 0, 0, 252},
				Nature:          GetNatureTable()["hardy"],
				Stats:           [6]uint{20, 13, 10, 10, 10, 13},
			},
		},
		{
			pkmn: GeneratePokemon(uint16(393), uint8(5), [6]uint8{0, 0, 0, 0, 0, 0}, [6]uint8{0, 0, 0, 0, 0, 0}, GetNatureTable()["adamant"]), // constructor w/ dex number, level, ivs, evs, nature
			want: Pokemon{
				NatDex:          393,
				Level:           5,
				TotalExperience: 135,
				CurrentHP:       20,
				IVs:             [6]uint8{0, 0, 0, 0, 0, 0},
				EVs:             [6]uint8{0, 0, 0, 0, 0, 0},
				Nature:          GetNatureTable()["adamant"],
				Stats:           [6]uint{20, 11, 10, 9, 10, 9},
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Pokemon constructor test"), func(t *testing.T) {
			t.Parallel()
			got := tt.pkmn
			if !got.Equals(tt.want) {
				t.Errorf("Pokemon constructor got %v, want %v", got.VerboseString(), tt.want.VerboseString())
			}
		})
	}
}

func TestPokemonConstructorAccurateResult(t *testing.T) {
	tests := []struct {
		pkmn Pokemon
		want Pokemon
	}{
		{
			pkmn: GeneratePokemon(uint16(445), uint8(78), [6]uint8{24, 12, 30, 16, 23, 5}, [6]uint8{74, 190, 91, 48, 84, 23}, GetNatureTable()["adamant"]),
			want: Pokemon{
				NatDex:          445,
				Level:           78,
				TotalExperience: 593190,
				CurrentHP:       289,
				IVs:             [6]uint8{24, 12, 30, 16, 23, 5},
				EVs:             [6]uint8{74, 190, 91, 48, 84, 23},
				Nature:          GetNatureTable()["adamant"],
				Stats:           [6]uint{289, 278, 193, 135, 171, 171},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Pokemon accurate constructor test"), func(t *testing.T) {
			t.Parallel()
			got := tt.pkmn
			if !got.Equals(tt.want) {
				t.Errorf("Pokemon constructor got %v, want %v", got.VerboseString(), tt.want.VerboseString())
			}
		})
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
			t.Parallel()
			got := fmt.Sprintf("%s", tt.pkmn)
			if got != tt.want {
				t.Errorf("Pokemon Stringer %d got %v, want %v", tt.pkmn.NatDex, got, tt.want)
			}
		})
	}
}
