package pokemonbattlelib

import (
	"log"
	"math/rand"
	"reflect"
	"sort"
)

// A Pokemon battle. Enforces rules of the battle, and queries `Agent`s for turns.
type Battle struct {
	Weather  Weather // one of the 6 in-battle weather conditions
	ShiftSet bool    // shift or set battle style for NPC trainer battles
	State    BattleState
	rng      RNG

	parties  []*party                   // All parties participating in the battle
	metadata map[BattleMeta]interface{} // Metadata to be tracked during a battle

	tQueue     []Transaction
	tProcessed []Transaction
}

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
		State: BattleBeforeStart,
		rng:   RNG(&rng),
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

// Start the battle.
func (b *Battle) Start() error {
	// TODO: validate the battle, return error if invalid

	// Initiate the battle! Send out the first pokemon in the parties.
	b.State = BattleInProgress
	for _, party := range b.parties {
		party.SetActive(0)
	}
	return nil
}

// Handles all pre-turn logic
func (b *Battle) preRound() {
	// TODO
}

// Simulates a single round of the battle. Returns processed transactions for this turn and indicates whether the battle has ended.
func (b *Battle) SimulateRound() ([]Transaction, bool) {
	if b.State != BattleInProgress {
		log.Panic("battle is not currently in progress")
	}
	b.preRound()
	// Collects all turn info from each active Pokemon
	turns := make([]TurnContext, 0)
	for i, party := range b.parties {
		for j, pokemon := range party.activePokemon {
			ctx := b.getContext(party, pokemon)
			turn := (*party.Agent).Act(ctx)
			turns = append(turns, TurnContext{
				User: target{
					Pokemon:   *pokemon,
					party:     i,
					partySlot: j,
					Team:      party.team,
				},
				Turn:    turn,
				Context: ctx,
			})
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
				ftA := turnA.(FightTurn)
				ftB := turnB.(FightTurn)
				mvA := ctxA.Pokemon.Moves[ftA.Move]
				mvB := ctxB.Pokemon.Moves[ftB.Move]
				if mvA.Priority != mvB.Priority {
					return mvA.Priority > mvB.Priority
				}
				// speedy pokemon should go first
				return ctxA.Pokemon.Speed() > ctxB.Pokemon.Speed()
			}
		} else {
			// make higher priority turns go first
			return turnA.Priority() > turnB.Priority()
		}
		// fallthrough
		return false
	})
	// Run turns in sorted order and update battle state
	for len(turns) > 0 {
		turn := turns[0]
		turns = turns[1:]
		// here, we can't use the turn context's reference to the pokemon, because it is a copy of the ground truth pokemon
		self := b.getPokemon(turn.User.party, turn.User.partySlot)
		if self.CurrentHP == 0 {
			continue
		}
		switch t := turn.Turn.(type) {
		case FightTurn:
			user := turn.Context.Pokemon
			// pre-move checks
			if user.StatusEffects.check(StatusFreeze) || user.StatusEffects.check(StatusParalyze) {
				immobilize := false
				if user.StatusEffects.check(StatusFreeze) {
					immobilize = b.rng.Roll(4, 5)
				} else if user.StatusEffects.check(StatusParalyze) {
					immobilize = b.rng.Roll(1, 4)
				}
				if immobilize {
					b.QueueTransaction(ImmobilizeTransaction{
						Target: target{
							Pokemon: user,
						},
						StatusEffect: user.StatusEffects & StatusNonvolatileMask,
					})
					continue // forfeit turn
				}
			} else if user.StatusEffects.check(StatusSleep) {
				b.QueueTransaction(ImmobilizeTransaction{
					Target: target{
						Pokemon: user,
					},
					StatusEffect: StatusSleep,
				})
				continue // forfeit turn
			}

			// use the move
			move := user.Moves[t.Move]
			accuracy := float64(move.Accuracy)
			if b.Weather == WeatherFog {
				accuracy *= 3. / 5.
			}
			// Todo: account for receiver's evasion
			receiver := b.getPokemon(t.Target.party, t.Target.partySlot)
			if move.Accuracy != 0 && !b.rng.Roll(int(accuracy), 100) {
				b.QueueTransaction(EvadeTransaction{
					User: &user,
				})
				continue
			}
			// See: https://github.com/StevensSEC/pokemonbattlelib/wiki/Requirements#fight-using-a-move
			// Status Moves
			if move.Category == MoveCategoryStatus {
				switch move.ID {
				case MoveStunSpore:
					b.QueueTransaction(InflictStatusTransaction{
						Target:       receiver,
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
							Target:       receiver,
							StatusEffect: StatusInfatuation,
						})
						if receiver.HeldItem != nil && receiver.HeldItem.ID == ItemDestinyKnot {
							b.QueueTransaction(InflictStatusTransaction{
								Target:       self,
								StatusEffect: StatusInfatuation,
							})
						}
					}
				case MoveRainDance:
					turns := 5
					if self.HeldItem != nil && self.HeldItem.ID == ItemDampRock {
						turns = 8
					}
					b.QueueTransaction(WeatherTransaction{
						Weather: WeatherHarshSunlight,
						Turns:   turns,
					})
				case MoveSunnyDay:
					turns := 5
					if self.HeldItem != nil && self.HeldItem.ID == ItemHeatRock {
						turns = 8
					}
					b.QueueTransaction(WeatherTransaction{
						Weather: WeatherHarshSunlight,
						Turns:   turns,
					})
				case MoveHail:
					turns := 5
					if self.HeldItem != nil && self.HeldItem.ID == ItemIcyRock {
						turns = 8
					}
					b.QueueTransaction(WeatherTransaction{
						Weather: WeatherHarshSunlight,
						Turns:   turns,
					})
				case MoveHowl:
					b.QueueTransaction(ModifyStatTransaction{
						Target: self,
						Stat:   StatAtk,
						Stages: +1,
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
							Target: self,
							Amount: self.MaxHP() / 4,
						})
					}
				}
			} else {
				// Physical/Special Moves
				weather := 1.0
				if rain, sun := b.Weather == WeatherRain, b.Weather == WeatherHarshSunlight; (rain && move.Type == TypeWater) || (sun && move.Type == TypeFire) {
					weather = 1.5
				} else if (rain && move.Type == TypeFire) || (sun && move.Type == TypeWater) {
					weather = 0.5
				}
				crit := 1.0
				if b.rng.Roll(1, CritChances[user.StatModifiers[StatCritChance]]) {
					crit = 2.0
				}
				stab := 1.0
				if move != nil && user.Type&move.Type != 0 {
					stab = 1.5
					if user.Ability != nil && user.Ability.ID == 91 { // Adaptability
						stab = 2.0
					}
				}
				modifier := weather * crit * stab
				levelEffect := float64((2 * user.Level / 5) + 2)
				movePower := float64(move.Power)
				attack := float64(user.Attack())
				defense := float64(receiver.Defense())
				// Move modifiers
				if move.Category == MoveCategorySpecial {
					attack = float64(user.SpecialAttack())
					defense = float64(receiver.SpecialDefense())
				}
				// Weather modifiers
				if b.Weather == WeatherSandstorm {
					if receiver.Type&TypeRock != 0 {
						defense *= 1.5
					}
					if move.ID == MoveSolarBeam {
						movePower /= 2
					}
				}
				if b.Weather == WeatherHail && move.ID == MoveSolarBeam {
					movePower /= 2
				}
				if b.Weather == WeatherFog {
					if move.ID == MoveWeatherBall {
						movePower *= 2
					}
					if move.ID == MoveSolarBeam {
						movePower /= 2
					}
				}
				// Item modifiers
				if user.HeldItem != nil {
					switch user.HeldItem.ID {
					case ItemLifeOrb:
						modifier *= 1.30
					}
				}
				damage := (((levelEffect * movePower * attack / defense) / 50) + 2) * modifier
				b.QueueTransaction(DamageTransaction{
					User:   &user,
					Target: t.Target,
					Move:   user.Moves[t.Move],
					Damage: uint(damage),
				})
				// Handle draining moves (Absorb, Mega Drain, Giga Drain, Drain Punch, etc.)
				if move.metadata.Drain != 0 {
					drain := damage * float64(move.metadata.Drain) / 100
					if user.HeldItem != nil && user.HeldItem.ID == ItemBigRoot {
						drain *= 1.30 // 30% more HP than normal
					}
					if drain == 0 {
						// Min 1 HP drain
						drain = 1
					}
					b.QueueTransaction(HealTransaction{
						Target: &user,
						Amount: uint(drain),
					})
				}
				if user.HeldItem != nil {
					switch user.HeldItem.ID {
					case ItemKingsRock:
						// King's Rock makes non-flinching moves have a 10% to cause flinch
						// TODO: ensure only certain moves are affected -> https://bulbapedia.bulbagarden.net/wiki/King%27s_Rock
						if move.metadata.FlinchChance == 0 && b.rng.Roll(1, 10) {
							b.QueueTransaction(InflictStatusTransaction{
								Target:       receiver,
								StatusEffect: StatusFlinch,
							})
						}
					case ItemLifeOrb:
						b.QueueTransaction(DamageTransaction{
							Target: turn.User,
							Damage: self.MaxHP() / 10,
						})
					}
				}
			}
		case ItemTurn:
			receiver := b.getPokemon(t.Target.party, t.Target.partySlot)
			move := receiver.Moves[t.Move]
			b.QueueTransaction(ItemTransaction{
				Target: receiver,
				Item:   t.Item,
				Move:   move,
			})
		default:
			log.Panicf("Unknown turn of type %v", t)
		}

		b.ProcessQueue()
		if b.State == BattleEnd {
			break
		}
	}
	b.postRound()

	b.ProcessQueue()
	if len(b.tQueue) > 0 {
		log.Panic("FATAL: There are still unprocessed transactions at the end of the round.")
	}
	transactions := b.tProcessed
	b.tProcessed = []Transaction{}
	return transactions, b.State == BattleEnd
}

