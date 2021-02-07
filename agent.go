package pokemonbattlelib

import (
	"log"
	"time"
)

type Agent struct {
	ID         int64
	Party      []Pokemon
	EventHooks map[BattleState]func()
	HandleTurn TurnHandler
}
type TurnHandler func(*Battle) interface{}

// Creates a new Agent with a unique ID
func NewAgent(handleTurn TurnHandler) *Agent {
	return &Agent{
		ID:         time.Now().Unix(),
		Party:      make([]Pokemon, 0),
		EventHooks: make(map[BattleState]func()),
		HandleTurn: handleTurn,
	}
}

// Registers a custom handler for a change in battle state
func (a *Agent) RegisterHook(state BattleState, handler func()) {
	if _, ok := a.EventHooks[state]; ok {
		log.Panicf("hook already registered for %v", state)
	}
	a.EventHooks[state] = handler
}

// Unregisters a custom handler for a change in battle state
func (a *Agent) UnregisterHook(state BattleState) {
	if _, ok := a.EventHooks[state]; !ok {
		log.Panicf("no event hook found for %v", state)
	}
	delete(a.EventHooks, state)
}

const MAX_PARTY_SIZE = 6

// Adds Pokemon to the agent's party
func (a *Agent) AddPokemon(pokemon ...Pokemon) {
	if len(a.Party)+len(pokemon) <= MAX_PARTY_SIZE {
		a.Party = append(a.Party, pokemon...)
	} else {
		log.Panicf("player party size cannot exceed %d", MAX_PARTY_SIZE)
	}
}

// Removes a Pokemon by their position in the party
func (a *Agent) RemovePokemon(i int) {
	if i < 0 || i > len(a.Party)-1 {
		log.Panicf("cannot remove Pokemon at position %v", i)
	}
	a.Party = append(a.Party[:i], a.Party[i+1:]...)
}
