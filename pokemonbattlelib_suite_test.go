package pokemonbattlelib_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPokemonbattlelib(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pokemonbattlelib Suite")
}
