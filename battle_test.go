package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var defaultMoveOpt = WithMoves(GetMove(MovePound))

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
	pkmn1 := GeneratePokemon(PkmnCharmander, defaultMoveOpt)
	pkmn2 := GeneratePokemon(PkmnSquirtle, defaultMoveOpt)
	party1 := NewOccupiedParty(&_a1, 0, pkmn1)
	party2 := NewOccupiedParty(&_a2, 1, pkmn2)
	b := NewBattle()
	b.AddParty(party1, party2)
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
			party1 := NewOccupiedParty(&agent1, 0, GeneratePokemon(PkmnCharmander, defaultMoveOpt))
			party2 := NewOccupiedParty(&agent2, 1, GeneratePokemon(PkmnSquirtle, defaultMoveOpt))
			b := NewBattle()
			b.AddParty(party1, party2)
			b.SetSeed(849823)
		})

		It("panics when getting an invalid Pokemon", func() {
			party := NewOccupiedParty(&agent1, 0, GeneratePokemon(PkmnBulbasaur, defaultMoveOpt))
			b := NewBattle()
			b.AddParty(party)
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
			b := NewBattle()
			party1 := NewOccupiedParty(&agent1, 0)
			party2 := NewOccupiedParty(&agent2, 1, GeneratePokemon(PkmnBulbasaur, defaultMoveOpt))
			b.AddParty(party1, party2)
			Expect(b.Start()).NotTo(Succeed())
		})

		It("should fail when both parties are on the same team", func() {
			b := NewBattle()
			party1 := NewOccupiedParty(&agent1, 0, GeneratePokemon(PkmnBulbasaur, defaultMoveOpt))
			party2 := NewOccupiedParty(&agent2, 0, GeneratePokemon(PkmnBulbasaur, defaultMoveOpt))
			b.AddParty(party1, party2)
			Expect(b.Start()).NotTo(Succeed())
		})
	})
})

