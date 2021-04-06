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

				// TODO: test the difference in damage between the transactions rather than the exact values of the transactions
				// TODO: make it so that target doesn't need to include `Pokemon` or `Team`
				a1 <- FightTurn{Move: 1, Target: target{Pokemon: pkmn2, party: 1, partySlot: 0, Team: 1}}
				a2 <- FightTurn{Move: 0, Target: target{Pokemon: pkmn1, party: 0, partySlot: 0, Team: 0}}
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(DamageTransaction{
					User: pkmn1,
					Target: target{
						Pokemon:   pkmn2,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
					Move:   GetMove(MoveTackle),
					Damage: 3,
				}))

				b.QueueTransaction(HealTransaction{
					Target: pkmn2,
					Amount: 200,
				})
				b.ProcessQueue()

				a1 <- FightTurn{Move: 0, Target: target{Pokemon: pkmn2, party: 1, partySlot: 0, Team: 1}}
				a2 <- FightTurn{Move: 0, Target: target{Pokemon: pkmn1, party: 0, partySlot: 0, Team: 0}}
				t, _ = b.SimulateRound()
				Expect(t).To(HaveTransaction(DamageTransaction{
					User: pkmn1,
					Target: target{
						Pokemon:   pkmn2,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
					Move:   GetMove(MoveFireFang),
					Damage: 8,
				}))
			})

			It("should have no effect", func() {
				pkmn1 := GeneratePokemon(PkmnGastly, WithMoves(MoveShadowBall))
				pkmn2 := GeneratePokemon(PkmnBidoof, WithMoves(MoveTackle))
				b := New1v1Battle(pkmn1, &_a1, pkmn2, &_a2)
				b.rng = SimpleRNG()
				Expect(b.Start()).To(Succeed())

				// TODO: make it so that target doesn't need to include `Pokemon` or `Team`
				a1 <- FightTurn{Move: 0, Target: target{Pokemon: pkmn2, party: 1, partySlot: 0, Team: 1}}
				a2 <- FightTurn{Move: 0, Target: target{Pokemon: pkmn1, party: 0, partySlot: 0, Team: 0}}
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(DamageTransaction{
					User: pkmn1,
					Target: target{
						Pokemon:   pkmn2,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
					Move:   GetMove(MoveShadowBall),
					Damage: 0,
				}))
				Expect(t).To(HaveTransaction(DamageTransaction{
					User: pkmn2,
					Target: target{
						Pokemon:   pkmn1,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					Move:   GetMove(MoveTackle),
					Damage: 0,
				}))
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
		DescribeTable("Changing Pokemon stat modifiers",
			func(id MoveId, stat, stages int) {
				charmander.Moves[0] = GetMove(id)
				Expect(battle.Start()).To(Succeed())
				t, _ := battle.SimulateRound()
				Expect(t).To(HaveTransaction(ModifyStatTransaction{
					Target: charmander,
					Stat:   stat,
					Stages: stages,
				}))
				// Bound by min/max stat modifier
				charmander.StatModifiers[stat] = MaxStatModifier
				t, _ = battle.SimulateRound()
				Expect(t).To(HaveTransaction(ModifyStatTransaction{
					Target: charmander,
					Stat:   stat,
					Stages: stages,
				}))
				Expect(charmander.StatModifiers[stat]).To(BeEquivalentTo(MaxStatModifier))
			},
			Entry("Howl", MoveHowl, StatAtk, +1),
			Entry("Double Team", MoveDoubleTeam, StatEvasion, +1),
		)

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
		pkmn   *Pokemon
		pkmn2  *Pokemon
		battle *Battle
	)

	BeforeEach(func() {
		pkmn = GeneratePokemon(PkmnVenusaur, WithLevel(50), defaultMoveOpt)
		pkmn.CurrentHP = 10
		pkmn2 = GeneratePokemon(PkmnWartortle, WithLevel(50), defaultMoveOpt)
		battle = New1v1Battle(pkmn, &agent, pkmn2, &agent)
	})

	Context("when the battle processes item turns", func() {
		It("should create ItemTransaction(s) properly", func() {
			Expect(battle.Start()).To(Succeed())
			t, _ := battle.SimulateRound()
			Expect(t).To(HaveTransaction(
				ItemTransaction{
					Target: target{
						party:     0,
						partySlot: 0,
						Team:      0,
						Pokemon:   pkmn,
					},
					Item: ItemPotion,
					Move: nil,
				},
			))
		})

		It("should heal the Pokemon by 20 HP", func() {
			Expect(battle.Start()).To(Succeed())
			battle.SimulateRound()
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
					User: bulbasaur,
					Target: target{
						Pokemon:   charmander,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
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
			p2 := GeneratePokemon(PkmnCharmander, WithLevel(5), WithMoves(MoveFakeOut))
			p2.Stats[StatSpeed] = 10
			b := New1v1Battle(p1, &a1, p2, &a2)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransactionsInOrder(
				DamageTransaction{
					User: p2,
					Target: target{
						Pokemon:   p1,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					Damage: 5,
					Move:   GetMove(MoveFakeOut),
				},
				DamageTransaction{
					User: p1,
					Target: target{
						Pokemon:   p2,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
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
					User: ninjask,
					Target: target{
						Pokemon:   charmander,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					Move:   GetMove(MovePound),
					Damage: 3,
				},
				DamageTransaction{
					User: charmander,
					Target: target{
						Pokemon:   ninjask,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
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
		solarBeam := GetMove(MoveSolarBeam)
		weatherBall := GetMove(MoveWeatherBall)
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
				Expect(t).To(HaveTransaction(
					DamageTransaction{
						User: machamp,
						Target: target{
							Pokemon:   bidoof,
							party:     1,
							partySlot: 0,
							Team:      1,
						},
						Damage: 183,
					},
				))
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
				Expect(t).To(HaveTransaction(
					DamageTransaction{
						User: lileep,
						Target: target{
							Pokemon:   bidoof,
							party:     1,
							partySlot: 0,
							Team:      1,
						},
						Damage: 41,
					},
				))
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
				Expect(t).To(HaveTransaction(
					DamageTransaction{
						User: machamp,
						Target: target{
							Pokemon:   bidoof,
							party:     1,
							partySlot: 0,
							Team:      1,
						},
						Damage: 61,
					},
				))
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
				Expect(t).To(HaveTransaction(
					DamageTransaction{
						User: lileep,
						Target: target{
							Pokemon:   bidoof,
							party:     1,
							partySlot: 0,
							Team:      1,
						},
						Damage: 124,
					},
				))
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

				// TODO: compare doing solar beam WITH sandstorm deals less damage than WITHOUT sandstorm.
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(DamageTransaction{
					User: bulbasaur,
					Target: target{
						Pokemon:   bidoof,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					Move:   solarBeam,
					Damage: 18,
				}))

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
				Expect(t).To(HaveTransaction(DamageTransaction{
					User: bulbasaur,
					Target: target{
						Pokemon:   bidoof,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					Move:   solarBeam,
					Damage: 10,
				}))
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
					User: nil,
					Target: target{
						Pokemon:   bulbasaur,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
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

				// TODO: compare doing solar beam WITH hail deals less damage than WITHOUT hail.
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(DamageTransaction{
					User: bulbasaur,
					Target: target{
						Pokemon:   bidoof,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					Move:   solarBeam,
					Damage: 18,
				}))

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
				Expect(t).To(HaveTransaction(DamageTransaction{
					User: bulbasaur,
					Target: target{
						Pokemon:   bidoof,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					Move:   solarBeam,
					Damage: 10,
				}))
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
					User: nil,
					Target: target{
						Pokemon:   bulbasaur,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
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
			b.Weather = WeatherFog
			b.metadata[MetaWeatherTurns] = 5
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			// TODO: Accuracy decreases from fog
			// Solar beam weakened
			Expect(t).To(HaveTransaction(DamageTransaction{
				User: bulbasaur,
				Target: target{
					Pokemon:   castform,
					party:     0,
					partySlot: 0,
					Team:      0,
				},
				Move:   solarBeam,
				Damage: 12,
			}))
			bulbasaur.Moves[0] = moonlight
			bulbasaur.CurrentHP = bulbasaur.MaxHP()
			t, _ = b.SimulateRound()
			// Moonlight heals 1/4 max HP, weather ball boosted
			Expect(t).To(HaveTransaction(DamageTransaction{
				User: castform,
				Target: target{
					Pokemon:   bulbasaur,
					party:     1,
					partySlot: 0,
					Team:      1,
				},
				Move:   weatherBall,
				Damage: 21,
			}))
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
					Target: target{
						Pokemon:   b.parties[0].pokemon()[0],
						party:     0,
						partySlot: 0,
						Team:      0,
					},
				}, SendOutTransaction{
					Target: target{
						Pokemon:   b.parties[0].pokemon()[1],
						party:     0,
						partySlot: 1,
						Team:      0,
					},
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
					Target: target{
						Pokemon:   pkmn1,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
				},
				SendOutTransaction{
					Target: target{
						Pokemon:   pkmn3,
						party:     0,
						partySlot: 1,
						Team:      0,
					},
				},
			))
			Expect(t).ToNot(HaveTransaction(
				DamageTransaction{
					User: pkmn1,
					Target: target{
						Pokemon:   pkmn2,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
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
			holderTarget := target{
				Pokemon:   b.getPokemonInBattle(0, 0),
				party:     0,
				partySlot: 0,
				Team:      0,
			}
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
			target := target{
				Pokemon:   b.getPokemonInBattle(0, 0),
				party:     0,
				partySlot: 0,
				Team:      0,
			}
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
					Target: target{
						Pokemon:   pkmn1,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
				},
				EndBattleTransaction{
					Reason: EndKnockout,
					Winner: 1,
				},
			))
			Expect(t).ToNot(HaveTransaction(
				DamageTransaction{
					User: pkmn1,
					Target: target{
						Pokemon:   pkmn2,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
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

	Context("when a Pokemon has a status effect, it affects the Pokemon in battle", func() {
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
				Target: target{
					Pokemon:   pkmn1,
					party:     0,
					partySlot: 0,
					Team:      0,
				},
				Damage:       1,
				StatusEffect: StatusPoison,
			}))
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{
					Pokemon:   pkmn2,
					party:     1,
					partySlot: 0,
					Team:      1,
				},
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
				Target: target{
					Pokemon:   pkmn1,
					party:     0,
					partySlot: 0,
					Team:      0,
				},
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
					Target: target{
						Pokemon:   pkmn1,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
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
					Target: target{
						Pokemon:   pkmn1,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
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
						Target: target{
							Pokemon:   pkmn1,
							party:     0,
							partySlot: 0,
							Team:      0,
						},
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
						Target: target{
							Pokemon:   pkmn1,
							party:     0,
							partySlot: 0,
							Team:      0,
						},
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
							Target: target{
								Pokemon:   pkmn1,
								party:     0,
								partySlot: 0,
								Team:      0,
							},
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
				Target: target{
					Pokemon:   pkmn1,
					party:     0,
					partySlot: 0,
					Team:      0,
				},
				StatusEffect: StatusParalyze,
			})
			b.ProcessQueue()
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				CureStatusTransaction{
					Target: target{
						Pokemon:   pkmn1,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					StatusEffect: StatusParalyze,
				},
			))
		})
	})
})

var _ = Describe("Misc/held items", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))
	var holder *Pokemon

	setup := func(item Item, pkmn int) (*Battle, *Pokemon) {
		p1 := GeneratePokemon(PkmnSnorlax, WithLevel(25), WithMoves(MoveSplash))
		holder = GeneratePokemon(pkmn, WithLevel(25), WithMoves(MoveSplash))
		holder.HeldItem = item
		b := New1v1Battle(p1, &a1, holder, &a2)
		b.rng = SimpleRNG()
		Expect(b.Start()).To(Succeed())
		return b, holder
	}

	Context("when Pokemon hold certain misc. items in battle", func() {
		It("handles Black Sludge", func() {
			// Heal poison types for 1/16 HP
			b, holder := setup(ItemBlackSludge, PkmnGrimer)
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(HealTransaction{
				Target: holder,
				Amount: holder.MaxHP() / 16,
			}))
			// Damage non-poison types for 1/8 HP
			b, holder = setup(ItemBlackSludge, PkmnAerodactyl)
			t, _ = b.SimulateRound()
			Expect(t).ToNot(HaveTransaction(HealTransaction{
				Target: holder,
				Amount: holder.MaxHP() / 16,
			}))
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{
					party:     1,
					partySlot: 0,
					Team:      1,
					Pokemon:   holder,
				},
				Damage: holder.MaxHP() / 8,
			}))
		})

		It("handles Destiny Knot", func() {
			b, holder := setup(ItemDestinyKnot, PkmnMimeJr)
			attacker := b.getPokemonInBattle(0, 0)
			attacker.Moves[0] = GetMove(MoveAttract)
			attacker.Gender = GenderMale
			holder.Gender = GenderFemale
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(InflictStatusTransaction{
				Target:       attacker,
				StatusEffect: StatusInfatuation,
			}))
		})

		It("handles Expert Belt", func() {
			b, holder := setup(ItemExpertBelt, PkmnMachamp)
			holder.Moves[0] = GetMove(MoveCloseCombat)
			t, _ := b.SimulateRound()
			// Damage boosted by 20%
			Expect(t).To(HaveTransaction(
				DamageTransaction{
					User: holder,
					Target: target{
						Pokemon:   b.getPokemonInBattle(0, 0),
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					Damage: 201,
				},
			))
		})

		It("handles Leftovers", func() {
			b, holder := setup(ItemLeftovers, PkmnSnorlax)
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(HealTransaction{
				Target: holder,
				Amount: holder.MaxHP() / 16,
			}))
		})

		It("handles Life Orb", func() {
			b, holder := setup(ItemLifeOrb, PkmnSnorlax)
			holder.Moves[0] = GetMove(MoveTackle)
			t, _ := b.SimulateRound()
			// Boost damage by 30%
			Expect(t).To(HaveTransaction(DamageTransaction{
				User: holder,
				Target: target{
					party:     0,
					partySlot: 0,
					Team:      0,
					Pokemon:   b.getPokemonInBattle(0, 0),
				},
				Damage: 32,
			}))
			// Take 10% of max HP
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{
					party:     1,
					partySlot: 0,
					Team:      1,
					Pokemon:   holder,
				},
				Damage: holder.MaxHP() / 10,
			}))
		})

		It("handles Muscle Band", func() {
			b, holder := setup(ItemMuscleBand, PkmnSnorlax)
			holder.Moves[0] = GetMove(MoveTackle)
			t, _ := b.SimulateRound()
			// Boost physical move damage by 10%
			Expect(t).To(HaveTransaction(DamageTransaction{
				User: holder,
				Target: target{
					party:     0,
					partySlot: 0,
					Team:      0,
					Pokemon:   b.getPokemonInBattle(0, 0),
				},
				Damage: 27,
			}))
		})

		It("handles Shell Bell", func() {
			b, holder := setup(ItemShellBell, PkmnSnorlax)
			holder.Moves[0] = GetMove(MoveTackle)
			t, _ := b.SimulateRound()
			// Self-inflict 1/8 of dealt damage
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{
					party:     1,
					partySlot: 0,
					Team:      1,
					Pokemon:   holder,
				},
				Damage: 3,
			}))
		})

		It("handles White Herb", func() {
			b, holder := setup(ItemWhiteHerb, PkmnSnorlax)
			holder.StatModifiers = [9]int{-1, -1, -1, -1, -1, -1, 0, -1, -1}
			b.SimulateRound()
			// Consumed after use
			Expect(holder.HeldItem).To(Equal(ItemNone))
			// Reset all lowered stat modifiers
			for _, stage := range holder.StatModifiers {
				if stage < 0 {
					Fail("Expected all lowered stats to be reset")
				}
			}
		})

		It("handles Wise Glasses", func() {
			b, holder := setup(ItemWiseGlasses, PkmnSnorlax)
			holder.Moves[0] = GetMove(MoveSurf)
			t, _ := b.SimulateRound()
			// Boost special move damage by 10%
			Expect(t).To(HaveTransaction(DamageTransaction{
				User: holder,
				Target: target{
					party:     0,
					partySlot: 0,
					Team:      0,
					Pokemon:   b.getPokemonInBattle(0, 0),
				},
				Damage: 16,
			}))
		})

		DescribeTable("Status curing held items",
			func(item Item, status StatusCondition) {
				b, holder := setup(item, PkmnSnorlax)
				holder.StatusEffects.apply(status)
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(CureStatusTransaction{
					Target: target{
						party:     1,
						partySlot: 0,
						Team:      1,
						Pokemon:   holder,
					},
					StatusEffect: status,
				}))
				// Item should be consumed after use
				Expect(holder.HeldItem).To(Equal(ItemNone))
			},
			Entry("Mental Herb", ItemMentalHerb, StatusInfatuation),
		)

		DescribeTable("Flinch inducing items",
			func(item Item) {
				b, holder := setup(ItemKingsRock, PkmnLucario)
				holder.Moves[0] = GetMove(MoveTackle)
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(InflictStatusTransaction{
					Target:       b.getPokemonInBattle(0, 0),
					StatusEffect: StatusFlinch,
				}))
			},
			Entry("King's Rock", ItemKingsRock),
			Entry("Razor Fang", ItemRazorFang),
		)

		DescribeTable("Weather duration boosting rocks",
			func(item Item, weather Weather, move MoveId) {
				b, holder := setup(item, PkmnCastform)
				holder.Moves[0] = GetMove(move)
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(WeatherTransaction{
					Weather: weather,
					Turns:   8,
				}))
			},
			Entry("Damp Rock", ItemDampRock, WeatherRain, MoveRainDance),
			Entry("Heat Rock", ItemHeatRock, WeatherHarshSunlight, MoveSunnyDay),
			Entry("Icy Rock", ItemIcyRock, WeatherHail, MoveHail),
			Entry("Smooth Rock", ItemSmoothRock, WeatherSandstorm, MoveSandstorm),
		)
	})
})

