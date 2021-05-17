package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

type dumbAgent struct{}

// Blindly uses the first move on the first opponent pokemon.
func (a dumbAgent) Act(ctx *BattleContext) Turn {
	// You can use `a` (reference to self) for self-targeting turns
	for _, target := range ctx.Opponents() {
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
	for _, target := range ctx.Allies() {
		return ItemTurn{
			Item:   ItemPotion,
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

func (a rcAgent) newTarget(party, slot uint) AgentTarget {
	return AgentTarget{
		target: target{
			party: party,
			slot:  slot,
		},
	}
}

func newRcAgent() rcAgent {
	return rcAgent(make(chan Turn, 1))
}

// Example usage for rcAgent
var _ = Describe("RC Agent", func() {
	a1 := newRcAgent()
	a2 := newRcAgent()
	_a1 := Agent(a1)
	_a2 := Agent(a2)
	party1 := NewOccupiedParty(PkmnDefault())
	party2 := NewOccupiedParty(PkmnDefault())
	b := NewSingleBattle(party1, &_a1, party2, &_a2)
	Expect(b.Start()).To(Succeed())
	a1 <- FightTurn{
		Move:   0,
		Target: a1.newTarget(1, 0),
	}
	a2 <- FightTurn{
		Move:   0,
		Target: a1.newTarget(0, 0),
	}
})

var _ = Describe("Battle initialization", func() {
	agent1 := Agent(new(dumbAgent))
	agent2 := Agent(new(dumbAgent))

	Context("when creating a new battle", func() {
		It("runs without panicking", func() {
			party1 := NewOccupiedParty(PkmnDefault())
			party2 := NewOccupiedParty(PkmnDefault())
			b := NewSingleBattle(party1, &agent1, party2, &agent2)
			b.SetSeed(849823)
		})

		It("panics when getting an invalid Pokemon", func() {
			party := NewOccupiedParty(PkmnDefault())
			b := NewBattle()
			b.AddParty(party, &agent1, 0)
			Expect(func() {
				b.getPokemon(target{1, 5})
			}).To(Panic())
			Expect(func() {
				b.getPokemon(target{0, 5})
			}).To(Panic())
		})
	})

	Context("team validation", func() {
		It("should fail when party has no pokemon", func() {
			b := NewSingleBattle(
				NewParty(), &agent1,
				NewOccupiedParty(PkmnDefault()), &agent2,
			)
			Expect(b.Start()).NotTo(Succeed())
		})

		It("should fail when both parties are on the same team", func() {
			b := NewBattle()
			b.AddParty(NewOccupiedParty(PkmnDefault()), &agent1, 0)
			b.AddParty(NewOccupiedParty(PkmnDefault()), &agent2, 0)
			Expect(b.Start()).NotTo(Succeed())
		})
	})
})

var _ = Describe("One round of battle", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))

	var (
		b  *Battle
		p1 *Pokemon
		p2 *Pokemon
	)

	BeforeEach(func() {
		p1 = PkmnDefault()
		p2 = PkmnDefault()
		b = New1v1Battle(p1, &a1, p2, &a2)
		b.ruleset &= ^BattleRuleFaint
		b.rng = SimpleRNG()
	})

	Context("when simulating a round between two agents", func() {
		It("panics if battle is not in progress", func() {
			Expect(func() {
				b.SimulateRound()
			}).To(Panic())
		})

		It("should create 2 damage transactions", func() {
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(UseMoveTransaction{
				User:   target{0, 0},
				Target: target{1, 0},
			}))
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{1, 0},
				Damage: 2,
			}))
			Expect(t).To(HaveTransaction(UseMoveTransaction{
				User:   target{1, 0},
				Target: target{0, 0},
			}))
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{0, 0},
				Damage: 2,
			}))
		})

		It("should cause Pokemon to have reduced HP", func() {
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(p1.CurrentHP < p1.MaxHP()).To(BeTrue())
			Expect(p2.CurrentHP < p2.MaxHP()).To(BeTrue())
		})
	})

	Context("when dealing damage to a Pokemon", func() {
		It("should account for same-type attack bonus", func() {
			p1 = PkmnWithType(TypeFire)
			p1.Moves[0] = GetMove(registerMoveWithType(TypeFire))
			p2 := PkmnNoDamage()
			b = New1v1Battle(p1, &a1, p2, &a2)
			b.rng = NeverRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			damage := DamageDealt(t, target{0, 0})
			// Adaptability increases stab from 1.5x to 2x
			p1.Ability = AbilityAdaptability
			t, _ = b.SimulateRound()
			Expect(DamageDealt(t, target{0, 0})).To(BeNumerically(">", damage))
		})

		Context("Type Matchups", func() {
			var (
				a1  rcAgent
				a2  Agent
				_a1 Agent
				b   *Battle
			)

			BeforeEach(func() {
				a1 = newRcAgent()
				a2 = Agent(dumbAgent{})
				_a1 = Agent(a1)
				b = NewBattle()
				b.ruleset &= ^BattleRuleFaint
				b.rng = SimpleRNG()
			})

			It("should account for supereffective type matchups", func() {
				pkmn1 := PkmnWithMoves(MoveFireFang, MoveTackle)
				pkmn2 := PkmnWithType(TypeGrass)
				b = New1v1Battle(pkmn1, &_a1, pkmn2, &a2)
				Expect(b.Start()).To(Succeed())
				a1 <- FightTurn{Move: 1, Target: a1.newTarget(1, 0)}
				t, _ := b.SimulateRound()
				damage := DamageDealt(t, target{0, 0})
				a1 <- FightTurn{Move: 0, Target: a1.newTarget(1, 0)}
				t, _ = b.SimulateRound()
				Expect(DamageDealt(t, target{0, 0})).To(BeNumerically(">", damage))
			})

			It("should have no effect", func() {
				pkmn1 := PkmnWithMoves(MoveShadowBall)
				pkmn1.Type = TypeGhost
				pkmn2 := PkmnWithMoves(MoveTackle)
				pkmn2.Type = TypeNormal
				b = New1v1Battle(pkmn1, &_a1, pkmn2, &a2)
				Expect(b.Start()).To(Succeed())
				a1 <- FightTurn{Move: 0, Target: a1.newTarget(1, 0)}
				t, _ := b.SimulateRound()
				Expect(DamageDealt(t, target{0, 0})).To(Equal(0))
			})
		})

		It("should account for critical hits", func() {
			pkmn := PkmnDefault()
			b = New1v1Battle(pkmn, &a1, PkmnNoDamage(), &a2)
			b.rng = NeverRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			damage := DamageDealt(t, target{0, 0})
			b.rng = AlwaysRNG()
			t, _ = b.SimulateRound()
			Expect(DamageDealt(t, target{0, 0})).To(BeNumerically(">", damage))
		})

		It("should miss moves randomly based on accuracy/evasion", func() {
			b.rng = NeverRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(MoveFailTransaction{
				User:   target{0, 0},
				Reason: FailMiss,
			}))
			b.rng = SimpleRNG()
			t, _ = b.SimulateRound()
			Expect(t).ToNot(HaveTransaction(MoveFailTransaction{
				User:   target{0, 0},
				Reason: FailMiss,
			}))
		})
	})
})

