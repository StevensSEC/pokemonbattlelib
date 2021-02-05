package pokemonbattlelib

import (
	"fmt"
	"log"
)

// TEMP
type _Turn interface{}

type Engine struct {
	battle   *Battle
	history  []_Turn
	listener Listener
	party    map[Player][]Pokemon
	state    EngineState
}

type Player string

const (
	PLAYER1 Player = "PLAYER 1"
	PLAYER2 Player = "PLAYER 2"
)

type EngineState string

const (
	BEFORE_READY        EngineState = "BEFORE_READY"
	READY               EngineState = "READY"
	BATTLE_START        EngineState = "BATTLE_START"
	BEFORE_PLAYER1_TURN EngineState = "BEFORE_PLAYER1_TURN"
	AFTER_PLAYER1_TURN  EngineState = "AFTER_PLAYER1_TURN"
	BEFORE_PLAYER2_TURN EngineState = "BEFORE_PLAYER2_TURN"
	AFTER_PLAYER2_TURN  EngineState = "AFTER_PLAYER2_TURN"
	BATTLE_END          EngineState = "BATTLE_END"
)

type Listener interface {
	OnBeforeReady()
	OnReady()
	OnBattleStart()
	OnBeforePlayer1Turn()
	OnAfterPlayer1Turn()
	OnBeforePlayer2Turn()
	OnAfterPlayer2Turn()
	OnBattleEnd()
}

func NewEngine(listener Listener) *Engine {
	history := make([]_Turn, 0)
	party := make(map[Player][]Pokemon)
	battle := &Battle{}
	engine := &Engine{
		battle:   battle,
		history:  history,
		listener: listener,
		party:    party,
		state:    BEFORE_READY,
	}
	if engine.listener != nil {
		engine.listener.OnBeforeReady()
	}
	return engine
}

const MAX_PARTY_SIZE = 6

func (e *Engine) AddPokemon(player Player, pokemon Pokemon) {
	if len(e.party[player]) < MAX_PARTY_SIZE {
		e.party[player] = append(e.party[player], pokemon)
	} else {
		log.Fatalf("player party size cannot exceed %d", MAX_PARTY_SIZE)
	}
	if len(e.party[PLAYER1]) >= 1 && len(e.party[PLAYER2]) >= 1 {
		e.state = READY
		if e.listener != nil {
			e.listener.OnReady()
		}
	}
}

func (e *Engine) StartBattle(settings interface{}) {
	if e.state != READY {
		log.Fatalf("not enough Pokemon added to start a battle")
	}
	e.state = BATTLE_START
	if e.listener != nil {
		e.listener.OnBattleStart()
	}
}

func (e *Engine) String() string {
	return fmt.Sprintf("<Engine: %s - Turn %d>", e.state, len(e.history)+1)
}
