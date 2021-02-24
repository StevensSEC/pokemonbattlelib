package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LCRNG type", func() {
	var (
		gen LCRNG
	)

	It("should allow a user to set the seed", func() {
		gen = LCRNG(0)
		gen.SetSeed(1234)
		Expect(int(gen)).To(Equal(1234))
	})

	It("should generate a value in bounds", func() {
		gen = LCRNG(678576)
		got := gen.Get(1, 10)
		Expect(got <= 10).To(BeTrue())
		Expect(got >= 1).To(BeTrue())
	})

	It("should roll predictably", func() {
		gen = LCRNG(1337)
		Expect(gen.Roll(10)).To(BeTrue())
		Expect(gen.Roll(10)).To(BeFalse())
	})
})
