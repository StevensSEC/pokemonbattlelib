package pokemonbattlelib

import (
	"encoding/json"
	"fmt"
	"math"
)

// Maximum number of moves a Pokemon can have in its moveset
const MaxMoves = 4

type Pokemon struct {
	NatDex            uint16                      // National Pokedex Number
	Level             uint8                       // value from 1-100 influencing stats
	Ability           Ability                     // name of this Pokemon's ability
	TotalExperience   uint                        // the total amount of exp this Pokemon has gained, influencing its level
	Gender            Gender                      // this Pokemon's gender
	IVs               [6]uint8                    // values from 0-31 that represents a Pokemon's 'genetic' potential
	EVs               [6]uint8                    // values from 0-255 that represents a Pokemon's training in a particular stat
	Nature            Nature                      // represents a Pokemon's disposition and affects stats
	Stats             [6]uint                     // the actual stats of a Pokemon determined from the above data
	StatModifiers     [9]int                      // ranges from +6 (buffing) to -6 (debuffing) a stat
	StatusEffects     StatusCondition             // the current status effects inflicted on a Pokemon
	CurrentHP         uint                        // the remaining HP of this Pokemon
	HeldItem          Item                        // the item a Pokemon is holding
	Moves             [MaxMoves]*Move             // the moves the Pokemon currenly knows
	Friendship        int                         // how close this Pokemon is to its Trainer
	OriginalTrainerID uint16                      // a number associated with the first Trainer who caught this Pokemon
	Type              Type                        // Indicates what type(s) (up to 2 simultaneously) this pokemon has
	metadata          map[PokemonMeta]interface{} // Data that is conditionally needed in a battle
}

// Metadata for a Pokemon to keep track of
type PokemonMeta int

const (
	MetaLastMove PokemonMeta = iota
	MetaSleepTime
)

// Keeps track of base stats and EV yield for a Pokemon
type PokemonBaseStats struct {
	Stats   [6]int // base stats of a Pokemon
	EVYield [6]int // effort points gained when Pokemon is defeated
}

// Constants for growth rates of a Pokemon
const (
	GrowthSlow = iota + 1
	GrowthMediumFast
	GrowthFast
	GrowthMediumSlow
	GrowthErratic
	GrowthFluctuating
)

// Constants for IVs and EVs
const (
	MaxFriendship = 255
	MaxEV         = 252
	MaxIV         = 31
	TotalEV       = 510
	MinLevel      = 1
	MaxLevel      = 100
)

type GeneratePokemonOption func(p *Pokemon)

// Creates a new Pokemon given its national dex number and other options.
func GeneratePokemon(natdex int, opts ...GeneratePokemonOption) *Pokemon {
	p := &Pokemon{
		NatDex:          uint16(natdex),
		Level:           1,
		TotalExperience: 0,
		IVs:             [6]uint8{0, 0, 0, 0, 0, 0},
		EVs:             [6]uint8{0, 0, 0, 0, 0, 0},
		StatModifiers:   [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0},
		Stats:           [6]uint{1, 4, 4, 4, 4, 4},
		Nature:          NatureHardy, // this nature is neutral and has no effect
		metadata:        make(map[PokemonMeta]interface{}),
	}
	p.Type = pokemonData[p.NatDex].Type
	for _, opt := range opts {
		opt(p)
	}
	p.computeStats()
	return p
}

func WithLevel(level uint8) GeneratePokemonOption {
	return func(p *Pokemon) {
		p.GainLevels(int(level - p.Level))
	}
}

func WithTotalExp(totalExp uint) GeneratePokemonOption {
	return func(p *Pokemon) {
		p.GainExperience(int(totalExp))
	}
}

func WithIVs(ivs [6]uint8) GeneratePokemonOption {
	return func(p *Pokemon) {
		p.IVs = ivs
	}
}

func WithEVs(evs [6]uint8) GeneratePokemonOption {
	return func(p *Pokemon) {
		p.EVs = evs
	}
}

func WithNature(nature Nature) GeneratePokemonOption {
	return func(p *Pokemon) {
		p.Nature = nature
	}
}

func WithMoves(moves ...*Move) GeneratePokemonOption {
	return func(p *Pokemon) {
		if len(moves) > MaxMoves {
			panic(fmt.Sprintf("A Pokemon cannot have more than %d moves", MaxMoves))
		}

		var limited_moves [4]*Move
		copy(limited_moves[:], moves)
		p.Moves = limited_moves
	}
}

func (p *Pokemon) GetName() string {
	return pokemonData[p.NatDex].Name
}

func (p *Pokemon) GetGrowthRate() int {
	return pokemonGrowthRates[int(p.NatDex)]
}

func (p *Pokemon) GetBaseStats() [6]int {
	return pokemonBaseStats[int(p.NatDex)].Stats
}

func (p *Pokemon) GetEVYield() [6]int {
	return pokemonBaseStats[int(p.NatDex)].EVYield
}

func (p *Pokemon) HasValidLevel() bool {
	return p.Level >= MinLevel && p.Level <= MaxLevel
}

func (p *Pokemon) HasValidIVs() bool {
	for _, IV := range p.IVs {
		if IV > MaxIV {
			return false
		}
	}
	return true
}

func (p *Pokemon) HasValidEVs() bool {
	totalEVs := 0
	for _, EV := range p.EVs {
		if EV > MaxEV {
			return false
		}
		totalEVs += int(EV)
	}
	return totalEVs <= TotalEV
}

