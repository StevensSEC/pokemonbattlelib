package pokemonbattlelib

import "testing"

func TestGetEffect(t *testing.T) {
	move := Ice
	def := Grass

	expectedOutput := 200
	actual, err := GetEffect(move, def)

	if err != nil {
		t.Errorf("expected no error, but %v", err)
	}

	if actual != expectedOutput {
		t.Errorf("exepcted effect is %v, but got %v", expectedOutput, actual)
	}
}
