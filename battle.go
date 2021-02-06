package pokemonbattlelib

import "fmt"

type Battle struct {
	// TODO
	State   BattleState
	Player1 *Client
	Player2 *Client
}
type BattleState int

const (
	BEFORE_START BattleState = iota
	BATTLE_START
	BEFORE_PLAYER1_TURN
	AFTER_PLAYER1_TURN
	BEFORE_PLAYER2_TURN
	AFTER_PLAYER2_TURN
	BATTLE_END
)

// Creates a new battle instance that clients can connect to
func NewBattle() *Battle {
	return &Battle{State: BEFORE_START}
}

// Updates the battle state and dispatches calls to custom event hooks
func (b *Battle) Dispatch(state BattleState) {
	b.State = state
	if handler, ok := b.Player1.EventHooks[state]; ok {
		handler()
	}
	if handler, ok := b.Player2.EventHooks[state]; ok {
		handler()
	}
}

// Adds a player to the battle
// Fails if two players are already connected
func (b *Battle) AddClient(c *Client) error {
	if b.Player1 == nil {
		b.Player1 = c
	} else if b.Player2 == nil {
		b.Player2 = c
	} else {
		return fmt.Errorf("battle already has two clients connected")
	}
	return nil
}

// Removes a player by their ID from the battle
func (b *Battle) RemoveClient(clientID int64) error {
	if b.Player1 != nil && b.Player1.ID == clientID {
		b.Player1 = nil
	} else if b.Player2 != nil && b.Player2.ID == clientID {
		b.Player2 = nil
	} else {
		return fmt.Errorf("no client with ID '%v' found", clientID)
	}
	return nil
}

// Swaps player 1 with player 2 in a battle
func (b *Battle) SwapPlayers() error {
	if b.Player1 == nil || b.Player2 == nil {
		return fmt.Errorf("battle does not have two clients/players")
	}
	b.Player1, b.Player2 = b.Player2, b.Player1
	return nil
}

// Starts a battle if all pre-conditions are satisfied
func (b *Battle) Start() error {
	if b.Player1 == nil || b.Player2 == nil {
		return fmt.Errorf("not enough players to start a battle")
	}
	if len(b.Player1.Party) < 1 || len(b.Player2.Party) < 1 {
		return fmt.Errorf("not enough Pokemon in player parties to start a battle")
	}
	return b.Simulate()
}

// Simulates a battle using an event loop to handle all I/O and event dispatching
func (b *Battle) Simulate() error {
	b.Dispatch(BATTLE_START)
	for {
		switch b.State {
		case BATTLE_START:
			// Initialize battle conditions (weather, abilities, etc)
			b.Dispatch(BEFORE_PLAYER1_TURN)
		case BEFORE_PLAYER1_TURN:
			// action := b.Player1.HandleTurn()
			// Fix: Handle action
			// fmt.Println("player 1 action:", action)
			b.Dispatch(AFTER_PLAYER1_TURN)
		case AFTER_PLAYER1_TURN:
			// Fix: Handle events that occur between turns
			// This includes weather, status effects, abilities, swaps, etc
			b.Dispatch(BEFORE_PLAYER2_TURN)
		case BEFORE_PLAYER2_TURN:
			// action := b.Player2.HandleTurn()
			// Fix: Handle action
			// fmt.Println("player 2 action:", action)
			b.Dispatch(AFTER_PLAYER2_TURN)
		case AFTER_PLAYER2_TURN:
			// Fix: Handle events that occur after both players move
			// Also check if battle is over, otherwise go back to player 1
			// Temporarily end battle after 1 round
			if true {
				b.Dispatch(BATTLE_END)
				return nil
			}
			b.Dispatch(BEFORE_PLAYER1_TURN)
		}
	}
}
