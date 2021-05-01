package pokemonbattlelib

//go:generate go run ./scripts/transaction_marshall/gen_transaction_marshall.go

// Transactions describes a change to battle state.
// A sequence of transactions should be able to describe an entire battle.
type Transaction interface {
	Mutate(b *Battle) // Modifies the battle to apply the transaction. Can also queue additional transactions via b.QueueTransaction().
}

type UseMoveTransaction struct {
	User   target
	Target target
	Move   *Move
}

func (t UseMoveTransaction) Mutate(b *Battle) {
	user := b.getPokemon(t.User)
	receiver := b.getPokemon(t.Target)
	accuracy := CalcAccuracy(b.Weather, user, receiver, t.Move)
	b.QueueTransaction(PPTransaction{
		Move:   t.Move,
		Amount: -1,
	})
	if t.Move.Accuracy() != 0 && !b.rng.Roll(int(accuracy), 100) {
		b.QueueTransaction(MoveFailTransaction{
			User:   t.User,
			Reason: FailMiss,
		})
		return
	}
	// See: https://github.com/StevensSEC/pokemonbattlelib/wiki/Requirements#fight-using-a-move
	// Status Moves
	if t.Move.Category() == MoveCategoryStatus {
		switch t.Move.Id {
		case MoveStunSpore:
			b.QueueTransaction(InflictStatusTransaction{
				Target:       t.Target,
				StatusEffect: StatusParalyze,
			})
		case MoveSpite:
			if m := receiver.metadata[MetaLastMove]; m != nil {
				b.QueueTransaction(PPTransaction{
					Move:   m.(*Move),
					Amount: -4,
				})
			}
		case MoveAttract:
			g1, g2 := user.Gender, receiver.Gender
			// Only applies when Pokemon are opposite gender
			if g1 != GenderGenderless && g2 != GenderGenderless && g1 != g2 {
				b.QueueTransaction(InflictStatusTransaction{
					Target:       t.Target,
					StatusEffect: StatusInfatuation,
				})
				if receiver.HeldItem == ItemDestinyKnot {
					b.QueueTransaction(InflictStatusTransaction{
						Target:       t.User,
						StatusEffect: StatusInfatuation,
					})
				}
			}
		case MoveRainDance:
			turns := 5
			if user.HeldItem == ItemDampRock {
				turns = 8
			}
			b.QueueTransaction(WeatherTransaction{
				Weather: WeatherRain,
				Turns:   turns,
			})
		case MoveSunnyDay:
			turns := 5
			if user.HeldItem == ItemHeatRock {
				turns = 8
			}
			b.QueueTransaction(WeatherTransaction{
				Weather: WeatherHarshSunlight,
				Turns:   turns,
			})
		case MoveHail:
			turns := 5
			if user.HeldItem == ItemIcyRock {
				turns = 8
			}
			b.QueueTransaction(WeatherTransaction{
				Weather: WeatherHail,
				Turns:   turns,
			})
		case MoveSandstorm:
			turns := 5
			if user.HeldItem == ItemSmoothRock {
				turns = 8
			}
			b.QueueTransaction(WeatherTransaction{
				Weather: WeatherSandstorm,
				Turns:   turns,
			})
		case MoveSplash:
			b.QueueTransaction(MoveFailTransaction{
				User:   t.User,
				Reason: FailOther,
			})
		case MoveDefog:
			if b.Weather == WeatherFog {
				b.QueueTransaction(WeatherTransaction{
					Weather: WeatherClearSkies,
				})
			}
		case MoveMoonlight, MoveSynthesis, MoveMorningSun:
			if b.Weather == WeatherFog {
				b.QueueTransaction(HealTransaction{
					Target: t.User,
					Amount: user.MaxHP() / 4,
				})
			}
		default:
			if t.Move.StatStages() != 0 {
				if t.Move.Targets() == MoveTargetUser {
					b.QueueTransaction(ModifyStatTransaction{
						Target:        t.User,
						SelfInflicted: true,
						Stat:          int(t.Move.AffectedStat()),
						Stages:        int(t.Move.StatStages()),
					})
				} else if t.Move.Targets() == MoveTargetSelected || t.Move.Targets() == MoveTargetAllOpponents {
					b.QueueTransaction(ModifyStatTransaction{
						Target: t.Target,
						Stat:   int(t.Move.AffectedStat()),
						Stages: int(t.Move.StatStages()),
					})
				} else {
					blog.Printf("Unknown target for stat modifying move: %s: %v", t.Move.Name(), t.Move.Targets())
				}
			} else {
				blog.Printf("Unimplemented status move: %s", t.Move.Name())
			}
		}
	} else {
		// Physical/Special Moves
		damage := CalcMoveDamage(b.Weather, user, receiver, t.Move)
		var crit uint = 1
		if b.rng.Roll(1, user.CritChance()) {
			crit = 2
		}
		// Target effects
		if receiver.HeldItem != ItemNone {
			switch receiver.HeldItem {
			case ItemStickyBarb:
				b.QueueTransaction(DamageTransaction{
					Target: t.User,
					Damage: user.MaxHP() / 8,
				})
				if t.Move.Flags()&FlagContact != 0 && user.HeldItem == ItemNone {
					b.QueueTransaction(
						GiveItemTransaction{
							Target: t.User,
							Item:   receiver.HeldItem,
						},
						GiveItemTransaction{
							Target: t.Target,
							Item:   ItemNone,
						},
					)
				}
			}
		}
		damage *= crit
		b.QueueTransaction(DamageTransaction{
			Target: t.Target,
			Move:   t.Move,
			Damage: uint(damage),
		})
		// Handle draining moves (Absorb, Mega Drain, Giga Drain, Drain Punch, etc.)
		// However, if Drain is negative, it's actually recoil damage.
		if t.Move.Drain() != 0 {
			drain := int(damage) * t.Move.Drain() / 100
			if drain > 0 {
				// These multiplers only apply to draining moves, not recoil moves
				if user.HeldItem == ItemBigRoot {
					drain = drain * 130 / 100 // 30% more HP than normal
				}
			}
			if drain == 0 {
				// Min 1 HP drain
				drain = 1
			}
			if drain > 0 {
				b.QueueTransaction(HealTransaction{
					Target: t.User,
					Amount: uint(drain),
				})
			} else {
				// recoil damage
				b.QueueTransaction(DamageTransaction{
					Target: t.User,
					Damage: uint(-drain),
				})
			}
		}
		if t.Move.FlinchChance() > 0 && b.rng.Roll(t.Move.FlinchChance(), 100) {
			b.QueueTransaction(InflictStatusTransaction{
				Target:       t.Target,
				StatusEffect: StatusFlinch,
			})
		}
		// Other item effects in battle
		switch user.HeldItem {
		case ItemKingsRock, ItemRazorFang:
			// King's Rock makes non-flinching moves have a 10% to cause flinch
			// TODO: ensure only certain moves are affected -> https://bulbapedia.bulbagarden.net/wiki/King%27s_Rock
			if t.Move.FlinchChance() == 0 && b.rng.Roll(1, 10) {
				b.QueueTransaction(InflictStatusTransaction{
					Target:       t.Target,
					StatusEffect: StatusFlinch,
				})
			}
		case ItemLifeOrb:
			b.QueueTransaction(DamageTransaction{
				Target: t.User,
				Damage: user.MaxHP() / 10,
			})
		case ItemShellBell:
			b.QueueTransaction(DamageTransaction{
				Target: t.User,
				Damage: uint(damage / 8),
			})
		}
	}
	user.metadata[MetaLastMove] = t.Move
}

