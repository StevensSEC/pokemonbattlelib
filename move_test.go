package pokemonbattlelib

import (
	"fmt"
	"testing"

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

func TestMoveString(t *testing.T) {
	tests := []struct {
		move Move
		want string
	}{
		{
			move: Move{
				Name:     "Shadow Ball",
				Type:     Ghost,
				Category: Special,
				MaxPP:    15,
				Priority: 0,
				Power:    80,
				Accuracy: 100,
			},
			want: "Shadow Ball\nType: [Ghost], Power: 80, Accuracy: 100\n",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Move Stringer: %s", tt.move.Name), func(t *testing.T) {
			got := tt.move.String()
			if got != tt.want {
				t.Errorf("Move Stringer %s got %v, want %v", tt.move.Name, got, tt.want)
			}
		})
	}
}

var _ = Describe("Move category", func() {
	It("should show correct string for move category", func() {
		Expect(Status.String()).To(Equal("Status"))
		Expect(Physical.String()).To(Equal("Physical"))
		Expect(Special.String()).To(Equal("Special"))
		Expect(func() {
			var _ = MoveCategory(99).String()
		}).To(Panic())
	})
})
