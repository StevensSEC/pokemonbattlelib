package pokemonbattlelib

import (
	"fmt"
	"log"
)

type Battle struct {
	// TODO
	State   BattleState
	Players []*Agent
}
type BattleState int

type Player int

const (
	PLAYER1 Player = iota
	PLAYER2
)
const (
	BEFORE_START BattleState = iota
	BATTLE_START
	BEFORE_PLAYER1_TURN
	AFTER_PLAYER1_TURN
	BEFORE_PLAYER2_TURN
	AFTER_PLAYER2_TURN
	BATTLE_END
)

// Creates a new battle instance that agents can connect to
func NewBattle() *Battle {
	return &Battle{State: BEFORE_START}
}

// Updates the battle state and dispatches calls to custom event hooks
func (b *Battle) Dispatch(state BattleState) {
	b.State = state
	for _, player := range b.Players {
		if handler, ok := player.EventHooks[state]; ok {
			handler()
		}
	}
}

// Adds player(s) to the battle
// Fails if two players are already connected
func (b *Battle) AddAgents(a ...*Agent) {
	if len(b.Players)+len(a) > 2 {
		log.Panicf("battle cannot have more than two agents connected")
	}
	b.Players = append(b.Players, a...)
}

// Removes a player by their ID from the battle
func (b *Battle) RemoveAgent(agentID int64) {
	players := make([]*Agent, 0)
	for _, p := range b.Players {
		if p.ID != agentID {
			players = append(players, p)
		}
	}
	if len(players) == len(b.Players) {
		log.Panicf("no agent with ID '%v' found", agentID)
	}
	b.Players = players
}

// Starts a battle if all pre-conditions are satisfied
func (b *Battle) Start() {
	if len(b.Players) < 2 {
		log.Panicf("not enough players to start a battle")
	}
	if len(b.Players[PLAYER1].Party) < 1 || len(b.Players[PLAYER2].Party) < 1 {
		log.Panicf("not enough Pokemon in player parties to start a battle")
	}
	b.Simulate()
}

// Simulates a battle using an event loop to handle all I/O and event dispatching
func (b *Battle) Simulate() {
	b.Dispatch(BATTLE_START)
	for b.State != BATTLE_END {
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
				break
			}
			b.Dispatch(BEFORE_PLAYER1_TURN)
		}
	}
	// Cleanup after battle is over / ask to start new battle
}

func (b *Battle) getActivePokemon(idx int) *Pokemon {
	// TODO: implment
	// 1. find active pokemon in all parties
	// 2. should return the pokemon at the given index
	return nil
}

// Tells the battle that the active this pokemon is making this move.
func (b *Battle) processTurn(activePokemon int, t Turn) {
	switch t := t.(type) {
	case FightTurn:
		fmt.Printf("TODO: Implement fight %v", t)
	default:
		panic("Unknown turn")
	}
}

// An abstration over all possible actions an `Agent` can make in one round. Each Pokemon gets one turn.
type Turn interface {
	Priority() int // Gets the turn's priority. Higher values go first.
}

type FightTurn struct {
	moveIdx   int // Denotes which of the pokemon's moves to use.
	targetIdx int // The active pokemon that on the receiving end of the move.
}

func (turn FightTurn) Priority() int {
	return 0
}
