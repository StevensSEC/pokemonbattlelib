package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pokemon generation", func() {
	It("generates a Pokemon given just a dex number", func() {
		p := GeneratePokemon(8)
		Expect(int(p.NatDex)).To(Equal(8))
	})

	It("gets the name of a generated Pokemon", func() {
		p := GeneratePokemon(1)
		Expect(p.GetName()).To(Equal("Bulbasaur"))
	})

	It("generates a Pokemon with a given level", func() {
		p := GeneratePokemon(393, WithLevel(5))
		want := &Pokemon{
			NatDex:          393, // piplup if you're curious
			Level:           5,
			TotalExperience: 135,
			CurrentHP:       20,
			IVs:             [6]uint8{0, 0, 0, 0, 0, 0},
			EVs:             [6]uint8{0, 0, 0, 0, 0, 0},
			Nature:          GetNature(HARDY),
			Stats:           [6]uint{20, 10, 10, 11, 10, 9},
		}
		Expect(p).To(Equal(want))
	})

	It("generates a Pokemon with a given total experience", func() {
		p := GeneratePokemon(393, WithTotalExp(135))
		want := &Pokemon{
			NatDex:          393,
			Level:           5,
			TotalExperience: 135,
			CurrentHP:       20,
			IVs:             [6]uint8{0, 0, 0, 0, 0, 0},
			EVs:             [6]uint8{0, 0, 0, 0, 0, 0},
			Nature:          GetNature(HARDY),
			Stats:           [6]uint{20, 10, 10, 11, 10, 9},
		}
		Expect(p).To(Equal(want))
	})

	It("generates a Pokemon with a given set of IVs", func() {
		pkmn := GeneratePokemon(393, WithLevel(5), WithIVs([6]uint8{31, 31, 31, 31, 31, 31}))
		want := &Pokemon{
			NatDex:          393,
			Level:           5,
			TotalExperience: 135,
			CurrentHP:       21,
			IVs:             [6]uint8{31, 31, 31, 31, 31, 31},
			EVs:             [6]uint8{0, 0, 0, 0, 0, 0},
			Nature:          GetNature(HARDY),
			Stats:           [6]uint{21, 11, 11, 12, 12, 10},
		}
		Expect(pkmn).To(Equal(want))
	})

	It("generates a Pokemon with a given set of EVs", func() {
		pkmn := GeneratePokemon(393, WithLevel(5), WithEVs([6]uint8{0, 252, 6, 0, 0, 252}))
		want := &Pokemon{
			NatDex:          393,
			Level:           5,
			TotalExperience: 135,
			CurrentHP:       20,
			IVs:             [6]uint8{0, 0, 0, 0, 0, 0},
			EVs:             [6]uint8{0, 252, 6, 0, 0, 252},
			Nature:          GetNature(HARDY),
			Stats:           [6]uint{20, 13, 10, 11, 10, 12},
		}
		Expect(pkmn).To(Equal(want))
	})

	It("generates a Pokemon with a given Nature", func() {
		pkmn := GeneratePokemon(393, WithLevel(5), WithNature(GetNature(ADAMANT)))
		want := &Pokemon{
			NatDex:          393,
			Level:           5,
			TotalExperience: 135,
			CurrentHP:       20,
			IVs:             [6]uint8{0, 0, 0, 0, 0, 0},
			EVs:             [6]uint8{0, 0, 0, 0, 0, 0},
			Nature:          GetNature(ADAMANT),
			Stats:           [6]uint{20, 11, 10, 9, 10, 9},
		}
		Expect(pkmn).To(Equal(want))
	})

	It("generates a Pokemon with a given moveset", func() {
		pound := GetMove(MOVE_POUND)
		pursuit := GetMove(MOVE_PURSUIT)
		pkmn := GeneratePokemon(393, WithMoves(&pound, &pursuit))
		Expect(pkmn.Moves).To(BeEquivalentTo([MAX_MOVES]*Move{&pound, &pursuit, nil, nil}))
	})

	It("creates Pokemon with accurate stats reflecting its given values", func() {
		// see: https://bulbapedia.bulbagarden.net/wiki/Stat, scroll down to 'Example'
		pkmn := GeneratePokemon(
			445,
			WithLevel(78),
			WithIVs([6]uint8{24, 12, 30, 16, 23, 5}),
			WithEVs([6]uint8{74, 190, 91, 48, 84, 23}),
			WithNature(GetNature(ADAMANT)))
		want := &Pokemon{
			NatDex:          445, // garchomp
			Level:           78,
			TotalExperience: 593190,
			CurrentHP:       289,
			IVs:             [6]uint8{24, 12, 30, 16, 23, 5},
			EVs:             [6]uint8{74, 190, 91, 48, 84, 23},
			Nature:          GetNature(ADAMANT),
			Stats:           [6]uint{289, 278, 193, 135, 171, 171},
		}
		Expect(pkmn).To(Equal(want))
	})

	It("panics when computing stats of illegal Pokemon", func() {
		p := GeneratePokemon(1)
		p.Level = 105
		Expect(func() { p.computeStats() }).To(Panic())
	})

	It("panics when trying to create a Pokemon out of level bounds", func() {
		Expect(func() { GeneratePokemon(396, WithLevel(MAX_LEVEL+1)) }).To(Panic())
		Expect(func() { GeneratePokemon(396, WithLevel(MIN_LEVEL-1)) }).To(Panic())
	})

	It("panics when creating a Pokemon with higher than max IVs", func() {
		Expect(func() { GeneratePokemon(396, WithIVs([6]uint8{32, 32, 32, 32, 32, 32})) }).To(Panic())
	})

	It("panics when creating a Pokemon with higher than max EVs", func() {
		Expect(func() { GeneratePokemon(396, WithEVs([6]uint8{255, 255, 255, 255, 255, 255})) }).To(Panic())
	})

	It("panics when creating a Pokemon with more than the maximum allowed moves", func() {
		pound := GetMove(MOVE_POUND)
		Expect(func() { GeneratePokemon(396, WithMoves(&pound, &pound, &pound, &pound, &pound)) }).To(Panic())
	})

})

var _ = Describe("Test leveling methods", func() {
	It("panics when leveling beyond the max level", func() {
		pkmn := GeneratePokemon(6, WithLevel(MAX_LEVEL))
		Expect(func() { pkmn.GainLevels(1) }).To(Panic())
	})

	It("panics when trying to level down", func() {
		pkmn := GeneratePokemon(393, WithLevel(5))
		Expect(func() { pkmn.GainLevels(-1) }).To(Panic())
	})

	It("panics when trying to lose experience", func() {
		pkmn := GeneratePokemon(393, WithLevel(5))
		Expect(func() { pkmn.GainExperience(-135) }).To(Panic())
	})

	It("prevents a Pokemon from gaining experience beyond the max", func() {
		pkmn := GeneratePokemon(493, WithLevel(MAX_LEVEL))
		pkmn.GainExperience(100000000000)
		Expect(int(pkmn.Level)).To(Equal(MAX_LEVEL))
	})
})

var _ = Describe("Stringer interface", func() {
	var (
		pkmn *Pokemon
		want string
	)

	It("prints as expected", func() {
		pkmn = GeneratePokemon(1, WithLevel(5))
		pkmn.Gender = Female
		want = "Bulbasaur♀\tLv5\nHP: 19/19\n"
		Expect(pkmn.String()).To(Equal(want))
	})
})
