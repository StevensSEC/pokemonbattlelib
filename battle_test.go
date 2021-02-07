package pokemonbattlelib

import (
	"fmt"
	"testing"
)

func ExampleBattle() {
	b := NewBattle()
	c1 := NewAgent(func(b *Battle) interface{} { return "P1" })
	c1.RegisterHook(BEFORE_PLAYER1_TURN, func() {
		fmt.Println("called before player 1 turn")
	})
	c2 := NewAgent(func(b *Battle) interface{} { return "P2" })
	c2.RegisterHook(BATTLE_END, func() {
		fmt.Println("called after battle ended")
	})
	b.AddAgents(c1, c2)
	b.Start()
}

func TestBattle(t *testing.T) {
	// Write test cases
}
