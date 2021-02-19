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
	StatModifiers     [6]int        // ranges from +6 (buffing) to -6 (debuffing) a stat
	StatusEffects     uint          // the current status effects inflicted on a Pokemon
	CurrentHP         uint          // the remaining HP of this Pokemon
	HeldItem          *Item         // the item a Pokemon is holding
	Moves             [4]*Move      // the moves the Pokemon currenly knows
	Friendship        uint8         // how close this Pokemon is to its Trainer
	OriginalTrainerID uint16        // a number associated with the first Trainer who caught this Pokemon
	Elemental         ElementalType // Indicates what type(s) (up to 2 simultaneously) this pokemon has
}

// Constants for IVs and EVs
const (
	MAX_FRIENDSHIP = 255
	MAX_EV         = 255
	MAX_IV         = 31
	TOTAL_EV       = 510
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

// implement Stringer

// display a Pokemon close to how it would appear in a Pokemon battle
func (p Pokemon) String() string {
	return fmt.Sprintf("%v%v\tLv%d\nHP: %d/%d\n", p.GetName(),
		p.Gender, p.Level, p.CurrentHP, p.Stats[STAT_HP])
}

// Applies damage to a Pokemon. May cause Pokemon to faint.
func (p *Pokemon) damage(amount uint) {
	if p.CurrentHP >= amount {
		p.CurrentHP -= amount
	} else {
		// prevent underflow
		p.CurrentHP = 0
	}
}

// Restore HP to a Pokemon. Can also be used to revive a fainted Pokemon.
func (p *Pokemon) heal(amount uint) {
	if diff := p.Stats[STAT_HP] - p.CurrentHP; diff <= amount {
		amount = diff
	}
	p.CurrentHP += amount
}
