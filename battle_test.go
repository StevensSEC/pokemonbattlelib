package pokemonbattlelib

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type dumbAgent struct{}

// Blindly uses the first move on the first opponent pokemon.
func (a dumbAgent) Act(ctx *BattleContext) Turn {
	// You can use `a` (reference to self) for self-targeting turns
	for _, target := range ctx.Opponents {
		return FightTurn{
			Move:   0,
			Target: target,
		}
	}
	panic("no opponents found")
}

type healAgent struct{}

// Always uses a potion on first Pokemon
func (a healAgent) Act(ctx *BattleContext) Turn {
	item := GetItem(ITEM_POTION)
	for _, target := range ctx.Allies {
		return ItemTurn{
			Item:   &item,
			Target: target,
		}
	}
	panic("no allies found")
}

// An agent that can be given turns to take via a channel.
type rcAgent chan Turn

func (a rcAgent) Act(ctx *BattleContext) Turn {
	return <-a
}

func newRcAgent() rcAgent {
	return rcAgent(make(chan Turn, 1))
}

// Example usage for rcAgent
func TestRcAgent(t *testing.T) {
	a1 := newRcAgent()
	a2 := newRcAgent()
	_a1 := Agent(a1)
	_a2 := Agent(a2)
	pound := GetMove(1)
	pkmn1 := GeneratePokemon(4)
	pkmn1.Moves[0] = &pound
	pkmn2 := GeneratePokemon(7)
	pkmn2.Moves[0] = &pound
	party1 := NewOccupiedParty(&_a1, 0, pkmn1)
	party2 := NewOccupiedParty(&_a2, 1, pkmn2)
	b := NewBattle()
	b.AddParty(party1, party2)

	b.Start()
	a1 <- FightTurn{
		Move: 0,
		Target: target{
			party:     1,
			partySlot: 0,
		},
	}
	a2 <- FightTurn{
		Move: 0,
		Target: target{
			party:     0,
			partySlot: 0,
		},
	}
	transactions, _ := b.SimulateRound()
	for _, t := range transactions {
		fmt.Printf("%s\n", t.BattleLog())
	}
}

var _ = Describe("Battle", func() {
	var (
		agent1 Agent
		agent2 Agent
	)

	BeforeEach(func() {
		agent1 = Agent(dumbAgent{})
		agent2 = Agent(dumbAgent{})
	})

	Context("Battle setup", func() {
		It("runs without panicking", func() {
			party1 := NewOccupiedParty(&agent1, 0, GeneratePokemon(4))
			party2 := NewOccupiedParty(&agent2, 1, GeneratePokemon(7))
			b := NewBattle()
			b.AddParty(party1, party2)
			b.SetSeed(849823)
		})
	})
})

