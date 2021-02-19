package pokemonbattlelib

import (
	"fmt"
	"io"
	"log"
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

// Constants for looking up Pokemon stats
const (
	STAT_HP = iota
	STAT_ATK
	STAT_DEF
	STAT_SPATK
	STAT_SPDEF
	STAT_SPD
)

// Constants for status effects on a Pokemon
const (
	STATUS_BURN = 1 << iota
	STATUS_FREEZE
	STATUS_PARALYZE
	STATUS_POISON
	STATUS_BADLY_POISON
	STATUS_SLEEP
)

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
	MAX_FRIENDSHIP    = 255
	MAX_EV            = 255
	MAX_IV            = 31
	MAX_STAT_MODIFIER = 6
	MIN_STAT_MODIFIER = -6
	TOTAL_EV          = 510
	MIN_LEVEL         = 1
	MAX_LEVEL         = 100
)

type GeneratePokemonOption func(p *Pokemon)

// Creates a new Pokemon given its national dex number
func GeneratePokemon(natdex uint16, opts ...GeneratePokemonOption) *Pokemon {
	p := &Pokemon{
		NatDex:          natdex,
		Level:           1,
		TotalExperience: 0,
		IVs:             [6]uint8{0, 0, 0, 0, 0, 0},
		EVs:             [6]uint8{0, 0, 0, 0, 0, 0},
		Nature:          GetNatureTable()["hardy"],
	}
	for _, opt := range opts {
		opt(p)
	}
	p.computeStats()
	return p
}

func WithLevel(level uint8) GeneratePokemonOption {
	return func(p *Pokemon) {
		p.Level = level
		p.TotalExperience = computeExpFromLevel(p.Level, p.GetGrowthRate())
	}
}

func WithTotalExp(totalExp uint) GeneratePokemonOption {
	return func(p *Pokemon) {
		p.TotalExperience = totalExp
		p.Level = computeLevelFromExp(p.TotalExperience, p.GetGrowthRate())
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

func computeLevelFromExp(exp uint, growth_rate int) uint8 {
	experience_csv := getCsvReader("data/experience.csv")

	var lastExp uint = 0
	for {
		record, err := experience_csv.Read()

		if err == io.EOF {
			break
		}

		nextExp := uint(parseInt(record[2]))
		if parseInt(record[0]) == growth_rate {
			if exp > lastExp && exp <= nextExp {
				return uint8(parseInt(record[1]))
			}
		}

		lastExp = nextExp
	}
	panic("There was a problem computing the level.")
}

func computeExpFromLevel(level uint8, growth_rate int) uint {
	experience_csv := getCsvReader("data/experience.csv")

	for {
		record, err := experience_csv.Read()

		if err == io.EOF {
			break
		}

		if parseInt(record[0]) == growth_rate {
			record_level := uint8(parseInt(record[1]))
			if level == record_level {
				return uint(parseInt(record[2]))
			}
		}
	}
	panic("There was a problem computing the experience.")
}

func (p *Pokemon) GetName() string {
	return PokemonNames[p.NatDex]
}

func (p *Pokemon) GetGrowthRate() int {
	return PokemonGrowthRates[int(p.NatDex)]
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

func (p *Pokemon) computeStats() {
	if !p.HasValidLevel() {
		log.Println(fmt.Sprintf("Failed to compute stats for %s, level %d invalid", p.GetName(), p.Level))
	}

	if !p.HasValidIVs() {
		log.Println(fmt.Sprintf("Failed to compute stats for %s, ivs %d invalid", p.GetName(), p.IVs))
	}

	if !p.HasValidEVs() {
		log.Println(fmt.Sprintf("Failed to compute stats for %s, evs %d invalid", p.GetName(), p.EVs))
	}

	// get base stats
	pokemon_stats_csv := getCsvReader("data/pokemon_stats.csv")
	var base_stats [6]uint
	already_read_stats := 0
	for {
		record, err := pokemon_stats_csv.Read()
		if p.NatDex == uint16(parseInt(record[0])) {
			base_stats[parseInt(record[1])-1] = uint(parseInt(record[2]))
			already_read_stats++
		}

		if already_read_stats == 6 {
			break
		}

		if err == io.EOF {
			break
		}
	}

	// compute HP
	hp_ev_term := math.Floor(float64(p.EVs[STAT_HP]) / 4)
	base_iv_ev_level_hp_term := (2*base_stats[STAT_HP] + uint(p.IVs[STAT_HP]) + uint(hp_ev_term)) * uint(p.Level)
	hp_floor_term := math.Floor(float64(base_iv_ev_level_hp_term) / 100)
	p.Stats[STAT_HP] = uint(hp_floor_term + float64(p.Level) + 10)
	p.CurrentHP = p.Stats[STAT_HP]

	// compute all other stats
	natureModifiers := p.Nature.getNatureModifers()
	for _, stat := range [5]int{STAT_ATK, STAT_DEF, STAT_SPATK, STAT_SPDEF, STAT_SPD} {
		ev_term := math.Floor(float64(p.EVs[stat]) / 4)
		base_iv_ev_level_term := (2*base_stats[stat] + uint(p.IVs[stat]) + uint(ev_term)) * uint(p.Level)
		floor_term := math.Floor(float64(base_iv_ev_level_term) / 100)
		nature_term := float64(floor_term+5) * natureModifiers[stat]
		p.Stats[stat] = uint(math.Floor(nature_term))
	}
}

func (p *Pokemon) VerboseString() string {
	outputString := ""
	outputString += "{\n" +
		fmt.Sprintf("\tNational dex number: %d\n", p.NatDex) +
		fmt.Sprintf("\tLevel: %d\n", p.Level) +
		fmt.Sprintf("\tAbility: %s\n", p.Ability) +
		fmt.Sprintf("\tTotal experience: %d\n", p.TotalExperience) +
		fmt.Sprintf("\tGender: %s\n", p.Gender) +
		fmt.Sprintf("\tIVs: %v\n", p.IVs) +
		fmt.Sprintf("\tEVs: %v\n", p.EVs) +
		fmt.Sprintf("\tNature: %s\n", p.Nature) +
		fmt.Sprintf("\tStats: %v\n", p.Stats) +
		fmt.Sprintf("\tCurrent HP: %d\n", p.CurrentHP) +
		fmt.Sprintf("\tHeld item: %v\n", p.HeldItem) +
		fmt.Sprintf("\tMoves: %v\n", p.Moves) +
		fmt.Sprintf("\tFriendship: %d\n", p.Friendship) +
		fmt.Sprintf("\tOriginal trainer ID: %d\n", p.OriginalTrainerID) +
		"}"
	return outputString
}

// implement Stringer

// display a Pokemon close to how it would appear in a Pokemon battle
func (p Pokemon) String() string {
	return fmt.Sprintf("%v%v\tLv%d\nHP: %d/%d\n", p.GetName(),
		p.Gender, p.Level, p.CurrentHP, p.Stats[STAT_HP])
}

// Restore HP to a Pokemon. Can also be used to revive a fainted Pokemon.
func (p *Pokemon) RestoreHP(amount uint) Transaction {
	if diff := p.Stats[STAT_HP] - p.CurrentHP; diff <= amount {
		amount = diff
	}
	return HealTransaction{Target: p, Amount: amount}
}
