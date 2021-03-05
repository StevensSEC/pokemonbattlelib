package pokemonbattlelib

import (
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
	item := GetItem(ItemPotion)
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
var _ = Describe("RC Agent", func() {
	a1 := newRcAgent()
	a2 := newRcAgent()
	_a1 := Agent(a1)
	_a2 := Agent(a2)
	pkmn1 := GeneratePokemon(PkmnCharmander, WithMoves(GetMove(MovePound)))
	pkmn2 := GeneratePokemon(PkmnSquirtle, WithMoves(GetMove(MovePound)))
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
			party1 := NewOccupiedParty(&agent1, 0, GeneratePokemon(PkmnCharmander))
			party2 := NewOccupiedParty(&agent2, 1, GeneratePokemon(PkmnSquirtle))
			b := NewBattle()
			b.AddParty(party1, party2)
			b.SetSeed(849823)
		})

		It("panics when getting an invalid Pokemon", func() {
			party := NewOccupiedParty(&agent1, 0, GeneratePokemon(PkmnBulbasaur))
			b := NewBattle()
			b.AddParty(party)
			Expect(func() {
				b.getPokemon(1, 5)
			}).To(Panic())
			Expect(func() {
				b.getPokemon(0, 5)
			}).To(Panic())
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
		charmander = GeneratePokemon(PkmnCharmander, WithMoves(GetMove(MovePound)))
		party1 = NewOccupiedParty(&agent1, 0, charmander)
		squirtle = GeneratePokemon(PkmnSquirtle, WithMoves(GetMove(MovePound)))
		party2 = NewOccupiedParty(&agent2, 1, squirtle)
		battle = NewBattle()
		battle.AddParty(party1, party2)
		battle.rng = &SimpleRNG
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
			pound := GetMove(MovePound)
			Expect(t).To(HaveTransaction(DamageTransaction{
				User: charmander,
				Target: target{
					Pokemon:   *squirtle,
					party:     1,
					partySlot: 0,
					Team:      1,
				},
				Move:   pound,
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
				Move:   pound,
				Damage: 3,
			}))
		})

		It("should cause Pokemon to have reduced HP", func() {
			Expect(battle.Start()).To(Succeed())
			battle.SimulateRound()
			Expect(charmander.CurrentHP < charmander.HP()).To(BeTrue())
			Expect(squirtle.CurrentHP < squirtle.HP()).To(BeTrue())
		})
	})

	Context("when dealing damage to a Pokemon", func() {
		It("should account for same-type attack bonus", func() {
			// TODO: remove when elemental type added
			charmander.Type = TypeFire
			charmander.Moves[0] = GetMove(MoveEmber)
			Expect(battle.Start()).To(Succeed())
			battle.SimulateRound()
			Expect(squirtle.CurrentHP).To(BeEquivalentTo(6))
			squirtle.CurrentHP = 100
			adaptability := Ability{ID: 91}
			charmander.Ability = &adaptability
			battle.SimulateRound()
			Expect(squirtle.CurrentHP).To(BeEquivalentTo(93))
		})

		It("should account for critical hits", func() {
			battle.rng = &AlwaysRNG
			Expect(battle.Start()).To(Succeed())
			battle.SimulateRound()
			Expect(squirtle.CurrentHP).To(BeEquivalentTo(4))
		})

		It("should miss moves randomly based on accuracy/evasion", func() {
			battle.rng = &NeverRNG
			Expect(battle.Start()).To(Succeed())
			t, _ := battle.SimulateRound()
			Expect(t).To(HaveTransaction(EvadeTransaction{
				User: charmander,
			}))
			battle.rng = &SimpleRNG
			t, _ = battle.SimulateRound()
			Expect(t).ToNot(HaveTransaction(EvadeTransaction{
				User: charmander,
			}))
		})
	})

	Context("when certain moves are used in battle", func() {
		It("should change a Pokemon's stat modifiers", func() {
			charmander.Moves[0] = GetMove(MoveHowl)
			Expect(battle.Start()).To(Succeed())
			t, _ := battle.SimulateRound()
			Expect(t).To(HaveTransaction(ModifyStatTransaction{
				Target: charmander,
				Stat:   StatAtk,
				Stages: +1,
			}))
			// Bound by min/max stat modifier
			charmander.StatModifiers[StatAtk] = MaxStatModifier
			t, _ = battle.SimulateRound()
			Expect(t).To(HaveTransaction(ModifyStatTransaction{
				Target: charmander,
				Stat:   StatAtk,
				Stages: +1,
			}))
			Expect(charmander.StatModifiers[StatAtk]).To(BeEquivalentTo(MaxStatModifier))
		})
	})
})

