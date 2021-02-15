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
	Parties  []*Party
	teams    [][]int // An array of teams, which are arrays of Party ids. Used to derive allies and opponents
}

// A Pokemon party. Can hold up to 6 Pokemon. Also manages how many pokemon are out on the battlefield.
type Party struct {
	Pokemon []*Pokemon
	active  []int  // Which pokemon in the party are out on the battlefield
	Agent   *Agent // The agent that has control over this party
}

func (p *Party) AddPokemon(pkmn ...*Pokemon) {
	p.Pokemon = append(p.Pokemon, pkmn...)
}

// Set the indeces of which pokemon are on the battlefield.
func (p *Party) SetActive(idx int) {
	if len(p.active) == 0 {
		p.active = append(p.active, idx)
	} else {
		p.active[0] = idx
	}
}

// Get the indeces of which pokemon are on the battlefield.
func (p *Party) GetActive() []int {
	return p.active
}

// Set which agents have which allies. Not to be confused with `Party`.
func (b *Battle) SetTeams(t [][]int) error {
	// TODO: validate
	b.teams = t
	return nil
}

type BattleState int

const (
	BEFORE_START BattleState = iota
	BATTLE_IN_PROGRESS
	BATTLE_END
)

// Adds Parties to the battle.
func (b *Battle) AddParty(p ...*Party) {
	b.Parties = append(b.Parties, p...)
}

func (b *Battle) Start() error {
	// TODO: validate the battle, return error if invalid

	// Initiate the battle! Send out the first pokemon in the parties.
	b.State = BATTLE_IN_PROGRESS
	for _, p := range b.Parties {
		p.SetActive(0)
	}
	return nil
}

// Simulates a single round of the battle.
func (b *Battle) SimulateRound() {
	active := b.getActivePokemon()

	context := battleContext{
		Context:       *b,
		ActivePokemon: active,
	}

	// get turns from agents
	// map of active pokemon indexes to their turns
	turns := map[int]Turn{}
	for i, ap := range active {
		context.SetSelf(i)
		party := *b.Parties[ap.PartyIdx]
		turn := (*party.Agent).Act(&context)
		turns[i] = turn
	}

	turnOrder := sortTurns(b, active, turns)

	for _, apIdx := range turnOrder {
		switch t := turns[apIdx].(type) {
		case FightTurn:
			user := b.derefActivePokemon(active[apIdx])
			target := b.derefActivePokemon(active[t.targetIdx])
			// See: https://github.com/StevensSEC/pokemonbattlelib/wiki/Requirements#fight-using-a-move
			modifier := uint(1) // TODO: damage multiplers
			damage := (((2*uint(user.Level)/5)+2)*uint(user.Moves[t.moveIdx].Power)*user.Stats[STAT_ATK]/target.Stats[STAT_DEF]/50 + 2) * modifier
			(*target).CurrentHP -= damage
		default:
			panic("Unknown turn")
		}
	}
}

// Get references to all Pokemon that are active on the battlefield.
func (b *Battle) getActivePokemon() []activePokemon {
	active := []activePokemon{}
	for p, party := range b.Parties {
		for _, idx := range party.GetActive() {
			ap := activePokemon{
				PartyIdx:   p,
				PokemonIdx: idx,
			}
			active = append(active, ap)
		}
	}
	return active
}

// Get a pointer to the actual Pokemon that `ap` is referencing.
func (b *Battle) derefActivePokemon(ap activePokemon) *Pokemon {
	return (*b.Parties[ap.PartyIdx]).Pokemon[ap.PokemonIdx]
}

// A type that is necessary in order to implement the `Interface` interface, which is used by the sort package to sort.
type apTurnOrder struct {
	battle *Battle
	active []activePokemon
	turns  map[int]Turn
	order  []int
}

func newTurnOrder(battle *Battle, ap []activePokemon, turns map[int]Turn) apTurnOrder {
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
}

// Used to provide data to Agents such that the Agent can't maliciously modify the state of the battle.
type BattleInfo interface {
	GetPokemon(int) Pokemon
	Self() int
	Allies() []int
	Opponents() []int
}

// Implements the BattleInfo interface. Passed to Agents so they can decide a turn.
type battleContext struct {
	Context       Battle
	ActivePokemon []activePokemon
	self          int
}

func (c *battleContext) getTeam() []int {
	selfParty := c.ActivePokemon[c.self].PartyIdx
	var team []int
	for _, t := range c.Context.teams {
		if contains(t, selfParty) {
			team = t
			break
		}
	}
	return team
}
func (c *battleContext) GetPokemon(idx int) Pokemon {
	p := c.Context.derefActivePokemon(c.ActivePokemon[idx])
	return *p
}
func (c *battleContext) SetSelf(idx int) {
	c.self = idx
}

// Get the index of the active Pokemon that this battle context is referencing.
// In other words, the Pokemon you are acting as.
func (c *battleContext) Self() int {
	return c.self
}
func (c *battleContext) Allies() []int {
	team := c.getTeam()

	allies := []int{}
	for i, ap := range c.ActivePokemon {
		if i != c.self && contains(team, ap.PartyIdx) {
			allies = append(allies, i)
		}
	}

	return allies
}
func (c *battleContext) Opponents() []int {
	team := c.getTeam()

	opponents := []int{}
	for i, ap := range c.ActivePokemon {
		if !contains(team, ap.PartyIdx) {
			opponents = append(opponents, i)
		}
	}

	return opponents
}

// An abstration over all possible actions an `Agent` can make in one round. Each Pokemon gets one turn.
type Turn interface {
	Priority() int // Gets the turn's priority. Higher values go first.
}

// Indicates that the Pokemon will fight. The active Pokemon indicated by `Self()` in the battle context will use
// the move at `moveIdx` on the active Pokemon indicated by `targetIdx`.
type FightTurn struct {
	moveIdx   int // Denotes which of the pokemon's moves to use.
	targetIdx int // The active pokemon that on the receiving end of the move.
}

func (turn FightTurn) Priority() int {
	return 0
}
