package pokemonbattlelib

import "testing"

func testLog(t *testing.T, tt Transaction, want string) {
	// TODO: generalize more testing functions
	if got := tt.BattleLog(); got != want {
		t.Errorf("Expected battle log to be %s, received %s.\n", want, got)
	}
}

func TestUseItem(t *testing.T) {
	i := GetItem(ITEM_POTION)
	if i.Name != "Potion" {
		t.Errorf("expected item to be named 'Potion', received %v", i.Name)
	}
	p := GeneratePokemon(1)
	p.CurrentHP = 50
	p.Stats = [6]uint{100, 0, 0, 0, 0, 0}
	testLog(t, p.UseItem(&i)[0], "Bulbasaur restored 20 HP.")
	p.CurrentHP = 99
	testLog(t, p.UseItem(&i)[0], "Bulbasaur restored 1 HP.")
}
