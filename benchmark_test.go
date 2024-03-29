package pokemonbattlelib

import (
	"fmt"
	"math/rand"
	"testing"
)

type smarterAgent struct{}

func (smarterAgent) Act(ctx *BattleContext) Turn {
	best := 0
	for i, target := range ctx.Opponents() {
		if target.Pokemon.Defense() < ctx.Opponents()[best].Pokemon.Defense() {
			best = i
		}
	}
	self := ctx.Self().Pokemon
	opponent := ctx.Opponents()[best].Pokemon
	var strongestMove int
	var strongestDamage uint
	for i, move := range self.Moves {
		if move == nil {
			continue
		}
		damage := CalcMoveDamage(ctx.Battle.Weather, &self, &opponent, move)
		if damage > strongestDamage {
			strongestDamage = damage
			strongestMove = i
		}
	}
	return FightTurn{
		Move:   strongestMove,
		Target: ctx.Opponents()[best],
	}
}

func randParty() *Party {
	party := NewParty()
	count := rand.Intn(5) + 1
	for j := 0; j < count; j++ {
		p := GeneratePokemon(rand.Intn(493)+1,
			WithLevel(uint8(47+rand.Intn(6))),
			WithMoves(
				MoveId(rand.Intn(len(AllMoves))+1),
				MoveId(rand.Intn(len(AllMoves))+1),
				MoveId(rand.Intn(len(AllMoves))+1),
				MoveId(rand.Intn(len(AllMoves))+1),
			),
		)
		party.AddPokemon(p)
	}
	return party
}

func dumpBattle(b *Battle) {
	fmt.Printf("Battle: %v\n", b)
	for i, p := range b.parties {
		fmt.Printf("party%d := NewOccupiedParty(\n", i)
		for _, pkmn := range p.pokemon() {
			fmt.Printf("\tGeneratePokemon(%d, WithLevel(%d), ", pkmn.NatDex, pkmn.Level)
			fmt.Printf("WithMoves(")
			for _, m := range pkmn.Moves {
				fmt.Printf("%d, ", m.Id)
			}
			fmt.Printf("),")
			fmt.Printf("WithIVs(%#v),", pkmn.IVs)
			fmt.Printf("),\n")
		}
		fmt.Printf(")\n")
	}
}

func BenchmarkBattle(b *testing.B) {
	var battle *Battle
	// bl := log.New(os.Stdout, "[battle] ", log.Lshortfile)
	// SetLogger(bl)
	log := GetLogger()
	rand.Seed(8778723)

	a1 := Agent(smarterAgent{})
	stuck := 0
	for i := 0; i < b.N; i++ {
		p1 := randParty()
		p2 := randParty()
		battle = NewSingleBattle(p1, &a1, p2, &a1)
		rng := battle.rng.(*LCRNG)
		seed := uint32(*rng)
		err := battle.Start()
		if err != nil {
			panic(err)
		}
		rounds := 0
		for {
			transactions, ended := battle.SimulateRound()
			for _, t := range transactions {
				switch tt := t.(type) {
				case UseMoveTransaction:
					log.Printf("%s used %s on %s", tt.User, tt.Move, tt.Target)
				case DamageTransaction:
					log.Printf(" %s took %d damage", tt.Target, tt.Damage)
				case HealTransaction:
					log.Printf("%s healed for %d HP", tt.Target, tt.Amount)
				case FaintTransaction:
					log.Printf("%s fainted", tt.Target)
				default:
					log.Printf("Transaction: %T - %v", t, t)
				}
			}
			// HACK: sometimes we generate battles with unimplemented status moves. this avoids getting stuck
			if len(transactions) == 0 {
				ended = true
			}
			if ended {
				break
			}
			rounds++

			if rounds > 100 {
				fmt.Printf("\n\nBATTLE GOT STUCK:\n")
				dumpBattle(battle)
				fmt.Printf("b.SetSeed(%d)\n", seed)
				// log.Panic("Battle is stuck!")
				stuck++
				break
			}
		}
	}
	log.Printf("%d battles got stuck", stuck)
}
