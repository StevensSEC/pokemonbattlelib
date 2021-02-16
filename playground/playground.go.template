package main

import (
	"fmt"

	lib "github.com/StevensSEC/pokemonbattlelib"
)

func check(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func main() {
	b := lib.Battle{}
	err := b.Start()
	check(err)
}
