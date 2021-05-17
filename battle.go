package pokemonbattlelib

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

// A Pokemon battle. Enforces rules of the battle, and queries `Agent`s for turns.
type Battle struct {
	Weather  Weather // one of the 6 in-battle weather conditions
	ShiftSet bool    // shift or set battle style for NPC trainer battles
	State    BattleState
	rng      RNG
	ruleset  BattleRule

	parties  []*battleParty             // All parties participating in the battle
	metadata map[BattleMeta]interface{} // Metadata to be tracked during a battle

	tQueue     []Transaction
	tProcessed []Transaction
	results    BattleResults
}

type BattleRule int

const (
	BattleRuleFaint BattleRule = 1 << iota
	BattleRuleStruggle
)
const BattleRuleSetDefault = BattleRuleFaint | BattleRuleStruggle

type BattleMeta int

const (
	MetaWeatherTurns BattleMeta = iota
)

type BattleState int

const (
	BattleBeforeStart BattleState = iota
	BattleInProgress
	BattleEnd
)

// Creates a new battle instance, setting initial conditions
func NewBattle() *Battle {
	rng := LCRNG(rand.Uint32())
	b := Battle{
		State:   BattleBeforeStart,
		rng:     RNG(&rng),
		ruleset: BattleRuleSetDefault,
		metadata: map[BattleMeta]interface{}{
			MetaWeatherTurns: 0,
		},
	}
	return &b
}

// Sets the seed of the underlying random number generator used for the battle.
func (b *Battle) SetSeed(seed uint) {
	b.rng.SetSeed(seed)
}

// Add a party to the battle, controlled by an agent. This method is preferred over `AddBattleParty`.
func (b *Battle) AddParty(p *Party, a *Agent, team int) {
	b.AddBattleParty(&battleParty{
		Party:         p,
		Agent:         a,
		activePokemon: make(map[uint]uint),
		team:          team,
	})
}

// Adds one or more parties to a team in the battle
func (b *Battle) AddBattleParty(p ...*battleParty) {
	b.parties = append(b.parties, p...)
}

// Gets the battle party for a given target
func (b *Battle) GetParty(t target) *battleParty {
	return b.parties[t.party]
}

// Gets a reference to a Pokemon from a target
func (b *Battle) getPokemon(t target) *Pokemon {
	if t.party >= uint(len(b.parties)) {
		panic(ErrorPartyIndex)
	}
	party := b.parties[t.party]
	pokemon := party.pokemon()
	if t.slot >= uint(len(pokemon)) {
		panic(ErrorPartyIndex)
	}
	slot := t.slot
	if party.IsActivePokemon(t.slot) {
		slot = party.activePokemon[t.slot]
	}
	return pokemon[slot]
}

// Gets all the active Pokemon (targets) in the battle
func (b *Battle) AllTargets() []target {
	targets := make([]target, 0)
	for party, p := range b.parties {
		for slot := range p.activePokemon {
			t := target{
				party: uint(party),
				slot:  slot,
			}
			targets = append(targets, t)
		}
	}
	return targets
}

// Gets all active ally Pokemon for a party
func (b *Battle) getAllies(p *battleParty) []target {
	allies := make([]target, 0)
	targets := b.AllTargets()
	for _, target := range targets {
		party := b.parties[target.party]
		if party.team == p.team {
			allies = append(allies, target)
		}
	}
	return allies
}

// Gets all active opponent Pokemon for a party
func (b *Battle) getOpponents(p *battleParty) []target {
	opponents := make([]target, 0)
	targets := b.AllTargets()
	for _, target := range targets {
		party := b.parties[target.party]
		if party.team != p.team {
			opponents = append(opponents, target)
		}
	}
	return opponents
}

// Start the battle.
func (b *Battle) Start() error {
	// validate team count
	teams := map[int]int{}
	for i, party := range b.parties {
		if len(party.pokemon()) == 0 {
			return fmt.Errorf("Party (index: %d) has no pokemon.", i)
		}
		teams[party.team]++
	}
	if len(teams) != 2 {
		return fmt.Errorf("Parties have invalid teams. There should be 2 teams with 1 party each, got %d teams", len(teams))
	}

	// Initiate the battle! Send out the first pokemon in the parties.
	b.State = BattleInProgress
	for _, party := range b.parties {
		party.SetActive(0)
	}
	return nil
}

// Targets act as a composite key to determine which Pokemon is being referenced.
// It composes of the party (index into the battle's parties),
// and the slot of an active Pokemon within that party.
type target struct {
	party uint // Identifier for a party (index in battle parties, or "party ID")
	slot  uint // The slot of the active Pokemon
}

