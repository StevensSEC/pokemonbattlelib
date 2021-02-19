package pokemonbattlelib

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGeneratePokemon(t *testing.T) {
	p := GeneratePokemon(8)
	if p.NatDex != 8 {
		t.Errorf("expected Pokemon national dex to be 8, received %v", p.NatDex)
	}
}

func TestPokemonName(t *testing.T) {
	p := GeneratePokemon(1)
	if p.GetName() != "Bulbasaur" {
		t.Errorf("expected Pokemon name to be Bulbasaur, received %v", p.GetName())
	}
}

func TestPokemonConstructor(t *testing.T) {
	tests := []struct {
		pkmn *Pokemon
		want *Pokemon
	}{
		{
			pkmn: GeneratePokemon(393, WithLevel(5)), // constructor w/ dex number and level
			want: &Pokemon{
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
			pkmn: GeneratePokemon(393, WithTotalExp(135)), // constructor w/ dex number and total exp
			want: &Pokemon{
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
			pkmn: GeneratePokemon(393, WithLevel(5), WithIVs([6]uint8{31, 31, 31, 31, 31, 31})), // constructor w/ dex number, level, ivs
			want: &Pokemon{
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
			pkmn: GeneratePokemon(393, WithLevel(5), WithIVs([6]uint8{0, 0, 0, 0, 0, 0}), WithEVs([6]uint8{0, 252, 6, 0, 0, 252})), // constructor w/ dex number, level, ivs, evs
			want: &Pokemon{
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
			pkmn: GeneratePokemon(393, WithLevel(5), WithIVs([6]uint8{0, 0, 0, 0, 0, 0}), WithEVs([6]uint8{0, 0, 0, 0, 0, 0}), WithNature(GetNatureTable()["adamant"])), // constructor w/ dex number, level, ivs, evs, nature
			want: &Pokemon{
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
			got := tt.pkmn
			if !reflect.DeepEqual(got, tt.pkmn) {
				t.Errorf("Pokemon constructor got %v, want %v", got.VerboseString(), tt.want.VerboseString())
			}
		})
	}
}

func TestPokemonConstructorAccurateResult(t *testing.T) {
	tests := []struct {
		pkmn *Pokemon
		want *Pokemon
	}{
		{
			// see: https://bulbapedia.bulbagarden.net/wiki/Stat, scroll down to 'Example'
			pkmn: GeneratePokemon(445, WithLevel(78), WithIVs([6]uint8{24, 12, 30, 16, 23, 5}), WithEVs([6]uint8{74, 190, 91, 48, 84, 23}), WithNature(GetNatureTable()["adamant"])),
			want: &Pokemon{
				NatDex:          445, // garchomp
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
			got := tt.pkmn
			if got.TotalExperience != tt.want.TotalExperience {
				t.Errorf("Pokemon constructor produced %v for experience, want %v", got.TotalExperience, tt.want.TotalExperience)
			}

			if !reflect.DeepEqual(got.Stats, tt.want.Stats) {
				t.Errorf("Pokemon constructor produced %v for stats, want %v", got.Stats, tt.want.Stats)
			}
		})
	}
}

func TestPokemonCannotLevelBeyondMax(t *testing.T) {
	pkmn := GeneratePokemon(6, WithLevel(MAX_LEVEL))
	defer func() { recover() }()
	pkmn.GainLevels(1)
	t.Errorf("GainLevels did not panic when leveling past max")
}

func TestPokemonCannotDecreaseLevel(t *testing.T) {
	pkmn := GeneratePokemon(393, WithLevel(5))
	defer func() { recover() }()
	pkmn.GainLevels(-1)
	t.Errorf("GainLevels did not panic when attempting to remove levels")
}

func TestPokemonCannotDecreaseExperience(t *testing.T) {
	pkmn := GeneratePokemon(393, WithLevel(5))
	defer func() { recover() }()
	pkmn.GainExperience(-135)
	t.Errorf("GainExperience did not panic when attempting to remove experience points")
}

func TestPokemonLevelsToMaxWhenGainingExpBeyondMax(t *testing.T) {
	pkmn := GeneratePokemon(493, WithLevel(MAX_LEVEL))
	pkmn.GainExperience(100000000000)
	if pkmn.Level != MAX_LEVEL {
		t.Errorf("Expected level to be %d, got %d", MAX_LEVEL, pkmn.Level)
	}
}

func TestPokemonCannotHaveHigherThanMaxLevel(t *testing.T) {
	defer func() { recover() }()
	GeneratePokemon(396, WithLevel(MAX_LEVEL+1))
	t.Errorf("GeneratePokemon did not panic when making a Pokemon with an invalid level")
}

func TestPokemonCannotHaveHigherThanMaxIVs(t *testing.T) {
	defer func() { recover() }()
	GeneratePokemon(396, WithIVs([6]uint8{32, 32, 32, 32, 32, 32}))
	t.Errorf("GeneratePokemon did not panic when making a Pokemon with invalid IVs")
}
func TestPokemonCannotHaveHigherThanTotalEVs(t *testing.T) {
	defer func() { recover() }()
	GeneratePokemon(396, WithEVs([6]uint8{255, 255, 255, 255, 255, 255}))
	t.Errorf("GeneratePokemon did not panic when making a Pokemon with an invalid EVs")
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
