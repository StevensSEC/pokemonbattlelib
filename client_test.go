package pokemonbattlelib

func handleTurn() interface{} {
	return "123"
}

func ExampleClient() {
	c := NewClient(handleTurn)
	p := Pokemon{1}
	c.AddPokemon(p)
	c.AddPokemon(Pokemon{7})
	// Output: 123
}
