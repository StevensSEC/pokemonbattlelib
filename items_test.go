package pokemonbattlelib

import "testing"

func TestUseItem(t *testing.T) {
	i := GetItem(ITEM_POTION)
	if i.Name != "Potion" {
		t.Errorf("expected item to be named 'Potion', received %v", i.Name)
	}
	p := GeneratePokemon(1)
	p.CurrentHP = 50
	p.Stats = [6]uint{100, 0, 0, 0, 0, 0}
	transactions := p.UseItem(&i)
	if v, ok := transactions[0].(HealTransaction); !ok {
		t.Fatal("expected potion to create 1 heal transaction")
	} else {
		if v.Target != p {
			t.Errorf("expect item to be used on correct Pokemon")
		}
		if v.Amount != 20 {
			t.Errorf("expect potion to create heal transaction for 20 HP")
		}
		want := "Bulbasaur restored 20 HP."
		if got := v.BattleLog(); want != got {
			t.Errorf("Expected battle log to be %v, received %v\n.", want, got)
		}
	}
}
