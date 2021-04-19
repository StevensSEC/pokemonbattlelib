package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var defaultMoveOpt = WithMoves(MovePound)

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
	for _, target := range ctx.Allies {
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

func newRcAgent() rcAgent {
	return rcAgent(make(chan Turn, 1))
}

// Example usage for rcAgent
var _ = Describe("RC Agent", func() {
	a1 := newRcAgent()
	a2 := newRcAgent()
	_a1 := Agent(a1)
	_a2 := Agent(a2)
	party1 := NewOccupiedParty(GeneratePokemon(PkmnCharmander, WithMoves(MoveSplash)))
	party2 := NewOccupiedParty(GeneratePokemon(PkmnSquirtle, WithMoves(MoveSplash)))
	b := NewSingleBattle(party1, &_a1, party2, &_a2)
	Expect(b.Start()).To(Succeed())
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
})

var _ = Describe("Battle initialization", func() {
	agent1 := Agent(new(dumbAgent))
	agent2 := Agent(new(dumbAgent))

	Context("when creating a new battle", func() {
		It("runs without panicking", func() {
			party1 := NewOccupiedParty(GeneratePokemon(PkmnCharmander, defaultMoveOpt))
			party2 := NewOccupiedParty(GeneratePokemon(PkmnSquirtle, defaultMoveOpt))
			b := NewSingleBattle(party1, &agent1, party2, &agent2)
			b.SetSeed(849823)
		})

		It("panics when getting an invalid Pokemon", func() {
			party := NewOccupiedParty(GeneratePokemon(PkmnBulbasaur, defaultMoveOpt))
			b := NewBattle()
			b.AddParty(party, &agent1, 0)
			Expect(func() {
				b.getPokemonInBattle(1, 5)
			}).To(Panic())
			Expect(func() {
				b.getPokemonInBattle(0, 5)
			}).To(Panic())
		})
	})

	Context("team validation", func() {
		It("should fail when party has no pokemon", func() {
			b := NewSingleBattle(
				NewParty(), &agent1,
				NewOccupiedParty(GeneratePokemon(PkmnBulbasaur, defaultMoveOpt)), &agent2,
			)
			Expect(b.Start()).NotTo(Succeed())
		})

		It("should fail when both parties are on the same team", func() {
			b := NewBattle()
			b.AddParty(NewOccupiedParty(GeneratePokemon(PkmnBulbasaur, defaultMoveOpt)), &agent1, 0)
			b.AddParty(NewOccupiedParty(GeneratePokemon(PkmnBulbasaur, defaultMoveOpt)), &agent2, 0)
			Expect(b.Start()).NotTo(Succeed())
		})
	})
})

var _ = Describe("Battle memory management", func() {
	agent1 := Agent(new(dumbAgent))

	It("should produce transactions that reference the ground truth pointers", func() {
		party1 := NewOccupiedParty(GeneratePokemon(PkmnCharmander, defaultMoveOpt))
		party2 := NewOccupiedParty(GeneratePokemon(PkmnSquirtle, defaultMoveOpt))
		b := NewSingleBattle(party1, &agent1, party2, &agent1)
		Expect(b.Start()).To(Succeed())

		t, _ := b.SimulateRound()
		for i := range t {
			switch tt := t[i].(type) {
			case DamageTransaction:
				real := b.getPokemon(tt.Target)
				Expect(tt.Target.Pokemon).To(BeIdenticalTo(real))
			}
		}
	})
})

var _ = Describe("One round of battle", func() {
	//FIXME: this test suite needs to be separated into multiple suites
	agent1 := Agent(new(dumbAgent))
	agent2 := Agent(new(dumbAgent))

	var (
		battle     *Battle
		charmander *Pokemon
		squirtle   *Pokemon
	)

	BeforeEach(func() {
		charmander = GeneratePokemon(PkmnCharmander, defaultMoveOpt)
		squirtle = GeneratePokemon(PkmnSquirtle, defaultMoveOpt)
		battle = New1v1Battle(charmander, &agent1, squirtle, &agent2)
		battle.rng = SimpleRNG()
	})

	Context("when simulating a round between two agents", func() {
		It("panics if battle is not in progress", func() {
			Expect(func() {
				battle.SimulateRound()
			}).To(Panic())
		})

		It("should create 2 damage transactions", func() {
			Expect(battle.Start()).To(Succeed())
			t, _ := battle.SimulateRound()
			Expect(t).To(HaveLen(2))
			Expect(t).To(HaveTransaction(DamageTransaction{
				User: charmander,
				Target: target{
					Pokemon:   squirtle,
					party:     1,
					partySlot: 0,
					Team:      1,
				},
				Move:   GetMove(MovePound),
				Damage: 3,
			}))
			Expect(t).To(HaveTransaction(DamageTransaction{
				User: squirtle,
				Target: target{
					Pokemon:   charmander,
					party:     0,
					partySlot: 0,
					Team:      0,
				},
				Move:   GetMove(MovePound),
				Damage: 3,
			}))
		})

		It("should cause Pokemon to have reduced HP", func() {
			Expect(battle.Start()).To(Succeed())
			battle.SimulateRound()
			Expect(charmander.CurrentHP < charmander.MaxHP()).To(BeTrue())
			Expect(squirtle.CurrentHP < squirtle.MaxHP()).To(BeTrue())
		})
	})

	Context("when dealing damage to a Pokemon", func() {
		It("should account for same-type attack bonus", func() {
			charmander = GeneratePokemon(PkmnCharmander, WithMoves(MovePound))
			bidoof := GeneratePokemon(PkmnBidoof, WithMoves(MoveTackle))
			battle = New1v1Battle(charmander, &agent1, bidoof, &agent2)
			battle.rng = SimpleRNG()

			charmander.Moves[0] = GetMove(MoveEmber)
			Expect(battle.Start()).To(Succeed())
			battle.SimulateRound()
			Expect(bidoof.CurrentHP).To(BeEquivalentTo(8))
			bidoof.CurrentHP = 100
			charmander.Ability = AbilityAdaptability
			battle.SimulateRound()
			Expect(bidoof.CurrentHP).To(BeEquivalentTo(94))
		})

		Context("Type Matchups", func() {
			var (
				a1  rcAgent
				a2  rcAgent
				_a1 Agent
				_a2 Agent
				b   *Battle
			)

			BeforeEach(func() {
				a1 = newRcAgent()
				a2 = newRcAgent()
				_a1 = Agent(a1)
				_a2 = Agent(a2)
				b = NewBattle()
				b.rng = SimpleRNG()
			})

			It("should account for supereffective type matchups", func() {
				pkmn1 := GeneratePokemon(
					PkmnMightyena,
					WithIVs([6]uint8{31, 0, 31, 0, 31, 31}),
					WithMoves(
						MoveFireFang,
						MoveTackle,
					),
				)
				pkmn2 := GeneratePokemon(
					PkmnTurtwig,
					WithMoves(MoveTackle),
					WithIVs([6]uint8{31, 31, 31, 31, 31, 0}),
				)
				b := New1v1Battle(pkmn1, &_a1, pkmn2, &_a2)
				b.rng = SimpleRNG()
				Expect(b.Start()).To(Succeed())

				a1 <- FightTurn{Move: 1, Target: b.getTarget(1, 0)}
				a2 <- FightTurn{Move: 0, Target: b.getTarget(0, 0)}
				t, _ := b.SimulateRound()
				damage := DamageDealt(t, pkmn1)
				Expect(damage).To(Equal(3))

				b.QueueTransaction(HealTransaction{
					Target: pkmn2,
					Amount: 200,
				})
				b.ProcessQueue()

				a1 <- FightTurn{Move: 0, Target: b.getTarget(1, 0)}
				a2 <- FightTurn{Move: 0, Target: b.getTarget(0, 0)}
				t, _ = b.SimulateRound()
				Expect(DamageDealt(t, pkmn1)).To(BeNumerically(">", damage))
			})

			It("should have no effect", func() {
				pkmn1 := GeneratePokemon(PkmnGastly, WithMoves(MoveShadowBall))
				pkmn2 := GeneratePokemon(PkmnBidoof, WithMoves(MoveTackle))
				b := New1v1Battle(pkmn1, &_a1, pkmn2, &_a2)
				b.rng = SimpleRNG()
				Expect(b.Start()).To(Succeed())

				a1 <- FightTurn{Move: 0, Target: b.getTarget(1, 0)}
				a2 <- FightTurn{Move: 0, Target: b.getTarget(0, 0)}
				t, _ := b.SimulateRound()
				Expect(DamageDealt(t, pkmn1)).To(Equal(0))
				Expect(DamageDealt(t, pkmn2)).To(Equal(0))
			})
		})

		It("should account for critical hits", func() {
			battle.rng = AlwaysRNG()
			Expect(battle.Start()).To(Succeed())
			battle.SimulateRound()
			Expect(squirtle.CurrentHP).To(BeEquivalentTo(5))
		})

		It("should miss moves randomly based on accuracy/evasion", func() {
			battle.rng = NeverRNG()
			Expect(battle.Start()).To(Succeed())
			t, _ := battle.SimulateRound()
			Expect(t).To(HaveTransaction(MoveFailTransaction{
				User:   charmander,
				Reason: FailMiss,
			}))
			battle.rng = SimpleRNG()
			t, _ = battle.SimulateRound()
			Expect(t).ToNot(HaveTransaction(MoveFailTransaction{
				User:   charmander,
				Reason: FailMiss,
			}))
		})
	})

	Context("when certain moves are used in battle", func() {
		It("should change a move's PP", func() {
			battle.rng = AlwaysRNG()
			charmander.Moves[0] = GetMove(MoveSpite)
			Expect(battle.Start()).To(Succeed())
			battle.SimulateRound() // set Pokemon's last move
			charmander.CurrentHP = charmander.MaxHP()
			squirtle.CurrentHP = squirtle.MaxHP()
			squirtle.Moves[0].CurrentPP = 1
			t, _ := battle.SimulateRound()
			Expect(t).To(HaveTransaction(PPTransaction{
				Move:   squirtle.Moves[0],
				Amount: -4,
			}))
			// Ensure that PP stays in bounds
			Expect(squirtle.Moves[0].CurrentPP).To(BeEquivalentTo(0))
		})
	})
})

var _ = Describe("Using items in battle", func() {
	agent := Agent(new(healAgent))
	var (
		pkmn  *Pokemon
		pkmn2 *Pokemon
		b     *Battle
	)

	BeforeEach(func() {
		pkmn = GeneratePokemon(PkmnVenusaur, WithLevel(50), defaultMoveOpt)
		pkmn.CurrentHP = 10
		pkmn2 = GeneratePokemon(PkmnWartortle, WithLevel(50), defaultMoveOpt)
		b = New1v1Battle(pkmn, &agent, pkmn2, &agent)
	})

	Context("when the battle processes item turns", func() {
		It("should create ItemTransaction(s) properly", func() {
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				ItemTransaction{
					Target: b.getTarget(0, 0),
					Item:   ItemPotion,
					Move:   nil,
				},
			))
		})

		It("should heal the Pokemon by 20 HP", func() {
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(pkmn.CurrentHP).To(BeEquivalentTo(30))
		})
	})
})

