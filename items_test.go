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
	p := GetPokemon(1)
	p.CurrentHP = 50
	p.Stats = [6]uint{100, 0, 0, 0, 0, 0}
	testLog(t, p.UseItem(&i)[0], "Bulbasaur restored 20 HP.")
	i = GetItem(ITEM_ANTIDOTE)
	p.StatusEffects = STATUS_POISON
	testLog(t, p.UseItem(&i)[0], "Bulbasaur's <STATUS> was cured.")
}

func TestUseMoveItem(t *testing.T) {
	p := GetPokemon(1)
	pound := GetMove(1)
	pound.CurrentPP = 15
	pound.MaxPP = 30
	p.Moves[0] = &pound
	i := GetItem(ITEM_ETHER)
	testLog(t, p.UseMoveItem(&i, p.Moves[0])[0], "Pound restored 10 PP.")
}

func TestUseVitamins(t *testing.T) {
	p := GetPokemon(1)
	i := GetItem(ITEM_CALCIUM)
	tt := p.UseVitamins(&i)
	testLog(t, tt[0], "Bulbasaur gained 10 <STAT> effort.")
	testLog(t, tt[1], "Bulbasaur gained 5 friendship.")
	p.Friendship = 200
	p.EVs[STAT_ATK] = 200
	tt = p.UseVitamins(&i)
	testLog(t, tt[0], "Bulbasaur gained 2 friendship.")
}

func TestHeldItemStatBoost(t *testing.T) {
	p := GetPokemon(149)
	outrage := GetMove(200)
	i := GetItem(ITEM_DRACO_PLATE)
	if n := p.GetDamageMultiplier(&i, &outrage); n != 1.20 {
		t.Errorf("expected Draco Plate to boost damage to 1.2x, received %vx\n", n)
	}
}
