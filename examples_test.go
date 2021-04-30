// This contains examples that will appear in the generated documentation. For that reason, they don't use gomega, or really perform any assertions.

package pokemonbattlelib

import (
	"fmt"
)

func ExampleBattle() {
	p1 := GeneratePokemon(PkmnPikachu,
		WithLevel(20),
		WithMoves(MoveThunderShock))
	p2 := GeneratePokemon(PkmnBulbasaur,
		WithLevel(20),
		WithMoves(MoveTackle))
	a1 := Agent(new(dumbAgent))
	b := New1v1Battle(p1, &a1, p2, &a1)
	transactions, _ := b.SimulateRound()
	for _, t := range transactions {
		switch tt := t.(type) {
		case DamageTransaction:
			fmt.Printf("%s used %s on %s for %d damage", tt.User, tt.Move, tt.Target.Pokemon, tt.Damage)
		case HealTransaction:
			fmt.Printf("%s healed for %d HP", tt.Target, tt.Amount)
		case FaintTransaction:
			fmt.Printf("%s fainted", tt.Target.Pokemon)
		default:
			fmt.Printf("Transaction: %T - %v", t, t)
		}
	}
}
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
