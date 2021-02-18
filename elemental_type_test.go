package pokemonbattlelib

import (
	"fmt"
	"testing"
)

func TestGetEffect(t *testing.T) {
	tests := []struct {
		move ElementalType
		def  ElementalType
		want Effectiveness
	}{
		{
			move: Water,
			def:  Fire,
			want: SuperEffective,
		},
		{
			move: Grass,
			def:  Water,
			want: SuperEffective,
		},
		{
			move: Fire,
			def:  Water,
			want: Ineffective,
		},
		{
			move: Ghost,
			def:  Normal,
			want: NoEffect,
		},
		{
			move: Steel,
			def:  Ground,
			want: NormalEffect,
		},
		{
			move: Fire,
			def:  Fire | Water,
			want: VeryIneffective,
		},
		{
			move: Water,
			def:  Ground | Rock,
			want: VerySuperEffective,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Type Effectiveness: %v vs %v", tt.move, tt.def), func(t *testing.T) {
			got := GetEffect(tt.move, tt.def)
			if got != tt.want {
				t.Errorf("Type Effectiveness: %v vs %v should be %f, got %f", tt.move, tt.def, tt.want, got)
			}
		})
	}
}

func TestElementalTypeString(t *testing.T) {
	tests := []struct {
		value ElementalType
		want  string
	}{
		{
			value: Water,
			want:  "[Water]",
		},
		{
			value: Normal,
			want:  "[Normal]",
		},
		{
			value: Fire,
			want:  "[Fire]",
		},
		{
			value: Fire | Grass,
			want:  "[Fire][Grass]",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("ElementalType Stringer: %s", tt.want), func(t *testing.T) {
			got := fmt.Sprintf("%s", tt.value)
			if got != tt.want {
				t.Errorf("ElementalType (%d) got %v, want %v", tt.value, got, tt.want)
			}
		})
	}
}
