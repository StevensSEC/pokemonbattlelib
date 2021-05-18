package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scrappy", func() {
	a := Agent(new(dumbAgent))

	PDescribeTable("should allow damage to ghost types",
		func(element Type) {
			b := New1v1Battle(
				GeneratePokemon(PkmnBidoof, WithAbility(AbilityScrappy), WithMoves(registerMoveWithType(element))), &a,
				GeneratePokemon(PkmnGengar, WithMoves(TestMoveNoDamage)), &a,
			)
			Expect(b.Start()).To(Succeed())
			t, _ := b.SimulateRound()
			dealt := DamageDealt(t, target{1, 0})
			Expect(dealt).To(BeNumerically(">", 0))
		},
		Entry("normal", TypeNormal),
		Entry("fighting", TypeFighting),
	)
})
