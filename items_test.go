package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Using items", func() {
	DescribeTable("should correctly retrieve items by ID",
		func(item Item, name string) {
			Expect(item.Data().Name).To(Equal(name))
		},
		Entry("No item", ItemNone, ""),
		Entry("Potion", ItemPotion, "Potion"),
		Entry("Flame orb", ItemFlameOrb, "Flame Orb"),
	)

	It("should use the item and produce transactions", func() {
		p := GeneratePokemon(PkmnBulbasaur)
		p.Stats[StatHP] = 100
		t := p.UseItem(ItemPotion)
		Expect(t).To(HaveTransaction(
			HealTransaction{
				Target: p,
				Amount: 20,
			},
		))
	})
})
