package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type cheater struct{}

func (a cheater) Act(ctx *BattleContext) Turn {
	pkmn := ctx.Self()
	pkmn.Pokemon.Type |= TypeDragon
	opponent := ctx.Opponents()[0]
	return FightTurn{Move: 0, Target: opponent}
}

type funkyAgent func(*BattleContext) Turn

func (a funkyAgent) Act(ctx *BattleContext) Turn {
	return a(ctx)
}

var _ = Describe("Agents", func() {
	It("should never receive a ground truth reference to pokemon in the battle", func() {
		b := NewBattle()
		a1 := newRcAgent()
		_a1 := Agent(a1)
		a2 := Agent(funkyAgent(func(ctx *BattleContext) Turn {
			real := b.getPokemon(ctx.Opponents()[0].target)
			Expect(&ctx.Opponents()[0].Pokemon).ToNot(BeIdenticalTo(real))
			Expect(&ctx.Opponents()[0].Pokemon.Moves[0]).ToNot(BeIdenticalTo(real.Moves[0])) // See #294
			for _, t := range ctx.Targets() {
				real := b.getPokemon(t.target) // because range makes a copy of the array
				Expect(t.Pokemon).ToNot(BeIdenticalTo(real))
			}
			return FightTurn{Move: 0, Target: ctx.Opponents()[0]}
		}))
		party1 := NewOccupiedParty(PkmnDefault())
		party2 := NewOccupiedParty(PkmnDefault())
		b.AddParty(party1, &_a1, 0)
		b.AddParty(party2, &a2, 1)
		Expect(b.Start()).To(Succeed())
		a1 <- FightTurn{
			Move: 0,
			Target: AgentTarget{
				target: target{
					party: 1,
					slot:  0,
				},
			},
		}
		b.SimulateRound()
	})
})
