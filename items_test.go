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
	entries := []TableEntry{
		Entry("None", ItemNone),
		Entry("Potion", ItemPotion),
		Entry("Oran Berry", ItemOranBerry),
		Entry("Ultra Ball", ItemUltraBall),
	}

	DescribeTable("should marshall and unmarshall items",
		func(i Item) {
			b, err := json.Marshal(i)
			Expect(err).To(Succeed())
			var got Item
			err = json.Unmarshal(b, &got)
			Expect(err).To(Succeed())
			Expect(got).To(BeEquivalentTo(i))
		},
		entries...,
	)

	DescribeTable("marshalled items should include data",
		func(i Item) {
			b, err := json.Marshal(i)
			Expect(err).To(Succeed())
			var got ItemData
			err = json.Unmarshal(b, &got)
			Expect(err).To(Succeed())
			Expect(got.Name).To(BeEquivalentTo(i.Name()))
			Expect(got.Flags).To(BeEquivalentTo(i.Flags()))
		},
		entries...,
	)

	DescribeTable("unmarshalling items should accept just the ID",
		func(i Item) {
			b, err := json.Marshal(uint16(i))
			Expect(err).To(Succeed())
			var got Item
			err = json.Unmarshal(b, &got)
			Expect(err).To(Succeed())
			Expect(got).To(BeEquivalentTo(i))
		},
		entries...,
	)

	DescribeTable("unmarshalling items should accept a struct with an Id field",
		func(i Item) {
			item := struct {
				Id uint16
			}{
				Id: uint16(i),
			}
			b, err := json.Marshal(item)
			Expect(err).To(Succeed())
			var got Item
			err = json.Unmarshal(b, &got)
			Expect(err).To(Succeed())
			Expect(got).To(BeEquivalentTo(i))
		},
		entries...,
	)
})
