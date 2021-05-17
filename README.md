# Pokemon Battle Library

A portable library for accurately simulating pokemon battles.

![CI Test](https://github.com/StevensSEC/pokemonbattlelib/workflows/CI%20Test/badge.svg)
[![codecov](https://codecov.io/gh/StevensSEC/pokemonbattlelib/branch/main/graph/badge.svg?token=lFGcKzL3Cp)](https://codecov.io/gh/StevensSEC/pokemonbattlelib)

# Usage

[Documentation](https://pkg.go.dev/github.com/StevensSEC/pokemonbattlelib)

### Basic Example Battle

```go
p1 := GeneratePokemon(PkmnPikachu,
	WithLevel(20),
	WithMoves(MoveThunderShock))
p2 := GeneratePokemon(PkmnBulbasaur,
	WithLevel(20),
	WithMoves(MoveTackle))
a1 := Agent(new(dumbAgent))
b := New1v1Battle(p1, &a1, p2, &a1)
b.Start()
transactions, _ := b.SimulateRound()
for _, t := range transactions {
	switch tt := t.(type) {
	case UseMoveTransaction:
		fmt.Printf("%s used %s on %s\n", b.getPokemon(tt.User), tt.Move, b.getPokemon(tt.Target))
	case DamageTransaction:
		fmt.Printf("%s took %d damage\n", b.getPokemon(tt.Target), tt.Damage)
	case HealTransaction:
		fmt.Printf("%s healed for %d HP\n", b.getPokemon(tt.Target), tt.Amount)
	case FaintTransaction:
		fmt.Printf("%s fainted\n", b.getPokemon(tt.Target))
	case PPTransaction:
		continue
	default:
		fmt.Printf("Transaction: %T - %v\n", t, t)
	}
}
```

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

