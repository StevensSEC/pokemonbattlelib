package pokemonbattlelib

import (
	"fmt"
	"io"
	"log"
	"math"
	"reflect"
)

type Pokemon struct {
	NatDex            uint16   // National Pokedex Number
	Level             uint8    // value from 1-100 influencing stats
	Ability           *Ability // name of this Pokemon's ability
	TotalExperience   uint     // the total amount of exp this Pokemon has gained, influencing its level
	Gender            Gender   // this Pokemon's gender
	IVs               [6]uint8 // values from 0-31 that represents a Pokemon's 'genetic' potential
	EVs               [6]uint8 // values from 0-255 that represents a Pokemon's training in a particular stat
	Nature            *Nature  // represents a Pokemon's disposition and affects stats
	Stats             [6]uint  // the actual stats of a Pokemon determined from the above data
	CurrentHP         uint     // the remaining HP of this Pokemon
	HeldItem          *Item    // the item a Pokemon is holding
	Moves             [4]*Move // the moves the Pokemon currenly knows
	Friendship        uint8    // how close this Pokemon is to its Trainer
	OriginalTrainerID uint16   // a number associated with the first Trainer who caught this Pokemon
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

// Creates a new Pokemon given its national dex number
func NewPokemon(args ...interface{}) Pokemon {
	// Mandatory parameters
	var natdex uint16
	var level uint8

	// Optional parameters
	var totalExp uint // used to compute level if provided
	var ivs [6]uint8 = [6]uint8{0, 0, 0, 0, 0, 0}
	var evs [6]uint8 = [6]uint8{0, 0, 0, 0, 0, 0}
	var nature *Nature = GetNatureTable()["hardy"] // this nature provides no bonuses/debuffs

	if len(args) < 2 {
		panic("Not enough parameters.")
	}

	for i, p := range args {
		switch i {
		case 0: // natdex
			param, ok := p.(uint16)
			if !ok {
				panic(fmt.Sprintf("Parameter %d: expected type uint16, got %v", i, reflect.TypeOf(p)))
			}
			natdex = param
		case 1: // level or total exp gained
			switch p.(type) {
			case uint8: // level
				level = p.(uint8)
				totalExp = computeExpFromLevel(level, getGrowthRateFromDexNumber(natdex))
			case uint: // total experience
				totalExp = p.(uint)
				level = computeLevelFromExp(totalExp, getGrowthRateFromDexNumber(natdex))
			default:
				panic(fmt.Sprintf("Parameter %d: expected type uint8 or uint, got %v", i, reflect.TypeOf(p)))
			}
		case 2: // ivs
			param, ok := p.([6]uint8)
			if !ok {
				panic(fmt.Sprintf("Parameter %d: expected type [6]uint8, got %v", i, reflect.TypeOf(p)))
			}
			ivs = param
		case 3: //evs
			param, ok := p.([6]uint8)
			if !ok {
				panic(fmt.Sprintf("Parameter %d: expected type [6]uint8, got %v", i, reflect.TypeOf(p)))
			}
			evs = param
		case 4: // nature
			param, ok := p.(*Nature)
			if !ok {
				panic(fmt.Sprintf("Parameter %d: expected type Nature, got %v", i, reflect.TypeOf(p)))
			}
			nature = param
		}
	}

	for _, p := range ALL_POKEMON {
		if p.NatDex == natdex {
			p.Level = level
			p.TotalExperience = totalExp
			p.IVs = ivs
			p.EVs = evs
			p.Nature = nature
			p.computeStats()
			return p
		}
	}
	// Not exactly the best way to handle this
	panic(fmt.Sprintf("unknown Pokedex number %v\n", natdex))
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

func getGrowthRateFromDexNumber(natdex uint16) int {
	pokemon_species_csv := getCsvReader("data/pokemon_species.csv")
	for {
		record, err := pokemon_species_csv.Read()

		if err == io.EOF {
			break
		}

		if uint16(parseInt(record[0])) == natdex {
			return parseInt(record[14])
		}
	}
	panic(fmt.Sprintf("Unknown national dex number %d", natdex))
}

func (p *Pokemon) GetName() string {
	return PokemonNames[p.NatDex]
}

func (p *Pokemon) HasValidLevel() bool {
	return p.Level > 1 && p.Level <= 100
}

func (p *Pokemon) HasValidIVs() bool {
	for _, IV := range p.IVs {
		if IV > 31 {
			return false
		}
	}
	return true
}

func (p *Pokemon) HasValidEVs() bool {
	totalEVs := 0
	for _, EV := range p.EVs {
		totalEVs += int(EV)
	}
	return totalEVs <= 510
}

func (p *Pokemon) computeStats() {
	if !p.HasValidLevel() {
		log.Println(fmt.Printf("Failed to compute stats for %d, level %d invalid", p.NatDex, p.Level))
	}

	if !p.HasValidIVs() {
		log.Println(fmt.Printf("Failed to compute stats for %d, ivs %d invalid", p.NatDex, p.IVs))
	}

	if !p.HasValidEVs() {
		log.Println(fmt.Printf("Failed to compute stats for %d, evs %d invalid", p.NatDex, p.EVs))
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

func (p *Pokemon) Equals(other Pokemon) bool {
	return p.NatDex == other.NatDex &&
		p.Level == other.Level &&
		p.TotalExperience == other.TotalExperience &&
		p.Gender == other.Gender &&
		p.IVs == other.IVs &&
		p.EVs == other.EVs &&
		p.Nature.Equals(other.Nature) &&
		p.Stats == other.Stats &&
		p.CurrentHP == other.CurrentHP
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