var _ = Describe("Getting pokemon from parties", func() {
	agent1 := Agent(new(dumbAgent))
	agent2 := Agent(new(dumbAgent))
	var (
		battle *Battle
	)

	BeforeEach(func() {
		party1 := NewOccupiedParty(
			GeneratePokemon(PkmnCharmander, defaultMoveOpt),
			GeneratePokemon(PkmnSquirtle, defaultMoveOpt),
			GeneratePokemon(PkmnMetapod, defaultMoveOpt),
		)
		party2 := NewOccupiedParty(GeneratePokemon(PkmnBeedrill, defaultMoveOpt))
		battle = NewSingleBattle(party1, &agent1, party2, &agent2)
	})

	Context("when getting Pokemon by party/slot", func() {
		It("should get the Pokemon the user expects", func() {
			pkmn := battle.getPokemonInBattle(0, 1)
			Expect(pkmn.NatDex).To(BeEquivalentTo(PkmnSquirtle))
		})
	})

	Context("when getting ally Pokemon", func() {
		It("should return targets whose team matches the passed party", func() {
			Expect(battle.Start()).To(Succeed())
			for _, party := range battle.parties {
				allies := battle.GetAllies(party)
				Expect(allies).To(HaveLen(1))
			}
		})
	})

	Context("when getting opponent Pokemon", func() {
		It("should return targets whose team does not match the passed party ", func() {
			Expect(battle.Start()).To(Succeed())
			for _, party := range battle.parties {
				opponents := battle.GetOpponents(party)
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
		)

		It("should order turns properly based on priority", func() {
			a2 := Agent(new(healAgent))
			bulbasaur := GeneratePokemon(PkmnBulbasaur, WithMoves(MovePound))
			charmander := GeneratePokemon(PkmnCharmander, WithMoves(MovePound))
			b := New1v1Battle(bulbasaur, &a1, charmander, &a2)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransactionsInOrder(
				HealTransaction{
					Target: charmander,
					Amount: 0,
				},
				DamageTransaction{
					User:   bulbasaur,
					Target: b.getTarget(1, 0),
					Move:   GetMove(MovePound),
					Damage: 3,
				},
			))
		})
	})

	Context("when determining priority for equal turn types", func() {
		It("should handle moves with higher priority first", func() {
			p1 := GeneratePokemon(PkmnBulbasaur, WithLevel(5), WithMoves(MovePound))
			p1.Stats[StatSpeed] = 100
			p2 := GeneratePokemon(PkmnCharmander, WithLevel(5), WithMoves(MoveQuickAttack))
			p2.Stats[StatSpeed] = 10
			b := New1v1Battle(p1, &a1, p2, &a2)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransactionsInOrder(
				DamageTransaction{
					User:   p2,
					Target: b.getTarget(0, 0),
					Damage: 5,
					Move:   GetMove(MoveQuickAttack),
				},
				DamageTransaction{
					User:   p1,
					Target: b.getTarget(1, 0),
					Damage: 5,
					Move:   GetMove(MovePound),
				},
			))
		})

		It("should handle faster Pokemon first", func() {
			charmander := GeneratePokemon(PkmnCharmander, defaultMoveOpt)
			ninjask := GeneratePokemon(PkmnNinjask, defaultMoveOpt) // ninjask is faster than charmander
			b := New1v1Battle(charmander, &a1, ninjask, &a2)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransactionsInOrder(
				DamageTransaction{
					User:   ninjask,
					Target: b.getTarget(0, 0),
					Move:   GetMove(MovePound),
					Damage: 3,
				},
				DamageTransaction{
					User:   charmander,
					Target: b.getTarget(1, 0),
					Move:   GetMove(MovePound),
					Damage: 3,
				},
			))
		})
	})
})

