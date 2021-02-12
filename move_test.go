package pokemonbattlelib

import (
	"fmt"
	"testing"
)

func TestMoveString(t *testing.T) {
	m := Move{
		Name:     "Shadow Ball",
		Type:     8,
		Category: Special,
		MaxPP:    15,
		Priority: 0,
		Power:    80,
		Accuracy: 100,
	}
	if m.String() != "Shadow Ball\nType: 8, Power: 80, Accuracy: 100\n" {
		t.Fail()
	}
}

func TestMoveCategoryString(t *testing.T) {
	tests := []struct {
		value MoveCategory
		want  string
	}{
		{
			value: Status,
			want:  "Status",
		},
		{
			value: Physical,
			want:  "Physical",
		},
		{
			value: Special,
			want:  "Special",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("MoveCategory Stringer: %s", tt.want), func(t *testing.T) {
			t.Parallel()
			got := tt.value.String()
			if got != tt.want {
				t.Errorf("Move Category (%d) got %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
