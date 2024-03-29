package pokemonbattlelib

import (
	"encoding/json"
	"fmt"
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get move by ID", func() {
	It("should get the correct move", func() {
		m := GetMove(MovePound)
		Expect(m.Name()).To(Equal("Pound"))
	})

	It("should panic when a move does not exist", func() {
		Expect(func() {
			GetMove(MoveId(math.MaxUint16))
		}).To(Panic())
	})
})

var _ = Describe("Move Marshalling and Unmarshalling", func() {
	DescribeTable("should marshall and unmarshall moves",
		func(m MoveId) {
			move := GetMove(m)
			b, err := json.Marshal(move)
			Expect(err).To(Succeed())
			var got Move
			err = json.Unmarshal(b, &got)
			Expect(err).To(Succeed())
			Expect(got).To(BeEquivalentTo(*move))
		},
		Entry("Pound", MovePound),
		Entry("Healing Wish", MoveHealingWish),
		Entry("Water Gun", MoveWaterGun),
		Entry("Whirlpool", MoveWhirlpool),
		Entry("Aerial Ace", MoveAerialAce),
		Entry("Ember", MoveEmber),
		Entry("Lava Plume", MoveLavaPlume),
	)
})

var _ = Describe("Move string representation", func() {
	tests := []struct {
		move Move
		want string
	}{
		{
			move: *GetMove(MoveShadowBall),
			want: "Shadow Ball",
		},
	}
	It("should show the correct string for moves", func() {
		for _, tt := range tests {
			Expect(fmt.Sprintf("%s", tt.move)).To(Equal(tt.want))
		}
	})
})

var _ = Describe("Move category", func() {
	It("should show correct string for move category", func() {
		Expect(MoveCategoryStatus.String()).To(Equal("Status"))
		Expect(MoveCategoryPhysical.String()).To(Equal("Physical"))
		Expect(MoveCategorySpecial.String()).To(Equal("Special"))
		Expect(func() {
			var _ = MoveCategory(99).String()
		}).To(Panic())
	})
})
