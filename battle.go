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
		switch t := turn.Turn.(type) {
		case FightTurn:
			user := turn.Context.Pokemon
			receiver := b.getPokemon(t.Target.party, t.Target.partySlot)
			// See: https://github.com/StevensSEC/pokemonbattlelib/wiki/Requirements#fight-using-a-move
			modifier := uint(1) // TODO: damage multiplers
			damage := (((2*uint(user.Level)/5)+2)*uint(user.Moves[t.Move].Power)*user.Stats[STAT_ATK]/receiver.Stats[STAT_DEF]/50 + 2) * modifier
			transactions = append(transactions, DamageTransaction{
				User:   &user,
				Target: receiver,
				Move:   user.Moves[t.Move],
				Damage: damage,
			})
		case ItemTurn:
			receiver := b.getPokemon(t.Target.party, t.Target.partySlot)
			move := receiver.Moves[t.Move]
			transactions = append(transactions, ItemTransaction{
				Target: receiver,
				Item:   t.Item,
				Move:   move,
			})
			if t.Item.Category == ItemPPRecovery {
				transactions = append(transactions, receiver.UseMoveItem(t.Item, receiver.Moves[t.Move])...)
			} else {
				transactions = append(transactions, receiver.UseItem(t.Item)...)
			}
		default:
			log.Panicf("Unknown turn of type %v", t)
		}
	}
	// process transations
	for _, transaction := range transactions {
		switch t := transaction.(type) {
		case DamageTransaction:
			t.Target.CurrentHP -= t.Damage
		case ItemTransaction:
			// TODO: do not consume certain items
			if t.Target.HeldItem == t.Item {
				t.Target.HeldItem = nil
			}
		case HealTransaction:
			t.Target.CurrentHP += t.Amount
		case PPTransaction:
			t.Move.CurrentPP += t.Amount
		case CureStatusTransaction:
			t.Target.StatusEffects &= ^t.StatusEffects
		case EVTransaction:
			t.Target.EVs[t.Stat] += t.Amount
		case FriendshipTransaction:
			t.Target.Friendship += t.Amount
		case ModifyStatTransaction:
			t.Target.StatModifiers[t.Stat] += t.Stages
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

type DamageTransaction struct {
	User          *Pokemon
	Target        *Pokemon
	Move          *Move
	Damage        uint
	StatusEffects uint
}

func (t DamageTransaction) BattleLog() string {
	return fmt.Sprintf("%s used %s on %s for %d damage.",
		t.User.GetName(),
		t.Move.Name,
		t.Target.GetName(),
		t.Damage,
	)
}

type ItemTransaction struct {
	Target *Pokemon
	Item   *Item
	Move   *Move
}

func (t ItemTransaction) BattleLog() string {
	return fmt.Sprintf("%s used on %s.", t.Item.Name, t.Target.GetName())
}

type HealTransaction struct {
	Target *Pokemon
	Amount uint
}

func (t HealTransaction) BattleLog() string {
	return fmt.Sprintf("%s restored %d HP.", t.Target.GetName(), t.Amount)
}

type PPTransaction struct {
	Move   *Move
	Amount int
}

func (t PPTransaction) BattleLog() string {
	return fmt.Sprintf("%s restored %d PP.", t.Move.Name, t.Amount)
}

type EVTransaction struct {
	Target *Pokemon
	Stat   int
	Amount uint8
}

func (t EVTransaction) BattleLog() string {
	// TODO: add stat string representation
	return fmt.Sprintf("%s gained %d <STAT> effort.", t.Target.GetName(), t.Amount)
}

type FriendshipTransaction struct {
	Target *Pokemon
	Amount uint8
}

func (t FriendshipTransaction) BattleLog() string {
	return fmt.Sprintf("%s gained %d friendship.", t.Target.GetName(), t.Amount)
}

type CureStatusTransaction struct {
	Target        *Pokemon
	StatusEffects uint
}

func (t CureStatusTransaction) BattleLog() string {
	// TODO: add status string representation
	return fmt.Sprintf("%s's <STATUS> was cured.", t.Target.GetName())
}

type ModifyStatTransaction struct {
	Target *Pokemon
	Stat   int
	Stages int
}

func (t ModifyStatTransaction) BattleLog() string {
	// TODO: add stat string representation
	msg := "<STAT> "
	if t.Stages == 1 {
		msg += "rose!"
	} else if t.Stages >= 2 {
		msg += "sharply rose!"
	} else if t.Stages == -1 {
		msg += "fell!"
	} else if t.Stages <= -2 {
		msg += "harshly fell!"
	} else {
		msg += "???"
	}
	return fmt.Sprintf("%s's %s.", t.Target.GetName(), msg)
}
