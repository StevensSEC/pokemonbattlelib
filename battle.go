package pokemonbattlelib

import (
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
	turns := make([]Turn, 0)
	for _, party := range b.parties {
		for _, pokemon := range party.activePokemon {
			ctx := b.GetContext(party, pokemon)
			turn := (*party.Agent).Act(ctx)
			turns = append(turns, turn)
		}
	}
	sort.SliceStable(turns, func(i, j int) bool {
		tA := turns[i]
		tB := turns[j]
		if reflect.TypeOf(tA) == reflect.TypeOf(tB) {
			switch tA.(type) {
			case FightTurn:
				// speedy pokemon should go first
				return pA.Stats[5] > pB.Stats[5]
			}
		} else {
			// make higher priority turns go first
			return tA.Priority() > tB.Priority()
		}
		// fallthrough
		return false
	})
	// turnOrder := sortTurns(b, active, turns)
	// for _, apIdx := range turnOrder {
	// 	switch t := turns[apIdx].(type) {
	// 	case FightTurn:
	// 		fmt.Printf("TODO: Implement fight %v\n", t)
	// 	default:
	// 		panic("Unknown turn")
	// 	}
	// }
}

type Context struct {
	Battle    Battle
	Pokemon   Pokemon
	Allies    AgentParties
	Opponents AgentParties
}

// Gets the current context for a pokemon to act (perform a turn)
func (b *Battle) GetContext(party *party, pokemon *Pokemon) *Context {
	return &Context{
		Battle:    *b,
		Pokemon:   *pokemon,
		Allies:    b.GetAllies(party),
		Opponents: b.GetOpponents(party),
	}
}

// A type that is necessary in order to implement the `Interface` interface, which is used by the sort package to sort.
// type apTurnOrder struct {
// 	battle *Battle
// 	active []activePokemon
// 	turns  map[int]Turn
// 	order  []int
// }

/* func newTurnOrder(battle *Battle, ap []activePokemon, turns map[int]Turn) apTurnOrder {
	o := []int{}
	for i := range ap {
		o = append(o, i)
	}

	ord := apTurnOrder{
		battle: battle,
		order:  o,
		active: ap,
		turns:  turns,
	}

	return ord
}

func (t *apTurnOrder) GetOrder() []int {
	return t.order
}

func (t *apTurnOrder) Len() int {
	return len(t.order)
}

func (t *apTurnOrder) Swap(i, j int) {
	t.order[i], t.order[j] = t.order[j], t.order[i]
}

// Determine if turn a should happen before turn b.
func (t *apTurnOrder) Less(a, b int) bool {
	pA, pB := t.battle.derefActivePokemon(t.active[a]), t.battle.derefActivePokemon(t.active[b])
	tA, tB := t.turns[a], t.turns[b]
	if reflect.TypeOf(tA) == reflect.TypeOf(tB) {
		switch tA.(type) {
		case FightTurn:
			// speedy pokemon should go first
			return pA.Stats[5] > pB.Stats[5]
		}
	} else {
		// make higher priority turns go first
		return tA.Priority() > tB.Priority()
	}
	// fallthrough
	return false
}

// Returns the indexes of the active pokemon in the order that their turns should take place.
func sortTurns(battle *Battle, ap []activePokemon, turns map[int]Turn) []int {
	t := newTurnOrder(battle, ap, turns)
	sort.Sort(&t)
	return t.GetOrder()
}

// References a Pokemon currently on the battlefield.
type activePokemon struct {
	PartyIdx   int
	PokemonIdx int
} */

// An abstration over all possible actions an `Agent` can make in one round. Each Pokemon gets one turn.
type Turn interface {
	Priority() int // Gets the turn's priority. Higher values go first.
}

type FightTurn struct {
	agent  *Agent // The agent which the move is targetting.
	move   int    // Denotes the index (0-3) of the pokemon's which of the pokemon's moves to use.
	target int    // The active pokemon of the agent on the receiving end of the move.
}

func (turn FightTurn) Priority() int {
	return 0
}
