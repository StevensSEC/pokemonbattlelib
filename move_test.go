package pokemonbattlelib

import "testing"

func TestMoveString(t *testing.T) {
	m := Move{
		Name: "Shadow Ball",
		Type: "Ghost",
		Category: "Special",
		Max_PP: 15,
		Priority: 0,
		Power: 80,
		Accuracy: 100,
	}
	if m.String() != "Shadow Ball\nType: Ghost, Power: 80, Accuracy: 100\n"{
		t.Fail();
	}
}