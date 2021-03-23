package pokemonbattlelib

import (
	"encoding/json"
	"math"

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

	It("should panic if item is invalid", func() {
		Expect(func() { Item(math.MaxUint16).Data() }).To(Panic())
	})
})

var _ = Describe("Item Marshalling and Unmarshalling", func() {
	DescribeTable("should marshall and unmarshall items",
		func(i Item) {
			b, err := json.Marshal(i)
			Expect(err).To(Succeed())
			var got Item
			err = json.Unmarshal(b, &got)
			Expect(err).To(Succeed())
			Expect(got).To(BeEquivalentTo(i))
		},
		Entry("None", ItemNone),
		Entry("Potion", ItemPotion),
		Entry("Oran Berry", ItemOranBerry),
		Entry("Ultra Ball", ItemUltraBall),
	)
})