var _ = Describe("Using items in battle", func() {
	agent := Agent(new(healAgent))
	var (
		pkmn   *Pokemon
		party  *party
		battle *Battle
	)

	BeforeEach(func() {
		pkmn = GeneratePokemon(PkmnVenusaur, WithLevel(50))
		pkmn.CurrentHP = 10
		party = NewOccupiedParty(&agent, 0, pkmn)
		battle = NewBattle()
		battle.AddParty(party)
	})

	Context("when the battle processes item turns", func() {
		It("should create ItemTransaction(s) properly", func() {
			Expect(battle.Start()).To(Succeed())
			t, _ := battle.SimulateRound()
			potion := GetItem(ItemPotion)
			Expect(t).To(HaveTransaction(
				ItemTransaction{
					Target: pkmn,
					Item:   &potion,
					Move:   nil,
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
			GeneratePokemon(PkmnCharmander),
			GeneratePokemon(PkmnSquirtle),
			GeneratePokemon(PkmnMetapod),
		)
		party2 = NewOccupiedParty(&agent2, 1, GeneratePokemon(PkmnBeedrill))
		battle = NewBattle()
		battle.AddParty(party1, party2)
	})

	Context("when getting Pokemon by party/slot", func() {
		It("should get the Pokemon the user expects", func() {
			pkmn := battle.getPokemon(0, 1)
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
		It("should have a priority of 0 for FightTurn", func() {
			turn := FightTurn{}
			Expect(turn.Priority()).To(Equal(0))
		})

		It("should have a priority of 1 for ItemTurn", func() {
			turn := ItemTurn{}
			Expect(turn.Priority()).To(Equal(1))
		})

		It("should order turns properly based on priority", func() {
			a2 := Agent(new(healAgent))
			bulbasaur := GeneratePokemon(PkmnBulbasaur, WithMoves(GetMove(MovePound)))
			charmander := GeneratePokemon(PkmnCharmander, WithMoves(GetMove(MovePound)))
			p1 := NewOccupiedParty(&a1, 0, bulbasaur)
			p2 := NewOccupiedParty(&a2, 1, charmander)
			b := NewBattle()
			b.AddParty(p1, p2)
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
			p1 := GeneratePokemon(PkmnBulbasaur, WithLevel(5), WithMoves(GetMove(MovePound)))
			p1.Stats[StatSpeed] = 100
			party1 := NewOccupiedParty(&a1, 0, p1)
			p2 := GeneratePokemon(PkmnCharmander, WithLevel(5), WithMoves(GetMove(MoveFakeOut)))
			p2.Stats[StatSpeed] = 10
			party2 := NewOccupiedParty(&a2, 1, p2)
			b := NewBattle()
			b.AddParty(party1, party2)
			b.rng = &SimpleRNG
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveLen(2))
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
			pound := GetMove(MovePound)
			charmander := GeneratePokemon(PkmnCharmander, WithMoves(pound))
			ninjask := GeneratePokemon(PkmnNinjask, WithMoves(pound))
			p1 := NewOccupiedParty(&a1, 0, charmander)
			p2 := NewOccupiedParty(&a2, 1, ninjask) // ninjask is faster than charmander
			b := NewBattle()
			b.AddParty(p1, p2)
			b.rng = &SimpleRNG
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
					Move:   pound,
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
					Move:   pound,
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
		b.rng = &SimpleRNG
	})

	Context("when using certain moves/certain abilities cause weather", func() {
		// TODO: https://bulbapedia.bulbagarden.net/wiki/Weather#Causing_weather
		It("should clear fog when using MoveDefog", func() {
			poke1 := GeneratePokemon(PkmnBulbasaur, WithMoves(GetMove(MoveDefog)))
			poke2 := GeneratePokemon(PkmnBulbasaur, WithMoves(GetMove(MovePound)))
			p1.AddPokemon(poke1)
			p2.AddPokemon(poke2)
			b.Weather = WeatherFog
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(WeatherTransaction{
				Weather: WeatherClearSkies,
			}))
		})
	})

	Context("when weather is present, battles are affected", func() {
		ember := GetMove(MoveEmber)
		bubble := GetMove(MoveBubble)
		tackle := GetMove(MoveTackle)
		solarBeam := GetMove(MoveSolarBeam)
		weatherBall := GetMove(MoveWeatherBall)
		moonlight := GetMove(MoveMoonlight)
		It("should affect fire/water attacks during harsh sunlight", func() {
			charmander := GeneratePokemon(PkmnCharmander, WithLevel(100), WithMoves(ember))
			squirtle := GeneratePokemon(PkmnSquirtle, WithLevel(100), WithMoves(bubble))
			p1.AddPokemon(charmander)
			p2.AddPokemon(squirtle)
			b.Weather = WeatherHarshSunlight
			Expect(b.Start()).To(Succeed())
			transactions, _ := b.SimulateRound()
			// Fire boosted
			Expect(transactions).To(HaveTransaction(
				DamageTransaction{
					User: charmander,
					Target: target{
						Pokemon:   *squirtle,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
					Move:   ember,
					Damage: 50,
				},
			))
			// Water weakened
			Expect(transactions).To(HaveTransaction(
				DamageTransaction{
					User: squirtle,
					Target: target{
						Pokemon:   *charmander,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					Move:   bubble,
					Damage: 17,
				},
			))
		})
		It("should affect fire/water attacks during rain", func() {
			charmander := GeneratePokemon(PkmnCharmander, WithLevel(100), WithMoves(ember))
			squirtle := GeneratePokemon(PkmnSquirtle, WithLevel(100), WithMoves(bubble))
			p1.AddPokemon(charmander)
			p2.AddPokemon(squirtle)
			b.Weather = WeatherRain
			Expect(b.Start()).To(Succeed())
			transactions, _ := b.SimulateRound()
			// Fire weakened
			Expect(transactions).To(HaveTransaction(
				DamageTransaction{
					User: charmander,
					Target: target{
						Pokemon:   *squirtle,
						party:     1,
						partySlot: 0,
						Team:      1,
					},
					Move:   ember,
					Damage: 16,
				},
			))
			// Water boosted
			Expect(transactions).To(HaveTransaction(
				DamageTransaction{
					User: squirtle,
					Target: target{
						Pokemon:   *charmander,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					Move:   bubble,
					Damage: 53,
				},
			))
		})
		It("should damage/cause side effects during sandstorm", func() {
			geodude := GeneratePokemon(PkmnGeodude, WithLevel(50), WithMoves(tackle))
			geodude.Type = TypeRock | TypeGround
			bulbasaur := GeneratePokemon(PkmnBulbasaur, WithLevel(50), WithMoves(solarBeam))
			p1.AddPokemon(geodude)
			p2.AddPokemon(bulbasaur)
			b.Weather = WeatherSandstorm
			Expect(b.Start()).To(Succeed())
			transactions, _ := b.SimulateRound()
			Expect(transactions).To(HaveLen(3))
			// Weaken solar beam, SPDEF of rock type
			Expect(transactions).To(HaveTransaction(DamageTransaction{
				User: bulbasaur,
				Target: target{
					Pokemon:   *geodude,
					party:     0,
					partySlot: 0,
					Team:      0,
				},
				Move:   solarBeam,
				Damage: 37,
			}))
			// Damage from sandstorm
			Expect(transactions).To(HaveTransaction(DamageTransaction{
				User: nil,
				Target: target{
					Pokemon:   *bulbasaur,
					party:     1,
					partySlot: 0,
					Team:      1,
				},
				Move:   nil,
				Damage: 6,
			}))
		})
		It("should damage/cause side effects during hail", func() {
			articuno := GeneratePokemon(PkmnArticuno, WithMoves(tackle))
			articuno.Type = TypeIce
			bulbasaur := GeneratePokemon(PkmnBulbasaur, WithMoves(solarBeam))
			p1.AddPokemon(articuno)
			p2.AddPokemon(bulbasaur)
			b.Weather = WeatherHail
			Expect(b.Start()).To(Succeed())
			transactions, _ := b.SimulateRound()
			Expect(transactions).To(HaveLen(3))
			// Weaken solar beam
			Expect(transactions).To(HaveTransaction(DamageTransaction{
				User: bulbasaur,
				Target: target{
					Pokemon:   *articuno,
					party:     0,
					partySlot: 0,
					Team:      0,
				},
				Move:   solarBeam,
				Damage: 4,
			}))
			// Damage from hail
			Expect(transactions).To(HaveTransaction(DamageTransaction{
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
		It("should cause side effects during fog", func() {
			castform := GeneratePokemon(PkmnCastform, WithLevel(50), WithMoves(weatherBall))
			bulbasaur := GeneratePokemon(PkmnBulbasaur, WithLevel(50), WithMoves(solarBeam))
			p1.AddPokemon(castform)
			p2.AddPokemon(bulbasaur)
			b.Weather = WeatherFog
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
				Damage: 26,
			}))
			bulbasaur.Moves[0] = moonlight
			t, _ = b.SimulateRound()
			// Moonlight heals 1/4 max HP
			Expect(t).To(HaveTransaction(DamageTransaction{
				User: castform,
				Target: target{
					Pokemon:   *bulbasaur,
					party:     1,
					partySlot: 0,
					Team:      1,
				},
				Move:   weatherBall,
				Damage: 49,
			}))
			Expect(t).To(HaveTransaction(HealTransaction{
				Target: bulbasaur,
				Amount: 26,
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
		scary_monster := GeneratePokemon(PkmnSquirtle, WithLevel(100), WithMoves(GetMove(MovePound)))
		scary_monster.Stats[StatSpeed] = 1
		p1 = NewOccupiedParty(&a1, 0,
			GeneratePokemon(PkmnCharmander, WithMoves(GetMove(MovePound))),
			GeneratePokemon(PkmnTurtwig, WithMoves(GetMove(MovePound))),
		)
		p2 = NewOccupiedParty(&a2, 1, scary_monster)
		b = NewBattle()
		b.AddParty(p1, p2)
		b.rng = &SimpleRNG
	})

	Context("after a Pokemon faints in battle", func() {
		It("should switch to the next available Pokemon", func() {
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveLen(5))
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
			pkmn1 := GeneratePokemon(PkmnCharmander, WithLevel(3), WithMoves(GetMove(MovePound)))
			pkmn2 := GeneratePokemon(PkmnSquirtle, WithLevel(10), WithMoves(GetMove(MovePound)))
			pkmn3 := GeneratePokemon(PkmnTurtwig, WithLevel(3), WithMoves(GetMove(MovePound)))
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
			Expect(t).To(HaveLen(4), "Expected 4 transactions to occur")
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
			dies := GeneratePokemon(PkmnBulbasaur, WithLevel(1), WithMoves(GetMove(MovePound)))
			dies.Friendship = 100
			p1 := NewOccupiedParty(&a1, 0, dies)
			p2 := NewOccupiedParty(&a2, 1, GeneratePokemon(PkmnCharmander, WithLevel(25), WithMoves(GetMove(MovePound))))
			b = NewBattle()
			b.AddParty(p1, p2)
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(dies.Friendship).To(Equal(99))
		})

		It("should lose 5 or 10 friendship when fainting", func() {
			dies := GeneratePokemon(PkmnBulbasaur, WithLevel(1), WithMoves(GetMove(MovePound)))
			dies.Friendship = 100
			dies2 := GeneratePokemon(PkmnBulbasaur, WithLevel(1), WithMoves(GetMove(MovePound)))
			dies2.Friendship = 200
			p1 := NewOccupiedParty(&a1, 0, dies, dies2)
			p2 := NewOccupiedParty(&a2, 1, GeneratePokemon(PkmnCharmander, WithLevel(100), WithMoves(GetMove(MovePound))))
			b = NewBattle()
			b.AddParty(p1, p2)
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(dies.Friendship).To(Equal(95))
			b.SimulateRound()
			Expect(dies2.Friendship).To(Equal(190))
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
		pkmn1 = GeneratePokemon(PkmnCharmander, WithLevel(3), WithMoves(GetMove(MovePound)))
		pkmn1.CurrentHP = 1
		party1.AddPokemon(pkmn1)
		party2 := NewParty(&a2, 1)
		pkmn2 = GeneratePokemon(PkmnSquirtle, WithLevel(10), WithMoves(GetMove(MovePound)))
		pkmn2.Stats[StatSpeed] = 255
		party2.AddPokemon(pkmn2)
		b = NewBattle()
		b.AddParty(party1, party2)
	})

	Context("when all Pokemon faint on one team", func() {
		It("should end the battle", func() {
			Expect(b.Start()).To(Succeed())
			t, ended := b.SimulateRound()
			Expect(ended).To(BeTrue(), "Expected SimulateRound to indicate that the battle has ended, but it did not.")
			Expect(t).To(HaveLen(4))
			Expect(t).To(HaveTransactionsInOrder(
				FaintTransaction{
					Target: target{
						Pokemon:   *pkmn1,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
				},
				EndBattleTransaction{},
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

	Context("when using certain moves in battle causes status effects", func() {
		It("should inflict paralysis from MoveStunSpore", func() {
			p1 := GeneratePokemon(PkmnBulbasaur, WithMoves(GetMove(MoveSplash)))
			p2 := GeneratePokemon(PkmnBulbasaur, WithMoves(GetMove(MoveStunSpore)))
			party1 := NewOccupiedParty(&a1, 0, p1)
			party2 := NewOccupiedParty(&a2, 1, p2)
			b := NewBattle()
			b.AddParty(party1, party2)
			b.rng = &AlwaysRNG
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				InflictStatusTransaction{
					Target:       p1,
					StatusEffect: StatusParalyze,
				},
			))
		})
	})

	Context("when a Pokemon has a status effect, it affects the Pokemon in battle", func() {
		It("should inflict burn and poison damage", func() {
			p1 := GeneratePokemon(PkmnBulbasaur, WithMoves(GetMove(MovePound)))
			p1.StatusEffects = StatusPoison
			party1 := NewOccupiedParty(&a1, 0, p1)
			p2 := GeneratePokemon(PkmnIvysaur, WithMoves(GetMove(MovePound)))
			p2.StatusEffects = StatusBurn
			party2 := NewOccupiedParty(&a2, 1, p2)
			b := NewBattle()
			b.AddParty(party1, party2)
			b.SetSeed(1337)
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveLen(4), "Expected only 4 transactions to occur in a round")
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{
					Pokemon:   *p1,
					party:     0,
					partySlot: 0,
					Team:      0,
				},
				Damage:       1,
				StatusEffect: StatusPoison,
			}))
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{
					Pokemon:   *p2,
					party:     1,
					partySlot: 0,
					Team:      1,
				},
				Damage:       1,
				StatusEffect: StatusBurn,
			}))
		})

		It("should inflict badly poisoned damage", func() {
			p1 := GeneratePokemon(PkmnBulbasaur, WithLevel(100), WithMoves(GetMove(MovePound)))
			p1.StatusEffects = StatusBadlyPoison
			party1 := NewOccupiedParty(&a1, 0, p1)
			p2 := GeneratePokemon(PkmnIvysaur, WithLevel(100), WithMoves(GetMove(MovePound)))
			party2 := NewOccupiedParty(&a2, 1, p2)
			b := NewBattle()
			b.AddParty(party1, party2)
			b.rng = &AlwaysRNG
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveLen(3), "Expected only 3 transactions to occur in a round")
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{
					Pokemon:   *p1,
					party:     0,
					partySlot: 0,
					Team:      0,
				},
				Damage:       12,
				StatusEffect: StatusBadlyPoison,
			}))
		})

		It("should immobilize paralyzed Pokemon", func() {
			p1 := GeneratePokemon(PkmnBulbasaur, WithLevel(8), WithMoves(GetMove(MovePound)))
			p1.StatusEffects = StatusParalyze
			party1 := NewOccupiedParty(&a1, 0, p1)
			p2 := GeneratePokemon(PkmnCharmander, WithLevel(4), WithMoves(GetMove(MovePound)))
			party2 := NewOccupiedParty(&a2, 1, p2)
			b := NewBattle()
			b.AddParty(party1, party2)
			b.rng = &AlwaysRNG
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				ImmobilizeTransaction{
					Target: target{
						Pokemon:   *p1,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					StatusEffect: StatusParalyze,
				},
			))
		})

		It("should immobilize frozen Pokemon", func() {
			p1 := GeneratePokemon(PkmnBulbasaur, WithLevel(8), WithMoves(GetMove(MovePound)))
			p1.StatusEffects = StatusFreeze
			party1 := NewOccupiedParty(&a1, 0, p1)
			p2 := GeneratePokemon(PkmnCharmander, WithLevel(4), WithMoves(GetMove(MovePound)))
			party2 := NewOccupiedParty(&a2, 1, p2)
			b := NewBattle()
			b.AddParty(party1, party2)
			b.rng = &AlwaysRNG
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				ImmobilizeTransaction{
					Target: target{
						Pokemon:   *p1,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					StatusEffect: StatusFreeze,
				},
			))
		})

		It("should immobilize sleeping Pokemon", func() {
			p1 := GeneratePokemon(PkmnBulbasaur, WithLevel(8), WithMoves(GetMove(MovePound)))
			p1.StatusEffects = StatusSleep
			party1 := NewOccupiedParty(&a1, 0, p1)
			p2 := GeneratePokemon(PkmnCharmander, WithLevel(4), WithMoves(GetMove(MovePound)))
			party2 := NewOccupiedParty(&a2, 1, p2)
			b := NewBattle()
			b.AddParty(party1, party2)
			b.SetSeed(1337)
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(
				ImmobilizeTransaction{
					Target: target{
						Pokemon:   *p1,
						party:     0,
						partySlot: 0,
						Team:      0,
					},
					StatusEffect: StatusSleep,
				},
			))
		})

		It("should cure paralysis", func() {
			p1 := GeneratePokemon(PkmnBulbasaur, WithLevel(8), WithMoves(GetMove(MovePound)))
			p1.StatusEffects = StatusParalyze
			party1 := NewOccupiedParty(&a1, 0, p1)
			p2 := GeneratePokemon(PkmnCharmander, WithLevel(4), WithMoves(GetMove(MovePound)))
			party2 := NewOccupiedParty(&a2, 1, p2)
			b := NewBattle()
			b.AddParty(party1, party2)
			Expect(b.Start()).To(Succeed())
			b.QueueTransaction(CureStatusTransaction{
				Target: target{
					Pokemon:   *p1,
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
						Pokemon:   *p1,
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
