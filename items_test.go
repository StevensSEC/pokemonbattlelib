package pokemonbattlelib

import "testing"

func TestUseItem(t *testing.T) {
	i := GetItem(ITEM_POTION)
	if i.Name != "Potion" {
		t.Errorf("expected item to be named 'Potion', received %v", i.Name)
	}
	p := Pokemon{CurrentHP: 50, Stats: [6]uint{100, 0, 0, 0, 0, 0}}
	i.UseItem(&p)
	if p.CurrentHP != 70 {
		t.Errorf("expected item to heal Pokemon by 20HP")
	}
	i = GetItem(ITEM_ANTIDOTE)
	p.StatusEffects = STATUS_POISON
	i.UseItem(&p)
	if p.StatusEffects&STATUS_POISON != 0 {
		t.Errorf("expected item to cure poison status effect")
	}
}

func TestUseMoveItem(t *testing.T) {
	p := GetPokemon(1)
	p.Moves[0] = &Move{CurrentPP: 15, MaxPP: 30}
	i := GetItem(ITEM_ETHER)
	i.UseMoveItem(&p, p.Moves[0])
	if p.Moves[0].CurrentPP != 25 {
		t.Errorf("expected item to restore first move PP by 10")
	}
}

func TestHeldItemStatBoost(t *testing.T) {
	p := GetPokemon(149)
	outrage := GetMove(200)
	i := GetItem(ITEM_DRACO_PLATE)
	if n := i.GetDamageMultiplier(&p, &outrage); n != 1.20 {
		t.Errorf("expected Draco Plate to boost damage to 1.2x, received %vx\n", n)
	}
}