// A transaction to deal damage to an opponent Pokemon.
type DamageTransaction struct {
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
	receiver := b.getPokemon(t.Target)
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
	}
}

// A transaction to change the friendship level of a Pokemon.
type FriendshipTransaction struct {
	Target target // The target Pokemon
	Amount int    // The amount of friendship to increase/decrease
}

func (t FriendshipTransaction) Mutate(b *Battle) {
	pkmn := b.getPokemon(t.Target)
	pkmn.Friendship += t.Amount
}

// A transaction to change the EVs of a Pokemon.
type EVTransaction struct {
	Target target
	Stat   int
	Amount uint8
}

func (t EVTransaction) Mutate(b *Battle) {
	pkmn := b.getPokemon(t.Target)
	pkmn.EVs[t.Stat] += t.Amount
}

// A transaction to use and possibly consume an item.
type ItemTransaction struct {
	Target target
	IsHeld bool
	Item   Item
	Move   *Move
}

func (t ItemTransaction) Mutate(b *Battle) {
	receiver := b.getPokemon(t.Target)
	if t.Item.Flags()&FlagConsumable > 0 {
		if t.IsHeld {
			t.Item = receiver.HeldItem // auto-correct if the value is not present or does not match
			receiver.HeldItem = ItemNone
		}
		// TODO: remove consumed item from party's inventory
	}
	switch t.Item.Category() {
	case ItemCategoryFlutes:
		b.QueueTransaction(CureStatusTransaction{
			Target:       t.Target,
			StatusEffect: battleItemFlutes[t.Item],
		})
	case ItemCategoryStatBoost:
		if t.Item == ItemDireHit {
			b.QueueTransaction(ModifyStatTransaction{
				Target: t.Target,
				Stat:   StatCritChance,
				Stages: +2,
			})
		} else if t.Item == ItemGuardSpec {
			receiver.metadata[MetaStatChangeImmune] = 5
		} else {
			b.QueueTransaction(ModifyStatTransaction{
				Target: t.Target,
				Stat:   battleItemStats[t.Item],
				Stages: +1,
			})
		}
		b.QueueTransaction(FriendshipTransaction{
			Target: t.Target,
			Amount: [3]int{1, 1, 0}[receiver.Friendship/100],
		})
	case ItemCategoryHealing, ItemCategoryRevival, ItemCategoryStatusCures:
		if t.Item.Category() == ItemCategoryRevival && receiver.CurrentHP != 0 {
			return
		}
		fIndex := receiver.Friendship / 100
		data, ok := medicineData[t.Item]
		if ok {
			if data.Heal > 0 {
				b.QueueTransaction(HealTransaction{
					Target: t.Target,
					Amount: data.Heal,
				})
			}
			if amount := data.Friendship[fIndex]; amount != 0 {
				b.QueueTransaction(FriendshipTransaction{
					Target: t.Target,
					Amount: amount,
				})
			}
			if data.Cure != StatusNone {
				b.QueueTransaction(CureStatusTransaction{
					Target:       t.Target,
					StatusEffect: data.Cure,
				})
			}
		}
		// Other medicine item effects
		switch t.Item {
		case ItemSacredAsh:
			for slot := range b.parties[t.Target.party].pokemon() {
				pkmnTarget := target{t.Target.party, uint(slot)}
				pkmn := b.getPokemon(pkmnTarget)
				if pkmn.CurrentHP == 0 {
					b.QueueTransaction(HealTransaction{
						Target: pkmnTarget,
						Amount: pkmn.MaxHP(),
					})
				}
			}
		case ItemFullRestore, ItemMaxPotion, ItemMaxRevive, ItemRevivalHerb:
			b.QueueTransaction(HealTransaction{
				Target: t.Target,
				Amount: receiver.MaxHP(),
			})
		case ItemRevive:
			b.QueueTransaction(HealTransaction{
				Target: t.Target,
				Amount: receiver.MaxHP() / 2,
			})
		}
	}
	switch t.Item {
	// ItemCategoryPPRecovery
	case ItemElixir:
		for _, m := range receiver.Moves {
			if m == nil {
				continue
			}
			b.QueueTransaction(PPTransaction{
				Move:   m,
				Amount: 10,
			})
		}
	case ItemEther:
		b.QueueTransaction(PPTransaction{
			Move:   t.Move,
			Amount: 10,
		})
	case ItemMaxElixir:
		for _, m := range receiver.Moves {
			if m == nil {
				continue
			}
			b.QueueTransaction(PPTransaction{
				Move:   m,
				Amount: int8(m.MaxPP),
			})
		}
	case ItemMaxEther:
		b.QueueTransaction(PPTransaction{
			Move:   t.Move,
			Amount: int8(t.Move.MaxPP),
		})
	// ItemCategoryInAPinch
	case ItemApicotBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: t.Target,
			Stat:   StatSpDef,
			Stages: 1,
		})
	case ItemCustapBerry:
		// TODO: Force pokemon to go first
	case ItemGanlonBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: t.Target,
			Stat:   StatDef,
			Stages: 1,
		})
	case ItemLansatBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: t.Target,
			Stat:   StatCritChance,
			Stages: 2,
		})
	case ItemLiechiBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: t.Target,
			Stat:   StatAtk,
			Stages: 1,
		})
	case ItemMicleBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: t.Target,
			Stat:   StatAccuracy,
			Stages: 1,
		})
	case ItemPetayaBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: t.Target,
			Stat:   StatSpAtk,
			Stages: 1,
		})
	case ItemSalacBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: t.Target,
			Stat:   StatSpeed,
			Stages: 1,
		})
	case ItemStarfBerry:
		b.QueueTransaction(ModifyStatTransaction{
			Target: t.Target,
			Stat:   b.rng.Get(StatAtk, StatSpeed),
			Stages: 2,
		})
	case ItemBlackSludge:
		if receiver.Type&TypePoison != 0 {
			b.QueueTransaction(HealTransaction{
				Target: t.Target,
				Amount: receiver.MaxHP() / 16,
			})
		} else {
			b.QueueTransaction(DamageTransaction{
				Target: t.Target,
				Damage: receiver.MaxHP() / 8,
			})
		}
	case ItemLeftovers:
		b.QueueTransaction(HealTransaction{
			Target: t.Target,
			Amount: receiver.MaxHP() / 16,
		})
	case ItemMentalHerb:
		b.QueueTransaction(CureStatusTransaction{
			Target:       t.Target,
			StatusEffect: StatusInfatuation,
		})
	case ItemWhiteHerb:
		for stat, stages := range receiver.StatModifiers {
			if stages < 0 {
				b.QueueTransaction(ModifyStatTransaction{
					Target: t.Target,
					Stat:   stat,
					Stages: -stages,
				})
			}
		}
	// ItemCategoryBadHeldItems
	case ItemFlameOrb:
		b.QueueTransaction(InflictStatusTransaction{
			Target:       t.Target,
			StatusEffect: StatusBurn,
		})
	case ItemStickyBarb:
		b.QueueTransaction(DamageTransaction{
			Target: t.Target,
			Damage: receiver.MaxHP() / 8,
		})
	case ItemToxicOrb:
		b.QueueTransaction(InflictStatusTransaction{
			Target:       t.Target,
			StatusEffect: StatusBadlyPoison,
		})
	}
	// In a pinch consumption
	if receiver.HeldItem.Category() == ItemCategoryInAPinch && receiver.CurrentHP <= receiver.Stats[StatHP]/4 {
		b.QueueTransaction(ItemTransaction{
			Target: t.Target,
			IsHeld: true,
			Item:   receiver.HeldItem,
		})
	}
}

