package pokemonbattlelib

import (
	"fmt"
	"math"
)

type Pokemon struct {
	NatDex            uint16        // National Pokedex Number
	Level             uint8         // value from 1-100 influencing stats
	Ability           *Ability      // name of this Pokemon's ability
	TotalExperience   uint          // the total amount of exp this Pokemon has gained, influencing its level
	Gender            Gender        // this Pokemon's gender
	IVs               [6]uint8      // values from 0-31 that represents a Pokemon's 'genetic' potential
	EVs               [6]uint8      // values from 0-255 that represents a Pokemon's training in a particular stat
	Nature            *Nature       // represents a Pokemon's disposition and affects stats
	Stats             [6]uint       // the actual stats of a Pokemon determined from the above data
	StatModifiers     [6]int        // ranges from +6 (buffing) to -6 (debuffing) a stat
	StatusEffects     uint          // the current status effects inflicted on a Pokemon
	CurrentHP         uint          // the remaining HP of this Pokemon
	HeldItem          *Item         // the item a Pokemon is holding
	Moves             [4]*Move      // the moves the Pokemon currenly knows
	Friendship        uint8         // how close this Pokemon is to its Trainer
	OriginalTrainerID uint16        // a number associated with the first Trainer who caught this Pokemon
	Elemental         ElementalType // Indicates what type(s) (up to 2 simultaneously) this pokemon has
}

// Constants for growth rates of a Pokemon
const (
	SLOW = iota + 1
	MEDIUM_FAST
	FAST
	MEDIUM_SLOW
	ERRATIC
	FLUCTUATING
)

// Constants for IVs and EVs
const (
	MAX_FRIENDSHIP = 255
	MAX_EV         = 255
	MAX_IV         = 31
	TOTAL_EV       = 510
	MIN_LEVEL      = 1
	MAX_LEVEL      = 100
)

type GeneratePokemonOption func(p *Pokemon)

// Creates a new Pokemon given its national dex number
func GeneratePokemon(natdex int, opts ...GeneratePokemonOption) *Pokemon {
	p := &Pokemon{
		NatDex:          uint16(natdex),
		Level:           1,
		TotalExperience: 0,
		IVs:             [6]uint8{0, 0, 0, 0, 0, 0},
		EVs:             [6]uint8{0, 0, 0, 0, 0, 0},
		Nature:          GetNature(HARDY), // this nature is neutral and has no effect
	}
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

func WithNature(nature *Nature) GeneratePokemonOption {
	return func(p *Pokemon) {
		p.Nature = nature
	}
}

func (p *Pokemon) GetName() string {
	return pokemonNames[p.NatDex]
}

func (p *Pokemon) GetGrowthRate() int {
	return pokemonGrowthRates[int(p.NatDex)]
}

func (p *Pokemon) GetBaseStats() [6]int {
	return pokemonBaseStats[int(p.NatDex)]
}

func (p *Pokemon) HasValidLevel() bool {
	return p.Level >= MIN_LEVEL && p.Level <= MAX_LEVEL
}

func (p *Pokemon) HasValidIVs() bool {
	for _, IV := range p.IVs {
		if IV > MAX_IV {
			return false
		}
	}
	return true
}

func (p *Pokemon) HasValidEVs() bool {
	totalEVs := 0
	for _, EV := range p.EVs {
		if EV > MAX_EV {
			return false
		}
		totalEVs += int(EV)
	}
	return totalEVs <= TOTAL_EV
}

// Increases a Pokemon's level by `levels` and sets their total experience to
// the minimum amount for that level.
func (p *Pokemon) GainLevels(levels int) {
	newLevel := int(p.Level) + levels
	if newLevel > MAX_LEVEL {
		panic(fmt.Sprintf("Level %d %s cannot level up to level %d!", p.Level, p.GetName(), newLevel))
	}

	if levels < 0 {
		panic(fmt.Sprintf("%s's level tried to decrease by %d", p.GetName(), levels))
	}

	if levels == 0 {
		return
	}

	p.Level = uint8(newLevel)
	p.TotalExperience = uint(EXP_TABLE[p.GetGrowthRate()][int(p.Level)])
	p.computeStats()
}

// Increases a Pokemon's experience points by `exp` and levels up the Pokemon accordingly.
func (p *Pokemon) GainExperience(exp int) {

	if exp < 0 {
		panic(fmt.Sprintf("%s's experience tried to decrease by %d", p.GetName(), exp))
	}

	if exp == 0 {
		return
	}

	max_exp := EXP_TABLE[p.GetGrowthRate()][MAX_LEVEL]

	// if would gain experience beyond leveling to 100, set level to 100
	if int(p.TotalExperience)+exp > max_exp {
		p.GainLevels(MAX_LEVEL - int(p.Level))
		return
	}

	remaining_exp := exp
	toNextLevel := EXP_TABLE[p.GetGrowthRate()][int(p.Level+1)] - int(p.TotalExperience)

	for remaining_exp > toNextLevel {
		p.GainLevels(1)
		remaining_exp -= toNextLevel
		toNextLevel = EXP_TABLE[p.GetGrowthRate()][int(p.Level+1)] - int(p.TotalExperience)
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
	hp_ev_term := math.Floor(float64(p.EVs[STAT_HP]) / 4)
	base_iv_ev_level_hp_term := (2*uint(base_stats[STAT_HP]) + uint(p.IVs[STAT_HP]) + uint(hp_ev_term)) * uint(p.Level)
	hp_floor_term := math.Floor(float64(base_iv_ev_level_hp_term) / 100)
	p.Stats[STAT_HP] = uint(hp_floor_term + float64(p.Level) + 10)
	p.CurrentHP = p.Stats[STAT_HP]

	// compute all other stats
	natureModifiers := p.Nature.getNatureModifers()
	for _, stat := range [5]int{STAT_ATK, STAT_DEF, STAT_SPATK, STAT_SPDEF, STAT_SPD} {
		ev_term := math.Floor(float64(p.EVs[stat]) / 4)
		base_iv_ev_level_term := (2*uint(base_stats[stat]) + uint(p.IVs[stat]) + uint(ev_term)) * uint(p.Level)
		floor_term := math.Floor(float64(base_iv_ev_level_term) / 100)
		nature_term := float64(floor_term+5) * natureModifiers[stat]
		p.Stats[stat] = uint(math.Floor(nature_term))
	}
}

// implement Stringer

// display a Pokemon close to how it would appear in a Pokemon battle
func (p Pokemon) String() string {
	return fmt.Sprintf("%v%v\tLv%d\nHP: %d/%d\n", p.GetName(),
		p.Gender, p.Level, p.CurrentHP, p.Stats[STAT_HP])
}
