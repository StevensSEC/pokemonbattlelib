// This contains examples that will appear in the generated documentation. For that reason, they don't use gomega, or really perform any assertions.

package pokemonbattlelib

import "fmt"

func ExampleGeneratePokemon() {
	pkmn := GeneratePokemon(PkmnPikachu,
		WithLevel(28),
		WithMoves(MoveThunderShock, MoveIronTail, MoveTailWhip, MoveVoltTackle),
	)
	fmt.Printf("%s\n", pkmn)
	fmt.Printf("%s\n", pkmn.Moves)
	// Output:
	// Pikachu
	// [Thunder Shock Iron Tail Tail Whip Volt Tackle]
}

func ExampleParty() {
	party := NewParty()
	err := party.AddPokemon(
		GeneratePokemon(PkmnPikachu, WithMoves(MoveGrowl)),
		GeneratePokemon(PkmnCharizard, WithMoves(MoveGrowl)),
	)
	if err != nil {
		fmt.Print(err)
		return
	}
	for _, p := range party.Pokemon {
		fmt.Printf("%s\n", p)
	}
	// Output:
	// Pikachu
	// Charizard
}