var _ = Describe("Switching Pokemon", func() {
	a1 := newRcAgent()
	a2 := newRcAgent()

	setup := func() *Battle {
		pkmn1 := PkmnDefault()
		pkmn2 := PkmnDefault()
		pkmn3 := PkmnDefault()
		_a1 := Agent(a1)
		_a2 := Agent(a2)
		b := NewSingleBattle(NewOccupiedParty(pkmn1), &_a1, NewOccupiedParty(pkmn2, pkmn3), &_a2)
		Expect(b.Start()).To(Succeed())
		return b
	}

	Context("when switching Pokemon in battle", func() {
		It("should allow switching before other turns", func() {
			b := setup()
			a1 <- FightTurn{Move: 0, Target: a1.newTarget(1, 0)}
			a2 <- SwitchTurn{Target: a2.newTarget(1, 1)}
			pkmn2 := b.getPokemon(target{1, 0})
			pkmn3 := b.getPokemon(target{1, 1})
			t, _ := b.SimulateRound()
			// Pokemon was switched out with other Pokemon
			Expect(b.getPokemon(target{1, 0})).ToNot(Equal(pkmn2))
			Expect(b.getPokemon(target{1, 0})).To(Equal(pkmn3))
			Expect(b.getPokemon(target{1, 1})).To(Equal(pkmn3))
			Expect(t).To(HaveTransaction(
				UseMoveTransaction{
					User:   target{0, 0},
					Target: target{1, 0},
				},
			))
		})

		It("should not allow switching to invalid Pokemon", func() {
			b := setup()
			a1 <- FightTurn{Move: 0, Target: a1.newTarget(1, 0)}
			a2 <- SwitchTurn{Target: a2.newTarget(1, 1)}
			b.getPokemon(target{1, 1}).CurrentHP = 0
			Expect(func() { b.SimulateRound() }).To(Panic())
		})
	})
})