var _ = Describe("In-a-pinch Berries", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))
	var holder *Pokemon

	setup := func(item Item) *Battle {
		holder = GeneratePokemon(
			PkmnGrotle,
			WithLevel(25),
			WithMoves(MoveSplash),
			WithIVs([6]uint8{1, 1, 1, 20, 1, 1}),
		)
		holder.HeldItem = item
		holder.CurrentHP = holder.MaxHP() / 4
		b := New1v1Battle(
			GeneratePokemon(PkmnCombusken, WithLevel(25), WithMoves(MoveSplash)), &a1,
			holder, &a2,
		)
		b.rng = SimpleRNG()
		Expect(b.Start()).To(Succeed())
		return b
	}

	DescribeTable("Stat changing in-a-pinch berries",
		func(item Item, stat, stages int) {
			b := setup(item)
			t, _ := b.SimulateRound()

			Expect(t).To(HaveTransaction(ItemTransaction{
				Target: target{
					party:     1,
					partySlot: 0,
					Team:      1,
					Pokemon:   holder,
				},
				IsHeld: true,
				Item:   holder.HeldItem,
			}))
			Expect(t).To(HaveTransaction(ModifyStatTransaction{
				Target: holder,
				Stat:   stat,
				Stages: stages,
			}))
			Expect(b.parties[0].activePokemon[0].HeldItem).To(Equal(ItemNone))
			Expect(b.parties[1].activePokemon[0].HeldItem).To(Equal(ItemNone))
		},
		Entry("Apicot Berry", ItemApicotBerry, StatSpDef, 1),
		Entry("Ganlon Berry", ItemGanlonBerry, StatDef, 1),
		Entry("Lansat Berry", ItemLansatBerry, StatCritChance, 2),
		Entry("Liechi Berry", ItemLiechiBerry, StatAtk, 1),
		Entry("Petaya Berry", ItemPetayaBerry, StatSpAtk, 1),
		Entry("Salac Berry", ItemSalacBerry, StatSpeed, 1),
	)
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
				User: b.getPokemonInBattle(0, 0),
				Target: target{
					Pokemon:   b.getPokemonInBattle(1, 0),
					party:     1,
					partySlot: 0,
					Team:      1,
				},
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
				User: b.getPokemonInBattle(0, 0),
				Target: target{
					Pokemon:   b.getPokemonInBattle(1, 0),
					party:     1,
					partySlot: 0,
					Team:      1,
				},
				Damage: 61,
			},
			HealTransaction{
				Target: b.getPokemonInBattle(0, 0),
				Amount: 39,
			},
		))
	})
})
