package pokemonbattlelib

import (
	"fmt"
	"testing"
)

func handleTurn(b *Battle) interface{} {
	fmt.Println("Called during player 1's turn")
	return nil
}

func onBattleEnd() {
	fmt.Println("Called when a player wins the battle")
}

func ExampleAgent() {
	p1 := NewAgent(handleTurn)
	p := Pokemon{1}
	p1.AddPokemon(p)
	p1.AddPokemon(Pokemon{NatDex: 7})
	p1.RegisterHook(BATTLE_END, onBattleEnd)
}

func TestAgent(t *testing.T) {
	p1 := NewAgent(func(b *Battle) interface{} {
		if b.State == BEFORE_PLAYER1_TURN {
			return "attack"
		}
		return "item"
	})
	p2 := NewAgent(func(b *Battle) interface{} {
		if b.State == BEFORE_PLAYER2_TURN {
			return "swap"
		}
		return "attack"
	})
	p1.AddPokemon(Pokemon{1}, Pokemon{11}, Pokemon{14})
	if len(p1.Party) != 3 {
		t.Errorf("expected 3 Pokemon in player 1 party, got %d", len(p1.Party))
	}
	if p1.Party[2].NatDex != 14 {
		t.Errorf("expected last Pokemon in player 1 party to be 14, got %d", p1.Party[2].NatDex)
	}
	p2.AddPokemon(Pokemon{4}, Pokemon{25}, Pokemon{47}, Pokemon{52})
	p2.RemovePokemon(2)
	if len(p2.Party) != 3 {
		t.Errorf("expected 3 Pokemon in player 2 party, got %d", len(p2.Party))
	}
	b := NewBattle()
	b.AddAgents(p1, p2)
	// Test event hooks
	b.Start()
}
