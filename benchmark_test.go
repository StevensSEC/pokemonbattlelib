package pokemonbattlelib

import (
	"fmt"
	"math/rand"
	"testing"
)

type smartAgent struct{}

func (smartAgent) Act(ctx *BattleContext) Turn {
	var strongestMove int
	power := 0
	for i, move := range ctx.Pokemon.Moves {
		if move == nil {
			continue
		}
		if move.Power > power {
			power = move.Power
			strongestMove = i
		}
	}
	best := 0
	for i, target := range ctx.Opponents {
		if target.Pokemon.Stats[StatDef] < ctx.Opponents[best].Pokemon.Stats[StatDef] {
			best = i
		}
	}
	return FightTurn{
		Move:   strongestMove,
		Target: ctx.Opponents[best],
	}
}

func randParty() *party {
	a1 := Agent(smartAgent{})
	party := NewParty(&a1, 0)
	count := rand.Intn(5) + 1
	for j := 0; j < count; j++ {
		p := GeneratePokemon(rand.Intn(493), WithLevel(uint8(47+rand.Intn(6))))
		p.Moves[0] = &AllMoves[rand.Intn(len(AllMoves))]
		p.Moves[1] = &AllMoves[rand.Intn(len(AllMoves))]
		p.Moves[2] = &AllMoves[rand.Intn(len(AllMoves))]
		p.Moves[3] = &AllMoves[rand.Intn(len(AllMoves))]
		party.AddPokemon(p)
	}
	return party
}

func BenchmarkBattle(b *testing.B) {
	rand.Seed(8778723)
	for i := 0; i < b.N; i++ {
		// fmt.Printf("%d/%d", i, b.N)
		p1 := randParty()
		p2 := randParty()
		p2.team = 1
		battle := NewBattle()
		battle.AddParty(p1, p2)
		err := battle.Start()
		if err != nil {
			panic(err)
		}
		for {
			transactions, ended := battle.SimulateRound()
			fmt.Println(len(transactions))
			if ended {
				break
			}
		}
	}
}
