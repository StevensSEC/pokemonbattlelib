package pokemonbattlelib

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get move by ID", func() {
	It("should get the correct move", func() {
		m := GetMove(MOVE_POUND)
		Expect(m.Name).To(Equal("Pound"))
	})

	It("should panic when a move does not exist", func() {
		Expect(func() {
			GetMove(-1)
		}).To(Panic())
	})
})

var _ = Describe("Move string representation", func() {
	tests := []struct {
		move Move
		want string
	}{
		{
			move: Move{
				Name:     "Shadow Ball",
				Type:     Ghost,
				Category: MoveCategorySpecial,
				MaxPP:    15,
				Priority: 0,
				Power:    80,
				Accuracy: 100,
			},
			want: "Shadow Ball\nType: [Ghost], Power: 80, Accuracy: 100\n",
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