var _ = Describe("Weather", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))

	Context("when using certain moves/certain abilities cause weather", func() {
		// TODO: https://bulbapedia.bulbagarden.net/wiki/Weather#Causing_weather
		It("should clear fog when using MoveDefog", func() {
			b := New1v1Battle(
				GeneratePokemon(PkmnBulbasaur, WithMoves(MoveDefog)), &a1,
				GeneratePokemon(PkmnMagikarp, WithMoves(MoveSplash)), &a2,
			)
			b.rng = SimpleRNG()
			b.Weather = WeatherFog
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(WeatherTransaction{
				Weather: WeatherClearSkies,
			}))
		})

		It("should cause harsh sunlight", func() {
			b := New1v1Battle(
				GeneratePokemon(PkmnCharmander, WithMoves(MoveSunnyDay)), &a1,
				GeneratePokemon(PkmnMagikarp, WithMoves(MoveSplash)), &a2,
			)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(WeatherTransaction{
				Weather: WeatherHarshSunlight,
				Turns:   5,
			}))
			Expect(b.metadata[MetaWeatherTurns]).To(Equal(4))
		})
	})

	Context("when weather is present, battles are affected", func() {
		moonlight := GetMove(MoveMoonlight)
		It("should use metadata to track weather and clear weather over time", func() {
			b := New1v1Battle(
				GeneratePokemon(PkmnCharmander, WithMoves(MoveSunnyDay)), &a1,
				GeneratePokemon(PkmnMagikarp, WithMoves(MoveSplash)), &a2,
			)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(b.metadata[MetaWeatherTurns]).To(Equal(4))
			Expect(b.Weather).ToNot(Equal(WeatherFog))
			b.getPokemonInBattle(0, 0).Moves[0] = GetMove(MoveSplash)
			b.SimulateRound()
			Expect(b.metadata[MetaWeatherTurns]).To(Equal(3))
			b.metadata[MetaWeatherTurns] = 0
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(WeatherTransaction{
				Weather: WeatherClearSkies,
			}))
		})

		When("harsh sunlight", func() {
			It("should boost fire type moves", func() {
				// intentional non water type pokemon, intentional no supereffective type matchup
				machamp := GeneratePokemon(PkmnMachamp, WithLevel(100), WithMoves(MoveFlamethrower))
				bidoof := GeneratePokemon(PkmnBidoof, WithLevel(100), WithMoves(MoveTackle))
				b := New1v1Battle(machamp, &a1, bidoof, &a2)
				b.rng = SimpleRNG()
				b.Weather = WeatherHarshSunlight
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()
				Expect(DamageDealt(t, machamp)).To(Equal(183))
			})

			It("should weaken water type moves", func() {
				// intentional non water type pokemon, intentional no supereffective type matchup
				lileep := GeneratePokemon(PkmnLileep, WithLevel(100), WithMoves(MoveBrine))
				bidoof := GeneratePokemon(PkmnBidoof, WithLevel(100), WithMoves(MoveTackle))
				b := New1v1Battle(lileep, &a1, bidoof, &a2)
				b.rng = SimpleRNG()
				b.Weather = WeatherHarshSunlight
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()
				Expect(DamageDealt(t, lileep)).To(Equal(41))
			})
		})

		When("raining", func() {
			It("should affect fire type moves", func() {
				// intentional non water type pokemon, intentional no supereffective type matchup
				machamp := GeneratePokemon(PkmnMachamp, WithLevel(100), WithMoves(MoveFlamethrower))
				bidoof := GeneratePokemon(PkmnBidoof, WithLevel(100), WithMoves(MoveTackle))
				b := New1v1Battle(machamp, &a1, bidoof, &a2)
				b.rng = SimpleRNG()
				b.Weather = WeatherRain
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()
				// Fire weakened
				Expect(DamageDealt(t, machamp)).To(Equal(61))
			})

			It("should affect water type moves", func() {
				// intentional non water type pokemon, intentional no supereffective type matchup
				lileep := GeneratePokemon(PkmnLileep, WithLevel(100), WithMoves(MoveBrine))
				bidoof := GeneratePokemon(PkmnBidoof, WithLevel(100), WithMoves(MoveTackle))
				b := New1v1Battle(lileep, &a1, bidoof, &a2)
				b.rng = SimpleRNG()
				b.Weather = WeatherRain
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()
				// Water boosted
				Expect(DamageDealt(t, lileep)).To(Equal(124))
			})
		})

		// sandstorm tests and hailing tests could be tablized
		When("sandstorm", func() {
			It("should weaken solar beam", func() {
				bidoof := GeneratePokemon(PkmnBidoof,
					WithLevel(5),
					WithIVs([6]uint8{31, 0, 31, 0, 31, 0}),
					WithEVs([6]uint8{200, 0, 31, 0, 31, 0}),
					WithMoves(MoveTackle),
				)
				bulbasaur := GeneratePokemon(PkmnBulbasaur,
					WithLevel(5),
					WithIVs([6]uint8{31, 0, 31, 0, 31, 0}),
					WithEVs([6]uint8{200, 0, 31, 0, 31, 0}),
					WithMoves(MoveSolarBeam),
				)
				b := New1v1Battle(bidoof, &a1, bulbasaur, &a2)
				b.rng = SimpleRNG()
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()
				solarbeamDamage := DamageDealt(t, bulbasaur)
				Expect(solarbeamDamage).To(Equal(18))
				Expect(t).ToNot(HaveTransaction(FaintTransaction{}))
				b.QueueTransaction(
					WeatherTransaction{
						Weather: WeatherSandstorm,
						Turns:   5,
					},
					HealTransaction{
						Target: b.getPokemonInBattle(0, 0),
						Amount: 100,
					},
					HealTransaction{
						Target: b.getPokemonInBattle(1, 0),
						Amount: 100,
					},
				)
				b.ProcessQueue()
				t, _ = b.SimulateRound()
				Expect(DamageDealt(t, bulbasaur)).To(BeNumerically("<", solarbeamDamage))
			})

			It("should cause sandstorm damage", func() {
				bidoof := GeneratePokemon(PkmnBidoof, WithMoves(MoveTackle))
				bulbasaur := GeneratePokemon(PkmnBulbasaur, WithMoves(MoveSolarBeam))
				b := New1v1Battle(bidoof, &a1, bulbasaur, &a2)
				b.rng = SimpleRNG()
				b.Weather = WeatherHail
				b.metadata[MetaWeatherTurns] = 5
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(DamageTransaction{
					User:   nil,
					Target: b.getTarget(1, 0),
					Move:   nil,
					Damage: 0,
				}))
			})
		})

		When("hailing", func() {
			It("should weaken solar beam", func() {
				bidoof := GeneratePokemon(PkmnBidoof,
					WithLevel(5),
					WithIVs([6]uint8{31, 0, 31, 0, 31, 0}),
					WithEVs([6]uint8{200, 0, 31, 0, 31, 0}),
					WithMoves(MoveTackle),
				)
				bulbasaur := GeneratePokemon(PkmnBulbasaur,
					WithLevel(5),
					WithIVs([6]uint8{31, 0, 31, 0, 31, 0}),
					WithEVs([6]uint8{200, 0, 31, 0, 31, 0}),
					WithMoves(MoveSolarBeam),
				)
				b := New1v1Battle(bidoof, &a1, bulbasaur, &a2)
				b.rng = SimpleRNG()
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()
				solarbeamDamage := DamageDealt(t, bulbasaur)
				Expect(solarbeamDamage).To(Equal(18))
				Expect(t).ToNot(HaveTransaction(FaintTransaction{}))
				b.QueueTransaction(
					WeatherTransaction{
						Weather: WeatherHail,
						Turns:   5,
					},
					HealTransaction{
						Target: b.getPokemonInBattle(0, 0),
						Amount: 100,
					},
					HealTransaction{
						Target: b.getPokemonInBattle(1, 0),
						Amount: 100,
					},
				)
				b.ProcessQueue()
				t, _ = b.SimulateRound()
				Expect(DamageDealt(t, bulbasaur)).To(BeNumerically("<", solarbeamDamage))
			})

			It("should cause hail damage", func() {
				bidoof := GeneratePokemon(PkmnBidoof, WithMoves(MoveTackle))
				bulbasaur := GeneratePokemon(PkmnBulbasaur, WithMoves(MoveSolarBeam))
				b := New1v1Battle(bidoof, &a1, bulbasaur, &a2)
				b.rng = SimpleRNG()
				b.Weather = WeatherHail
				b.metadata[MetaWeatherTurns] = 5
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()

				Expect(t).To(HaveTransaction(DamageTransaction{
					User:   nil,
					Target: b.getTarget(1, 0),
					Move:   nil,
					Damage: 0,
				}))
			})
		})

		It("should cause side effects during fog", func() {
			castform := GeneratePokemon(PkmnCastform, WithLevel(10), WithMoves(MoveWeatherBall))
			bulbasaur := GeneratePokemon(PkmnBulbasaur, WithLevel(10), WithMoves(MoveSolarBeam))
			b := New1v1Battle(castform, &a1, bulbasaur, &a2)
			b.rng = SimpleRNG()
			baseAccuracy := CalcAccuracy(WeatherClearSkies, castform, bulbasaur, GetMove(MovePound))
			fogAccuracy := CalcAccuracy(WeatherFog, castform, bulbasaur, GetMove(MovePound))
			Expect(fogAccuracy).To(BeNumerically("<", baseAccuracy))
			b.Weather = WeatherFog
			b.metadata[MetaWeatherTurns] = 5
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			// Solar beam weakened
			Expect(DamageDealt(t, bulbasaur)).To(Equal(12))
			bulbasaur.Moves[0] = moonlight
			bulbasaur.CurrentHP = bulbasaur.MaxHP()
			t, _ = b.SimulateRound()
			// Moonlight heals 1/4 max HP, weather ball boosted
			Expect(DamageDealt(t, castform)).To(Equal(21))
			Expect(t).To(HaveTransaction(HealTransaction{
				Target: bulbasaur,
				Amount: 7,
			}))
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

	BeforeEach(func() {
		scary_monster := GeneratePokemon(PkmnSquirtle, WithLevel(100), WithMoves(MovePound))
		scary_monster.Stats[StatSpeed] = 1
		p1 = NewOccupiedParty(
			GeneratePokemon(PkmnCharmander, WithMoves(MovePound)),
			GeneratePokemon(PkmnTurtwig, WithMoves(MovePound)),
		)
		p2 = NewOccupiedParty(scary_monster) // and nice sprites
		b = NewSingleBattle(p1, &a1, p2, &a2)
		b.rng = SimpleRNG()
	})

	Context("after a Pokemon faints in battle", func() {
		It("should switch to the next available Pokemon", func() {
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
					Target: b.getTarget(0, 0),
				}, SendOutTransaction{
					Target: b.getTarget(0, 1),
				}))
		})

		It("should not allow fainted Pokemon to take turns", func() {
			pkmn1 := GeneratePokemon(PkmnCharmander, WithLevel(3), defaultMoveOpt)
			pkmn2 := GeneratePokemon(PkmnSquirtle, WithLevel(10), defaultMoveOpt)
			pkmn3 := GeneratePokemon(PkmnTurtwig, WithLevel(3), defaultMoveOpt)
			party1 := NewOccupiedParty(pkmn1, pkmn3)
			pkmn1.CurrentHP = 1
			party1.AddPokemon(pkmn1, pkmn3)
			party2 := NewOccupiedParty(pkmn2)
			pkmn2.Stats[StatSpeed] = 255
			party2.AddPokemon(pkmn2)
			b := NewSingleBattle(party1, &a1, party2, &a2)
			Expect(b.Start()).To(Succeed())
			t, ended := b.SimulateRound()
			Expect(ended).To(BeFalse(), "Expected SimulateRound to NOT indicate that the battle has ended, but it did.")
			Expect(t).To(HaveTransactionsInOrder(
				FaintTransaction{
					Target: b.getTarget(0, 0),
				},
				SendOutTransaction{
					Target: b.getTarget(0, 1),
				},
			))
			Expect(t).ToNot(HaveTransaction(
				DamageTransaction{
					User:   pkmn1,
					Target: b.getTarget(1, 0),
				},
			))
		})

		It("should lose 1 friendship when fainting", func() {
			dies := GeneratePokemon(PkmnBulbasaur, WithLevel(1), defaultMoveOpt)
			dies.Friendship = 100
			p1 := NewOccupiedParty(dies)
			p2 := NewOccupiedParty(GeneratePokemon(PkmnCharmander, WithLevel(25), defaultMoveOpt))
			b = NewBattle()
			b.AddParty(p1, &a1, 0)
			b.AddParty(p2, &a2, 1)
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(dies.Friendship).To(Equal(99))
		})

		It("should lose 5 or 10 friendship when fainting", func() {
			dies := GeneratePokemon(PkmnBulbasaur, WithLevel(1), defaultMoveOpt)
			dies.Friendship = 100
			dies2 := GeneratePokemon(PkmnBulbasaur, WithLevel(1), defaultMoveOpt)
			dies2.Friendship = 200
			p1 := NewOccupiedParty(dies, dies2)
			p2 := NewOccupiedParty(GeneratePokemon(PkmnCharmander, WithLevel(100), defaultMoveOpt))
			b = NewSingleBattle(p1, &a1, p2, &a2)
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(dies.Friendship).To(Equal(95))
			b.SimulateRound()
			Expect(dies2.Friendship).To(Equal(190))
		})

		It("should gain EVs when defeating Pokemon", func() {
			winner := GeneratePokemon(PkmnBulbasaur, WithLevel(100), defaultMoveOpt)
			loser := GeneratePokemon(PkmnBulbasaur, defaultMoveOpt)
			p1 = NewOccupiedParty(winner)
			p2 = NewOccupiedParty(loser)
			b = NewSingleBattle(p1, &a1, p2, &a2)
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				EVTransaction{
					Target: winner,
					Stat:   StatSpAtk,
					Amount: 1,
				},
			))
		})
	})

	When("holding Focus Sash", func() {
		setup := func() *Battle {
			holder := GeneratePokemon(PkmnMachoke, WithLevel(26), WithMoves(MoveSplash))
			holder.CurrentHP = 4
			holder.HeldItem = ItemFocusSash
			b = New1v1Battle(holder, &a1, GeneratePokemon(PkmnGrotle, WithLevel(30), WithMoves(MoveRazorLeaf)), &a2)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			return b
		}

		It("should not let the holder die", func() {
			b := setup()
			t, _ := b.SimulateRound()
			holderTarget := b.getTarget(0, 0)
			Expect(t).To(HaveTransaction(DamageTransaction{
				User:   b.getPokemonInBattle(1, 0),
				Target: holderTarget,
			}))
			Expect(t).ToNot(HaveTransaction(FaintTransaction{
				Target: holderTarget,
			}))
			Expect(b.parties[0].activePokemon[0].CurrentHP).To(BeEquivalentTo(1))
		})

		It("should consume the focus sash after damage is applied", func() {
			b := setup()
			t, _ := b.SimulateRound()
			target := b.getTarget(0, 0)
			Expect(t).To(HaveTransactionsInOrder(
				DamageTransaction{
					User:   b.getPokemonInBattle(1, 0),
					Target: target,
				},
				ItemTransaction{
					Target: target,
					IsHeld: true,
					Item:   ItemFocusSash,
				},
			))
			Expect(b.getPokemonInBattle(0, 0).HeldItem).To(Equal(ItemNone))
			Expect(b.getPokemonInBattle(1, 0).HeldItem).To(Equal(ItemNone))
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
		pkmn1 = GeneratePokemon(PkmnCharmander, WithLevel(3), defaultMoveOpt)
		pkmn1.CurrentHP = 1
		party1 := NewOccupiedParty(pkmn1)
		pkmn2 = GeneratePokemon(PkmnSquirtle, WithLevel(10), defaultMoveOpt)
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
					Target: b.getTarget(0, 0),
				},
				EndBattleTransaction{
					Reason: EndKnockout,
					Winner: 1,
				},
			))
			Expect(t).ToNot(HaveTransaction(
				DamageTransaction{
					User:   pkmn1,
					Target: b.getTarget(1, 0),
				},
			))
			Expect(b.GetResults().Winner).To(Equal(1))
		})
	})
})

