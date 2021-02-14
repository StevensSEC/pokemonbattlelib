package pokemonbattlelib

import (
	"log"
	"reflect"
	"sort"
)

// A Pokemon battle. Enforces rules of the battle, and queries `Agent`s for turns.
type Battle struct {
	Weather  int  // one of the 6 in-battle weather conditions
	ShiftSet bool // shift or set battle style for NPC trainer battles
	State    BattleState

	parties []*party // All parties participating in the battle
}

type BattleState int

const (
	BEFORE_START BattleState = iota
	BATTLE_IN_PROGRESS
	BATTLE_END
)

// Creates a new battle instance, setting initial conditions
func NewBattle() *Battle {
	b := Battle{
		State: BEFORE_START,
	}
	return &b
}

// Maps agents to the indices of their Pokemon. Used for Allies and Opponents.
type AgentParties map[*Agent]map[int]Pokemon

// Adds one or more parties to a team in the battle
func (b *Battle) AddParty(p ...*party) {
	b.parties = append(b.parties, p...)
}

// Gets all active ally Pokemon for a party
func (b *Battle) GetAllies(p *party) AgentParties {
	allies := make(AgentParties)
	for _, party := range b.parties {
		if party.team == p.team {
			allies[party.Agent] = party.GetActivePokemon()
		}
	}
	return allies
}

// Gets all active opponent Pokemon for a party
func (b *Battle) GetOpponents(p *party) AgentParties {
	opponents := make(AgentParties)
	for _, party := range b.parties {
		if party.team != p.team {
			opponents[party.Agent] = party.GetActivePokemon()
		}
	}
	return opponents
}

func (b *Battle) Start() error {
	// TODO: validate the battle, return error if invalid

	// Initiate the battle! Send out the first pokemon in the parties.
	b.State = BATTLE_IN_PROGRESS
	for _, party := range b.parties {
		party.SetActive(0)
	}
	return nil
}

// Simulates a single round of the battle.
func (b *Battle) SimulateRound() {
	// Collects all turn info from each active Pokemon
	turns := make([]TurnContext, 0)
	for _, party := range b.parties {
		for _, pokemon := range party.activePokemon {
			ctx := b.getContext(party, pokemon)
			turn := (*party.Agent).Act(ctx)
			turns = append(turns, TurnContext{Turn: &turn, Context: ctx})
		}
	}
	// Sort turns using an in-place stable sort
	sort.SliceStable(turns, func(i, j int) bool {
		turnA := turns[i].Turn
		turnB := turns[j].Turn
		ctxA := turns[i].Context
		ctxB := turns[j].Context
		if reflect.TypeOf(*turnA) == reflect.TypeOf(*turnB) {
			switch (*turnA).(type) {
			case FightTurn:
				// speedy pokemon should go first
				return ctxA.Pokemon.Stats[STAT_SPD] > ctxB.Pokemon.Stats[STAT_SPD]
			}
		} else {
			// make higher priority turns go first
			return (*turnA).Priority() > (*turnB).Priority()
		}
		// fallthrough
		return false
	})
	// Run turns in sorted order and update battle state
	for _, turn := range turns {
		switch t := (*turn.Turn).(type) {
		case FightTurn:
			log.Printf("TODO: Implement fight\n")
		default:
			log.Panicf("Unknown turn of type %v", t)
		}
	}
}

type BattleContext struct {
	Battle    Battle       // A copy of the current Battle, including weather, state, etc.
	Pokemon   Pokemon      // A copy of the Pokemon that is acting in this context
	Allies    AgentParties // Map of acting Pokemon's allied agents to their parties
	Opponents AgentParties // Map of acting Pokemon's opponent agents to their parties
}

// Gets the current context for a pokemon to act (perform a turn)
func (b *Battle) getContext(party *party, pokemon *Pokemon) *BattleContext {
	return &BattleContext{
		Battle:    *b,
		Pokemon:   *pokemon,
		Allies:    b.GetAllies(party),
		Opponents: b.GetOpponents(party),
	}
}

// An abstration over all possible actions an `Agent` can make in one round. Each Pokemon gets one turn.
type Turn interface {
	Priority() int // Gets the turn's priority. Higher values go first.
}

// Wrapper used to determine turn order in a battle
type TurnContext struct {
	Turn    *Turn          // A reference to the turn that a Pokemon made using an Agent
	Context *BattleContext // The context in which the Pokemon took its turn
}

type FightTurn struct {
	agent  *Agent // The agent which the move is targetting.
	move   int    // Denotes the index (0-3) of the pokemon's which of the pokemon's moves to use.
	target int    // The active pokemon of the agent on the receiving end of the move.
}

func (turn FightTurn) Priority() int {
	return 0
}
