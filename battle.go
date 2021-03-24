package pokemonbattlelib

import (
	"encoding/json"
	"fmt"
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
		panic(ErrorPartyIndex)
	}
	p := b.parties[party].pokemon
	if slot >= len(p) {
		panic(ErrorPartyIndex)
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
	// validate
	teams := map[int]int{}
	for i, party := range b.parties {
		if len(party.pokemon) == 0 {
			return fmt.Errorf("Party (index: %d) has no pokemon.", i)
		}
		teams[party.team]++
	}
	if len(teams) != 2 {
		return fmt.Errorf("Parties have invalid teams. There should be 2 teams with 1 party each, got %d teams", len(teams))
	}

	// Initiate the battle! Send out the first pokemon in the parties.
	b.State = BattleInProgress
	for _, party := range b.parties {
		party.SetActive(0)
	}
	return nil
}

// Handles all pre-turn logic
func (b *Battle) preRound() {
	for _, t := range b.GetTargets() {
		if v, ok := t.Pokemon.metadata[MetaSleepTime]; ok && v.(int) == 0 && t.Pokemon.StatusEffects.check(StatusSleep) {
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
		ctxA := (*turns)[i].Context
		ctxB := (*turns)[j].Context
		if reflect.TypeOf(turnA) == reflect.TypeOf(turnB) {
			switch turnA.(type) {
			case FightTurn:
				ftA := turnA.(FightTurn)
				ftB := turnB.(FightTurn)
				mvA := ctxA.Pokemon.Moves[ftA.Move]
				mvB := ctxB.Pokemon.Moves[ftB.Move]
				if mvA.Priority() != mvB.Priority() {
					return mvA.Priority() > mvB.Priority()
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
	for i, party := range b.parties {
		for j, pokemon := range party.activePokemon {
			ctx := b.getContext(party, pokemon)
			blog.Printf("Requesting turn from agent %d for pokemon %d (%s)", i, j, pokemon)
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

	blog.Println("Sorting turns")
	// Sort turns using an in-place stable sort
	b.sortTurns(&turns)

	// Run turns in sorted order and update battle state
	for len(turns) > 0 {
		turn := turns[0]
		turns = turns[1:]
		blog.Printf("Processing Turn %T for %s", turn.Turn, turn.User.Pokemon)
		// here, we can't use the turn context's reference to the pokemon, because it is a copy of the ground truth pokemon
		self := b.getPokemon(turn.User.party, turn.User.partySlot)
		if self.CurrentHP == 0 {
			continue
		}
		switch t := turn.Turn.(type) {
		case FightTurn:
			user := turn.Context.Pokemon
			move := user.Moves[t.Move]
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
			} else if user.StatusEffects.check(StatusSleep) && move.Id != MoveSnore && move.Id != MoveSleepTalk {
				b.QueueTransaction(ImmobilizeTransaction{
					Target: target{
						Pokemon: user,
					},
					StatusEffect: StatusSleep,
				})
				continue // forfeit turn
			}

			// use the move
			receiver := b.getPokemon(t.Target.party, t.Target.partySlot)
			evasion := float64(receiver.Evasion() / 100)
			// Todo: account for user's accuracy stage
			accuracy := float64(move.Accuracy()) * evasion
			if b.Weather == WeatherFog {
				accuracy *= 3. / 5.
			}
			if move.Accuracy() != 0 && !b.rng.Roll(int(accuracy), 100) {
				b.QueueTransaction(MoveFailTransaction{
					User:   &user,
					Reason: FailMiss,
				})
				continue
			}
			// See: https://github.com/StevensSEC/pokemonbattlelib/wiki/Requirements#fight-using-a-move
			// Status Moves
			if move.Category() == MoveCategoryStatus {
				switch move.Id {
				case MoveDoubleTeam:
					b.QueueTransaction(ModifyStatTransaction{
						Target: &user,
						Stat:   StatEvasion,
						Stages: +1,
					})
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
				case MoveSunnyDay:
					b.QueueTransaction(WeatherTransaction{
						Weather: WeatherHarshSunlight,
						Turns:   5,
					})
				case MoveHowl:
					b.QueueTransaction(ModifyStatTransaction{
						Target: &user,
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
				damage := calcMoveDamage(b.Weather, &user, receiver, move)
				if b.rng.Roll(1, CritChances[user.StatModifiers[StatCritChance]]) {
					damage *= 2
				}
				b.QueueTransaction(DamageTransaction{
					User:   &user,
					Target: t.Target,
					Move:   user.Moves[t.Move],
					Damage: damage,
				})
			}
		case ItemTurn:
			receiver := b.getPokemon(t.Target.party, t.Target.partySlot)
			move := receiver.Moves[t.Move]
			b.QueueTransaction(ItemTransaction{
				Target: t.Target,
				Item:   t.Item,
				Move:   move,
			})
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
	for a, party := range b.parties {
		for ap, pkmn := range party.activePokemon {
			t := target{
				party:     a,
				partySlot: ap,
				Pokemon:   *pkmn,
				Team:      party.team,
			}
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
			if pkmn.HeldItem.Category() == ItemCategoryInAPinch && pkmn.CurrentHP <= pkmn.Stats[StatHP]/4 {
				b.QueueTransaction(ItemTransaction{
					Target: t,
					IsHeld: true,
					Item:   pkmn.HeldItem,
				})
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

type target struct {
	party     int     // Identifier for a party (index in battle parties, or "party ID")
	partySlot int     // The slot of the active Pokemon
	Team      int     // The team that the Pokemon belongs to
	Pokemon   Pokemon // Pokemon that is a candidate target
}

func (t target) String() string {
	return fmt.Sprintf("Party %d (Slot %d) | Team %d | Pokemon:\n%s",
		t.party, t.partySlot, t.Team, t.Pokemon)
}

func (t *target) MarshalJSON() ([]byte, error) {
	type alias target // required to not enter infinite recursive loop
	return json.Marshal(&struct {
		Party int
		Slot  int
		*alias
	}{
		alias: (*alias)(t),
	})
}

func (t *target) UnmarshalJSON(data []byte) error {
	type alias target // required to not enter infinite recursive loop
	aux := &struct {
		Party int
		Slot  int
		*alias
	}{
		alias: (*alias)(t),
	}
	err := json.Unmarshal(data, &aux)
	t.party = aux.Party
	t.partySlot = aux.Slot
	if err != nil {
		return err
	}
	return nil
}

func (b *Battle) GetParty(t *target) *party {
	return b.parties[t.party]
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

// Get the battle context that will be shared with the client
func (b *Battle) GetRoundContext(t target) *BattleContext {
	return b.getContext(b.parties[t.party], b.parties[t.party].activePokemon[t.partySlot])
}

// An abstraction over all possible actions an `Agent` can make in one round. Each Pokemon gets one turn.
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
	Item   Item   // Which item is being consumed
}

func (turn ItemTurn) Priority() int {
	return 1
}