var _ = Describe("Battle metadata", func() {
	Context("Pokemon metadata", func() {
		It("should record the last used move of Pokemon in battle", func() {
			a1 := Agent(new(dumbAgent))
			razorLeaf := GetMove(MoveRazorLeaf)
			p1 := GeneratePokemon(PkmnBulbasaur, WithMoves(MoveRazorLeaf))
			ember := GetMove(MoveEmber)
			p2 := GeneratePokemon(PkmnCharmander, WithMoves(MoveEmber))
			b := New1v1Battle(p1, &a1, p2, &a1)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			Expect(p1.metadata).ToNot(HaveKeyWithValue(MetaLastMove, razorLeaf))
			b.SimulateRound()
			Expect(p1.metadata).To(HaveKeyWithValue(MetaLastMove, razorLeaf))
			Expect(p2.metadata).To(HaveKeyWithValue(MetaLastMove, ember))
		})
	})
})

var _ = Describe("Status Conditions", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))

	Context("when using certain moves in battle causes status effects", func() {
		It("should inflict paralysis from MoveStunSpore", func() {
			pkmn1 := GeneratePokemon(PkmnBulbasaur, WithMoves(MoveSplash))
			pkmn2 := GeneratePokemon(PkmnBulbasaur, WithMoves(MoveStunSpore))
			b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
			b.rng = AlwaysRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				InflictStatusTransaction{
					Target:       pkmn1,
					StatusEffect: StatusParalyze,
				},
			))
		})
	})

	Context("when a Pokemon has a nonvolatile status effect, it affects the Pokemon in battle", func() {
		It("should inflict burn and poison damage", func() {
			pkmn1 := GeneratePokemon(PkmnBulbasaur, defaultMoveOpt)
			pkmn1.StatusEffects = StatusPoison
			pkmn2 := GeneratePokemon(PkmnIvysaur, defaultMoveOpt)
			pkmn2.StatusEffects = StatusBurn
			b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
			b.rng = AlwaysRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target:       b.getTarget(0, 0),
				Damage:       1,
				StatusEffect: StatusPoison,
			}))
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target:       b.getTarget(1, 0),
				Damage:       1,
				StatusEffect: StatusBurn,
			}))
		})

		It("should inflict badly poisoned damage", func() {
			pkmn1 := GeneratePokemon(PkmnBulbasaur, WithLevel(100), defaultMoveOpt)
			pkmn1.StatusEffects = StatusBadlyPoison
			pkmn2 := GeneratePokemon(PkmnIvysaur, WithLevel(100), defaultMoveOpt)
			b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
			b.rng = AlwaysRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target:       b.getTarget(0, 0),
				Damage:       12,
				StatusEffect: StatusBadlyPoison,
			}))
		})

		It("should immobilize paralyzed Pokemon", func() {
			pkmn1 := GeneratePokemon(PkmnBulbasaur, WithLevel(8), defaultMoveOpt)
			pkmn1.StatusEffects = StatusParalyze
			pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
			b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
			b.rng = AlwaysRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				ImmobilizeTransaction{
					Target:       b.getTarget(0, 0),
					StatusEffect: StatusParalyze,
				},
			))
		})

		It("should immobilize frozen Pokemon", func() {
			pkmn1 := GeneratePokemon(PkmnBulbasaur, WithLevel(8), defaultMoveOpt)
			pkmn1.StatusEffects = StatusFreeze
			pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
			b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
			b.rng = AlwaysRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				ImmobilizeTransaction{
					Target:       b.getTarget(0, 0),
					StatusEffect: StatusFreeze,
				},
			))
		})

		Context("when a pokemon is asleep", func() {
			It("should immobilize sleeping Pokemon", func() {
				pkmn1 := GeneratePokemon(PkmnBulbasaur, WithLevel(8), defaultMoveOpt)
				pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
				b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
				b.rng = AlwaysRNG()
				Expect(b.Start()).To(Succeed())
				b.QueueTransaction(InflictStatusTransaction{
					Target:       pkmn1,
					StatusEffect: StatusSleep,
				})
				b.ProcessQueue()
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(
					ImmobilizeTransaction{
						Target:       b.getTarget(0, 0),
						StatusEffect: StatusSleep,
					},
				))
			})

			It("should allow sleeping Pokemon to wake up", func() {
				pkmn1 := GeneratePokemon(PkmnBulbasaur, WithLevel(8), defaultMoveOpt)
				pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
				b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
				b.rng = AlwaysRNG()
				Expect(b.Start()).To(Succeed())
				b.QueueTransaction(InflictStatusTransaction{
					Target:       pkmn1,
					StatusEffect: StatusSleep,
				})
				b.ProcessQueue()
				pkmn1.metadata[MetaSleepTime] = 0
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(
					CureStatusTransaction{
						Target:       b.getTarget(0, 0),
						StatusEffect: StatusSleep,
					},
				))
				Expect(pkmn1.StatusEffects.check(StatusSleep)).To(BeFalse())
			})

			It("should decrement sleeping counter", func() {
				pkmn1 := GeneratePokemon(PkmnBulbasaur, WithLevel(8), defaultMoveOpt)
				pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
				b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
				b.rng = AlwaysRNG()
				Expect(b.Start()).To(Succeed())
				b.QueueTransaction(InflictStatusTransaction{
					Target:       pkmn1,
					StatusEffect: StatusSleep,
				})
				b.ProcessQueue()
				counter := pkmn1.metadata[MetaSleepTime].(int)
				b.SimulateRound()
				Expect(pkmn1.metadata[MetaSleepTime].(int)).To(Equal(counter - 1))
			})

			DescribeTable("Sleep walking",
				func(moveid MoveId) {
					pkmn1 := GeneratePokemon(PkmnBulbasaur, WithLevel(8), WithMoves(moveid, MoveRazorLeaf))
					pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
					b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
					b.rng = AlwaysRNG()
					Expect(b.Start()).To(Succeed())
					b.QueueTransaction(InflictStatusTransaction{
						Target:       pkmn1,
						StatusEffect: StatusSleep,
					})
					b.ProcessQueue()
					t, _ := b.SimulateRound()
					Expect(t).ToNot(HaveTransaction(
						ImmobilizeTransaction{
							Target:       b.getTarget(0, 0),
							StatusEffect: StatusSleep,
						},
					))
				},
				Entry("Snore", MoveSnore),
				Entry("Sleep Talk", MoveSleepTalk),
			)
		})

		It("should cure paralysis", func() {
			pkmn1 := GeneratePokemon(PkmnBulbasaur, WithLevel(8), defaultMoveOpt)
			pkmn1.StatusEffects = StatusParalyze
			pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
			b := New1v1Battle(pkmn1, &a1, pkmn2, &a2)
			b.rng = AlwaysRNG()
			Expect(b.Start()).To(Succeed())
			b.QueueTransaction(CureStatusTransaction{
				Target:       b.getTarget(0, 0),
				StatusEffect: StatusParalyze,
			})
			b.ProcessQueue()
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				CureStatusTransaction{
					Target:       b.getTarget(0, 0),
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
			pikachu := b.getPokemonInBattle(0, 0)
			b.QueueTransaction(InflictStatusTransaction{
				Target:       pikachu,
				StatusEffect: StatusFlinch,
			})
			b.ProcessQueue()
			Expect(pikachu.StatusEffects.check(StatusFlinch)).To(BeTrue())
			return b
		}

		It("should not allow flinching pokemon to attack", func() {
			b := setup()
			pikachu := b.getPokemonInBattle(0, 0)
			t, _ := b.SimulateRound()
			Expect(t).To(Not(HaveTransaction(DamageTransaction{
				User: pikachu,
			})))
			Expect(t).To(HaveTransaction(ImmobilizeTransaction{
				Target:       b.getTarget(0, 0),
				StatusEffect: StatusFlinch,
			}))
		})

		It("should not remain flinched after the round has ended", func() {
			b := setup()
			pikachu := b.getPokemonInBattle(0, 0)
			b.SimulateRound()
			Expect(pikachu.StatusEffects.check(StatusFlinch)).To(BeFalse())
		})

		It("should be immobilized because of flinching, not paralysis", func() {
			b := setup()
			pikachu := b.getPokemonInBattle(0, 0)
			b.QueueTransaction(InflictStatusTransaction{
				Target:       pikachu,
				StatusEffect: StatusParalyze,
			})
			b.ProcessQueue()
			b.rng = NeverRNG()
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(ImmobilizeTransaction{
				Target:       b.getTarget(0, 0),
				StatusEffect: StatusFlinch,
			}))
		})
	})
})

