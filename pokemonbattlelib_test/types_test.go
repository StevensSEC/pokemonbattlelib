package pokemonbattlelib_test

import ( 
	"testing"
    . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
 
    ".."

	"fmt"
	"reflect"
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

func TestStatusConditionCheck(t *testing.T) {
	tests := []struct {
		status StatusCondition
		check  StatusCondition
	}{
		{
			status: StatusNone,
			check:  StatusNone,
		},
		{
			status: StatusBurn,
			check:  StatusBurn,
		},
		{
			status: StatusSleep | StatusCursed,
			check:  StatusSleep,
		},
		{
			status: StatusSleep | StatusCursed,
			check:  StatusCursed,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Status Condition check: %b", tt.status), func(t *testing.T) {
			if !tt.status.check(tt.check) {
				t.Errorf("Status Condition check failed %b, should have %b", tt.status, tt.check)
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

func TestStatusConditionClear(t *testing.T) {
	tests := []struct {
		name  string
		from  StatusCondition
		clear StatusCondition
		want  StatusCondition
	}{
		{
			name:  "Should clear burn",
			from:  StatusBurn,
			clear: StatusBurn,
			want:  StatusNone,
		},
		{
			name:  "Should preserve bound, clear sleep",
			from:  StatusSleep | StatusBound,
			clear: StatusSleep,
			want:  StatusBound,
		},
		{
			name:  "Should preserve bound, apply bound",
			from:  StatusSleep | StatusBound,
			clear: StatusBound,
			want:  StatusSleep,
		},
		{
			name:  "Should preserve sleep, leech seed, clear bound",
			from:  StatusSleep | StatusLeechSeed | StatusBound,
			clear: StatusBound,
			want:  StatusSleep | StatusLeechSeed,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Status Condition clear: %s", tt.name), func(t *testing.T) {
			got := tt.from
			got.clear(tt.clear)
			if got != tt.want {
				t.Errorf("Status Condition clear %s got %b, want %b", tt.name, got, tt.want)
			}
		})
	}
}
