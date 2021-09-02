package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Double Battles", func() {
	a1 := Agent(new(dumbAgent))

	It("should send out the first 2 pokemon of each party", func() {
		b := New2v2Battle(
			PkmnNoDamage(),
			PkmnNoDamage(),
			&a1,
			PkmnNoDamage(),
			PkmnNoDamage(),
			&a1,
		)
		Expect(b.Start()).To(Succeed())
		for _, party := range b.parties {
			Expect(len(party.activePokemon)).To(Equal(2))
		}
	})

	It("should have 4 UseMoveTransactions", func() {
		b := New2v2Battle(
			PkmnNoDamage(),
			PkmnNoDamage(),
			&a1,
			PkmnNoDamage(),
			PkmnNoDamage(),
			&a1,
		)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		count := 0
		for _, tt := range t {
			switch tt.(type) {
			case UseMoveTransaction:
				count++
			}
		}
		Expect(count).To(Equal(4))
	})

	It("should send out 1 pokemon for each party in team 0", func() {
		b := NewBattleOfType(BattleTypeDouble)
		p0a := NewOccupiedParty(PkmnDefault())
		p0b := NewOccupiedParty(PkmnDefault())
		p1 := NewOccupiedParty(PkmnDefault(), PkmnDefault())
		b.AddParty(p0a, &a1, 0)
		b.AddParty(p0b, &a1, 0)
		b.AddParty(p1, &a1, 1)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		for _, party := range b.GetPartiesOnTeam(0) {
			Expect(len(party.activePokemon)).To(Equal(1))
		}
		for _, party := range b.GetPartiesOnTeam(1) {
			Expect(len(party.activePokemon)).To(Equal(2))
		}
	})

	It("should allow 1 pokemon parties to participate in a double battle by themselves", func() {
		b := NewBattleOfType(BattleTypeDouble)
		p0 := NewOccupiedParty(PkmnDefault())
		p1 := NewOccupiedParty(PkmnDefault(), PkmnDefault())
		b.AddParty(p0, &a1, 0)
		b.AddParty(p1, &a1, 1)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		for _, party := range b.GetPartiesOnTeam(0) {
			Expect(len(party.activePokemon)).To(Equal(1))
		}
		for _, party := range b.GetPartiesOnTeam(1) {
			Expect(len(party.activePokemon)).To(Equal(2))
		}
	})

	It("should not allow battle to start with invalid teams", func() {
		b := NewBattleOfType(BattleTypeDouble)
		p0 := NewOccupiedParty(PkmnDefault(), PkmnDefault())
		b.AddParty(p0, &a1, 0)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Not(Succeed()))
	})
})
