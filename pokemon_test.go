package pokemonbattlelib

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pokemon generation", func() {
	It("generates a Pokemon given just a dex number", func() {
		p := GeneratePokemon(PkmnWartortle)
		Expect(int(p.NatDex)).To(Equal(8))
	})

	It("gets the name of a generated Pokemon", func() {
		p := GeneratePokemon(PkmnBulbasaur)
		Expect(p.GetName()).To(Equal("Bulbasaur"))
	})

	It("generates a Pokemon with a given level", func() {
		p := GeneratePokemon(PkmnPiplup, WithLevel(5))
		Expect(p.Level).To(BeEquivalentTo(5))
	})

	It("generates a Pokemon with a given total experience", func() {
		p := GeneratePokemon(PkmnPiplup, WithTotalExp(135))
		Expect(p.Level).To(BeEquivalentTo(5))
		Expect(p.TotalExperience).To(BeEquivalentTo(135))
	})

	It("generates a Pokemon with a given set of IVs", func() {
		p := GeneratePokemon(PkmnPiplup, WithIVs([6]uint8{31, 31, 31, 31, 31, 31}))
		Expect(p.IVs).To(BeEquivalentTo([6]uint8{31, 31, 31, 31, 31, 31}))
	})

	It("generates a Pokemon with a given set of EVs", func() {
		p := GeneratePokemon(PkmnPiplup, WithEVs([6]uint8{0, 252, 6, 0, 0, 252}))
		Expect(p.EVs).To(BeEquivalentTo([6]uint8{0, 252, 6, 0, 0, 252}))
	})

	It("generates a Pokemon with a given Nature", func() {
		p := GeneratePokemon(PkmnPiplup, WithNature(NatureAdamant))
		Expect(p.Nature).To(Equal(NatureAdamant))
	})

	It("generates a Pokemon with a given Ability", func() {
		p := GeneratePokemon(PkmnPiplup, WithAbility(AbilityBadDreams))
		Expect(p.Ability).To(Equal(AbilityBadDreams))
	})

	It("generates a Pokemon with a given moveset", func() {
		pound := GetMove(MovePound)
		pursuit := GetMove(MovePursuit)
		pkmn := GeneratePokemon(PkmnPiplup, WithMoves(pound, pursuit))
		Expect(pkmn.Moves).To(BeEquivalentTo([MaxMoves]*Move{pound, pursuit, nil, nil}))
	})

	It("creates Pokemon with accurate stats reflecting its given values", func() {
		// see: https://bulbapedia.bulbagarden.net/wiki/Stat, scroll down to 'Example'
		p := GeneratePokemon(
			PkmnGarchomp,
			WithLevel(78),
			WithIVs([6]uint8{24, 12, 30, 16, 23, 5}),
			WithEVs([6]uint8{74, 190, 91, 48, 84, 23}),
		)
		Expect(p.Stats).To(BeEquivalentTo([6]uint{289, 253, 193, 151, 171, 171}))
	})

	It("panics when computing stats of illegal Pokemon", func() {
		p := GeneratePokemon(PkmnBulbasaur)
		p.Level = 105
		Expect(func() { p.computeStats() }).To(Panic())
	})

	DescribeTable("Panic when given invalid options",
		func(opts ...GeneratePokemonOption) {
			Expect(func() { GeneratePokemon(PkmnStarly, opts...) }).To(Panic())
		},
		Entry("level too high", WithLevel(MaxLevel+1)),
		Entry("level too low", WithLevel(MinLevel-1)),
		Entry("IVs too high", WithIVs([6]uint8{32, 32, 32, 32, 32, 32})),
		Entry("EVs too high", WithEVs([6]uint8{255, 255, 255, 255, 255, 255})),
		Entry("too many moves", WithMoves(GetMove(MoveFly), GetMove(MoveAerialAce), GetMove(MoveRoost), GetMove(MovePeck), GetMove(MovePound))),
	)

	Describe("Validation", func() {
		It("succeeds when the pokemon has moves", func() {
			p := GeneratePokemon(PkmnPikachu, WithMoves(GetMove(MoveThunder)))
			Expect(p.Validate(PkmnRuleSetDefault)).To(Succeed())
		})

		It("fails when the pokemon has no moves", func() {
			p := GeneratePokemon(PkmnPikachu)
			Expect(p.Validate(PkmnRuleSetDefault)).To(MatchError(ErrorValidationMissingMoves))
		})

		It("fails when the pokemon has invalid level", func() {
			p := GeneratePokemon(PkmnPikachu, WithMoves(GetMove(MoveThunder)))
			p.Ability = 0
			Expect(p.Validate(PkmnRuleSetDefault)).To(MatchError(ErrorValidationMissingAbility))
		})

		It("fails when the pokemon has invalid level", func() {
			p := GeneratePokemon(PkmnPikachu, WithMoves(GetMove(MoveThunder)))
			p.Level = 0
			Expect(p.Validate(PkmnRuleSetDefault)).To(MatchError(ErrorValidationInvalidLevel))
		})

		It("fails when the pokemon has invalid IVs", func() {
			p := GeneratePokemon(PkmnPikachu, WithMoves(GetMove(MoveThunder)))
			p.IVs[StatHP] = 255
			Expect(p.Validate(PkmnRuleSetDefault)).To(MatchError(ErrorValidationInvalidIvs))
		})

		It("fails when the pokemon has invalid EVs", func() {
			p := GeneratePokemon(PkmnPikachu, WithMoves(GetMove(MoveThunder)))
			p.EVs[StatHP] = 255
			Expect(p.Validate(PkmnRuleSetDefault)).To(MatchError(ErrorValidationInvalidEvs))
		})
	})
})

var _ = Describe("Test leveling methods", func() {
	It("panics when leveling beyond the max level", func() {
		pkmn := GeneratePokemon(PkmnCharizard, WithLevel(MaxLevel))
		Expect(func() { pkmn.GainLevels(1) }).To(Panic())
	})

	It("panics when trying to level down", func() {
		pkmn := GeneratePokemon(PkmnPiplup, WithLevel(5))
		Expect(func() { pkmn.GainLevels(-1) }).To(Panic())
	})

	It("panics when trying to lose experience", func() {
		pkmn := GeneratePokemon(PkmnPiplup, WithLevel(5))
		Expect(func() { pkmn.GainExperience(-135) }).To(Panic())
	})

	It("prevents a Pokemon from gaining experience beyond the max", func() {
		pkmn := GeneratePokemon(PkmnPiplup, WithLevel(MaxLevel))
		pkmn.GainExperience(100000000000)
		Expect(int(pkmn.Level)).To(Equal(MaxLevel))
	})
})

var _ = Describe("Pokemon string representation", func() {
	var (
		pkmn *Pokemon
		want string
	)

	It("should show the correct string for Pokemon", func() {
		pkmn = GeneratePokemon(PkmnBulbasaur, WithLevel(5))
		want = "Bulbasaur"
		Expect(fmt.Sprintf("%s", pkmn)).To(Equal(want))
	})
})

var _ = Describe("Pokemon Data", func() {
	It("should work for Bulbasaur", func() {
		pkmn := GeneratePokemon(PkmnBulbasaur)
		Expect(pkmn.Data().Name).To(Equal("Bulbasaur"))
	})

	It("should work for Arceus", func() {
		pkmn := GeneratePokemon(PkmnArceus)
		Expect(pkmn.Data().Name).To(Equal("Arceus"))
	})
})
