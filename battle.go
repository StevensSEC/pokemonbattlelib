package pokemonbattlelib

import (
	"fmt"
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

// Adds one or more parties to a team in the battle
func (b *Battle) AddParty(p ...*party) {
	b.parties = append(b.parties, p...)
}

// Gets a reference to a Pokemon using party ID and party slot
func (b *Battle) getPokemon(party, slot int) *Pokemon {
	if party >= len(b.parties) {
		panic(PartyIndexError)
	}
	p := b.parties[party].pokemon
	if slot >= len(p) {
		panic(PartyIndexError)
	}
	return p[slot]
}

// Gets all active ally Pokemon for a party
func (b *Battle) GetAllies(p *party) []target {
	allies := make([]target, 0)
	targets := b.GetTargets()
	for _, target := range targets {
		if target.Team == p.team {
			allies = append(allies, target)
		}
	}
	return allies
}

// Gets all active opponent Pokemon for a party
func (b *Battle) GetOpponents(p *party) []target {
	opponents := make([]target, 0)
	targets := b.GetTargets()
	for _, target := range targets {
		if target.Team != p.team {
			opponents = append(opponents, target)
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
func (b *Battle) SimulateRound() []Transaction {
	if b.State != BATTLE_IN_PROGRESS {
		log.Panic("battle is not currently in progress")
	}
	// Collects all turn info from each active Pokemon
	turns := make([]TurnContext, 0)
	for _, party := range b.parties {
		for _, pokemon := range party.activePokemon {
			ctx := b.getContext(party, pokemon)
			turn := (*party.Agent).Act(ctx)
			turns = append(turns, TurnContext{Turn: turn, Context: ctx})
		}
	}
	// Sort turns using an in-place stable sort
	sort.SliceStable(turns, func(i, j int) bool {
		turnA := turns[i].Turn
		turnB := turns[j].Turn
		ctxA := turns[i].Context
		ctxB := turns[j].Context
		if reflect.TypeOf(turnA) == reflect.TypeOf(turnB) {
			switch turnA.(type) {
			case FightTurn:
				// speedy pokemon should go first
				return ctxA.Pokemon.Stats[STAT_SPD] > ctxB.Pokemon.Stats[STAT_SPD]
			}
		} else {
			// make higher priority turns go first
			return turnA.Priority() > turnB.Priority()
		}
		// fallthrough
		return false
	})
	// Run turns in sorted order and update battle state
	transactions := []Transaction{}
	for _, turn := range turns {
		queue := []Transaction{}
		switch t := turn.Turn.(type) {
		case FightTurn:
			user := turn.Context.Pokemon
			target := t.Target
			receiver := b.getPokemon(t.Target.party, t.Target.partySlot)
			// See: https://github.com/StevensSEC/pokemonbattlelib/wiki/Requirements#fight-using-a-move
			modifier := uint(1) // TODO: damage multiplers
			damage := (((2*uint(user.Level)/5)+2)*uint(user.Moves[t.Move].Power)*user.Stats[STAT_ATK]/receiver.Stats[STAT_DEF]/50 + 2) * modifier
			queue = append(queue, DamageTransaction{
				User:            &user,
				Target:          receiver,
				TargetParty:     target.party,
				TargetPartySlot: target.partySlot,
				Move:            user.Moves[t.Move],
				Damage:          damage,
			})
		case ItemTurn:
			receiver := b.getPokemon(t.Target.party, t.Target.partySlot)
			move := receiver.Moves[t.Move]
			queue = append(queue, ItemTransaction{
				Target: receiver,
				Item:   t.Item,
				Move:   move,
			})
			queue = append(queue, receiver.UseItem(t.Item)...)
		default:
			log.Panicf("Unknown turn of type %v", t)
		}
		// process transations for this turn
		for len(queue) > 0 {
			next := queue[0]
			queue = queue[1:]
			switch t := next.(type) {
			case DamageTransaction:
				if t.Target.CurrentHP >= t.Damage {
					t.Target.CurrentHP -= t.Damage
				} else {
					// prevent underflow
					t.Target.CurrentHP = 0
				}
				if t.Target.CurrentHP == 0 {
					// pokemon has fainted
					queue = append(queue, FaintTransaction{
						Target:          t.Target,
						TargetParty:     t.TargetParty,
						TargetPartySlot: t.TargetPartySlot,
					})
				}
			case ItemTransaction:
				// TODO: handle held items/consumables
			case HealTransaction:
				amount := t.Amount
				if diff := t.Target.Stats[STAT_HP] - t.Target.CurrentHP; diff <= amount {
					amount = diff
				}
				t.Target.CurrentHP += amount
			case FaintTransaction:
				p := b.parties[t.TargetParty]
				p.SetInactive(t.TargetPartySlot)
				anyAlive := false
				for i, pkmn := range p.pokemon {
					if pkmn.CurrentHP > 0 {
						anyAlive = true
						// TODO: prompt Agent for which pokemon to send out next
						// auto send out next pokemon
						queue = append(queue, SendOutTransaction{
							Target:          b.getPokemon(t.TargetParty, i),
							TargetParty:     t.TargetParty,
							TargetPartySlot: i,
						})
						break
					}
				}
				if !anyAlive {
					// TODO: cause the battle to end by knockout
				}
			case SendOutTransaction:
				p := b.parties[t.TargetParty]
				p.SetActive(t.TargetPartySlot)
			}
			// add to the list of processed transactions
			transactions = append(transactions, next)
		}
	}
	return transactions
}

type target struct {
	party     int     // Identifier for a party (index in battle parties, or "party ID")
	partySlot int     // The slot of the active Pokemon
	Team      int     // The team that the Pokemon belongs to
	Pokemon   Pokemon // Pokemon that is a candidate target
}

type BattleContext struct {
	Battle    Battle   // A copy of the current Battle, including weather, state, etc.
	Pokemon   Pokemon  // A copy of the Pokemon that is acting in this context
	Team      int      // The team of the acting Pokemon
	Allies    []target // Targets that are allies of the acting Pokemon
	Opponents []target // Targets that are opponents of the acting Pokemon
	Targets   []target // An array of all possible targets that the Pokemon can act on
}

// Gets all the active Pokemon (targets) in the battle
func (b *Battle) GetTargets() []target {
	targets := make([]target, 0)
	for partyID, party := range b.parties {
		for slot, active := range party.activePokemon {
			target := target{
				party:     partyID,
				partySlot: slot,
				Team:      party.team,
				Pokemon:   *active,
			}
			targets = append(targets, target)
		}
	}
	return targets
}

// Gets the current context for a pokemon to act (perform a turn)
func (b *Battle) getContext(party *party, pokemon *Pokemon) *BattleContext {
	return &BattleContext{
		Battle:    *b,
		Pokemon:   *pokemon,
		Team:      party.team,
		Allies:    b.GetAllies(party),
		Opponents: b.GetOpponents(party),
		Targets:   b.GetTargets(),
	}
}

// An abstration over all possible actions an `Agent` can make in one round. Each Pokemon gets one turn.

type Turn interface {
	Priority() int // Gets the turn's priority. Higher values go first. Not to be confused with Move priority.
}

// Wrapper used to determine turn order in a battle
type TurnContext struct {
	Turn    Turn           // A copy of the turn that a Pokemon made using an Agent
	Context *BattleContext // The context in which the Pokemon took its turn
}

type FightTurn struct {
	Move   int    // Denotes the index (0-3) of the pokemon's which of the pokemon's moves to use.
	Target target // Info containing data determining the target of
}

func (turn FightTurn) Priority() int {
	return 0
}

// An item turn has the a higher priority than any move.
type ItemTurn struct {
	Move   int    // Denotes the index (0-3) of the pokemon's which of the pokemon's moves to use.
	Target target // Info containing data determining the target of
	Item   *Item  // Which item is being consumed
}

func (turn ItemTurn) Priority() int {
	return 1
}

// Describes a change to battle state.
type Transaction interface {
	BattleLog() string
}

// A transaction to deal damage to an opponent Pokemon.
type DamageTransaction struct {
	User            *Pokemon
	Target          *Pokemon
	TargetParty     int
	TargetPartySlot int
	Move            *Move
	Damage          uint
}

func (t DamageTransaction) BattleLog() string {
	return fmt.Sprintf("%s used %s on %s for %d damage.",
		t.User.GetName(),
		t.Move.Name,
		t.Target.GetName(),
		t.Damage,
	)
}

// A transaction to use and possibly consume an item.
type ItemTransaction struct {
	Target *Pokemon
	Item   *Item
	Move   *Move
}

func (t ItemTransaction) BattleLog() string {
	return fmt.Sprintf("%s used on %s.", t.Item.Name, t.Target.GetName())
}

// A transaction to restore HP to a Pokemon.
type HealTransaction struct {
	Target *Pokemon
	Amount uint
}

func (t HealTransaction) BattleLog() string {
	return fmt.Sprintf("%s restored %d HP.", t.Target.GetName(), t.Amount)
}

// A transaction that makes a pokemon faint, and returns the pokemon to the pokeball.
type FaintTransaction struct {
	Target          *Pokemon
	TargetParty     int
	TargetPartySlot int
}

func (t FaintTransaction) BattleLog() string {
	return fmt.Sprintf("%s fainted.",
		t.Target.GetName(),
	)
}

// A transaction that makes a party send out a pokemon.
type SendOutTransaction struct {
	Target          *Pokemon
	TargetParty     int
	TargetPartySlot int
}

func (t SendOutTransaction) BattleLog() string {
	return fmt.Sprintf("%s was sent out.",
		t.Target.GetName(),
	)
}
