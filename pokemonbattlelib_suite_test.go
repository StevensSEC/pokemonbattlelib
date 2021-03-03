package pokemonbattlelib

import (
	"log"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPokemonbattlelib(t *testing.T) {
	log.SetOutput(nil)
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pokemonbattlelib Suite")
}
