package pokemonbattlelib

import (
	"fmt"
	"testing"
)

func TestGetEffect(t *testing.T) {
	move := Ice
	def := Grass

	expectedOutput := 200
	actual := GetEffect(move, def)

	if actual != expectedOutput {
		t.Errorf("exepcted effect is %v, but got %v", expectedOutput, actual)
	}

	tests := []struct {
		move ElementalType
		def  ElementalType
		want int
	}{
		{
			move: Water,
			def:  Fire,
			want: 200,
		},
		{
			move: Grass,
			def:  Water,
			want: 200,
		},
		{
			move: Fire,
			def:  Water,
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Type Effectiveness: %v vs %v", tt.move, tt.def), func(t *testing.T) {
			t.Parallel()
			got := GetEffect(tt.move, tt.def)
			if got != tt.want {
				t.Errorf("Type Effectiveness: %v vs %v should be %d, got %d", tt.move, tt.def, tt.want, got)
			}
		})
	}
}