var _ = Describe("Using items in battle", func() {
	a := Agent(new(healAgent))
	var (
		pkmn1 *Pokemon
		pkmn2 *Pokemon
		b     *Battle
	)

	BeforeEach(func() {
		pkmn1 = PkmnDefault()
		pkmn1.Stats[StatHP] = 100
		pkmn1.CurrentHP = 10
		pkmn2 = PkmnNoDamage()
		b = New1v1Battle(pkmn1, &a, pkmn2, &a)
	})

	Context("when the battle processes item turns", func() {
		It("should create ItemTransaction(s) properly", func() {
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				ItemTransaction{
					Target: target{0, 0},
					Item:   ItemPotion,
					Move:   nil,
				},
			))
		})

		It("should heal the Pokemon by 20 HP", func() {
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(pkmn1.CurrentHP).To(BeEquivalentTo(30))
		})
	})
})

var _ = Describe("Getting pokemon from parties", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))
	var (
		b *Battle
	)

	BeforeEach(func() {
		p1 := NewOccupiedParty(
			GeneratePokemon(PkmnCharmander, WithMoves(TestMoveDefault)),
			GeneratePokemon(PkmnSquirtle, WithMoves(TestMoveDefault)),
			GeneratePokemon(PkmnMetapod, WithMoves(TestMoveDefault)),
		)
		p2 := NewOccupiedParty(GeneratePokemon(PkmnBeedrill, WithMoves(TestMoveDefault)))
		b = NewSingleBattle(p1, &a1, p2, &a2)
	})

	Context("when getting Pokemon by party/slot", func() {
		It("should get the Pokemon the user expects", func() {
			pkmn := b.getPokemon(target{0, 1})
			Expect(pkmn.NatDex).To(BeEquivalentTo(PkmnSquirtle))
		})
	})

	Context("when getting ally Pokemon", func() {
		It("should return targets whose team matches the passed party", func() {
			Expect(b.Start()).To(Succeed())
			for _, party := range b.parties {
				allies := b.getAllies(party)
				Expect(allies).To(HaveLen(1))
			}
		})
	})

	Context("when getting opponent Pokemon", func() {
		It("should return targets whose team does not match the passed party ", func() {
			Expect(b.Start()).To(Succeed())
			for _, party := range b.parties {
				opponents := b.getOpponents(party)
				Expect(opponents).To(HaveLen(1))
			}
		})
	})
})