// A transaction to change the PP of a move.
type PPTransaction struct {
	Move   *Move
	Amount int8
}

func (t PPTransaction) Mutate(b *Battle) {
	t.Move.CurrentPP += uint8(t.Amount)
	if t.Move.CurrentPP >= t.Move.MaxPP {
		if t.Amount > 0 {
			t.Move.CurrentPP = t.Move.MaxPP
		} else {
			t.Move.CurrentPP = 0
		}
	}
}

// A transaction to change the held item of a Pokemon
type GiveItemTransaction struct {
	Target target
	Item   Item
}

func (t GiveItemTransaction) Mutate(b *Battle) {
	pkmn := b.getPokemon(t.Target)
	pkmn.HeldItem = t.Item
}

// A transaction to restore HP to a Pokemon.
type HealTransaction struct {
	Target target
	Amount uint
}

func (t HealTransaction) Mutate(b *Battle) {
	pkmn := b.getPokemon(t.Target)
	pkmn.CurrentHP += t.Amount
	if pkmn.CurrentHP > pkmn.MaxHP() {
		pkmn.CurrentHP = pkmn.MaxHP()
	}
}

// A transaction to apply a status effect to a Pokemon.
type InflictStatusTransaction struct {
	Target       target
	StatusEffect StatusCondition
}

