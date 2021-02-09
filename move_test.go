package pokemonbattlelib

import "testing"

func TestMoveString(t *testing.T) {
	m := Move{
		Name:     "Shadow Ball",
		Type:     8,
		Category: 3,
		Max_PP:   15,
		Priority: 0,
		Power:    80,
		Accuracy: 100,
	}
	if m.String() != "Shadow Ball\nType: 8, Power: 80, Accuracy: 100\n" {
		t.Fail()
	}
}

func TestMoveCategoryString(t *testing.T) {
	var c1, c2, c3 MoveCategory = 1, 2, 3

	if c1.String() != "Status" {
		t.Fail()
	}
	if c2.String() != "Physical" {
		t.Fail()
	}
	if c3.String() != "Special" {
		t.Fail()
	}
}