var _ = Describe("Turn priority", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))

	Context("each turn type should have a priority level", func() {
		DescribeTable("Turn priority",
			func(turn Turn, want int) {
				Expect(turn.Priority()).To(Equal(want))
			},
			Entry("FightTurn", FightTurn{}, 0),
			Entry("ItemTurn", ItemTurn{}, 1),
			Entry("SwitchTurn", SwitchTurn{}, 2),
		)

		It("should order turns properly based on priority", func() {
			a2 := Agent(new(healAgent))
			p1 := PkmnDefault()
			p2 := PkmnDefault()
			b := New1v1Battle(p1, &a1, p2, &a2)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransactionsInOrder(
				HealTransaction{
					Target: target{1, 0},
					Amount: 0,
				},
				UseMoveTransaction{
					User:   target{0, 0},
					Target: target{1, 0},
					Move:   GetMove(TestMoveDefault),
				},
				DamageTransaction{
					Target: target{1, 0},
					Move:   GetMove(TestMoveDefault),
					Damage: 2,
				},
			))
		})
	})

	Context("when determining priority for equal turn types", func() {
		It("should handle moves with higher priority first", func() {
			p1 := PkmnDefault()
			p1.Stats[StatSpeed] = 100
			p2 := PkmnWithMoves(MoveQuickAttack)
			p2.Stats[StatSpeed] = 10
			b := New1v1Battle(p1, &a1, p2, &a2)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransactionsInOrder(
				UseMoveTransaction{
					User:   target{1, 0},
					Target: target{0, 0},
					Move:   GetMove(MoveQuickAttack),
				},
				DamageTransaction{
					Target: target{0, 0},
					Damage: 3,
					Move:   GetMove(MoveQuickAttack),
				},
				UseMoveTransaction{
					User:   target{0, 0},
					Target: target{1, 0},
					Move:   GetMove(TestMoveDefault),
				},
				DamageTransaction{
					Target: target{1, 0},
					Damage: 2,
					Move:   GetMove(TestMoveDefault),
				},
			))
		})

		It("should handle faster Pokemon first", func() {
			p1 := GeneratePokemon(PkmnCharmander, WithMoves(TestMoveDefault))
			p2 := GeneratePokemon(PkmnNinjask, WithMoves(TestMoveDefault)) // ninjask is faster than charmander
			b := New1v1Battle(p1, &a1, p2, &a2)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransactionsInOrder(
				UseMoveTransaction{
					User:   target{1, 0},
					Target: target{0, 0},
					Move:   GetMove(TestMoveDefault),
				},
				DamageTransaction{
					Target: target{0, 0},
					Move:   GetMove(TestMoveDefault),
					Damage: 2,
				},
				UseMoveTransaction{
					User:   target{0, 0},
					Target: target{1, 0},
					Move:   GetMove(TestMoveDefault),
				},
				DamageTransaction{
					Target: target{1, 0},
					Move:   GetMove(TestMoveDefault),
					Damage: 2,
				},
			))
		})
	})
})

