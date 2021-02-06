package pokemonbattlelib

import (
	"fmt"
	"time"
)

type Client struct {
	ID         int64
	Party      []Pokemon
	EventHooks map[BattleState]func()
	HandleTurn TurnHandler
}
type TurnHandler func() interface{}

// Creates a new client with a unique ID
func NewClient(handleTurn TurnHandler) *Client {
	return &Client{
		ID:         time.Now().Unix(),
		Party:      make([]Pokemon, 6),
		EventHooks: make(map[BattleState]func()),
		HandleTurn: handleTurn,
	}
}

// Registers a custom handler for a change in battle state
func (c *Client) RegisterHook(state BattleState, handler func()) error {
	if _, ok := c.EventHooks[state]; ok {
		return fmt.Errorf("hook already registered for %v", state)
	}
	c.EventHooks[state] = handler
	return nil
}

// Unregisters a custom handler for a change in battle state
func (c *Client) UnregisterHook(state BattleState) error {
	if _, ok := c.EventHooks[state]; !ok {
		return fmt.Errorf("no event hook found for %v", state)
	}
	delete(c.EventHooks, state)
	return nil
}

const MAX_PARTY_SIZE = 6

func (c *Client) AddPokemon(pokemon Pokemon) error {
	if len(c.Party) < MAX_PARTY_SIZE {
		c.Party = append(c.Party, pokemon)
	} else {
		return fmt.Errorf("player party size cannot exceed %d", MAX_PARTY_SIZE)
	}
	return nil
}

func (c *Client) RemovePokemon(pokemonID string) error {
	// Fix: remove Pokemon by ID
	return fmt.Errorf("not implemented")
}