var _ = Describe("Draining moves", func() {
	a1 := Agent(new(dumbAgent))
	var b *Battle

	BeforeEach(func() {
		b = NewSingleBattle(
			NewOccupiedParty(
				GeneratePokemon(PkmnRoselia,
					WithLevel(25),
					WithMoves(MoveGigaDrain),
				),
			),
			&a1,
			NewOccupiedParty(
				GeneratePokemon(PkmnBidoof,
					WithLevel(25),
					WithMoves(MoveSplash),
				),
			),
			&a1,
		)
		b.rng = SimpleRNG()
		Expect(b.Start()).To(Succeed())
	})

	It("should damage the target and heal the user", func() {
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransactionsInOrder(
			DamageTransaction{
				User:   b.getPokemonInBattle(0, 0),
				Target: b.getTarget(1, 0),
				Damage: 61,
			},
			HealTransaction{
				Target: b.getPokemonInBattle(0, 0),
				Amount: 30,
			},
		))
	})

	It("should heal more when the user is holding a big root", func() {
		b.getPokemonInBattle(0, 0).HeldItem = ItemBigRoot
		t, _ := b.SimulateRound()

		Expect(t).To(HaveTransactionsInOrder(
			DamageTransaction{
				User:   b.getPokemonInBattle(0, 0),
				Target: b.getTarget(1, 0),
				Damage: 61,
			},
			HealTransaction{
				Target: b.getPokemonInBattle(0, 0),
				Amount: 39,
			},
		))
	})
})