var _ = Describe("Fainting", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))
	var (
		p1 *Party
		p2 *Party
		b  *Battle
	)

	Context("after a Pokemon faints in battle", func() {
		It("should switch to the next available Pokemon", func() {
			pkmn1 := PkmnDefault()
			pkmn2 := PkmnDefault()
			pkmn2.CurrentHP = 1
			pkmn3 := PkmnDefault()
			p1 = NewOccupiedParty(
				pkmn2,
				pkmn3,
			)
			p2 = NewOccupiedParty(pkmn1) // and nice sprites
			b = NewSingleBattle(p1, &a1, p2, &a2)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			// Charmander smashed his nubby little fist into Squirtle as
			// hard as he could. Spectators gasped and winced when the
			// impact created a very audible crack. But it was not
			// Squirtle's shell that broke, it was Charmanders knuckles.
			// The Squirtle was unfazed.
			// Ash watched in horror as his Charmander was obliterated from the
			// battlefield. "Critical hit!" echoed the automated announcer. The
			// Squirtle snarled, now covered in the entrails of his previous
			// opponent. "OH GOD, WHAT THE FUCK!?" sobbed Ash, "Is my friend
			// really gone forever? Please tell me I'm dreaming, this can't be real!"
			Expect(t).To(HaveTransactionsInOrder(
				FaintTransaction{
					Target: target{0, 0},
				}, SendOutTransaction{
					Target: target{0, 1},
				}))
		})

		It("should not allow fainted Pokemon to take turns", func() {
			pkmn1 := PkmnDefault()
			pkmn2 := PkmnDefault()
			pkmn3 := PkmnDefault()
			p1 := NewOccupiedParty(pkmn1, pkmn3)
			pkmn1.CurrentHP = 1
			p1.AddPokemon(pkmn1, pkmn3)
			p2 := NewOccupiedParty(pkmn2)
			pkmn2.Stats[StatSpeed] = 255
			p2.AddPokemon(pkmn2)
			b := NewSingleBattle(p1, &a1, p2, &a2)
			Expect(b.Start()).To(Succeed())
			t, ended := b.SimulateRound()
			Expect(ended).To(BeFalse(), "Expected SimulateRound to NOT indicate that the battle has ended, but it did.")
			Expect(t).To(HaveTransactionsInOrder(
				FaintTransaction{
					Target: target{0, 0},
				},
				SendOutTransaction{
					Target: target{0, 1},
				},
			))
			Expect(t).ToNot(HaveTransaction(
				DamageTransaction{
					Target: target{1, 0},
				},
			))
		})

		It("should lose 1 friendship when fainting", func() {
			dies := PkmnDefault()
			dies.CurrentHP = 1
			dies.Friendship = 100
			p1 := NewOccupiedParty(dies)
			p2 := NewOccupiedParty(PkmnDefault())
			b = NewBattle()
			b.AddParty(p1, &a1, 0)
			b.AddParty(p2, &a2, 1)
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(b.getPokemon(target{0, 0}).Friendship).To(Equal(99))
		})

		It("should lose 5 or 10 friendship when fainting", func() {
			dies := PkmnDefault()
			dies.CurrentHP = 1
			dies.Friendship = 100
			dies2 := PkmnDefault()
			dies2.CurrentHP = 1
			dies2.Friendship = 200
			pkmn := PkmnDefault()
			pkmn.Level = 50
			p1 := NewOccupiedParty(dies, dies2)
			p2 := NewOccupiedParty(pkmn)
			b = NewSingleBattle(p1, &a1, p2, &a2)
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(b.getPokemon(target{0, 0}).Friendship).To(Equal(95))
			b.SimulateRound()
			Expect(b.getPokemon(target{0, 1}).Friendship).To(Equal(190))
		})

		It("should gain EVs when defeating Pokemon", func() {
			winner := PkmnDefault()
			loser := GeneratePokemon(PkmnBulbasaur, WithMoves(TestMoveDefault))
			loser.CurrentHP = 1
			p1 = NewOccupiedParty(winner)
			p2 = NewOccupiedParty(loser)
			b = NewSingleBattle(p1, &a1, p2, &a2)
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				EVTransaction{
					Target: target{0, 0},
					Stat:   StatSpAtk,
					Amount: 1,
				},
			))
		})
	})

	When("holding Focus Sash", func() {
		setup := func() (*Battle, *Pokemon) {
			holder := PkmnNoDamage()
			holder.CurrentHP = 2
			holder.HeldItem = ItemFocusSash
			b = New1v1Battle(holder, &a1, PkmnDefault(), &a2)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			return b, holder
		}

		It("should not let the holder die", func() {
			b, holder := setup()
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{0, 0},
			}))
			Expect(t).ToNot(HaveTransaction(FaintTransaction{
				Target: target{0, 0},
			}))
			Expect(holder.CurrentHP).To(BeEquivalentTo(1))
		})

		It("should consume the focus sash after damage is applied", func() {
			b, holder := setup()
			t, _ := b.SimulateRound()
			target := target{0, 0}
			Expect(t).To(HaveTransactionsInOrder(
				DamageTransaction{
					Target: target,
				},
				ItemTransaction{
					Target: target,
					IsHeld: true,
					Item:   ItemFocusSash,
				},
			))
			Expect(holder.HeldItem).To(Equal(ItemNone))
			Expect(holder.HeldItem).To(Equal(ItemNone))
		})
	})
})

var _ = Describe("Battle end", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))
	var (
		b     *Battle
		pkmn1 *Pokemon
		pkmn2 *Pokemon
	)

	BeforeEach(func() {
		pkmn1 = PkmnDefault()
		pkmn1.CurrentHP = 1
		party1 := NewOccupiedParty(pkmn1)
		pkmn2 = PkmnDefault()
		pkmn2.Stats[StatSpeed] = 255
		party2 := NewOccupiedParty(pkmn2)
		b = NewSingleBattle(party1, &a1, party2, &a2)
		b.rng = SimpleRNG()
	})

	Context("when all Pokemon faint on one team", func() {
		It("should end the battle", func() {
			Expect(b.Start()).To(Succeed())
			t, ended := b.SimulateRound()
			Expect(ended).To(BeTrue(), "Expected SimulateRound to indicate that the battle has ended, but it did not.")
			Expect(t).To(HaveTransactionsInOrder(
				FaintTransaction{
					Target: target{0, 0},
				},
				EndBattleTransaction{
					Reason: EndKnockout,
					Winner: 1,
				},
			))
			Expect(t).ToNot(HaveTransaction(
				DamageTransaction{
					Target: target{1, 0},
				},
			))
			Expect(b.GetResults().Winner).To(Equal(1))
		})
	})
})

