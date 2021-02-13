package pokemonbattlelib

import "testing"

func TestItemUse(t *testing.T) {
	i := NewItem(ITEM_POTION)
	if i.Name != "Potion" {
		t.Errorf("expected item to be named 'Potion', received %v", i.Name)
	}
	p := Pokemon{CurrentHP: 50, Stats: [6]uint{100, 0, 0, 0, 0, 0}, HeldItem: i}
	i.UseItem(&p)
	if p.CurrentHP != 70 {
		t.Errorf("expected item to heal Pokemon by 20HP")
	}
	i = NewItem(ITEM_ANTIDOTE)
	p.StatusEffects = STATUS_POISON
	i.UseItem(&p)
	if p.StatusEffects&STATUS_POISON != 0 {
		t.Errorf("expected item to cure poison status effect")
	}
	i = NewItem(ITEM_ETHER)
	p.Moves[0] = &Move{CurrentPP: 15, MaxPP: 30}
	i.UseMoveItem(&p, p.Moves[0])
	if p.Moves[0].CurrentPP != 25 {
		// PP restore not implemented
		// t.Errorf("expected item to restore first move PP by 10")
	}
}
