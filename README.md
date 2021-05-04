# Pokemon Battle Library

A portable library for accurately simulating pokemon battles.

![CI Test](https://github.com/StevensSEC/pokemonbattlelib/workflows/CI%20Test/badge.svg)
[![codecov](https://codecov.io/gh/StevensSEC/pokemonbattlelib/branch/main/graph/badge.svg?token=lFGcKzL3Cp)](https://codecov.io/gh/StevensSEC/pokemonbattlelib)

# Usage

[Documentation](https://pkg.go.dev/github.com/StevensSEC/pokemonbattlelib)

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
		case UseMoveTransaction:
			fmt.Printf("%s used %s on %s", tt.User, tt.Move, tt.Target.Pokemon)
		case DamageTransaction:
			fmt.Printf(" %s took %d damage", tt.Target.Pokemon, tt.Damage)
		case HealTransaction:
			fmt.Printf("%s healed for %d HP", tt.Target, tt.Amount)
		case FaintTransaction:
			fmt.Printf("%s fainted", tt.Target.Pokemon)
		default:
			fmt.Printf("Transaction: %T - %v", t, t)
		}
	}
	//Output
	//Pikachu, Thunder Shock, Bulbasaur
	//Bulbasaur, Thunder Shock Damage Amount
	//Pikachu or Bulbasaur, Amount of HP healed
	//Bulbasaur fainted
}
//Displaying how a battle works. Two Pokemon are created and keep battling and their stats change such as health, number of moves, etc.
# Contributing

1. Make sure that you at least [have Go 1.15](https://golang.org/dl/)
2. Clone this repo

```bash
git clone git@github.com:StevensSEC/pokemonbattlelib.git
```

3. Make and commit your changes.
4. Create a Pull Request with your changes.

Using the github cli:
```bash
gh pr create
```