var _ = Describe("Recoil moves", func() {
	a1 := Agent(new(dumbAgent))
	var b *Battle

	BeforeEach(func() {
		b = NewSingleBattle(
			NewOccupiedParty(
				GeneratePokemon(PkmnPikachu,
					WithLevel(25),
					WithMoves(MoveVoltTackle),
				),
			),
			&a1,
			NewOccupiedParty(
				GeneratePokemon(PkmnBidoof,
					WithLevel(25),
					WithMoves(MoveSplash),
				),
			),
			&a1,
		)
		b.rng = SimpleRNG()
		Expect(b.Start()).To(Succeed())
	})

	It("should damage the target and damage the user", func() {
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransactionsInOrder(
			DamageTransaction{
				User:   b.getPokemonInBattle(0, 0),
				Target: b.getTarget(1, 0),
				Damage: 57,
			},
			DamageTransaction{
				Target: b.getTarget(0, 0),
				Damage: 18,
			},
		))
	})
})

var _ = Describe("Move Effects", func() {
	a1 := Agent(new(dumbAgent))

	It("should cause flinching", func() {
		b := New1v1Battle(
			GeneratePokemon(PkmnMightyena, WithLevel(20), WithIVs([6]uint8{0, 0, 0, 0, 0, 31}), WithEVs([6]uint8{0, 0, 0, 0, 0, 252}), WithMoves(MoveBite)), &a1,
			GeneratePokemon(PkmnPonyta, WithLevel(20), WithIVs([6]uint8{31, 0, 31, 0, 31, 0}), WithEVs([6]uint8{252, 0, 0, 0, 0, 0}), WithMoves(MoveTackle)), &a1,
		)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransaction(DamageTransaction{
			User:   b.getPokemonInBattle(0, 0),
			Target: b.getTarget(1, 0),
		}))
		Expect(t).To(HaveTransactionsInOrder(
			InflictStatusTransaction{
				Target:       b.getPokemonInBattle(1, 0),
				StatusEffect: StatusFlinch,
			},
			ImmobilizeTransaction{
				Target:       b.getTarget(1, 0),
				StatusEffect: StatusFlinch,
			},
		))
	})

	It("should raise speed on flinch for pokemon with steadfast ability", func() {
		b := New1v1Battle(
			GeneratePokemon(PkmnMightyena, WithLevel(20), WithIVs([6]uint8{0, 0, 0, 0, 0, 31}), WithEVs([6]uint8{0, 0, 0, 0, 0, 252}), WithMoves(MoveBite)), &a1,
			GeneratePokemon(PkmnPonyta, WithLevel(20), WithIVs([6]uint8{31, 0, 31, 0, 31, 0}), WithEVs([6]uint8{252, 0, 0, 0, 0, 0}), WithMoves(MoveTackle), WithAbility(AbilitySteadfast)), &a1,
		)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransaction(DamageTransaction{
			User:   b.getPokemonInBattle(0, 0),
			Target: b.getTarget(1, 0),
		}))
		Expect(t).To(HaveTransactionsInOrder(
			InflictStatusTransaction{
				Target:       b.getPokemonInBattle(1, 0),
				StatusEffect: StatusFlinch,
			},
			ModifyStatTransaction{
				Target: b.getPokemonInBattle(1, 0),
				Stat:   StatSpeed,
				Stages: 1,
			},
			ImmobilizeTransaction{
				Target:       b.getTarget(1, 0),
				StatusEffect: StatusFlinch,
			},
		))
	})
})
