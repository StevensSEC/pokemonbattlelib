package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Using items", func() {
	It("should correctly retrieve items by ID", func() {
		i := GetItem(ITEM_POTION)
		Expect(i.Name).To(Equal(("Potion")))
		Expect(func() {
			GetItem(-1)
		}).To(Panic())
	})

	It("should use the item and produce transactions", func() {
		i := GetItem(ITEM_POTION)
		p := GeneratePokemon(1)
		p.Stats[StatHP] = 100
		logs := p.UseItem(&i)
		Expect(logs).To(HaveLen(1))
		Expect(logs[0].BattleLog()).To(Equal("Bulbasaur restored 20 HP."))
	})
})