var _ = Describe("One round of battle", func() {
	var (
		agent1     Agent
		agent2     Agent
		party1     *party
		party2     *party
		battle     *Battle
		charmander *Pokemon
		squirtle   *Pokemon
		pound      Move
	)

	BeforeEach(func() {
		agent1 = Agent(dumbAgent{})
		agent2 = Agent(dumbAgent{})
		pound = GetMove(1)
		charmander = GeneratePokemon(4, WithMoves(&pound))
		party1 = NewOccupiedParty(&agent1, 0, charmander)
		squirtle = GeneratePokemon(7, WithMoves(&pound))
		party2 = NewOccupiedParty(&agent2, 1, squirtle)
		battle = NewBattle()
		battle.AddParty(party1, party2)
	})

	It("starts without error", func() {
		err := battle.Start()
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("when simulating a round between two agents", func() {
		It("should return two transactions", func() {
			battle.Start()
			transactions, _ := battle.SimulateRound()
			Expect(transactions).To(HaveLen(2))
		})
		It("should log FightTurns correctly", func() {
			battle.Start()
			transactions, _ := battle.SimulateRound()
			log0 := transactions[0].BattleLog()
			Expect(log0).To(Equal("Charmander used Pound on Squirtle for 3 damage."))
			log1 := transactions[1].BattleLog()
			Expect(log1).To(Equal("Squirtle used Pound on Charmander for 3 damage."))
		})
		It("should cause Pokemon to have reduced HP", func() {
			battle.Start()
			battle.SimulateRound()
			Expect(charmander.CurrentHP < charmander.Stats[STAT_HP]).To(BeTrue())
			Expect(squirtle.CurrentHP < squirtle.Stats[STAT_HP]).To(BeTrue())
		})
	})
})

var _ = Describe("Using items in battle", func() {
	var (
		agent  Agent
		pkmn   *Pokemon
		party  *party
		battle *Battle
	)

	BeforeEach(func() {
		agent = Agent(healAgent{})
		pkmn = GeneratePokemon(3, WithLevel(50))
		pkmn.CurrentHP = 10
		party = NewOccupiedParty(&agent, 0, pkmn)
		battle = NewBattle()
		battle.AddParty(party)
	})

	Context("an Agent uses an ItemTurn to use a Potion on a Pokemon", func() {
		It("should run without error", func() {
			err := battle.Start()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should log ItemTurns correctly", func() {
			battle.Start()
			transactions, _ := battle.SimulateRound()
			log0 := transactions[0].BattleLog()
			Expect(log0).To(Equal("Potion used on Venusaur."))
		})
		It("should heal the Pokemon by 20 HP", func() {
			battle.Start()
			battle.SimulateRound()
			Expect(int(pkmn.CurrentHP)).To(Equal(30))
		})
	})
})

var _ = Describe("Active pokemon in battle", func() {
	var (
		agent Agent
		party *party
	)

	BeforeEach(func() {
		agent = Agent(dumbAgent{})
		party = NewOccupiedParty(&agent, 0, GeneratePokemon(7), GeneratePokemon(9))
	})

	It("should add an active Pokemon when SetActive is called", func() {
		party.SetActive(0)
		Expect(party.GetActivePokemon()).To(HaveLen(1))
	})

	It("should remove an active Pokemon when SetInactive is called", func() {
		party.SetActive(0)
		party.SetInactive(0)
		Expect(party.GetActivePokemon()).To(HaveLen(0))
	})

	It("should add the active Pokemon the user expects", func() {
		party.SetActive(1)
		pkmn := party.GetActivePokemon()[1]
		Expect(int(pkmn.NatDex)).To(Equal(9))
	})
})

var _ = Describe("Getting party Pokemon", func() {
	var (
		agent1 Agent
		agent2 Agent
		party1 *party
		party2 *party
		battle *Battle
	)

	BeforeEach(func() {
		agent1 = Agent(dumbAgent{})
		agent2 = Agent(dumbAgent{})
		party1 = NewOccupiedParty(&agent1, 0, GeneratePokemon(4), GeneratePokemon(7), GeneratePokemon(11))
		party2 = NewOccupiedParty(&agent2, 1, GeneratePokemon(15))
		battle = NewBattle()
		battle.AddParty(party1, party2)
	})

	Context("Calling GetPokemon()", func() {
		It("should get the Pokemon the user expects", func() {
			pkmn := battle.getPokemon(0, 1)
			Expect(int(pkmn.NatDex)).To(Equal(7))
		})
	})

	Context("Getting ally Pokemon", func() {
		It("should start the battle without error", func() {
			err := battle.Start()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should return targets whose team matches the passed party ", func() {
			battle.Start()
			for _, party := range []*party{party1, party2} {
				allies := battle.GetAllies(party)
				Expect(allies).To(HaveLen(1))
			}
		})
	})

	Context("Getting opposing Pokemon", func() {
		It("should start the battle without error", func() {
			err := battle.Start()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should return targets whose team does not match the passed party ", func() {
			battle.Start()
			for _, party := range []*party{party1, party2} {
				opponents := battle.GetOpponents(party)
				Expect(opponents).To(HaveLen(1))
			}
		})
	})
})

var _ = Describe("Move priority", func() {
	var (
		a1 Agent
		a2 Agent
	)

	BeforeEach(func() {
		a1 = Agent(dumbAgent{})
		a2 = Agent(dumbAgent{})
	})

	Specify("Moves with higher priority should go first", func() {
		pound := GetMove(1)
		p1 := GeneratePokemon(1, WithLevel(5), WithMoves(&pound))
		p1.Stats[STAT_SPD] = 100
		party1 := NewOccupiedParty(&a1, 0, p1)
		fakeout := GetMove(252)
		p2 := GeneratePokemon(4, WithLevel(5), WithMoves(&fakeout))
		p2.Stats[STAT_SPD] = 10
		party2 := NewOccupiedParty(&a2, 1, p2)
		b := NewBattle()
		b.AddParty(party1, party2)
		b.Start()
		transactions, _ := b.SimulateRound()
		Expect(len(transactions)).To(Equal(2))

		logtest := []struct {
			turn Transaction
			want string
		}{
			{
				turn: transactions[0],
				want: "Charmander used Fake Out on Bulbasaur for 5 damage.",
			},
			{
				turn: transactions[1],
				want: "Bulbasaur used Pound on Charmander for 5 damage.",
			},
		}
		for _, tt := range logtest {
			Expect(tt.turn.BattleLog()).To(Equal(tt.want))
		}
	})
})

var _ = Describe("Pokemon speed", func() {
	var (
		agent1 Agent
		agent2 Agent
		party1 *party
		party2 *party
		pound  Move
		battle *Battle
	)

	BeforeEach(func() {
		agent1 = Agent(dumbAgent{})
		agent2 = Agent(dumbAgent{})
		pound = GetMove(1)
		party1 = NewOccupiedParty(&agent1, 0, GeneratePokemon(4, WithMoves(&pound)))
		party2 = NewOccupiedParty(&agent2, 1, GeneratePokemon(291, WithMoves(&pound))) // ninjask is faster than charmander
		battle = NewBattle()
		battle.AddParty(party1, party2)
	})

	Context("A faster Pokemon is fighting a slower one", func() {
		It("should not error when starting", func() {
			err := battle.Start()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should create two transactions when simulating a round", func() {
			battle.Start()
			transactions, _ := battle.SimulateRound()
			Expect(transactions).To(HaveLen(2))
		})

		Specify("faster Pokemon should go first", func() {
			battle.Start()
			transactions, _ := battle.SimulateRound()
			log0 := transactions[0].BattleLog()
			Expect(log0).To(Equal("Ninjask used Pound on Charmander for 3 damage."))
			log1 := transactions[1].BattleLog()
			Expect(log1).To(Equal("Charmander used Pound on Ninjask for 3 damage."))
		})
	})
})

var _ = Describe("Turn priority", func() {
	Context("Fight turns", func() {
		It("should have a priority of 0", func() {
			turn := FightTurn{}
			Expect(turn.Priority()).To(Equal(0))
		})
	})

	Context("Item turns", func() {
		It("should have a priority of 1", func() {
			turn := ItemTurn{}
			Expect(turn.Priority()).To(Equal(1))
		})
	})
})

var _ = Describe("Fainting", func() {
	var (
		agent1 Agent
		agent2 Agent
		party1 *party
		party2 *party
		battle *Battle
		pound  Move
	)

	BeforeEach(func() {
		agent1 = Agent(dumbAgent{})
		agent2 = Agent(dumbAgent{})
		pound = GetMove(1)
		scary_monster := GeneratePokemon(7, WithLevel(100), WithMoves(&pound))
		scary_monster.Stats[STAT_SPD] = 1
		party1 = NewOccupiedParty(&agent1, 0, GeneratePokemon(4, WithMoves(&pound)), GeneratePokemon(387, WithMoves(&pound)))
		party2 = NewOccupiedParty(&agent2, 1, scary_monster)
		battle = NewBattle()
		battle.AddParty(party1, party2)
	})

	Context("Switch is forced after a Pokemon faints", func() {
		It("should start without error", func() {
			err := battle.Start()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("causes 4 transactions to occur", func() {
			battle.Start()
			transactions, _ := battle.SimulateRound()
			Expect(transactions).To(HaveLen(4))
		})
		It("should log all transactions as expected", func() {
			battle.Start()
			transactions, _ := battle.SimulateRound()
			log0 := transactions[0].BattleLog()
			Expect(log0).To(Equal("Charmander used Pound on Squirtle for 2 damage."))
			// Charmander smashed his nubby little fist into Squirtle as
			// hard as he could. Spectators gasped and winced when the
			// impact created a very audible crack. But it was not
			// Squirtle's shell that broke, it was Charmanders knuckles.
			// The Squirtle was unfazed.
			log1 := transactions[1].BattleLog()
			Expect(log1).To(Equal("Squirtle used Pound on Charmander for 674 damage."))
			// Ash watched in horror as his Charmander was obliterated from the
			// battlefield. "Critical hit!" echoed the automated announcer. The
			// Squirtle snarled, now covered in the entrails of his previous
			// opponent. "OH GOD, WHAT THE FUCK!?" sobbed Ash, "Is my friend
			// really gone forever? Please tell me I'm dreaming, this can't be real!"
			log2 := transactions[2].BattleLog()
			Expect(log2).To(Equal("Charmander fainted."))
			log3 := transactions[3].BattleLog()
			Expect(log3).To(Equal("Turtwig was sent out."))
		})
	})

	Specify("Dead pokemon should not take turns", func() {
		a1 := Agent(dumbAgent{})
		a2 := Agent(dumbAgent{})
		party1 := NewParty(&a1, 0)
		pkmn1 := GeneratePokemon(4, WithLevel(3))
		pkmn1.CurrentHP = 1
		pound := GetMove(1)
		pkmn1.Moves[0] = &pound
		party1.AddPokemon(pkmn1)
		party2 := NewParty(&a2, 1)
		pkmn2 := GeneratePokemon(7, WithLevel(10))
		pkmn2.Stats[STAT_SPD] = 255
		pkmn2.Moves[0] = &pound
		party2.AddPokemon(pkmn2)
		b := NewBattle()
		b.AddParty(party1, party2)
		err := b.Start()
		Expect(err).To(BeNil())

		transactions, ended := b.SimulateRound()
		Expect(ended).To(BeTrue(), "Expected SimulateRound to indicate that the battle has ended, but it did not.")
		Expect(len(transactions)).To(Equal(3), "Expected 3 transactions to occur")
		logtest := []struct {
			turn Transaction
			want string
		}{
			{
				turn: transactions[0],
				want: "Squirtle used Pound on Charmander for 11 damage.",
			},
			{
				turn: transactions[1],
				want: "Charmander fainted.",
			},
			{
				turn: transactions[2],
				want: "The battle has ended.",
			},
		}
		for _, tt := range logtest {
			Expect(tt.turn.BattleLog()).To(Equal(tt.want))
		}
	})

	Specify("Dead pokemon should not take turns", func() {
		a1 := Agent(dumbAgent{})
		a2 := Agent(dumbAgent{})
		party1 := NewParty(&a1, 0)
		pkmn1 := GeneratePokemon(4, WithLevel(3))
		pkmn3 := GeneratePokemon(387, WithLevel(3))
		pkmn1.CurrentHP = 1
		pound := GetMove(1)
		pkmn1.Moves[0] = &pound
		pkmn3.Moves[0] = &pound
		party1.AddPokemon(pkmn1, pkmn3)
		party2 := NewParty(&a2, 1)
		pkmn2 := GeneratePokemon(7, WithLevel(10))
		pkmn2.Stats[STAT_SPD] = 255
		pkmn2.Moves[0] = &pound
		party2.AddPokemon(pkmn2)
		b := NewBattle()
		b.AddParty(party1, party2)
		err := b.Start()
		Expect(err).To(BeNil())

		transactions, ended := b.SimulateRound()
		Expect(ended).To(BeFalse(), "Expected SimulateRound to NOT indicate that the battle has ended, but it did.")
		Expect(len(transactions)).To(Equal(3), "Expected 3 transactions to occur")
		logtest := []struct {
			turn Transaction
			want string
		}{
			{
				turn: transactions[0],
				want: "Squirtle used Pound on Charmander for 11 damage.",
			},
			{
				turn: transactions[1],
				want: "Charmander fainted.",
			},
			{
				turn: transactions[2],
				want: "Turtwig was sent out.",
			},
		}
		for _, tt := range logtest {
			Expect(tt.turn.BattleLog()).To(Equal(tt.want))
		}
	})
})

var _ = Describe("Ending a battle", func() {
	var (
		agent1 Agent
		agent2 Agent
		party1 *party
		party2 *party
		battle *Battle
		pound  Move
	)

	BeforeEach(func() {
		agent1 = Agent(dumbAgent{})
		agent2 = Agent(dumbAgent{})
		pound = GetMove(1)
		low_health_pkmn := GeneratePokemon(4, WithMoves(&pound))
		low_health_pkmn.CurrentHP = 1
		party1 = NewOccupiedParty(&agent1, 0, low_health_pkmn)
		party2 = NewOccupiedParty(&agent2, 1, GeneratePokemon(7, WithMoves(&pound)))
		battle = NewBattle()
		battle.AddParty(party1, party2)
	})

	Context("Battle ends by knockout", func() {
		It("should start without error", func() {
			err := battle.Start()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should end", func() {
			battle.Start()
			_, ended := battle.SimulateRound()
			Expect(ended).To(BeTrue())
		})

		It("should have 4 transactions occur", func() {
			battle.Start()
			transactions, _ := battle.SimulateRound()
			Expect(transactions).To(HaveLen(4))
		})

		It("should log all transaction correctly", func() {
			battle.Start()
			transactions, _ := battle.SimulateRound()
			log0 := transactions[0].BattleLog()
			Expect(log0).To(Equal("Charmander used Pound on Squirtle for 3 damage."))
			log1 := transactions[1].BattleLog()
			Expect(log1).To(Equal("Squirtle used Pound on Charmander for 3 damage."))
			log2 := transactions[2].BattleLog()
			Expect(log2).To(Equal("Charmander fainted."))
			log3 := transactions[3].BattleLog()
			Expect(log3).To(Equal("The battle has ended."))
		})
	})
})

var _ = Describe("Status Conditions", func() {
	var (
		a1 Agent
		a2 Agent
	)

	BeforeEach(func() {
		a1 = Agent(dumbAgent{})
		a2 = Agent(dumbAgent{})
	})

	It("should inflict burn and poison damage", func() {
		pound := GetMove(1)
		p1 := GeneratePokemon(1)
		p1.Moves[0] = &pound
		p1.StatusEffects = StatusPoison
		party1 := NewOccupiedParty(&a1, 0, p1)
		p2 := GeneratePokemon(2)
		p2.Moves[0] = &pound
		p2.StatusEffects = StatusBurn
		party2 := NewOccupiedParty(&a2, 1, p2)
		b := NewBattle()
		b.AddParty(party1, party2)
		b.Start()
		transactions, _ := b.SimulateRound()
		Expect(len(transactions)).To(Equal(4), "Expected only 4 transactions to occur in a round")

		logtest := []struct {
			turn Transaction
			want string
		}{
			{
				turn: transactions[2],
				want: "Bulbasaur took 1 damage from being poisoned.",
			},
			{
				turn: transactions[3],
				want: "Ivysaur took 1 damage from being burned.",
			},
		}
		for _, tt := range logtest {
			Expect(tt.turn.BattleLog()).To(Equal(tt.want))
		}
	})

	It("should inflict badly poisoned damage", func() {
		p1 := GeneratePokemon(1, WithLevel(8))
		pound := GetMove(1)
		p1.Moves[0] = &pound
		p1.StatusEffects = StatusBadlyPoison
		party1 := NewOccupiedParty(&a1, 0, p1)
		p2 := GeneratePokemon(2)
		p2.Moves[0] = &pound
		party2 := NewOccupiedParty(&a2, 1, p2)
		b := NewBattle()
		b.AddParty(party1, party2)
		b.Start()
		transactions, _ := b.SimulateRound()
		Expect(len(transactions)).To(Equal(3), "Expected only 3 transactions to occur in a round")

		logtest := []struct {
			turn Transaction
			want string
		}{
			{
				turn: transactions[2],
				want: "Bulbasaur took 1 damage from being badly poisoned.",
			},
		}
		for _, tt := range logtest {
			Expect(tt.turn.BattleLog()).To(Equal(tt.want))
		}
	})
})
