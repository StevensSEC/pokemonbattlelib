package pokemonbattlelib_test

import (
	"testing"

	. "github.com/StevensSEC/pokemonbattlelib"
)

func TestLCRNGSeed(t *testing.T) {
	gen := LCRNG(0)
	gen.SetSeed(1234)
	if int(gen) != 1234 {
		t.Fail()
	}
}

func TestLCRNGGet(t *testing.T) {
	gen := LCRNG(678576)
	value := gen.Get(1, 10)
	if value > 10 || value < 1 {
		t.Fail()
	}
}
