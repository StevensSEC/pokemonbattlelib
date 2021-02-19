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
			got := GetElementalEffect(tt.move, tt.def)
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

func TestGenderString(t *testing.T) {
	tests := []struct {
		name   string
		gender Gender
		want   string
	}{
		{
			name:   "Genderless",
			gender: Genderless,
			want:   "",
		},
		{
			name:   "Female",
			gender: Female,
			want:   "♀",
		},
		{
			name:   "Male",
			gender: Male,
			want:   "♂",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Gender Stringer: %s", tt.name), func(t *testing.T) {
			got := fmt.Sprintf("%s", tt.gender)
			if got != tt.want {
				t.Errorf("Gender Stringer %s got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestStatusConditionApply(t *testing.T) {
	tests := []struct {
		name  string
		from  StatusCondition
		apply StatusCondition
		want  StatusCondition
	}{
		{
			name:  "Should apply burn",
			from:  StatusNone,
			apply: StatusBurn,
			want:  StatusBurn,
		},
		{
			name:  "Should apply burn and remove poison",
			from:  StatusPoison,
			apply: StatusBurn,
			want:  StatusBurn,
		},
		{
			name:  "Should apply freeze and remove sleep",
			from:  StatusSleep,
			apply: StatusFreeze,
			want:  StatusFreeze,
		},
		{
			name:  "Should preserve bound, replace sleep with poison",
			from:  StatusSleep | StatusBound,
			apply: StatusPoison,
			want:  StatusPoison | StatusBound,
		},
		{
			name:  "Should preserve sleep, apply bound",
			from:  StatusSleep,
			apply: StatusBound,
			want:  StatusSleep | StatusBound,
		},
		{
			name:  "Should preserve sleep, leech seed, apply bound",
			from:  StatusSleep | StatusLeechSeed,
			apply: StatusBound,
			want:  StatusSleep | StatusLeechSeed | StatusBound,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Status Condition apply: %s", tt.name), func(t *testing.T) {
			got := tt.from
			got.apply(tt.apply)
			if got != tt.want {
				t.Errorf("Status Condition apply %s got %b, want %b", tt.name, got, tt.want)
			}
		})
	}
}