var _ = Describe("One round of battle", func() {
	agent1 := Agent(new(dumbAgent))
	agent2 := Agent(new(dumbAgent))

	var (
		party1     *party
		party2     *party
		battle     *Battle
		charmander *Pokemon
		squirtle   *Pokemon
	)

	BeforeEach(func() {
		charmander = GeneratePokemon(PkmnCharmander, defaultMoveOpt)
		party1 = NewOccupiedParty(&agent1, 0, charmander)
		squirtle = GeneratePokemon(PkmnSquirtle, defaultMoveOpt)
		party2 = NewOccupiedParty(&agent2, 1, squirtle)
		battle = NewBattle()
		battle.AddParty(party1, party2)
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
					Pokemon:   *squirtle,
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
					Pokemon:   *charmander,
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
			charmander = GeneratePokemon(PkmnCharmander, WithMoves(GetMove(MovePound)))
			party1 = NewOccupiedParty(&agent1, 0, charmander)
			bidoof := GeneratePokemon(PkmnBidoof, WithMoves(GetMove(MoveTackle)))
			party2 = NewOccupiedParty(&agent2, 1, bidoof)
			battle = NewBattle()
			battle.AddParty(party1, party2)
			battle.rng = SimpleRNG()

			charmander.Moves[0] = GetMove(MoveEmber)
			Expect(battle.Start()).To(Succeed())
			battle.SimulateRound()
			Expect(bidoof.CurrentHP).To(BeEquivalentTo(7))
			bidoof.CurrentHP = 100
			charmander.Ability = AbilityAdaptability
			battle.SimulateRound()
			Expect(bidoof.CurrentHP).To(BeEquivalentTo(93))
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
						GetMove(MoveFireFang),
						GetMove(MoveTackle),
					),
				)
				party1 = NewOccupiedParty(&_a1, 0, pkmn1)
				pkmn2 := GeneratePokemon(
					PkmnTurtwig,
					WithMoves(GetMove(MoveTackle)),
					WithIVs([6]uint8{31, 31, 31, 31, 31, 0}),
				)
				party2 = NewOccupiedParty(&_a2, 1, pkmn2)
				b.AddParty(party1, party2)
				Expect(b.Start()).To(Succeed())

				// TODO: test the difference in damage between the transactions rather than the exact values of the transactions
				// TODO: make it so that target doesn't need to include `Pokemon` or `Team`
				a1 <- FightTurn{Move: 1, Target: target{Pokemon: *pkmn2, party: 1, partySlot: 0, Team: 1}}
				a2 <- FightTurn{Move: 0, Target: target{Pokemon: *pkmn1, party: 0, partySlot: 0, Team: 0}}
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(DamageTransaction{
					User: pkmn1,
					Target: target{
						Pokemon:   *pkmn2,
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

				a1 <- FightTurn{Move: 0, Target: target{Pokemon: *pkmn2, party: 1, partySlot: 0, Team: 1}}
				a2 <- FightTurn{Move: 0, Target: target{Pokemon: *pkmn1, party: 0, partySlot: 0, Team: 0}}
				t, _ = b.SimulateRound()
				Expect(t).To(HaveTransaction(DamageTransaction{
					User: pkmn1,
					Target: target{
						Pokemon:   *pkmn2,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
					Move:   GetMove(MoveFireFang),
					Damage: 8,
				}))
			})

			It("should have no effect", func() {
				pkmn1 := GeneratePokemon(
					PkmnGastly,
					WithMoves(GetMove(MoveShadowBall)),
				)
				party1 = NewOccupiedParty(&_a1, 0, pkmn1)
				pkmn2 := GeneratePokemon(
					PkmnBidoof,
					WithMoves(GetMove(MoveTackle)),
				)
				party2 = NewOccupiedParty(&_a2, 1, pkmn2)
				b.AddParty(party1, party2)
				Expect(b.Start()).To(Succeed())

				// TODO: make it so that target doesn't need to include `Pokemon` or `Team`
				a1 <- FightTurn{Move: 0, Target: target{Pokemon: *pkmn2, party: 1, partySlot: 0, Team: 1}}
				a2 <- FightTurn{Move: 0, Target: target{Pokemon: *pkmn1, party: 0, partySlot: 0, Team: 0}}
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(DamageTransaction{
					User: pkmn1,
					Target: target{
						Pokemon:   *pkmn2,
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
						Pokemon:   *pkmn1,
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
		party1 *party
		party2 *party
		battle *Battle
	)

	BeforeEach(func() {
		pkmn = GeneratePokemon(PkmnVenusaur, WithLevel(50), defaultMoveOpt)
		pkmn.CurrentHP = 10
		party1 = NewOccupiedParty(&agent, 0, pkmn)
		pkmn2 = GeneratePokemon(PkmnWartortle, WithLevel(50), defaultMoveOpt)
		party2 = NewOccupiedParty(&agent, 1, pkmn2)
		battle = NewBattle()
		battle.AddParty(party1, party2)
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
						Pokemon:   *pkmn,
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
		party1 *party
		party2 *party
		battle *Battle
	)

	BeforeEach(func() {
		party1 = NewOccupiedParty(&agent1, 0,
			GeneratePokemon(PkmnCharmander, defaultMoveOpt),
			GeneratePokemon(PkmnSquirtle, defaultMoveOpt),
			GeneratePokemon(PkmnMetapod, defaultMoveOpt),
		)
		party2 = NewOccupiedParty(&agent2, 1, GeneratePokemon(PkmnBeedrill, defaultMoveOpt))
		battle = NewBattle()
		battle.AddParty(party1, party2)
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
			for _, party := range []*party{party1, party2} {
				allies := battle.GetAllies(party)
				Expect(allies).To(HaveLen(1))
			}
		})
	})

	Context("when getting opponent Pokemon", func() {
		It("should return targets whose team does not match the passed party ", func() {
			Expect(battle.Start()).To(Succeed())
			for _, party := range []*party{party1, party2} {
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
			bulbasaur := GeneratePokemon(PkmnBulbasaur, defaultMoveOpt)
			charmander := GeneratePokemon(PkmnCharmander, defaultMoveOpt)
			p1 := NewOccupiedParty(&a1, 0, bulbasaur)
			p2 := NewOccupiedParty(&a2, 1, charmander)
			b := NewBattle()
			b.AddParty(p1, p2)
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
						Pokemon:   *charmander,
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
			p1 := GeneratePokemon(PkmnBulbasaur, WithLevel(5), defaultMoveOpt)
			p1.Stats[StatSpeed] = 100
			party1 := NewOccupiedParty(&a1, 0, p1)
			p2 := GeneratePokemon(PkmnCharmander, WithLevel(5), WithMoves(GetMove(MoveFakeOut)))
			p2.Stats[StatSpeed] = 10
			party2 := NewOccupiedParty(&a2, 1, p2)
			b := NewBattle()
			b.AddParty(party1, party2)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransactionsInOrder(
				DamageTransaction{
					User: p2,
					Target: target{
						Pokemon:   *p1,
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
						Pokemon:   *p2,
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
			ninjask := GeneratePokemon(PkmnNinjask, defaultMoveOpt)
			p1 := NewOccupiedParty(&a1, 0, charmander)
			p2 := NewOccupiedParty(&a2, 1, ninjask) // ninjask is faster than charmander
			b := NewBattle()
			b.AddParty(p1, p2)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransactionsInOrder(
				DamageTransaction{
					User: ninjask,
					Target: target{
						Pokemon:   *charmander,
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
						Pokemon:   *ninjask,
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
	var (
		p1 *party
		p2 *party
		b  *Battle
	)

	BeforeEach(func() {
		p1 = NewParty(&a1, 0)
		p2 = NewParty(&a2, 1)
		b = NewBattle()
		b.AddParty(p1, p2)
		b.rng = SimpleRNG()
	})

	Context("when using certain moves/certain abilities cause weather", func() {
		// TODO: https://bulbapedia.bulbagarden.net/wiki/Weather#Causing_weather
		It("should clear fog when using MoveDefog", func() {
			pkmn1 := GeneratePokemon(PkmnBulbasaur, WithMoves(GetMove(MoveDefog)))
			pkmn2 := GeneratePokemon(PkmnMagikarp, WithMoves(GetMove(MoveSplash)))
			p1.AddPokemon(pkmn1)
			p2.AddPokemon(pkmn2)
			b.Weather = WeatherFog
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(WeatherTransaction{
				Weather: WeatherClearSkies,
			}))
		})
		It("should cause harsh sunlight", func() {
			p1.AddPokemon(GeneratePokemon(PkmnCharmander, WithMoves(GetMove(MoveSunnyDay))))
			p2.AddPokemon(GeneratePokemon(PkmnMagikarp, WithMoves(GetMove(MoveSplash))))
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
			p1.AddPokemon(GeneratePokemon(PkmnCharmander, WithMoves(GetMove(MoveSunnyDay))))
			p2.AddPokemon(GeneratePokemon(PkmnMagikarp, WithMoves(GetMove(MoveSplash))))
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(b.metadata[MetaWeatherTurns]).To(Equal(4))
			Expect(b.Weather != WeatherFog)
			p1.pokemon[0].Moves[0] = GetMove(MoveSplash)
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
				machamp := GeneratePokemon(PkmnMachamp, WithLevel(100), WithMoves(GetMove(MoveFlamethrower)))
				bidoof := GeneratePokemon(PkmnBidoof, WithLevel(100), WithMoves(GetMove(MoveTackle)))
				p1.AddPokemon(machamp)
				p2.AddPokemon(bidoof)
				b.Weather = WeatherHarshSunlight
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(
					DamageTransaction{
						User: machamp,
						Target: target{
							Pokemon:   *bidoof,
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
				lileep := GeneratePokemon(PkmnLileep, WithLevel(100), WithMoves(GetMove(MoveBrine)))
				bidoof := GeneratePokemon(PkmnBidoof, WithLevel(100), WithMoves(GetMove(MoveTackle)))
				p1.AddPokemon(lileep)
				p2.AddPokemon(bidoof)
				b.Weather = WeatherHarshSunlight
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(
					DamageTransaction{
						User: lileep,
						Target: target{
							Pokemon:   *bidoof,
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
				machamp := GeneratePokemon(PkmnMachamp, WithLevel(100), WithMoves(GetMove(MoveFlamethrower)))
				bidoof := GeneratePokemon(PkmnBidoof, WithLevel(100), WithMoves(GetMove(MoveTackle)))
				p1.AddPokemon(machamp)
				p2.AddPokemon(bidoof)
				b.Weather = WeatherRain
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()
				// Fire weakened
				Expect(t).To(HaveTransaction(
					DamageTransaction{
						User: machamp,
						Target: target{
							Pokemon:   *bidoof,
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
				lileep := GeneratePokemon(PkmnLileep, WithLevel(100), WithMoves(GetMove(MoveBrine)))
				bidoof := GeneratePokemon(PkmnBidoof, WithLevel(100), WithMoves(GetMove(MoveTackle)))
				p1.AddPokemon(lileep)
				p2.AddPokemon(bidoof)
				b.Weather = WeatherRain
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()
				// Water boosted
				Expect(t).To(HaveTransaction(
					DamageTransaction{
						User: lileep,
						Target: target{
							Pokemon:   *bidoof,
							party:     1,
							partySlot: 0,
							Team:      1,
						},
						Damage: 125,
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
					WithMoves(GetMove(MoveTackle)),
				)
				bulbasaur := GeneratePokemon(PkmnBulbasaur,
					WithLevel(5),
					WithIVs([6]uint8{31, 0, 31, 0, 31, 0}),
					WithEVs([6]uint8{200, 0, 31, 0, 31, 0}),
					WithMoves(GetMove(MoveSolarBeam)),
				)
				p1.AddPokemon(bidoof)
				p2.AddPokemon(bulbasaur)
				Expect(b.Start()).To(Succeed())

				// TODO: compare doing solar beam WITH sandstorm deals less damage than WITHOUT sandstorm.
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(DamageTransaction{
					User: bulbasaur,
					Target: target{
						Pokemon:   *bidoof,
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
						Pokemon:   *bidoof,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					Move:   solarBeam,
					Damage: 10,
				}))
			})

			It("should cause sandstorm damage", func() {
				bidoof := GeneratePokemon(PkmnBidoof, WithMoves(GetMove(MoveTackle)))
				bulbasaur := GeneratePokemon(PkmnBulbasaur, WithMoves(GetMove(MoveSolarBeam)))
				p1.AddPokemon(bidoof)
				p2.AddPokemon(bulbasaur)
				b.Weather = WeatherHail
				b.metadata[MetaWeatherTurns] = 5
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()

				Expect(t).To(HaveTransaction(DamageTransaction{
					User: nil,
					Target: target{
						Pokemon:   *bulbasaur,
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
					WithMoves(GetMove(MoveTackle)),
				)
				bulbasaur := GeneratePokemon(PkmnBulbasaur,
					WithLevel(5),
					WithIVs([6]uint8{31, 0, 31, 0, 31, 0}),
					WithEVs([6]uint8{200, 0, 31, 0, 31, 0}),
					WithMoves(GetMove(MoveSolarBeam)),
				)
				p1.AddPokemon(bidoof)
				p2.AddPokemon(bulbasaur)
				Expect(b.Start()).To(Succeed())

				// TODO: compare doing solar beam WITH hail deals less damage than WITHOUT hail.
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(DamageTransaction{
					User: bulbasaur,
					Target: target{
						Pokemon:   *bidoof,
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
						Pokemon:   *bidoof,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					Move:   solarBeam,
					Damage: 10,
				}))
			})

			It("should cause hail damage", func() {
				bidoof := GeneratePokemon(PkmnBidoof, WithMoves(GetMove(MoveTackle)))
				bulbasaur := GeneratePokemon(PkmnBulbasaur, WithMoves(GetMove(MoveSolarBeam)))
				p1.AddPokemon(bidoof)
				p2.AddPokemon(bulbasaur)
				b.Weather = WeatherHail
				b.metadata[MetaWeatherTurns] = 5
				Expect(b.Start()).To(Succeed())
				t, _ := b.SimulateRound()

				Expect(t).To(HaveTransaction(DamageTransaction{
					User: nil,
					Target: target{
						Pokemon:   *bulbasaur,
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
			castform := GeneratePokemon(PkmnCastform, WithLevel(10), WithMoves(weatherBall))
			bulbasaur := GeneratePokemon(PkmnBulbasaur, WithLevel(10), WithMoves(solarBeam))
			p1.AddPokemon(castform)
			p2.AddPokemon(bulbasaur)
			b.Weather = WeatherFog
			b.metadata[MetaWeatherTurns] = 5
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			// TODO: Accuracy decreases from fog
			// Solar beam weakened
			Expect(t).To(HaveTransaction(DamageTransaction{
				User: bulbasaur,
				Target: target{
					Pokemon:   *castform,
					party:     0,
					partySlot: 0,
					Team:      0,
				},
				Move:   solarBeam,
				Damage: 13,
			}))
			bulbasaur.Moves[0] = moonlight
			bulbasaur.CurrentHP = bulbasaur.MaxHP()
			t, _ = b.SimulateRound()
			// Moonlight heals 1/4 max HP, weather ball boosted
			Expect(t).To(HaveTransaction(DamageTransaction{
				User: castform,
				Target: target{
					Pokemon:   *bulbasaur,
					party:     1,
					partySlot: 0,
					Team:      1,
				},
				Move:   weatherBall,
				Damage: 22,
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
		p1 *party
		p2 *party
		b  *Battle
	)

	BeforeEach(func() {
		scary_monster := GeneratePokemon(PkmnSquirtle, WithLevel(100), defaultMoveOpt)
		scary_monster.Stats[StatSpeed] = 1
		p1 = NewOccupiedParty(&a1, 0,
			GeneratePokemon(PkmnCharmander, defaultMoveOpt),
			GeneratePokemon(PkmnTurtwig, defaultMoveOpt),
		)
		p2 = NewOccupiedParty(&a2, 1, scary_monster)
		b = NewBattle()
		b.AddParty(p1, p2)
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
						Pokemon:   *b.parties[0].pokemon[0],
						party:     0,
						partySlot: 0,
						Team:      0,
					},
				}, SendOutTransaction{
					Target: target{
						Pokemon:   *b.parties[0].pokemon[1],
						party:     0,
						partySlot: 1,
						Team:      0,
					},
				}))
		})

		It("should not allow fainted Pokemon to take turns", func() {
			party1 := NewParty(&a1, 0)
			pkmn1 := GeneratePokemon(PkmnCharmander, WithLevel(3), defaultMoveOpt)
			pkmn2 := GeneratePokemon(PkmnSquirtle, WithLevel(10), defaultMoveOpt)
			pkmn3 := GeneratePokemon(PkmnTurtwig, WithLevel(3), defaultMoveOpt)
			pkmn1.CurrentHP = 1
			party1.AddPokemon(pkmn1, pkmn3)
			party2 := NewParty(&a2, 1)
			pkmn2.Stats[StatSpeed] = 255
			party2.AddPokemon(pkmn2)
			b := NewBattle()
			b.AddParty(party1, party2)
			Expect(b.Start()).To(Succeed())
			t, ended := b.SimulateRound()
			Expect(ended).To(BeFalse(), "Expected SimulateRound to NOT indicate that the battle has ended, but it did.")
			Expect(t).To(HaveTransactionsInOrder(
				FaintTransaction{
					Target: target{
						Pokemon:   *pkmn1,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
				},
				SendOutTransaction{
					Target: target{
						Pokemon:   *pkmn3,
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
						Pokemon:   *pkmn2,
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
			p1 := NewOccupiedParty(&a1, 0, dies)
			p2 := NewOccupiedParty(&a2, 1, GeneratePokemon(PkmnCharmander, WithLevel(25), defaultMoveOpt))
			b = NewBattle()
			b.AddParty(p1, p2)
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(dies.Friendship).To(Equal(99))
		})

		It("should lose 5 or 10 friendship when fainting", func() {
			dies := GeneratePokemon(PkmnBulbasaur, WithLevel(1), defaultMoveOpt)
			dies.Friendship = 100
			dies2 := GeneratePokemon(PkmnBulbasaur, WithLevel(1), defaultMoveOpt)
			dies2.Friendship = 200
			p1 := NewOccupiedParty(&a1, 0, dies, dies2)
			p2 := NewOccupiedParty(&a2, 1, GeneratePokemon(PkmnCharmander, WithLevel(100), defaultMoveOpt))
			b = NewBattle()
			b.AddParty(p1, p2)
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(dies.Friendship).To(Equal(95))
			b.SimulateRound()
			Expect(dies2.Friendship).To(Equal(190))
		})

		It("should gain EVs when defeating Pokemon", func() {
			winner := GeneratePokemon(PkmnBulbasaur, WithLevel(100), defaultMoveOpt)
			loser := GeneratePokemon(PkmnBulbasaur, defaultMoveOpt)
			p1 = NewOccupiedParty(&a1, 0, winner)
			p2 = NewOccupiedParty(&a2, 1, loser)
			b = NewBattle()
			b.AddParty(p1, p2)
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
		party1 := NewParty(&a1, 0)
		pkmn1 = GeneratePokemon(PkmnCharmander, WithLevel(3), defaultMoveOpt)
		pkmn1.CurrentHP = 1
		party1.AddPokemon(pkmn1)
		party2 := NewParty(&a2, 1)
		pkmn2 = GeneratePokemon(PkmnSquirtle, WithLevel(10), defaultMoveOpt)
		pkmn2.Stats[StatSpeed] = 255
		party2.AddPokemon(pkmn2)
		b = NewBattle()
		b.AddParty(party1, party2)
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
						Pokemon:   *pkmn1,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
				},
				EndBattleTransaction{
					Reason: EndKnockout,
				},
			))
			Expect(t).ToNot(HaveTransaction(
				DamageTransaction{
					User: pkmn1,
					Target: target{
						Pokemon:   *pkmn2,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
				},
			))
		})
	})
})

var _ = Describe("Battle metadata", func() {
	Context("Pokemon metadata", func() {
		It("should record the last used move of Pokemon in battle", func() {
			a1 := Agent(new(dumbAgent))
			razorLeaf := GetMove(MoveRazorLeaf)
			p1 := NewOccupiedParty(&a1, 0,
				GeneratePokemon(PkmnBulbasaur, WithMoves(razorLeaf)),
			)
			ember := GetMove(MoveEmber)
			p2 := NewOccupiedParty(&a1, 1, GeneratePokemon(
				PkmnCharmander, WithMoves(ember)),
			)
			b := NewBattle()
			b.AddParty(p1, p2)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			Expect(p1.pokemon[0].metadata).ToNot(HaveKeyWithValue(MetaLastMove, razorLeaf))
			b.SimulateRound()
			Expect(p1.pokemon[0].metadata).To(HaveKeyWithValue(MetaLastMove, razorLeaf))
			Expect(p2.pokemon[0].metadata).To(HaveKeyWithValue(MetaLastMove, ember))
		})
	})
})

var _ = Describe("Status Conditions", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))
	var (
		p1 *party
		p2 *party
		b  *Battle
	)

	BeforeEach(func() {
		b = NewBattle()
		b.rng = AlwaysRNG()
	})

	Context("when using certain moves in battle causes status effects", func() {
		It("should inflict paralysis from MoveStunSpore", func() {
			pkmn1 := GeneratePokemon(PkmnBulbasaur, WithMoves(GetMove(MoveSplash)))
			pkmn2 := GeneratePokemon(PkmnBulbasaur, WithMoves(GetMove(MoveStunSpore)))
			p1 = NewOccupiedParty(&a1, 0, pkmn1)
			p2 = NewOccupiedParty(&a2, 1, pkmn2)
			b.AddParty(p1, p2)
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
			p1 = NewOccupiedParty(&a1, 0, pkmn1)
			pkmn2 := GeneratePokemon(PkmnIvysaur, defaultMoveOpt)
			pkmn2.StatusEffects = StatusBurn
			p2 = NewOccupiedParty(&a2, 1, pkmn2)
			b.AddParty(p1, p2)
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{
					Pokemon:   *pkmn1,
					party:     0,
					partySlot: 0,
					Team:      0,
				},
				Damage:       1,
				StatusEffect: StatusPoison,
			}))
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{
					Pokemon:   *pkmn2,
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
			p1 = NewOccupiedParty(&a1, 0, pkmn1)
			pkmn2 := GeneratePokemon(PkmnIvysaur, WithLevel(100), defaultMoveOpt)
			p2 = NewOccupiedParty(&a2, 1, pkmn2)
			b.AddParty(p1, p2)
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{
					Pokemon:   *pkmn1,
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
			p1 = NewOccupiedParty(&a1, 0, pkmn1)
			pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
			p2 = NewOccupiedParty(&a2, 1, pkmn2)
			b.AddParty(p1, p2)
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				ImmobilizeTransaction{
					Target: target{
						Pokemon:   *pkmn1,
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
			p1 = NewOccupiedParty(&a1, 0, pkmn1)
			pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
			p2 = NewOccupiedParty(&a2, 1, pkmn2)
			b.AddParty(p1, p2)
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				ImmobilizeTransaction{
					Target: target{
						Pokemon:   *pkmn1,
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
				p1 = NewOccupiedParty(&a1, 0, pkmn1)
				pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
				p2 = NewOccupiedParty(&a2, 1, pkmn2)
				b.AddParty(p1, p2)
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
							Pokemon:   *pkmn1,
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
				p1 = NewOccupiedParty(&a1, 0, pkmn1)
				pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
				p2 = NewOccupiedParty(&a2, 1, pkmn2)
				b.AddParty(p1, p2)
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
							Pokemon:   *pkmn1,
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
				p1 = NewOccupiedParty(&a1, 0, pkmn1)
				pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
				p2 = NewOccupiedParty(&a2, 1, pkmn2)
				b.AddParty(p1, p2)
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
					pkmn1 := GeneratePokemon(PkmnBulbasaur, WithLevel(8), WithMoves(GetMove(moveid), GetMove(MoveRazorLeaf)))
					p1 = NewOccupiedParty(&a1, 0, pkmn1)
					pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
					p2 = NewOccupiedParty(&a2, 1, pkmn2)
					b.AddParty(p1, p2)
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
								Pokemon:   *pkmn1,
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
			p1 = NewOccupiedParty(&a1, 0, pkmn1)
			pkmn2 := GeneratePokemon(PkmnCharmander, WithLevel(4), defaultMoveOpt)
			p2 = NewOccupiedParty(&a2, 1, pkmn2)
			b.AddParty(p1, p2)
			Expect(b.Start()).To(Succeed())
			b.QueueTransaction(CureStatusTransaction{
				Target: target{
					Pokemon:   *pkmn1,
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
						Pokemon:   *pkmn1,
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

var _ = Describe("In-a-pinch Berries", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))
	var holder *Pokemon

	setup := func(item Item) *Battle {
		p1 := NewOccupiedParty(&a1, 0, GeneratePokemon(
			PkmnCombusken,
			WithLevel(25),
			WithMoves(GetMove(MoveSplash)),
		))
		holder = GeneratePokemon(
			PkmnGrotle,
			WithLevel(25),
			WithMoves(GetMove(MoveSplash)),
			WithIVs([6]uint8{1, 1, 1, 20, 1, 1}),
		)
		holder.HeldItem = item
		holder.CurrentHP = holder.MaxHP() / 4
		p2 := NewOccupiedParty(&a2, 1, holder)
		b := NewBattle()
		b.rng = SimpleRNG()
		b.AddParty(p1, p2)
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
					Pokemon:   *holder,
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
