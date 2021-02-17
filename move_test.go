package pokemonbattlelib

import (
	"fmt"
	"testing"
)

func TestGetMove(t *testing.T) {
	m := GetMove(1)
	if m.Name != "Pound" {
		t.Errorf("expected move name to be Pound, got %v", m.Name)
	}
}

func TestMoveRestorePP(t *testing.T) {
	m := Move{Name: "Pound", CurrentPP: 15, MaxPP: 25}
	logtest := []struct {
		turn Transaction
		want string
	}{
		{
			turn: m.RestorePP(5),
			want: "Pound restored 5 PP.",
		},
		{
			turn: m.RestorePP(99),
			want: "Pound restored 10 PP.",
		},
	}
	for _, tt := range logtest {
		got := tt.turn.BattleLog()
		if got != tt.want {
			t.Errorf("Expected battle log to be %s, got %s", tt.want, got)
		}
	}
}

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
			t.Parallel()
			got := fmt.Sprintf("%s", tt.move)
			if got != tt.want {
				t.Errorf("Move Stringer %s got %v, want %v", tt.move.Name, got, tt.want)
			}
		})
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
