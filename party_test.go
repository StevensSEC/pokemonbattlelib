package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Party creation", func() {
	agent := Agent(new(dumbAgent))

	Context("when adding Pokemon to a party", func() {
		It("fails when adding too many Pokemon to a party", func() {
			party := newBattlePartyOld(&agent, 0)
			for i := 0; i < MaxPartySize; i += 1 {
				party.AddPokemon(PkmnWithMoves(MoveTackle))
			}
			Expect(party.AddPokemon(PkmnWithMoves(MoveTackle))).To(MatchError(ErrorPartyFull))
		})

		It("fails when adding invalid pokemon to a party", func() {
			party := newBattlePartyOld(&agent, 0)
			Expect(party.AddPokemon(GeneratePokemon(PkmnBulbasaur))).To(MatchError(ErrorValidationMissingMoves))
		})
	})
})

var _ = Describe("Active pokemon", func() {
	agent := Agent(new(dumbAgent))
	var (
		party *battleParty
	)

	BeforeEach(func() {
		party = newOccupiedBattleParty(&agent, 0,
			PkmnWithMoves(TestMoveDefault),
			PkmnWithMoves(TestMoveDefault),
		)
	})

	Context("when managing active Pokemon in a party", func() {
		It("should correctly add an active Pokemon", func() {
			party.SetActive(0)
			Expect(party.GetActivePokemon()).To(HaveLen(1))
			Expect(party.activePokemon[0].NatDex).To(BeEquivalentTo(PkmnSquirtle))
			party.SetActive(1)
			Expect(party.activePokemon[1].NatDex).To(BeEquivalentTo(PkmnBlastoise))
		})

		It("should remove an active Pokemon", func() {
			party.SetActive(0)
			party.SetInactive(0)
			Expect(party.GetActivePokemon()).To(HaveLen(0))
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
})
