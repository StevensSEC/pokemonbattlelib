package pokemonbattlelib

import "fmt"

// Transactions describes a change to battle state.
// A sequence of transactions should be able to describe an entire battle.
type Transaction interface {
	BattleLog() string
}

// Handler for all transactions to be used during a battle
func (b *Battle) ProcessTransaction(next Transaction) {
	switch t := next.(type) {
	case DamageTransaction:
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
		}
	case ItemTransaction:
		// TODO: do not consume certain items
		if t.Target.HeldItem == t.Item {
			t.Target.HeldItem = nil
		}
	case HealTransaction:
		t.Target.CurrentHP += t.Amount
	case FaintTransaction:
		p := b.parties[t.Target.party]
		p.SetInactive(t.Target.partySlot)
		anyAlive := false
		for i, pkmn := range p.pokemon {
			if pkmn.CurrentHP > 0 {
				anyAlive = true
				// TODO: prompt Agent for which pokemon to send out next
				// auto send out next pokemon
				target := target{
					Pokemon:   *b.getPokemon(t.Target.party, i),
					Team:      p.team,
					party:     t.Target.party,
					partySlot: t.Target.partySlot,
				}
				b.QueueTransaction(SendOutTransaction{
					Target: target,
				})
				break
			}
		}
		if !anyAlive {
			// cause the battle to end by knockout
			b.QueueTransaction(EndBattleTransaction{})
		}
	case SendOutTransaction:
		p := b.parties[t.Target.party]
		p.SetActive(t.Target.partySlot)
	case EndBattleTransaction:
		b.State = BATTLE_END
	}
	// add to the list of processed transactions
	b.transactionsProcessed = append(b.transactionsProcessed, next)
}

// A transaction to deal damage to an opponent Pokemon.
type DamageTransaction struct {
	User   *Pokemon
	Target target
	Move   *Move
	Damage uint
}

func (t DamageTransaction) BattleLog() string {
	return fmt.Sprintf("%s used %s on %s for %d damage.",
		t.User.GetName(),
		t.Move.Name,
		t.Target.Pokemon.GetName(),
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