var _ = Describe("Battle metadata", func() {
	Context("Pokemon metadata", func() {
		It("should record the last used move of Pokemon in battle", func() {
			a := Agent(new(dumbAgent))
			p1 := PkmnNoDamage()
			p2 := PkmnNoDamage()
			b := New1v1Battle(p1, &a, p2, &a)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			Expect(p1.metadata).ToNot(HaveKeyWithValue(MetaLastMove, p1.Moves[0]))
			b.SimulateRound()
			Expect(p1.metadata).To(HaveKeyWithValue(MetaLastMove, p1.Moves[0]))
			Expect(p2.metadata).To(HaveKeyWithValue(MetaLastMove, p2.Moves[0]))
		})

		It("should record the last item a Pokemon consumed", func() {
			a := Agent(new(dumbAgent))
			p1 := PkmnNoDamage()
			p1.HeldItem = ItemApicotBerry
			p1.CurrentHP = 20
			p2 := PkmnNoDamage()
			b := New1v1Battle(p1, &a, p2, &a)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			Expect(p1.metadata).ToNot(HaveKeyWithValue(MetaLastItem, ItemApicotBerry))
			b.SimulateRound()
			Expect(p1.metadata).To(HaveKeyWithValue(MetaLastItem, ItemApicotBerry))
		})
	})
})

