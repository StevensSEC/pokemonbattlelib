package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Using items", func() {
	DescribeTable("should correctly retrieve items by ID",
		func(item Item, name string) {
			Expect(item.Name()).To(Equal(name))
		},
		Entry("No item", ItemNone, ""),
		Entry("Potion", ItemPotion, "Potion"),
		Entry("Flame orb", ItemFlameOrb, "Flame Orb"),
	)
})
