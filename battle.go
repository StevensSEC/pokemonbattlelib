package pokemonbattlelib

import (
	"log"
	"math/rand"
	"reflect"
	"sort"
)

// A Pokemon battle. Enforces rules of the battle, and queries `Agent`s for turns.
type Battle struct {
	Weather  int  // one of the 6 in-battle weather conditions
	ShiftSet bool // shift or set battle style for NPC trainer battles
	State    BattleState
	rng      RNG

	parties []*party // All parties participating in the battle

	tQueue     []Transaction
	tProcessed []Transaction
}

type BattleState int

const (
	BEFORE_START BattleState = iota
	BATTLE_IN_PROGRESS
	BATTLE_END
)

// Creates a new battle instance, setting initial conditions
func NewBattle() *Battle {
	rng := LCRNG(rand.Uint32())
	b := Battle{
		State: BEFORE_START,
		rng:   RNG(&rng),
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

func (b *Battle) Start() error {
	// TODO: validate the battle, return error if invalid

	// Initiate the battle! Send out the first pokemon in the parties.
	b.State = BATTLE_IN_PROGRESS
	for _, party := range b.parties {
		party.SetActive(0)
	}
	return nil
}

// Simulates a single round of the battle. Returns processed transactions for this turn and indicates whether the battle has ended.
func (b *Battle) SimulateRound() ([]Transaction, bool) {
	if b.State != BATTLE_IN_PROGRESS {
		log.Panic("battle is not currently in progress")
	}
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
	for len(turns) > 0 {
		turn := turns[0]
		turns = turns[1:]
		// here, we can't use the turn context's reference to the pokemon, because it is a copy of the ground truth pokemon
		user := b.getPokemon(turn.User.party, turn.User.partySlot)
		if user.CurrentHP == 0 {
			continue
		}
		switch t := turn.Turn.(type) {
		case FightTurn:
			user := turn.Context.Pokemon

			// pre-move checks
			if user.StatusEffects.check(StatusFreeze) || user.StatusEffects.check(StatusParalyze) {
				var success bool
				if user.StatusEffects.check(StatusFreeze) {
					success = b.rng.Get(1, 5) == 1
				} else if user.StatusEffects.check(StatusParalyze) {
					success = b.rng.Get(1, 4) != 1
				}
				if !success {
					b.QueueTransaction(ImmobilizeTransaction{
						Target: target{
							Pokemon: user,
						},
						StatusEffect: user.StatusEffects & NONVOLATILE_STATUS_MASK,
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
			receiver := b.getPokemon(t.Target.party, t.Target.partySlot)
			// See: https://github.com/StevensSEC/pokemonbattlelib/wiki/Requirements#fight-using-a-move
			if move.Category == Status {
				if move.ID == MOVE_STUN_SPORE {
					b.QueueTransaction(InflictStatusTransaction{
						Target: receiver,
						Status: StatusParalyze,
					})
				}
			} else {
				stab := 1.0
				if move != nil && user.Elemental&move.Type != 0 {
					stab = 1.5
					if user.Ability != nil && user.Ability.ID == 91 { // Adaptability
						stab = 2.0
					}
				}
				modifier := stab // TODO: damage multiplers
				levelEffect := float64((2 * user.Level / 5) + 2)
				movePower := float64(move.Power)
				statRatio := float64(user.Stats[STAT_ATK] / receiver.Stats[STAT_DEF])
				if move.Category == Special {
					statRatio = float64(user.Stats[STAT_SPATK] / receiver.Stats[STAT_SPDEF])
				}
				damage := (((levelEffect * movePower * statRatio) / 50) + 2) * modifier
				b.QueueTransaction(DamageTransaction{
					User:   &user,
					Target: t.Target,
					Move:   user.Moves[t.Move],
					Damage: uint(damage),
				})
			}
		case ItemTurn:
			receiver := b.getPokemon(t.Target.party, t.Target.partySlot)
			move := receiver.Moves[t.Move]
			b.QueueTransaction(ItemTransaction{
				Target: receiver,
				Item:   t.Item,
				Move:   move,
			})
			b.QueueTransaction(receiver.UseItem(t.Item)...)
		default:
			log.Panicf("Unknown turn of type %v", t)
		}

		b.ProcessQueue()
		if b.State == BATTLE_END {
			break
		}
	}

	// handle post turn status effects
	for a, party := range b.parties {
		for ap, pkmn := range party.activePokemon {
			if pkmn.StatusEffects.check(StatusBurn) || pkmn.StatusEffects.check(StatusPoison) || pkmn.StatusEffects.check(StatusBadlyPoison) {
				t := target{
					party:     a,
					partySlot: ap,
					Pokemon:   *pkmn,
					Team:      party.team,
				}
				cond := pkmn.StatusEffects & NONVOLATILE_STATUS_MASK
				var damage uint
				switch cond {
				case StatusBurn, StatusPoison:
					damage = pkmn.Stats[STAT_HP] / 8
				case StatusBadlyPoison:
					// TODO: implement counter for increasing bad poison damage
					damage = pkmn.Stats[STAT_HP] / 16
				}
				b.QueueTransaction(DamageTransaction{
					Target:       t,
					Damage:       damage,
					StatusEffect: cond,
				})
			}
		}
	}
	b.ProcessQueue()

	if len(b.tQueue) > 0 {
		log.Panic("FATAL: There are still unprocessed transactions at the end of the round.")
	}
	transactions := b.tProcessed
	b.tProcessed = []Transaction{}
	return transactions, b.State == BATTLE_END
}

// Add Transactions to the queue.
func (b *Battle) QueueTransaction(t ...Transaction) {
	b.tQueue = append(b.tQueue, t...)
}

// Process Transactions that are in the queue until the queue is empty.
func (b *Battle) ProcessQueue() {
	for len(b.tQueue) > 0 {
		next := b.tQueue[0]
		b.tQueue = b.tQueue[1:]
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
		case ItemTransaction:
			// TODO: do not consume certain items
			if t.Target.HeldItem == t.Item {
				t.Target.HeldItem = nil
			}
		case FriendshipTransaction:
			t.Target.Friendship += t.Amount
		case HealTransaction:
			t.Target.CurrentHP += t.Amount
		case InflictStatusTransaction:
			t.Target.StatusEffects.apply(t.Status)
		case CureStatusTransaction:
			receiver := b.getPokemon(t.Target.party, t.Target.partySlot)
			receiver.StatusEffects.clear(t.Status)
		case FaintTransaction:
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
		case SendOutTransaction:
			p := b.parties[t.Target.party]
			p.SetActive(t.Target.partySlot)
		case EndBattleTransaction:
			b.State = BATTLE_END
		}
		// add to the list of processed transactions
		b.tProcessed = append(b.tProcessed, next)
		if b.State == BATTLE_END {
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
