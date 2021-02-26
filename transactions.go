package pokemonbattlelib

import "fmt"

// Transactions describes a change to battle state.
// A sequence of transactions should be able to describe an entire battle.
type Transaction interface {
	BattleLog() string
	// Mutate(b *Battle) // Modifies the battle to apply the transaction. Can also queue additional transactions via b.QueueTransaction().
}

// A transaction to deal damage to an opponent Pokemon.
type DamageTransaction struct {
	User         *Pokemon
	Target       target
	Move         *Move
	Damage       uint
	StatusEffect StatusCondition
}

func (t DamageTransaction) BattleLog() string {
	if t.User != nil && t.Move != nil {
		return fmt.Sprintf("%s used %s on %s for %d damage.",
			t.User.GetName(),
			t.Move.Name,
			t.Target.Pokemon.GetName(),
			t.Damage,
		)
	} else if t.StatusEffect != StatusNone {
		return fmt.Sprintf("%s took %d damage from being %s.",
			t.Target.Pokemon.GetName(),
			t.Damage,
			t.StatusEffect,
		)
	} else {
		panic("I don't know how to log this DamageTransaction.")
	}
}

func (t DamageTransaction) Mutate(b *Battle) {
	receiver := b.getPokemon(t.Target.party, t.Target.partySlot)
	if receiver.CurrentHP >= t.Damage {
		receiver.CurrentHP -= t.Damage
	} else {
		// prevent underflow
		receiver.CurrentHP = 0
	}
	if receiver.CurrentHP == 0 {
		// pokemon has fainted
		b.QueueTransaction(FaintTransaction{
			Target: t.Target,
		})
		// friendship is lowered based on level difference
		levelGap := t.User.Level - receiver.Level
		loss := -1
		if levelGap >= 30 {
			if receiver.Friendship < 200 {
				loss = -5
			} else {
				loss = -10
			}
		}
		b.QueueTransaction(FriendshipTransaction{
			Target: receiver,
			Amount: loss,
		})
	}
}

// A transaction to change the friendship level of a Pokemon.
type FriendshipTransaction struct {
	Target *Pokemon // The target Pokemon
	Amount int      // The amount of friendship to increase/decrease
}

func (t FriendshipTransaction) BattleLog() string {
	return fmt.Sprintf("%s's friendship changed by %v.", t.Target.GetName(), t.Amount)
}

func (t FriendshipTransaction) Mutate(b *Battle) {
	t.Target.Friendship += t.Amount
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

// A transaction to apply a status effect to a Pokemon.
type InflictStatusTransaction struct {
	Target *Pokemon
	Status StatusCondition
}

func (t InflictStatusTransaction) BattleLog() string {
	// TODO: add status string representation
	return fmt.Sprintf("%s now has <STATUS: %d>!", t.Target.GetName(), t.Status)
}

type CureStatusTransaction struct {
	Target target
	Status StatusCondition
}

func (t CureStatusTransaction) BattleLog() string {
	return fmt.Sprintf("%s is no longer %s.", t.Target.Pokemon.GetName(), t.Status)
}

// A transaction that makes a pokemon faint, and returns the pokemon to the pokeball.
type FaintTransaction struct {
	Target target
}

func (t FaintTransaction) BattleLog() string {
	return fmt.Sprintf("%s fainted.",
		t.Target.Pokemon.GetName(),
	)
}

// A transaction that makes a party send out a pokemon.
type SendOutTransaction struct {
	Target target
}

func (t SendOutTransaction) BattleLog() string {
	return fmt.Sprintf("%s was sent out.",
		t.Target.Pokemon.GetName(),
	)
}

type EndBattleTransaction struct{}

func (t EndBattleTransaction) BattleLog() string {
	// TODO: include reason the battle ended
	return "The battle has ended."
}

// Handles pre-turn status checks. (Paralysis, Sleeping, etc.)
type ImmobilizeTransaction struct {
	Target       target
	StatusEffect StatusCondition
}

func (t ImmobilizeTransaction) BattleLog() string {
	return fmt.Sprintf("%s is %s and is unable to move.",
		t.Target.Pokemon.GetName(),
		t.Target.Pokemon.StatusEffects&NONVOLATILE_STATUS_MASK)
}

// Handles evasion, misses, dodging, etc. when using moves
type EvadeTransaction struct {
	User *Pokemon
}

func (t EvadeTransaction) BattleLog() string {
	return fmt.Sprintf("%s's attack missed!", t.User.GetName())
}
