package pokemonbattlelib

import "fmt"

func ExampleBattle() {
	b := NewBattle()
	c1 := NewClient(func() interface{} { return "P1" })
	err := c1.RegisterHook(BEFORE_PLAYER1_TURN, func() {
		fmt.Println("called before player 1 turn")
	})
	if err != nil {
		fmt.Println("failed to add hook:", err)
	}
	c2 := NewClient(func() interface{} { return "P2" })
	c2.RegisterHook(BATTLE_END, func() {
		fmt.Println("called after battle ended")
	})
	b.AddClient(c1)
	b.AddClient(c2)
	err = b.Start()
	if err != nil {
		fmt.Println("error starting battle:", err)
	}
	// Output:
	// called before player 1 turn
	// called after battle ended
}
