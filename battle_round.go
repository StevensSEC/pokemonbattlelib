package pokemonbattlelib

import (
	"reflect"
	"sort"
)

// Handles all pre-turn logic
func (b *Battle) preRound() {
	for _, t := range b.AllTargets() {
		pkmn := b.GetPokemon(t)
		if v, ok := pkmn.metadata[MetaSleepTime]; ok && v.(int) == 0 && pkmn.StatusEffects.check(StatusSleep) {
			b.QueueTransaction(CureStatusTransaction{
				Target:       t,
				StatusEffect: StatusSleep,
			})
		}
	}
}

func (b *Battle) sortTurns(turns *[]TurnContext) {
	sort.SliceStable(*turns, func(i, j int) bool {
		turnA := (*turns)[i].Turn
		turnB := (*turns)[j].Turn
		pkmnA := b.GetPokemon((*turns)[i].User)
		pkmnB := b.GetPokemon((*turns)[j].User)
		if reflect.TypeOf(turnA) == reflect.TypeOf(turnB) {
			switch turnA.(type) {
			case FightTurn:
				ftA := turnA.(FightTurn)
				ftB := turnB.(FightTurn)
				mvA := pkmnA.Moves[ftA.Move]
				mvB := pkmnB.Moves[ftB.Move]
				if mvA.Priority() != mvB.Priority() {
					return mvA.Priority() > mvB.Priority()
				}
				// Held item priority
				itemLastA := 0
				itemLastB := 0
				switch pkmnA.HeldItem {
				case ItemFullIncense, ItemLaggingTail:
					itemLastA = 1
				}
				switch pkmnB.HeldItem {
				case ItemFullIncense, ItemLaggingTail:
					itemLastB = 1
				}
				if itemLastA != itemLastB {
					return itemLastA < itemLastB
				}
				// speedy pokemon should go first
				return pkmnA.Speed() > pkmnB.Speed()
			}
		} else {
			// make higher priority turns go first
			return turnA.Priority() > turnB.Priority()
		}
		// fallthrough
		return false
	})
}

// Simulates a single round of the battle. Returns processed transactions for this turn and indicates whether the battle has ended.
func (b *Battle) SimulateRound() ([]Transaction, bool) {
	if b.State != BattleInProgress {
		blog.Panic("battle is not currently in progress")
	}
	b.preRound()
	b.ProcessQueue()
	// Collects all turn info from each active Pokemon
	turns := make([]TurnContext, 0)
	for _, t := range b.AllTargets() {
		blog.Printf("Requesting turn for target: %v", t)
		party := b.parties[t.party]
		ctx := b.GetBattleContext(t)
		turn := (*party.Agent).Act(ctx)
		turns = append(turns, TurnContext{
			User: t,
			Turn: turn,
		})
	}
	blog.Println("Sorting turns")
	// Sort turns using an in-place stable sort
	b.sortTurns(&turns)
	// Run turns in sorted order and update battle state
	for len(turns) > 0 {
		turn := turns[0]
		turns = turns[1:]
		user := b.GetPokemon(turn.User)
		blog.Printf("Processing Turn %T for %s", turn.Turn, user)
		if user.CurrentHP == 0 {
			continue
		}
		switch t := turn.Turn.(type) {
		case FightTurn:
			move := user.Moves[t.Move]
			// pre-move checks
			if user.StatusEffects.check(StatusFreeze) || user.StatusEffects.check(StatusParalyze) || user.StatusEffects.check(StatusFlinch) {
				immobilize := false
				status := user.StatusEffects & StatusNonvolatileMask
				if user.StatusEffects.check(StatusFreeze) {
					immobilize = b.rng.Roll(4, 5)
				} else if user.StatusEffects.check(StatusParalyze) {
					immobilize = b.rng.Roll(1, 4)
				}
				if user.StatusEffects.check(StatusFlinch) {
					immobilize = true
					status = StatusFlinch
				}
				if immobilize {
					b.QueueTransaction(ImmobilizeTransaction{
						Target:       turn.User,
						StatusEffect: status,
					})
					continue // forfeit turn
				}
			} else if user.StatusEffects.check(StatusSleep) && move.Id != MoveSnore && move.Id != MoveSleepTalk {
				b.QueueTransaction(ImmobilizeTransaction{
					Target:       turn.User,
					StatusEffect: StatusSleep,
				})
				continue // forfeit turn
			}

			// use the move
			b.QueueTransaction(UseMoveTransaction{
				User:   turn.User,
				Target: t.Target.target,
				Move:   user.Moves[t.Move],
			})
		case ItemTurn:
			receiver := b.GetPokemon(t.Target.target)
			b.QueueTransaction(ItemTransaction{
				Target: t.Target.target,
				Item:   t.Item,
				Move:   receiver.Moves[t.Move],
			})
		case SwitchTurn:
			party := b.GetParty(turn.User)
			if !party.IsActivePokemon(turn.User.slot) {
				blog.Panic(ErrorCannotSwitch)
			}
			pkmn := b.GetPokemon(t.Target.target)
			if pkmn.CurrentHP == 0 {
				blog.Panic(ErrorCannotSwitch)
			}
			party.SetInactive(turn.User.slot)
			party.activePokemon[turn.User.slot] = t.Target.slot
		default:
			blog.Panicf("Unknown turn of type %v", t)
		}
		b.ProcessQueue()
		if b.State == BattleEnd {
			break
		}
	}
	b.postRound()
	b.ProcessQueue()
	if len(b.tQueue) > 0 {
		blog.Panic("FATAL: There are still unprocessed transactions at the end of the round.")
	}
	transactions := b.tProcessed
	b.tProcessed = []Transaction{}
	return transactions, b.State == BattleEnd
}

