package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Using items", func() {
	It("should correctly retrieve items by ID", func() {
		i := GetItem(ItemPotion)
		Expect(i.Name).To(Equal(("Potion")))
		Expect(func() {
			GetItem(-1)
		}).To(Panic())
	})

	It("should use the item and produce transactions", func() {
		i := GetItem(ItemPotion)
		p := GeneratePokemon(1)
		p.Stats[StatHP] = 100
		t := p.UseItem(&i)
		Expect(t).To(HaveLen(1))
		Expect(t).To(HaveTransaction(
			HealTransaction{
				Target: p,
				Amount: 20,
			},
		))
	})
})