// Increases a Pokemon's level by `levels` and sets their total experience to
// the minimum amount for that level.
func (p *Pokemon) GainLevels(levels int) {
	newLevel := int(p.Level) + levels
	if newLevel > MaxLevel {
		panic(fmt.Sprintf("Level %d %s cannot level up to level %d!", p.Level, p.GetName(), newLevel))
	}

	if levels < 0 {
		panic(fmt.Sprintf("%s's level tried to decrease by %d", p.GetName(), levels))
	}

	if levels == 0 {
		return
	}

	p.Level = uint8(newLevel)
	p.TotalExperience = uint(ExpTable[p.GetGrowthRate()][int(p.Level)])
	p.computeStats()
}

// Increases a Pokemon's experience points by `exp` and levels up the Pokemon accordingly.
func (p *Pokemon) GainExperience(exp int) {

	if exp < 0 {
		panic(fmt.Sprintf("%s's experience tried to decrease by %d", p.GetName(), exp))
	}

	max_exp := ExpTable[p.GetGrowthRate()][MaxLevel]

	// if would gain experience beyond leveling to 100, set level to 100
	if int(p.TotalExperience)+exp > max_exp {
		p.GainLevels(MaxLevel - int(p.Level))
		return
	}

	remaining_exp := exp
	toNextLevel := ExpTable[p.GetGrowthRate()][int(p.Level+1)] - int(p.TotalExperience)

	for remaining_exp >= toNextLevel {
		p.GainLevels(1)
		remaining_exp -= toNextLevel
		toNextLevel = ExpTable[p.GetGrowthRate()][int(p.Level+1)] - int(p.TotalExperience)
	}

	// add whats left
	p.TotalExperience += uint(remaining_exp)

}

func (p *Pokemon) computeStats() {
	if !p.HasValidLevel() {
		panic(fmt.Sprintf("Failed to compute stats for %s, level %d invalid", p.GetName(), p.Level))
	}

	if !p.HasValidIVs() {
		panic(fmt.Sprintf("Failed to compute stats for %s, ivs %d invalid", p.GetName(), p.IVs))
	}

	if !p.HasValidEVs() {
		panic(fmt.Sprintf("Failed to compute stats for %s, evs %d invalid", p.GetName(), p.EVs))
	}

	// get base stats
	base_stats := p.GetBaseStats()

	// compute HP
	hp_ev_term := math.Floor(float64(p.EVs[StatHP]) / 4)
	base_iv_ev_level_hp_term := (2*uint(base_stats[StatHP]) + uint(p.IVs[StatHP]) + uint(hp_ev_term)) * uint(p.Level)
	hp_floor_term := math.Floor(float64(base_iv_ev_level_hp_term) / 100)
	p.Stats[StatHP] = uint(hp_floor_term + float64(p.Level) + 10)
	p.CurrentHP = p.Stats[StatHP]

	// compute all other stats
	natureModifiers := p.Nature.getNatureModifers()
	for _, stat := range [5]int{StatAtk, StatDef, StatSpAtk, StatSpDef, StatSpeed} {
		ev_term := math.Floor(float64(p.EVs[stat]) / 4)
		base_iv_ev_level_term := (2*uint(base_stats[stat]) + uint(p.IVs[stat]) + uint(ev_term)) * uint(p.Level)
		floor_term := math.Floor(float64(base_iv_ev_level_term) / 100)
		nature_term := float64(floor_term+5) * natureModifiers[stat]
		p.Stats[stat] = uint(math.Floor(nature_term))
	}
}

// Stat getters return the effective stat values from modifiers
func (p *Pokemon) MaxHP() uint {
	// MaxHP is never modified by items/moves in battle?
	return p.Stats[StatHP]
}

func (p *Pokemon) Attack() uint {
	effective := float64(p.Stats[StatAtk])
	// TODO: attack modifiers
	return uint(effective)
}

func (p *Pokemon) Defense() uint {
	effective := float64(p.Stats[StatDef])
	// TODO: defense modifiers
	return uint(effective)
}

func (p *Pokemon) SpecialAttack() uint {
	effective := float64(p.Stats[StatSpAtk])
	// TODO: special attack modifiers
	return uint(effective)
}

func (p *Pokemon) SpecialDefense() uint {
	effective := float64(p.Stats[StatSpDef])
	// TODO: special defense modifiers
	return uint(effective)
}

func (p *Pokemon) Speed() uint {
	effective := float64(p.Stats[StatSpeed])
	// TODO: speed modifiers
	return uint(effective)
}

// display a Pokemon close to how it would appear in a Pokemon battle
func (p Pokemon) String() string {
	return p.GetName()
}

// Restore HP to a Pokemon. Can also be used to revive a fainted Pokemon.
func (p *Pokemon) RestoreHP(amount uint) Transaction {
	if diff := p.MaxHP() - p.CurrentHP; diff <= amount {
		amount = diff
	}
	return HealTransaction{Target: p, Amount: amount}
}

func (p *Pokemon) MarshalJSON() ([]byte, error) {
	type alias Pokemon // required to not enter infinite recursive loop
	return json.Marshal(&struct {
		Name string
		*alias
	}{
		Name:  p.GetName(),
		alias: (*alias)(p),
	})
}

func (p *Pokemon) UnmarshalJSON(data []byte) error {
	type alias Pokemon // required to not enter infinite recursive loop
	aux := &struct {
		*alias
	}{
		alias: (*alias)(p),
	}
	if p.metadata == nil {
		p.metadata = make(map[PokemonMeta]interface{})
	}
	return json.Unmarshal(data, &aux)
}
