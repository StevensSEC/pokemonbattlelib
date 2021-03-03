package pokemonbattlelib

// Transactions describes a change to battle state.
// A sequence of transactions should be able to describe an entire battle.
type Transaction interface {
	Mutate(b *Battle) // Modifies the battle to apply the transaction. Can also queue additional transactions via b.QueueTransaction().
}

// A transaction to deal damage to an opponent Pokemon.
type DamageTransaction struct {
	User         *Pokemon
	Target       target
	Move         *Move
	Damage       uint
	StatusEffect StatusCondition
}

func (t DamageTransaction) Mutate(b *Battle) {
	// Minimum 1HP attack
	if t.Damage == 0 {
		t.Damage = 1
	}
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

func (t FriendshipTransaction) Mutate(b *Battle) {
	t.Target.Friendship += t.Amount
}

// A transaction to use and possibly consume an item.
type ItemTransaction struct {
	Target *Pokemon
	Item   *Item
	Move   *Move
}

func (t ItemTransaction) Mutate(b *Battle) {
	// TODO: do not consume certain items
	if t.Target.HeldItem == t.Item {
		t.Target.HeldItem = nil
	}
}

// A transaction to restore HP to a Pokemon.
type HealTransaction struct {
	Target *Pokemon
	Amount uint
}

func (t HealTransaction) Mutate(b *Battle) {
	t.Target.CurrentHP += t.Amount
}

// A transaction to apply a status effect to a Pokemon.
type InflictStatusTransaction struct {
	Target       *Pokemon
	StatusEffect StatusCondition
}

func (t InflictStatusTransaction) Mutate(b *Battle) {
	t.Target.StatusEffects.apply(t.StatusEffect)
}

type CureStatusTransaction struct {
	Target       target
	StatusEffect StatusCondition
}

func (t CureStatusTransaction) Mutate(b *Battle) {
	receiver := b.getPokemon(t.Target.party, t.Target.partySlot)
	receiver.StatusEffects.clear(t.StatusEffect)
}

// A transaction that makes a pokemon faint, and returns the pokemon to the pokeball.
type FaintTransaction struct {
	Target target
}

func (t FaintTransaction) Mutate(b *Battle) {
	p := b.parties[t.Target.party]
	p.SetInactive(t.Target.partySlot)
	anyAlive := false
	for i, pkmn := range p.pokemon {
		if pkmn.CurrentHP > 0 {
			anyAlive = true
			// TODO: prompt Agent for which pokemon to send out next
			// auto send out next pokemon
			b.QueueTransaction(SendOutTransaction{
				Target: target{
					Pokemon:   *b.getPokemon(t.Target.party, i),
					party:     t.Target.party,
					partySlot: i,
					Team:      t.Target.Team,
				},
			})
			break
		}
	}
	if !anyAlive {
		// cause the battle to end by knockout
		b.QueueTransaction(EndBattleTransaction{})
	}
}

// A transaction that makes a party send out a pokemon.
type SendOutTransaction struct {
	Target target
}

func (t SendOutTransaction) Mutate(b *Battle) {
	p := b.parties[t.Target.party]
	p.SetActive(t.Target.partySlot)
}

// Changes the current weather in a battle
type WeatherTransaction struct {
	Weather Weather
}

func (t WeatherTransaction) Mutate(b *Battle) {
	b.Weather = t.Weather
}

// A transaction that ends the battle.
type EndBattleTransaction struct{}

func (t EndBattleTransaction) Mutate(b *Battle) {
	b.State = BattleEnd
}

// Handles pre-turn status checks. (Paralysis, Sleeping, etc.)
type ImmobilizeTransaction struct {
	Target       target
	StatusEffect StatusCondition
}

func (t ImmobilizeTransaction) Mutate(b *Battle) {
	// currently a no-op.
}

// Handles evasion, misses, dodging, etc. when using moves
type EvadeTransaction struct {
	User *Pokemon
}

func (t EvadeTransaction) Mutate(b *Battle) {
	// currently a no-op.
}
