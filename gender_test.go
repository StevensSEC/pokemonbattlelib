package pokemonbattlelib

import "testing"

func TestGenderString(t *testing.T) {
	var g1, g2, g3 Gender = Genderless, Female, Male

	if g1.String() != "" {
		t.Fail()
	}

	if g2.String() != "\u2640" {
		t.Fail()
	}

	if g3.String() != "\u2642" {
		t.Fail()
	}
}
