package pokemonbattlelib

import (
	"fmt"
	"testing"
)

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
