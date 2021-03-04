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
			party1 := NewOccupiedParty(&agent1, 0, GeneratePokemon(PkmnCharmander))
			party2 := NewOccupiedParty(&agent2, 1, GeneratePokemon(PkmnSquirtle))
			b := NewBattle()
			b.AddParty(party1, party2)
			b.SetSeed(849823)
		})

		It("panics when adding too many Pokemon to a party", func() {
			party := NewParty(&agent1, 0)
			for i := 0; i < MaxPartySize; i += 1 {
				party.AddPokemon(GeneratePokemon(PkmnBulbasaur))
			}
			Expect(func() {
				party.AddPokemon(GeneratePokemon(PkmnBulbasaur))
			}).To(Panic())
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
	var (
		agent1     Agent
		agent2     Agent
		party1     *party
		party2     *party
		battle     *Battle
		charmander *Pokemon
		squirtle   *Pokemon
	)

	BeforeEach(func() {
		agent1 = Agent(dumbAgent{})
		agent2 = Agent(dumbAgent{})
		charmander = GeneratePokemon(4, WithMoves(GetMove(MovePound)))
		party1 = NewOccupiedParty(&agent1, 0, charmander)
		squirtle = GeneratePokemon(7, WithMoves(GetMove(MovePound)))
		party2 = NewOccupiedParty(&agent2, 1, squirtle)
		battle = NewBattle()
		battle.AddParty(party1, party2)
		battle.rng = &SimpleRNG
	})

	It("starts without error", func() {
		err := battle.Start()
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("panics if battle is not in progress", func() {
		Expect(func() {
			battle.SimulateRound()
		}).To(Panic())
	})

	Context("when simulating a round between two agents", func() {
		It("should return two transactions", func() {
			Expect(battle.Start()).To(Succeed())
			transactions, _ := battle.SimulateRound()
			Expect(transactions).To(HaveLen(2))
		})
		It("should create 2 damage transactions", func() {
			Expect(battle.Start()).To(Succeed())
			transactions, _ := battle.SimulateRound()
			pound := GetMove(MovePound)
			Expect(transactions).To(HaveTransaction(DamageTransaction{
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
			Expect(transactions).To(HaveTransaction(DamageTransaction{
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
			Expect(charmander.CurrentHP < charmander.Stats[StatHP]).To(BeTrue())
			Expect(squirtle.CurrentHP < squirtle.Stats[StatHP]).To(BeTrue())
		})
	})

	Context("should deal the correct amount of damage", func() {
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
	})

	Context("should account for accuracy/evasion", func() {
		It("should miss moves randomly", func() {
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

	Context("Interacts with StatModifiers in battle", func() {
		It("should increase stat modifiers from certain moves", func() {
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
		battle.SetSeed(1337)
	})

	Context("an Agent uses an ItemTurn to use a Potion on a Pokemon", func() {
		It("should run without error", func() {
			err := battle.Start()
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("should log ItemTurns correctly", func() {
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
		party = NewOccupiedParty(&agent, 0, GeneratePokemon(PkmnSquirtle), GeneratePokemon(PkmnBlastoise))
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

	It("should panic when Pokemon should not change active state", func() {
		Expect(func() {
			party.SetInactive(0)
		}).To(Panic())
		party.SetActive(0)
		Expect(func() {
			party.SetActive(0)
		}).To(Panic())
	})

	It("should panic when Pokemon does not exist", func() {
		Expect(func() {
			party.IsActivePokemon(7)
		}).To(Panic())
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
		party1 = NewOccupiedParty(&agent1, 0,
			GeneratePokemon(PkmnCharmander),
			GeneratePokemon(PkmnSquirtle),
			GeneratePokemon(PkmnMetapod),
		)
		party2 = NewOccupiedParty(&agent2, 1, GeneratePokemon(PkmnBeedrill))
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
			Expect(battle.Start()).To(Succeed())
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
			Expect(battle.Start()).To(Succeed())
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
		p1 := GeneratePokemon(1, WithLevel(5), WithMoves(GetMove(MovePound)))
		p1.Stats[StatSpeed] = 100
		party1 := NewOccupiedParty(&a1, 0, p1)
		p2 := GeneratePokemon(4, WithLevel(5), WithMoves(GetMove(MoveFakeOut)))
		p2.Stats[StatSpeed] = 10
		party2 := NewOccupiedParty(&a2, 1, p2)
		b := NewBattle()
		b.AddParty(party1, party2)
		b.SetSeed(1337)
		Expect(b.Start()).To(Succeed())
		transactions, _ := b.SimulateRound()
		Expect(transactions).To(HaveLen(2))
		Expect(transactions).To(HaveTransactionsInOrder(
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
})

var _ = Describe("Pokemon speed", func() {
	var (
		agent1     Agent
		agent2     Agent
		party1     *party
		party2     *party
		pound      *Move
		battle     *Battle
		charmander *Pokemon
		ninjask    *Pokemon
	)

	BeforeEach(func() {
		agent1 = Agent(dumbAgent{})
		agent2 = Agent(dumbAgent{})
		pound = GetMove(MovePound)
		charmander = GeneratePokemon(4, WithMoves(pound))
		ninjask = GeneratePokemon(291, WithMoves(pound))
		party1 = NewOccupiedParty(&agent1, 0, charmander)
		party2 = NewOccupiedParty(&agent2, 1, ninjask) // ninjask is faster than charmander
		battle = NewBattle()
		battle.AddParty(party1, party2)
		battle.SetSeed(1337)
	})

	Context("A faster Pokemon is fighting a slower one", func() {
		It("should not error when starting", func() {
			err := battle.Start()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should create two transactions when simulating a round", func() {
			Expect(battle.Start()).To(Succeed())
			transactions, _ := battle.SimulateRound()
			Expect(transactions).To(HaveLen(2))
		})

		Specify("faster Pokemon should go first", func() {
			Expect(battle.Start()).To(Succeed())
			transactions, _ := battle.SimulateRound()
			Expect(transactions).To(HaveTransactionsInOrder(
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

var _ = Describe("Weather", func() {
	var (
		a1 Agent
		a2 Agent
		p1 *party
		p2 *party
		b  *Battle
	)

	BeforeEach(func() {
		a1 = Agent(dumbAgent{})
		a2 = Agent(dumbAgent{})
		p1 = NewParty(&a1, 0)
		p2 = NewParty(&a2, 1)
		b = NewBattle()
		b.AddParty(p1, p2)
		b.SetSeed(1337)
	})

	Context("Weather should be caused by moves/abilities", func() {
		// TODO: https://bulbapedia.bulbagarden.net/wiki/Weather#Causing_weather
		It("should change to clear skies from defog", func() {
			poke1 := GeneratePokemon(1, WithMoves(GetMove(MoveDefog)))
			poke2 := GeneratePokemon(1, WithMoves(GetMove(MovePound)))
			p1.AddPokemon(poke1)
			p2.AddPokemon(poke2)
			b.Weather = WeatherFog
			Expect(b.Start()).To(Succeed())
			transactions, _ := b.SimulateRound()
			Expect(transactions).To(HaveTransaction(WeatherTransaction{
				Weather: WeatherClearSkies,
			}))
		})
	})

	Context("Weather has effects in battle", func() {
		ember := GetMove(MoveEmber)
		bubble := GetMove(MoveBubble)
		tackle := GetMove(MoveTackle)
		solarBeam := GetMove(MoveSolarBeam)
		weatherBall := GetMove(MoveWeatherBall)
		moonlight := GetMove(MoveMoonlight)
		It("should affect fire/water attacks during harsh sunlight", func() {
			charmander := GeneratePokemon(4, WithLevel(100), WithMoves(ember))
			squirtle := GeneratePokemon(7, WithLevel(100), WithMoves(bubble))
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
			charmander := GeneratePokemon(4, WithLevel(100), WithMoves(ember))
			squirtle := GeneratePokemon(7, WithLevel(100), WithMoves(bubble))
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
			geodude := GeneratePokemon(74, WithLevel(50), WithMoves(tackle))
			geodude.Type = TypeRock | TypeGround
			bulbasaur := GeneratePokemon(1, WithLevel(50), WithMoves(solarBeam))
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
			articuno := GeneratePokemon(144, WithMoves(tackle))
			articuno.Type = TypeIce
			bulbasaur := GeneratePokemon(1, WithMoves(solarBeam))
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
			castform := GeneratePokemon(351, WithLevel(50), WithMoves(weatherBall))
			bulbasaur := GeneratePokemon(1, WithLevel(50), WithMoves(solarBeam))
			p1.AddPokemon(castform)
			p2.AddPokemon(bulbasaur)
			b.Weather = WeatherFog
			b.SetSeed(777777)
			Expect(b.Start()).To(Succeed())
			transactions, _ := b.SimulateRound()
			// Accuracy decreases from fog
			Expect(transactions).To(HaveTransaction(EvadeTransaction{
				User: castform,
			}))
			// Solar beam weakened
			Expect(transactions).To(HaveTransaction(DamageTransaction{
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
			transactions, _ = b.SimulateRound()
			Expect(transactions).To(HaveTransaction(EvadeTransaction{
				User: castform,
			}))
			// Moonlight heals 1/4 max HP
			Expect(transactions).To(HaveTransaction(HealTransaction{
				Target: bulbasaur,
				Amount: 26,
			}))
			transactions, _ = b.SimulateRound()
			Expect(transactions).To(HaveTransaction(DamageTransaction{
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
	)

	BeforeEach(func() {
		agent1 = Agent(dumbAgent{})
		agent2 = Agent(dumbAgent{})
		scary_monster := GeneratePokemon(7, WithLevel(100), WithMoves(GetMove(MovePound)))
		scary_monster.Stats[StatSpeed] = 1
		party1 = NewOccupiedParty(&agent1, 0,
			GeneratePokemon(4, WithMoves(GetMove(MovePound))),
			GeneratePokemon(387, WithMoves(GetMove(MovePound))),
		)
		party2 = NewOccupiedParty(&agent2, 1, scary_monster)
		battle = NewBattle()
		battle.AddParty(party1, party2)
		battle.SetSeed(1337)
	})

	Context("Switch is forced after a Pokemon faints", func() {
		It("should start without error", func() {
			err := battle.Start()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("causes 5 transactions to occur", func() {
			Expect(battle.Start()).To(Succeed())
			transactions, _ := battle.SimulateRound()
			Expect(transactions).To(HaveLen(5))
		})
		It("should log all transactions as expected", func() {
			Expect(battle.Start()).To(Succeed())
			transactions, _ := battle.SimulateRound()
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
			Expect(transactions).To(HaveTransactionsInOrder(
				FaintTransaction{
					Target: target{
						Pokemon:   *battle.parties[0].pokemon[0],
						party:     0,
						partySlot: 0,
						Team:      0,
					},
				}, SendOutTransaction{
					Target: target{
						Pokemon:   *battle.parties[0].pokemon[1],
						party:     0,
						partySlot: 1,
						Team:      0,
					},
				}))
		})
	})

	Context("Fainting causes friendship to be lost", func() {
		It("should lose 1 friendship when fainting", func() {
			dies := GeneratePokemon(1, WithLevel(1), WithMoves(GetMove(MovePound)))
			dies.Friendship = 100
			p1 := NewOccupiedParty(&agent1, 0, dies)
			p2 := NewOccupiedParty(&agent2, 1, GeneratePokemon(4, WithLevel(25), WithMoves(GetMove(MovePound))))
			battle = NewBattle()
			battle.AddParty(p1, p2)
			Expect(battle.Start()).To(Succeed())
			battle.SimulateRound()
			Expect(dies.Friendship).To(Equal(99))
		})
		It("should lose 5 or 10 friendship when fainting", func() {
			dies := GeneratePokemon(1, WithLevel(1), WithMoves(GetMove(MovePound)))
			dies.Friendship = 100
			dies2 := GeneratePokemon(1, WithLevel(1), WithMoves(GetMove(MovePound)))
			dies2.Friendship = 200
			p1 := NewOccupiedParty(&agent1, 0, dies, dies2)
			p2 := NewOccupiedParty(&agent2, 1, GeneratePokemon(4, WithLevel(100), WithMoves(GetMove(MovePound))))
			battle = NewBattle()
			battle.AddParty(p1, p2)
			Expect(battle.Start()).To(Succeed())
			battle.SimulateRound()
			Expect(dies.Friendship).To(Equal(95))
			battle.SimulateRound()
			Expect(dies2.Friendship).To(Equal(190))
		})
	})

	Specify("Dead pokemon should not take turns", func() {
		a1 := Agent(dumbAgent{})
		a2 := Agent(dumbAgent{})
		party1 := NewParty(&a1, 0)
		pkmn1 := GeneratePokemon(4, WithLevel(3), WithMoves(GetMove(MovePound)))
		pkmn1.CurrentHP = 1
		party1.AddPokemon(pkmn1)
		party2 := NewParty(&a2, 1)
		pkmn2 := GeneratePokemon(7, WithLevel(10), WithMoves(GetMove(MovePound)))
		pkmn2.Stats[StatSpeed] = 255
		party2.AddPokemon(pkmn2)
		b := NewBattle()
		b.AddParty(party1, party2)
		b.SetSeed(1337)
		Expect(b.Start()).To(Succeed())
		transactions, ended := b.SimulateRound()
		Expect(ended).To(BeTrue(), "Expected SimulateRound to indicate that the battle has ended, but it did not.")
		Expect(transactions).To(HaveLen(4))
		Expect(transactions).To(HaveTransactionsInOrder(
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
		Expect(transactions).ToNot(HaveTransaction(
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

	Specify("Dead pokemon should not take turns", func() {
		a1 := Agent(dumbAgent{})
		a2 := Agent(dumbAgent{})
		party1 := NewParty(&a1, 0)
		pkmn1 := GeneratePokemon(4, WithLevel(3), WithMoves(GetMove(MovePound)))
		pkmn2 := GeneratePokemon(7, WithLevel(10), WithMoves(GetMove(MovePound)))
		pkmn3 := GeneratePokemon(387, WithLevel(3), WithMoves(GetMove(MovePound)))
		pkmn1.CurrentHP = 1
		party1.AddPokemon(pkmn1, pkmn3)
		party2 := NewParty(&a2, 1)
		pkmn2.Stats[StatSpeed] = 255
		party2.AddPokemon(pkmn2)
		b := NewBattle()
		b.AddParty(party1, party2)
		b.SetSeed(1337)
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
})

var _ = Describe("Ending a battle", func() {
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
		low_health_pkmn := GeneratePokemon(4, WithMoves(GetMove(MovePound)))
		low_health_pkmn.CurrentHP = 1
		party1 = NewOccupiedParty(&agent1, 0, low_health_pkmn)
		party2 = NewOccupiedParty(&agent2, 1, GeneratePokemon(7, WithMoves(GetMove(MovePound))))
		battle = NewBattle()
		battle.AddParty(party1, party2)
		battle.SetSeed(1337)
	})

	Context("Battle ends by knockout", func() {
		It("should start without error", func() {
			err := battle.Start()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should end", func() {
			Expect(battle.Start()).To(Succeed())
			_, ended := battle.SimulateRound()
			Expect(ended).To(BeTrue())
		})

		It("should have 5 transactions occur", func() {
			Expect(battle.Start()).To(Succeed())
			transactions, _ := battle.SimulateRound()
			Expect(transactions).To(HaveLen(5))
		})

		It("should log all transaction correctly", func() {
			Expect(battle.Start()).To(Succeed())
			transactions, _ := battle.SimulateRound()
			Expect(transactions).To(HaveTransactionsInOrder(
				FaintTransaction{
					Target: target{
						Pokemon:   *battle.parties[0].pokemon[0],
						party:     0,
						partySlot: 0,
						Team:      0,
					},
				},
				EndBattleTransaction{},
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
	var (
		a1 Agent
		a2 Agent
	)

	BeforeEach(func() {
		a1 = Agent(dumbAgent{})
		a2 = Agent(dumbAgent{})
	})

	It("should cause status effects from moves", func() {
		p1 := GeneratePokemon(1, WithMoves(GetMove(MoveSplash)))
		p2 := GeneratePokemon(1, WithMoves(GetMove(MoveStunSpore)))
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

	It("should inflict burn and poison damage", func() {
		p1 := GeneratePokemon(1, WithMoves(GetMove(MovePound)))
		p1.StatusEffects = StatusPoison
		party1 := NewOccupiedParty(&a1, 0, p1)
		p2 := GeneratePokemon(2, WithMoves(GetMove(MovePound)))
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
		p1 := GeneratePokemon(1, WithLevel(100), WithMoves(GetMove(MovePound)))
		p1.StatusEffects = StatusBadlyPoison
		party1 := NewOccupiedParty(&a1, 0, p1)
		p2 := GeneratePokemon(2, WithLevel(100), WithMoves(GetMove(MovePound)))
		party2 := NewOccupiedParty(&a2, 1, p2)
		b := NewBattle()
		b.AddParty(party1, party2)
		b.SetSeed(1337)
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

	Specify("Paralysis", func() {
		p1 := GeneratePokemon(1, WithLevel(8), WithMoves(GetMove(MovePound)))
		p1.StatusEffects = StatusParalyze
		party1 := NewOccupiedParty(&a1, 0, p1)
		p2 := GeneratePokemon(4, WithLevel(4), WithMoves(GetMove(MovePound)))
		party2 := NewOccupiedParty(&a2, 1, p2)
		b := NewBattle()
		b.AddParty(party1, party2)
		b.SetSeed(1)
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

	Specify("Freeze", func() {
		p1 := GeneratePokemon(1, WithLevel(8), WithMoves(GetMove(MovePound)))
		p1.StatusEffects = StatusFreeze
		party1 := NewOccupiedParty(&a1, 0, p1)
		p2 := GeneratePokemon(4, WithLevel(4), WithMoves(GetMove(MovePound)))
		party2 := NewOccupiedParty(&a2, 1, p2)
		b := NewBattle()
		b.AddParty(party1, party2)
		b.SetSeed(34987)
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

	Specify("Sleep", func() {
		p1 := GeneratePokemon(1, WithLevel(8), WithMoves(GetMove(MovePound)))
		p1.StatusEffects = StatusSleep
		party1 := NewOccupiedParty(&a1, 0, p1)
		p2 := GeneratePokemon(4, WithLevel(4), WithMoves(GetMove(MovePound)))
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

	It("Should cure paralysis", func() {
		p1 := GeneratePokemon(1, WithLevel(8), WithMoves(GetMove(MovePound)))
		p1.StatusEffects = StatusParalyze
		party1 := NewOccupiedParty(&a1, 0, p1)
		p2 := GeneratePokemon(4, WithLevel(4), WithMoves(GetMove(MovePound)))
		party2 := NewOccupiedParty(&a2, 1, p2)
		b := NewBattle()
		b.AddParty(party1, party2)
		b.SetSeed(1337)
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
