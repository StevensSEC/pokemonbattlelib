package pokemonbattlelib

import (
	"fmt"
)

type Battle struct {
	Weather      int  // one of the 6 in-battle weather conditions
	ShiftSet     bool // shift or set battle style for NPC trainer battles
	State        BattleState
	Parties      []*Party
	Agents       []*Agent
	agentParties map[int]int // Maps agent indexes to party indexes
	teams        [][]int     // An array of teams, which are arrays of Agents. Used to derive allies and opponents
}

type Party struct {
	Pokemon []*Pokemon
	active  []int // Which pokemon in the party are out on the battlefield
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

// Gives an Agent control over a Party.
func (b *Battle) LinkAgentParty(agentIdx, partyIdx int) error {
	if b.agentParties == nil {
		b.agentParties = make(map[int]int)
	}
	// TODO: check if agentIdx, partyIdx are valid indexes
	// TODO: check if agent is already linked
	// TODO: check if party is already linked
	b.agentParties[agentIdx] = partyIdx
	return nil
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

// Adds Agent(s) to the battle.
func (b *Battle) AddAgent(a ...*Agent) {
	b.Agents = append(b.Agents, a...)
}

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
	// get the currently active pokemon
	active := []activePokemon{}
	for a := range b.Agents {
		party := b.Parties[b.agentParties[a]]
		for _, idx := range party.GetActive() {
			ap := activePokemon{
				AgentIdx:   a,
				PokemonIdx: idx,
			}
			active = append(active, ap)
		}
	}

	context := battleContext{
		Context:       *b,
		ActivePokemon: active,
	}

	// get turns from agents
	// map of active pokemon indexes to their turns
	turns := map[int]Turn{}
	for i, ap := range active {
		context.SetSelf(i)
		agent := *b.Agents[ap.AgentIdx]
		turn := agent.Act(&context)
		turns[i] = turn
	}

	turnOrder := []int{}
	// TODO: determine the correct order based on priority and other factors
	for i := range active {
		turnOrder = append(turnOrder, i)
	}

	for _, apIdx := range turnOrder {
		switch t := turns[apIdx].(type) {
		case FightTurn:
			fmt.Printf("TODO: Implement fight %v\n", t)
		default:
			panic("Unknown turn")
		}
	}
}

type activePokemon struct {
	AgentIdx   int
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
	selfAgent := c.ActivePokemon[c.self].AgentIdx
	var team []int
	for _, t := range c.Context.teams {
		if contains(t, selfAgent) {
			team = t
			break
		}
	}
	return team
}
func (c *battleContext) GetPokemon(idx int) Pokemon {
	ap := c.ActivePokemon[idx]
	b := c.Context
	p := (*b.Parties[b.agentParties[ap.AgentIdx]]).Pokemon[ap.PokemonIdx]
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
		if i != c.self && contains(team, ap.AgentIdx) {
			allies = append(allies, i)
		}
	}

	return allies
}
func (c *battleContext) Opponents() []int {
	team := c.getTeam()

	opponents := []int{}
	for i, ap := range c.ActivePokemon {
		if !contains(team, ap.AgentIdx) {
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
