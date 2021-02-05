package pokemonbattlelib

import "fmt"

// this will change to use a map for hooks e.g.
// engine.Register(BEFORE_BATTLE, handler_func)
type CustomListener struct{}

func (l *CustomListener) OnBeforeReady() {
	fmt.Println("called on before ready")
}

func (l *CustomListener) OnBattleStart() {
	fmt.Println("called right before battle starts")
}

func ExampleEngine() {
	l := &CustomListener{}
	engine := NewEngine(l)
	fmt.Println("engine:", engine)
	p := Pokemon{1}
	engine.AddPokemon(PLAYER1, p)
	engine.AddPokemon(PLAYER2, Pokemon{7})
	engine.StartBattle(nil)
	// Output: 123
}