func (t target) String() string {
	return fmt.Sprintf("target{%d, %d}", t.party, t.slot)
}

func (t target) MarshalJSON() ([]byte, error) {
	type alias target // required to not enter infinite recursive loop
	return json.Marshal(&struct {
		Party uint
		Slot  uint
		*alias
	}{
		Party: t.party,
		Slot:  t.slot,
		alias: (*alias)(&t),
	})
}

func (t *target) UnmarshalJSON(data []byte) error {
	type alias target // required to not enter infinite recursive loop
	aux := &struct {
		Party uint
		Slot  uint
		*alias
	}{
		alias: (*alias)(t),
	}
	err := json.Unmarshal(data, &aux)
	t.party = aux.Party
	t.slot = aux.Slot
	if err != nil {
		return err
	}
	return nil
}

type AgentTarget struct {
	target          // Inherit party/slot from `target`
	Team    int     // The team that the Pokemon belongs to
	Pokemon Pokemon // Copy of Pokemon for Agents to use
}

type BattleContext struct {
	Battle Battle // A copy of the current Battle, including weather, state, etc.
	Team   int    // The team of the acting Pokemon
	target target // The party/slot of the acting Pokemon
}

func (b *Battle) GetBattleContext(t target) *BattleContext {
	p := b.parties[t.party]
	return &BattleContext{
		Battle: *b,
		Team:   p.team,
		target: t,
	}
}

func (bc *BattleContext) Self() AgentTarget {
	return AgentTarget{
		target:  bc.target,
		Team:    bc.Team,
		Pokemon: *bc.Battle.getPokemon(bc.target), // Make this a copy!
	}
}

func (bc *BattleContext) Allies() []AgentTarget {
	b := bc.Battle
	p := b.parties[bc.target.party]
	targets := make([]AgentTarget, 0)
	for _, t := range b.getAllies(p) {
		targets = append(targets, AgentTarget{
			target:  t,
			Team:    b.parties[t.party].team,
			Pokemon: *b.getPokemon(t),
		})
	}
	return targets
}

func (bc *BattleContext) Opponents() []AgentTarget {
	b := bc.Battle
	p := b.parties[bc.target.party]
	targets := make([]AgentTarget, 0)
	for _, t := range b.getOpponents(p) {
		targets = append(targets, AgentTarget{
			target:  t,
			Team:    b.parties[t.party].team,
			Pokemon: *b.getPokemon(t),
		})
	}
	return targets
}

func (bc *BattleContext) Targets() []AgentTarget {
	b := bc.Battle
	targets := make([]AgentTarget, 0)
	for _, t := range b.AllTargets() {
		targets = append(targets, AgentTarget{
			target:  t,
			Team:    b.parties[t.party].team,
			Pokemon: *b.getPokemon(t),
		})
	}
	return targets
}

// Get the results of the battle. The battle must be in the `BattleEnd` state.
func (b *Battle) GetResults() BattleResults {
	if b.State != BattleEnd {
		blog.Panic("Unable to get results of a battle that has not ended.")
	}
	return b.results
}

// Results for a Battle.
type BattleResults struct {
	Winner  int // The team that won the battle.
	Parties []*Party
}

// An abstraction over all possible actions an `Agent` can make in one round. Each Pokemon gets one turn.
type Turn interface {
	Priority() int // Gets the turn's priority. Higher values go first. Not to be confused with Move priority.
}

// Wrapper used to determine turn order in a battle
type TurnContext struct {
	User target // The pokemon that made this turn.
	Turn Turn   // A copy of the turn that a Pokemon made using an Agent
}

// A turn to represent a Pokemon using a Move.
type FightTurn struct {
	Move   int         // Denotes the index (0-3) of the pokemon's which of the pokemon's moves to use.
	Target AgentTarget // Info containing data determining the target of
}

func (turn FightTurn) Priority() int {
	return 0
}

// A turn to represent using an item from the Party's inventory. An item turn has the a higher priority than any move.
type ItemTurn struct {
	Move   int         // Denotes the index (0-3) of the pokemon's which of the pokemon's moves to use.
	Target AgentTarget // Info containing data determining the target of
	Item   Item        // Which item is being consumed
}

func (turn ItemTurn) Priority() int {
	return 1
}

// A turn to represent switching an active Pokemon for a different, inactive Pokemon in battle.
type SwitchTurn struct {
	Target AgentTarget // The target to swap to
}

func (turn SwitchTurn) Priority() int {
	return 2
}
