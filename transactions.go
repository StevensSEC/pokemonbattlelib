package pokemonbattlelib

//go:generate go run ./scripts/transaction_marshall/gen_transaction_marshall.go

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
	receiver := t.Target.Pokemon
	if receiver.CurrentHP >= t.Damage {
		receiver.CurrentHP -= t.Damage
	} else {
		// prevent underflow
		receiver.CurrentHP = 0
	}
	if receiver.CurrentHP == 0 {
		// Prevent OHKO with Focus Sash
		if receiver.HeldItem == ItemFocusSash {
			receiver.CurrentHP = 1
			b.QueueTransaction(ItemTransaction{
				Target: t.Target,
				IsHeld: true,
				Item:   receiver.HeldItem,
			})
			return
		}
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
		// EVs are gained based on EV yield of defeated Pokemon
		evGain := receiver.GetEVYield()
		for stat, amount := range evGain {
			if amount == 0 {
				continue
			}
			b.QueueTransaction(EVTransaction{
				Target: t.User,
				Stat:   stat,
				Amount: uint8(amount),
			})
		}
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

// A transaction to change the EVs of a Pokemon.
type EVTransaction struct {
	Target *Pokemon
	Stat   int
	Amount uint8
}

func (t EVTransaction) Mutate(b *Battle) {
	t.Target.EVs[t.Stat] += t.Amount
}

// A transaction to use and possibly consume an item.
type ItemTransaction struct {
	Target target
	IsHeld bool
	Item   Item
	Move   *Move
}

func (t ItemTransaction) Mutate(b *Battle) {
	target := t.Target.Pokemon
	if t.Item.Flags()&FlagConsumable > 0 {
		if t.IsHeld {
			t.Item = target.HeldItem // auto-correct if the value is not present or does not match
			target.HeldItem = ItemNone
		}
		// TODO: remove consumed item from party's inventory
	}
	switch t.Item.Category() {
	case ItemCategoryStatBoost:
		if t.Item == ItemDireHit {
			b.QueueTransaction(ModifyStatTransaction{
				Target: target,
				Stat:   StatCritChance,
				Stages: +2,
			})
		} else if t.Item == ItemGuardSpec {
			target.metadata[MetaStatChangeImmune] = 5
		} else {
			b.QueueTransaction(ModifyStatTransaction{
				Target: target,
				Stat:   battleItemStats[t.Item],
				Stages: +1,
			})
		}
		b.QueueTransaction(FriendshipTransaction{
			Target: target,
			Amount: [3]int{1, 1, 0}[target.Friendship/100],
		})
	}
	switch t.Item {
	// ItemCategoryMedicine
	case ItemPotion:
		b.QueueTransaction(target.RestoreHP(20))
	// ItemCategoryInAPinch
	case ItemApicotBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: target,
			Stat:   StatSpDef,
			Stages: 1,
		})
	case ItemCustapBerry:
		// TODO: Force pokemon to go first
	case ItemGanlonBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: target,
			Stat:   StatDef,
			Stages: 1,
		})
	case ItemLansatBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: target,
			Stat:   StatCritChance,
			Stages: 2,
		})
	case ItemLiechiBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: target,
			Stat:   StatAtk,
			Stages: 1,
		})
	case ItemMicleBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: target,
			Stat:   StatAccuracy,
			Stages: 1,
		})
	case ItemPetayaBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: target,
			Stat:   StatSpAtk,
			Stages: 1,
		})
	case ItemSalacBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: target,
			Stat:   StatSpeed,
			Stages: 1,
		})
	case ItemStarfBerry:
		// TODO: boost random stat, requires battle RNG to be available.
	// ItemCategoryHeldItems
	case ItemBlackSludge:
		if target.Type&TypePoison != 0 {
			b.QueueTransaction(HealTransaction{
				Target: target,
				Amount: target.MaxHP() / 16,
			})
		} else {
			b.QueueTransaction(DamageTransaction{
				Target: t.Target,
				Damage: target.MaxHP() / 8,
			})
		}
	case ItemLeftovers:
		b.QueueTransaction(HealTransaction{
			Target: target,
			Amount: target.MaxHP() / 16,
		})
	case ItemMentalHerb:
		b.QueueTransaction(CureStatusTransaction{
			Target:       t.Target,
			StatusEffect: StatusInfatuation,
		})
	case ItemWhiteHerb:
		for stat, stages := range target.StatModifiers {
			if stages < 0 {
				b.QueueTransaction(ModifyStatTransaction{
					Target: target,
					Stat:   stat,
					Stages: -stages,
				})
			}
		}
	// ItemCategoryBadHeldItems
	case ItemFlameOrb:
		b.QueueTransaction(InflictStatusTransaction{
			Target:       t.Target.Pokemon,
			StatusEffect: StatusBurn,
		})
	case ItemStickyBarb:
		b.QueueTransaction(DamageTransaction{
			Target: t.Target,
			Damage: target.MaxHP() / 8,
		})
	case ItemToxicOrb:
		b.QueueTransaction(InflictStatusTransaction{
			Target:       t.Target.Pokemon,
			StatusEffect: StatusBadlyPoison,
		})
	}
	// In a pinch consumption
	if target.HeldItem.Category() == ItemCategoryInAPinch && target.CurrentHP <= target.Stats[StatHP]/4 {
		b.QueueTransaction(ItemTransaction{
			Target: t.Target,
			IsHeld: true,
			Item:   target.HeldItem,
		})
	}
}

