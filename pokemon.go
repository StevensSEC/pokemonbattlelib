package pokemonbattlelib

import (
	"fmt"
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

// Retrieves a Pokemon given its national dex number
func GetPokemon(natdex uint16) Pokemon {
	for _, p := range ALL_POKEMON {
		if p.NatDex == natdex {
			return p
		}
	}
	// Not exactly the best way to handle this
	panic(fmt.Sprintf("unknown Pokedex number %v\n", natdex))
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

// implement Stringer

// display a Pokemon close to how it would appear in a Pokemon battle
func (p Pokemon) String() string {
	return fmt.Sprintf("%v%v\tLv%d\nHP: %d/%d\n", p.GetName(),
		p.Gender, p.Level, p.CurrentHP, p.Stats[STAT_HP])
}