// Handles all post-round logic
func (b *Battle) postRound() {
	blog.Println("Post-round")
	// Effects on every Pokemon
	for _, t := range b.AllTargets() {
		pkmn := b.GetPokemon(t)
		// Status effects
		if pkmn.StatusEffects.check(StatusBurn) || pkmn.StatusEffects.check(StatusPoison) || pkmn.StatusEffects.check(StatusBadlyPoison) {
			cond := pkmn.StatusEffects & StatusNonvolatileMask
			var damage uint
			switch cond {
			case StatusBurn, StatusPoison:
				damage = pkmn.MaxHP() / 8
			case StatusBadlyPoison:
				// TODO: implement counter for increasing bad poison damage
				damage = pkmn.MaxHP() / 16
			}
			b.QueueTransaction(DamageTransaction{
				Target:       t,
				Damage:       damage,
				StatusEffect: cond,
			})
		}
		pkmn.StatusEffects.clear(StatusFlinch) // Flinching only occurs over the course of a single turn. It never bleeds over into the next turn.
		if v, ok := pkmn.metadata[MetaStatChangeImmune]; ok {
			turns := v.(int)
			pkmn.metadata[MetaStatChangeImmune] = turns - 1
			if turns == 0 {
				delete(pkmn.metadata, MetaStatChangeImmune)
			}
		}
		// Weather effects
		// TODO: check for weather resisting abilities
		if b.Weather == WeatherSandstorm {
			if pkmn.EffectiveType()&(TypeRock|TypeGround|TypeSteel) == 0 {
				damage := pkmn.MaxHP() / 16
				b.QueueTransaction(DamageTransaction{
					Target: t,
					Damage: damage,
				})
			}
		} else if b.Weather == WeatherHail {
			if pkmn.EffectiveType()&TypeIce == 0 {
				damage := pkmn.MaxHP() / 16
				b.QueueTransaction(DamageTransaction{
					Target: t,
					Damage: damage,
				})
			}
		}
		// Held item effects
		if pkmn.HeldItem != ItemNone {
			b.QueueTransaction(ItemTransaction{
				Target: t,
				IsHeld: true,
				Item:   pkmn.HeldItem,
			})
		}
	}
	// Effects on the battle
	// Decrease weather counter/clear weather over time
	if b.Weather != WeatherClearSkies && b.metadata[MetaWeatherTurns] == 0 {
		b.QueueTransaction(WeatherTransaction{
			Weather: WeatherClearSkies,
		})
	}
	if turns := b.metadata[MetaWeatherTurns].(int); turns > 0 {
		b.metadata[MetaWeatherTurns] = turns - 1
	}
}

// Add Transactions to the queue.
func (b *Battle) QueueTransaction(t ...Transaction) {
	b.tQueue = append(b.tQueue, t...)
}

// Process Transactions that are in the queue until the queue is empty.
func (b *Battle) ProcessQueue() {
	for len(b.tQueue) > 0 {
		t := b.tQueue[0]
		blog.Printf("Processing Transaction %T", t)
		b.tQueue = b.tQueue[1:]
		t.Mutate(b)

		// add to the list of processed transactions
		b.tProcessed = append(b.tProcessed, t)
		if b.State == BattleEnd {
			break
		}
	}
}