// A transaction to change the PP of a move.
type PPTransaction struct {
	Move   *Move
	Amount int8
}

func (t PPTransaction) Mutate(b *Battle) {
	// This is why we should just use int
	if t.Amount < 0 {
		n := uint8(t.Amount * -1)
		if t.Move.CurrentPP < n {
			n = t.Move.CurrentPP
		}
		t.Move.CurrentPP -= n
	} else {
		t.Move.CurrentPP += uint8(t.Amount)
		if t.Move.CurrentPP > t.Move.MaxPP {
			t.Move.CurrentPP = t.Move.MaxPP
		}
	}
}

// A transaction to change the held item of a Pokemon
type GiveItemTransaction struct {
	Target *Pokemon
	Item   Item
}

func (t GiveItemTransaction) Mutate(b *Battle) {
	t.Target.HeldItem = t.Item
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
	if t.StatusEffect.check(StatusSleep) {
		t.Target.metadata[MetaSleepTime] = b.rng.Get(1, 5)
	}
	if t.Target.Ability == AbilitySteadfast && t.StatusEffect.check(StatusFlinch) {
		b.QueueTransaction(ModifyStatTransaction{
			Target: t.Target,
			Stat:   StatSpeed,
			Stages: 1,
		})
	}
}

type CureStatusTransaction struct {
	Target       target
	StatusEffect StatusCondition
}

func (t CureStatusTransaction) Mutate(b *Battle) {
	t.Target.Pokemon.StatusEffects.clear(t.StatusEffect)
	if t.StatusEffect.check(StatusSleep) {
		delete(t.Target.Pokemon.metadata, MetaSleepTime)
	}
}

// A transaction that makes a pokemon faint, and returns the pokemon to the pokeball.
type FaintTransaction struct {
	Target target
}

func (t FaintTransaction) Mutate(b *Battle) {
	p := b.parties[t.Target.party]
	p.SetInactive(t.Target.partySlot)
	anyAlive := false
	for i, pkmn := range p.pokemon() {
		if pkmn.CurrentHP > 0 {
			anyAlive = true
			// TODO: prompt Agent for which pokemon to send out next
			// auto send out next pokemon
			b.QueueTransaction(SendOutTransaction{
				Target: target{
					Pokemon:   b.getPokemonInBattle(t.Target.party, i),
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
		b.QueueTransaction(EndBattleTransaction{
			Reason: EndKnockout,
			Winner: (p.team + 1) % 2, // HACK: because there is always 2 teams in a battle
		})
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
	Turns   int
}

func (t WeatherTransaction) Mutate(b *Battle) {
	b.Weather = t.Weather
	b.metadata[MetaWeatherTurns] = t.Turns
}

// A transaction that ends the battle.
type EndReason int

const (
	EndKnockout EndReason = iota
	EndForfeit
	EndFlee
)

type EndBattleTransaction struct {
	Reason EndReason
	Winner int
}

func (t EndBattleTransaction) Mutate(b *Battle) {
	b.State = BattleEnd
	b.results.Winner = t.Winner
	for _, p := range b.parties {
		b.results.Parties = append(b.results.Parties, p.Party)
	}
}

// Handles pre-turn status checks. (Paralysis, Sleeping, etc.)
type ImmobilizeTransaction struct {
	Target       target
	StatusEffect StatusCondition
}

func (t ImmobilizeTransaction) Mutate(b *Battle) {
	receiver := t.Target.Pokemon
	if t.StatusEffect.check(StatusSleep) {
		receiver.metadata[MetaSleepTime] = receiver.metadata[MetaSleepTime].(int) - 1
	}
}

// Handles evasion, misses, dodging, etc. when using moves
type MoveFailTransaction struct {
	User   *Pokemon
	Reason MoveFailReason
}

func (t MoveFailTransaction) Mutate(b *Battle) {
	// currently a no-op.
}

// Modifies a stat's stages in the interval [-6, 6]
type ModifyStatTransaction struct {
	Target *Pokemon
	Stat   int
	Stages int
}

func (t ModifyStatTransaction) Mutate(b *Battle) {
	stage := t.Target.StatModifiers[t.Stat] + t.Stages
	min := MinStatModifier
	max := MaxStatModifier
	// Bounds for crit chance are [0, 4]
	if t.Stat == StatCritChance {
		min = 0
		max = len(CritChances) - 1
	}
	if stage < min {
		stage = min
	}
	if stage > max {
		stage = max
	}
	t.Target.StatModifiers[t.Stat] = stage
}
