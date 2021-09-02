package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ability: Air Lock", func() {
	a := Agent(new(dumbAgent))

	It("should negate the effects of weather", func() {
		p := PkmnNoDamage()
		b := New1v1Battle(p, &a, PkmnNoDamage(), &a)
		b.Weather = WeatherHail
		b.rng = SimpleRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransaction(DamageTransaction{
			Target: target{1, 0},
			Move:   nil,
		}))
		p.Ability = AbilityAirLock
		t, _ = b.SimulateRound()
		Expect(t).ToNot(HaveTransaction(DamageTransaction{
			Target: target{1, 0},
			Move:   nil,
		}))
	})
})

var _ = Describe("Ability: Scrappy", func() {
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

var _ = Describe("Ability: Iron Fist", func() {
	a := Agent(new(dumbAgent))

	It("should have punching moves deal 1.2x damage", func() {
		b := New1v1Battle(
			GeneratePokemon(PkmnMachop, WithLevel(15), WithAbility(AbilityIronFist), WithMoves(MoveCometPunch)), &a,
			GeneratePokemon(PkmnBidoof, WithLevel(15), WithMoves(MoveCometPunch)), &a,
		)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		normalDamage := DamageDealt(t, target{1, 0})
		boostedDamage := DamageDealt(t, target{0, 0})
		Expect(normalDamage).To(BeNumerically(">", 0))
		Expect(boostedDamage).To(BeNumerically(">", 0))
		Expect(boostedDamage).To(BeNumerically(">", normalDamage))
		Expect(normalDamage * 120 / 100).To(BeNumerically("==", boostedDamage))
	})

})
