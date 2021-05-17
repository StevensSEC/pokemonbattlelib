package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Weather", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))

	Context("when using certain moves/certain abilities cause weather", func() {
		// TODO: https://bulbapedia.bulbagarden.net/wiki/Weather#Causing_weather
		It("should clear fog when using MoveDefog", func() {
			b := New1v1Battle(
				PkmnWithMoves(MoveDefog), &a1,
				PkmnNoDamage(), &a2,
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
				PkmnWithMoves(MoveSunnyDay), &a1,
				PkmnNoDamage(), &a2,
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
				PkmnWithMoves(MoveSunnyDay), &a1,
				PkmnNoDamage(), &a2,
			)
			b.rng = SimpleRNG()
			Expect(b.Start()).To(Succeed())
			b.SimulateRound()
			Expect(b.metadata[MetaWeatherTurns]).To(Equal(4))
			Expect(b.Weather).ToNot(Equal(WeatherFog))
			b.getPokemon(target{0, 0}).Moves[0] = GetMove(MoveSplash)
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
				Expect(DamageDealt(t, target{0, 0})).To(Equal(183))
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
				Expect(DamageDealt(t, target{0, 0})).To(Equal(41))
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
				Expect(DamageDealt(t, target{0, 0})).To(Equal(61))
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
				Expect(DamageDealt(t, target{0, 0})).To(Equal(124))
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
				solarbeamDamage := DamageDealt(t, target{1, 0})
				Expect(solarbeamDamage).To(Equal(18))
				Expect(t).ToNot(HaveTransaction(FaintTransaction{}))
				b.QueueTransaction(
					WeatherTransaction{
						Weather: WeatherSandstorm,
						Turns:   5,
					},
					HealTransaction{
						Target: target{0, 0},
						Amount: 100,
					},
					HealTransaction{
						Target: target{1, 0},
						Amount: 100,
					},
				)
				b.ProcessQueue()
				t, _ = b.SimulateRound()
				Expect(DamageDealt(t, target{1, 0})).To(BeNumerically("<", solarbeamDamage))
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
					Target: target{1, 0},
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
				solarbeamDamage := DamageDealt(t, target{1, 0})
				Expect(solarbeamDamage).To(Equal(18))
				Expect(t).ToNot(HaveTransaction(FaintTransaction{}))
				b.QueueTransaction(
					WeatherTransaction{
						Weather: WeatherHail,
						Turns:   5,
					},
					HealTransaction{
						Target: target{0, 0},
						Amount: 100,
					},
					HealTransaction{
						Target: target{1, 0},
						Amount: 100,
					},
				)
				b.ProcessQueue()
				t, _ = b.SimulateRound()
				Expect(DamageDealt(t, target{1, 0})).To(BeNumerically("<", solarbeamDamage))
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
					Target: target{1, 0},
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
			Expect(DamageDealt(t, target{1, 0})).To(Equal(12))
			bulbasaur.Moves[0] = moonlight
			bulbasaur.CurrentHP = bulbasaur.MaxHP()
			t, _ = b.SimulateRound()
			// Moonlight heals 1/4 max HP, weather ball boosted
			Expect(DamageDealt(t, target{0, 0})).To(Equal(21))
			Expect(t).To(HaveTransaction(HealTransaction{
				Target: target{1, 0},
				Amount: 7,
			}))
		})
	})
})