// Handles all post-round logic
func (b *Battle) postRound() {
	// Effects on every Pokemon
	for _, t := range b.GetTargets() {
		pkmn := b.getPokemon(t.party, t.partySlot)
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
		// Weather effects
		// TODO: check for weather resisting abilities
		if b.Weather == WeatherSandstorm {
			if pkmn.Type&(TypeRock|TypeGround|TypeSteel) == 0 {
				damage := pkmn.MaxHP() / 16
				b.QueueTransaction(DamageTransaction{
					Target: t,
					Damage: damage,
				})
			}
		} else if b.Weather == WeatherHail {
			if pkmn.Type&TypeIce == 0 {
				damage := pkmn.MaxHP() / 16
				b.QueueTransaction(DamageTransaction{
					Target: t,
					Damage: damage,
				})
			}
		}
		// Held item effects
		if pkmn.HeldItem != nil {
			if pkmn.HeldItem.Category == ItemCategoryInAPinch && pkmn.CurrentHP <= pkmn.Stats[StatHP]/4 {
				b.QueueTransaction(ItemTransaction{
					Target: pkmn,
					Item:   pkmn.HeldItem,
				})
			}
			switch pkmn.HeldItem.ID {
			case ItemLeftovers:
				b.QueueTransaction(HealTransaction{
					Target: pkmn,
					Amount: pkmn.MaxHP() / 16,
				})
			case ItemBlackSludge:
				if pkmn.Type&TypePoison != 0 {
					b.QueueTransaction(HealTransaction{
						Target: pkmn,
						Amount: pkmn.MaxHP() / 16,
					})
				} else {
					b.QueueTransaction(DamageTransaction{
						Target: t,
						Damage: pkmn.MaxHP() / 8,
					})
				}
			}
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
		b.tQueue = b.tQueue[1:]
		t.Mutate(b)

		// add to the list of processed transactions
		b.tProcessed = append(b.tProcessed, t)
		if b.State == BattleEnd {
			break
		}
	}
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
	User    target         // The pokemon that made this turn.
	Turn    Turn           // A copy of the turn that a Pokemon made using an Agent
	Context *BattleContext // The context in which the Pokemon took its turn
}

// A turn to represent a Pokemon using a Move.
type FightTurn struct {
	Move   int    // Denotes the index (0-3) of the pokemon's which of the pokemon's moves to use.
	Target target // Info containing data determining the target of
}

func (turn FightTurn) Priority() int {
	return 0
}

// A turn to represent using an item from the Party's inventory. An item turn has the a higher priority than any move.
type ItemTurn struct {
	Move   int    // Denotes the index (0-3) of the pokemon's which of the pokemon's moves to use.
	Target target // Info containing data determining the target of
	Item   *Item  // Which item is being consumed
}

func (turn ItemTurn) Priority() int {
	return 1
}