func (t InflictStatusTransaction) Mutate(b *Battle) {
	pkmn := b.getPokemon(t.Target)
	pkmn.StatusEffects.apply(t.StatusEffect)
	if t.StatusEffect.check(StatusSleep) {
		pkmn.metadata[MetaSleepTime] = b.rng.Get(1, 5)
	}
	if pkmn.Ability == AbilitySteadfast && t.StatusEffect.check(StatusFlinch) {
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
	pkmn := b.getPokemon(t.Target)
	pkmn.StatusEffects.clear(t.StatusEffect)
	if t.StatusEffect.check(StatusSleep) {
		delete(pkmn.metadata, MetaSleepTime)
	}
}

// A transaction that makes a pokemon faint, and returns the pokemon to the pokeball.
type FaintTransaction struct {
	Target target
}

func (t FaintTransaction) Mutate(b *Battle) {
	pkmn := b.getPokemon(t.Target)
	if b.ruleset&BattleRuleFaint == 0 {
		blog.Println("Fainting is disabled - Pokemon HP fully restored")
		pkmn.CurrentHP = pkmn.MaxHP()
		return
	}
	// EVs are gained based on EV yield of defeated Pokemon
	evGain := pkmn.GetEVYield()
	for _, opponent := range b.getOpponents(b.parties[t.Target.party]) {
		// Friendship is lowered based on level difference
		opponentPkmn := b.getPokemon(opponent)
		levelGap := opponentPkmn.Level - pkmn.Level
		loss := -1
		if levelGap >= 30 {
			if pkmn.Friendship < 200 {
				loss = -5
			} else {
				loss = -10
			}
		}
		b.QueueTransaction(FriendshipTransaction{
			Target: t.Target,
			Amount: loss,
		})

		for stat, amount := range evGain {
			if amount == 0 {
				continue
			}
			b.QueueTransaction(EVTransaction{
				Target: opponent,
				Stat:   stat,
				Amount: uint8(amount),
			})
		}
	}

	p := b.parties[t.Target.party]
	p.SetInactive(t.Target.slot)
	anyAlive := false
	for slot, pkmn := range p.pokemon() {
		if pkmn.CurrentHP > 0 {
			anyAlive = true
			// TODO: prompt Agent for which pokemon to send out next
			// auto send out next pokemon
			b.QueueTransaction(SendOutTransaction{
				Target: target{t.Target.party, uint(slot)},
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
	p.SetActive(t.Target.slot)
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
	receiver := b.getPokemon(t.Target)
	if t.StatusEffect.check(StatusSleep) {
		receiver.metadata[MetaSleepTime] = receiver.metadata[MetaSleepTime].(int) - 1
	}
}

// Handles evasion, misses, dodging, etc. when using moves
type MoveFailTransaction struct {
	User   target
	Reason MoveFailReason
}

func (t MoveFailTransaction) Mutate(b *Battle) {
	// currently a no-op.
}

// Modifies a stat's stages in the interval [-6, 6]
type ModifyStatTransaction struct {
	Target        target
	SelfInflicted bool
	Stat          int
	Stages        int
}

func (t ModifyStatTransaction) Mutate(b *Battle) {
	pkmn := b.getPokemon(t.Target)
	_, immune := pkmn.metadata[MetaStatChangeImmune]
	if immune && t.Stages < 0 && !t.SelfInflicted {
		return
	}
	stage := pkmn.StatModifiers[t.Stat] + t.Stages
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
	pkmn.StatModifiers[t.Stat] = stage
}