var _ = Describe("Status Conditions", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))

	Context("when a Pokemon has a nonvolatile status effect, it affects the Pokemon in battle", func() {
		It("should inflict burn and poison damage", func() {
			pkmn1 := PkmnNoDamage()
			pkmn1.StatusEffects = StatusPoison
			pkmn2 := PkmnNoDamage()
			pkmn2.StatusEffects = StatusBurn
			b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
			b.rng = AlwaysRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target:       target{0, 0},
				Damage:       12,
				StatusEffect: StatusPoison,
			}))
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target:       target{1, 0},
				Damage:       12,
				StatusEffect: StatusBurn,
			}))
		})

		It("should inflict badly poisoned damage", func() {
			pkmn1 := PkmnNoDamage()
			pkmn1.Stats[StatHP] = 100
			pkmn1.StatusEffects = StatusBadlyPoison
			pkmn2 := PkmnNoDamage()
			b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
			b.rng = AlwaysRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target:       target{0, 0},
				Damage:       6,
				StatusEffect: StatusBadlyPoison,
			}))
		})

		It("should immobilize paralyzed Pokemon", func() {
			pkmn1 := PkmnNoDamage()
			pkmn1.StatusEffects = StatusParalyze
			pkmn2 := PkmnNoDamage()
			b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
			b.rng = AlwaysRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				ImmobilizeTransaction{
					Target:       target{0, 0},
					StatusEffect: StatusParalyze,
				},
			))
		})

		It("should immobilize frozen Pokemon", func() {
			pkmn1 := PkmnNoDamage()
			pkmn1.StatusEffects = StatusFreeze
			pkmn2 := PkmnNoDamage()
			b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
			b.rng = AlwaysRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				ImmobilizeTransaction{
					Target:       target{0, 0},
					StatusEffect: StatusFreeze,
				},
			))
		})

		Context("when a pokemon is asleep", func() {
			It("should immobilize sleeping Pokemon", func() {
				pkmn1 := PkmnNoDamage()
				pkmn2 := PkmnNoDamage()
				b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
				b.rng = AlwaysRNG()
				Expect(b.Start()).To(Succeed())
				b.QueueTransaction(InflictStatusTransaction{
					Target:       target{0, 0},
					StatusEffect: StatusSleep,
				})
				b.ProcessQueue()
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(
					ImmobilizeTransaction{
						Target:       target{0, 0},
						StatusEffect: StatusSleep,
					},
				))
			})

			It("should allow sleeping Pokemon to wake up", func() {
				pkmn1 := PkmnNoDamage()
				pkmn2 := PkmnNoDamage()
				b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
				b.rng = AlwaysRNG()
				Expect(b.Start()).To(Succeed())
				b.QueueTransaction(InflictStatusTransaction{
					Target:       target{0, 0},
					StatusEffect: StatusSleep,
				})
				b.ProcessQueue()
				pkmn1.metadata[MetaSleepTime] = 0
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(
					CureStatusTransaction{
						Target:       target{0, 0},
						StatusEffect: StatusSleep,
					},
				))
				Expect(pkmn1.StatusEffects.check(StatusSleep)).To(BeFalse())
			})

			It("should decrement sleeping counter", func() {
				pkmn1 := PkmnNoDamage()
				pkmn2 := PkmnNoDamage()
				b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
				b.rng = AlwaysRNG()
				Expect(b.Start()).To(Succeed())
				b.QueueTransaction(InflictStatusTransaction{
					Target:       target{0, 0},
					StatusEffect: StatusSleep,
				})
				b.ProcessQueue()
				counter := pkmn1.metadata[MetaSleepTime].(int)
				b.SimulateRound()
				Expect(pkmn1.metadata[MetaSleepTime].(int)).To(Equal(counter - 1))
			})

			DescribeTable("Sleep walking",
				func(moveid MoveId) {
					pkmn1 := PkmnWithMoves(moveid, TestMoveNoDamage)
					pkmn2 := PkmnNoDamage()
					b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
					b.rng = AlwaysRNG()
					Expect(b.Start()).To(Succeed())
					b.QueueTransaction(InflictStatusTransaction{
						Target:       target{0, 0},
						StatusEffect: StatusSleep,
					})
					b.ProcessQueue()
					t, _ := b.SimulateRound()
					Expect(t).ToNot(HaveTransaction(
						ImmobilizeTransaction{
							Target:       target{0, 0},
							StatusEffect: StatusSleep,
						},
					))
				},
				Entry("Snore", MoveSnore),
				Entry("Sleep Talk", MoveSleepTalk),
			)
		})

		It("should cure paralysis", func() {
			pkmn1 := PkmnNoDamage()
			pkmn1.StatusEffects = StatusParalyze
			pkmn2 := PkmnNoDamage()
			b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
			b.rng = AlwaysRNG()
			Expect(b.Start()).To(Succeed())
			b.QueueTransaction(CureStatusTransaction{
				Target:       target{0, 0},
				StatusEffect: StatusParalyze,
			})
			b.ProcessQueue()
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				CureStatusTransaction{
					Target:       target{0, 0},
					StatusEffect: StatusParalyze,
				},
			))
		})
	})

	Context("Flinching", func() {
		setup := func() *Battle {
			b := New1v1Battle(
				GeneratePokemon(PkmnPikachu, WithLevel(5), WithMoves(MoveTackle)), &a1,
				GeneratePokemon(PkmnFlareon, WithLevel(5), WithMoves(MoveTackle)), &a1,
			)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			pikachu := b.getPokemon(target{0, 0})
			b.QueueTransaction(InflictStatusTransaction{
				Target:       target{0, 0},
				StatusEffect: StatusFlinch,
			})
			b.ProcessQueue()
			Expect(pikachu.StatusEffects.check(StatusFlinch)).To(BeTrue())
			return b
		}

		It("should not allow flinching pokemon to attack", func() {
			b := setup()
			t, _ := b.SimulateRound()
			Expect(t).To(Not(HaveTransaction(UseMoveTransaction{
				User: target{0, 0},
			})))
			Expect(t).To(HaveTransaction(ImmobilizeTransaction{
				Target:       target{0, 0},
				StatusEffect: StatusFlinch,
			}))
		})

		It("should not remain flinched after the round has ended", func() {
			b := setup()
			pikachu := b.getPokemon(target{0, 0})
			b.SimulateRound()
			Expect(pikachu.StatusEffects.check(StatusFlinch)).To(BeFalse())
		})

		It("should be immobilized because of flinching, not paralysis", func() {
			b := setup()
			b.QueueTransaction(InflictStatusTransaction{
				Target:       target{0, 0},
				StatusEffect: StatusParalyze,
			})
			b.ProcessQueue()
			b.rng = NeverRNG()
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(ImmobilizeTransaction{
				Target:       target{0, 0},
				StatusEffect: StatusFlinch,
			}))
		})
	})
})
