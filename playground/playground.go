package main

import (
	"fmt"

	lib "github.com/StevensSEC/pokemonbattlelib"
)

type dumbAgent struct{}

// Blindly uses the first move on the first opponent pokemon.
func (a dumbAgent) Act(ctx *lib.BattleContext) lib.Turn {
	// You can use a (reference to self) for self-targeting turns
	for _, target := range ctx.Targets {
		return lib.FightTurn{
			Move:   0,
			Target: target,
		}
	}
	panic("no opponents found")
}

type badBoy struct{}

func (badBoy) Act(ctx *lib.BattleContext) lib.Turn {
	for _, target := range ctx.Targets {
		p := target.Pokemon
		p.CurrentHP = 765
		p.Stats[lib.STAT_ATK] = uint(123)
		fmt.Printf("setting %s\n", p.GetName())
		return lib.FightTurn{
			Move:   0,
			Target: target,
		}
	}
	panic("no opponents found")
}

func check(err error) {
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func main() {
	a1 := lib.Agent(badBoy{})
	a2 := lib.Agent(dumbAgent{})
	b := lib.NewBattle()
	party1 := lib.NewParty(&a1, 0)
	pkmn1 := lib.NewPokemon(4)
	party1.AddPokemon(&pkmn1)
	party2 := lib.NewParty(&a2, 1)
	pkmn2 := lib.NewPokemon(7)
	party2.AddPokemon(&pkmn2)
	b.AddParty(party1, party2)
	err := b.Start()
	check(err)

	b.SimulateRound()

	tpkmn := b.GetTargets()[0].Pokemon
	fmt.Printf("results: %s\n", tpkmn.GetName())
	fmt.Printf("current hp: %d\n", tpkmn.CurrentHP)
	fmt.Printf("result: %d\n", tpkmn.Stats[lib.STAT_ATK])
}
