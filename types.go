package pokemonbattlelib

import (
	"fmt"
	"math"
	"math/bits"
)

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

const (
	MAX_STAT_MODIFIER = 6
	MIN_STAT_MODIFIER = -6
)

// Represents a bit mask of all elemental types
type ElementalType uint32

const (
	Normal ElementalType = 1 << iota
	Fighting
	Flying
	Poison
	Ground
	Rock
	Bug
	Ghost
	Steel
	Fire
	Water
	Grass
	Electric
	Psychic
	Ice
	Dragon
	Dark
)

var elementalTypeStrings = map[ElementalType]string{
	Normal:   "Normal",
	Fighting: "Fighting",
	Flying:   "Flying",
	Poison:   "Poison",
	Ground:   "Ground",
	Rock:     "Rock",
	Bug:      "Bug",
	Ghost:    "Ghost",
	Steel:    "Steel",
	Fire:     "Fire",
	Water:    "Water",
	Grass:    "Grass",
	Electric: "Electric",
	Psychic:  "Psychic",
	Ice:      "Ice",
	Dragon:   "Dragon",
	Dark:     "Dark",
}

// Represents effectiveness of an elemental type matchup.
type Effectiveness float64

const (
	NoEffect           Effectiveness = 0
	VeryIneffective    Effectiveness = 0.25
	Ineffective        Effectiveness = 0.5
	NormalEffect       Effectiveness = 1
	SuperEffective     Effectiveness = 2
	VerySuperEffective Effectiveness = 4
)

var noEffect = map[ElementalType]ElementalType{
	Normal:   Ghost,
	Fighting: Ghost,
	Poison:   Steel,
	Ground:   Flying,
	Ghost:    Normal,
	Electric: Ground,
	Psychic:  Dark,
}

var halfEffect = map[ElementalType]ElementalType{
	Normal:   Rock | Steel,
	Fighting: Flying | Poison | Bug | Psychic,
	Flying:   Rock | Steel | Electric,
	Poison:   Poison | Ground | Rock | Ghost,
	Ground:   Bug | Grass,
	Rock:     Fighting | Ground | Steel,
	Bug:      Fighting | Flying | Poison | Ghost | Steel | Fire,
	Ghost:    Steel | Dark,
	Steel:    Steel | Fire | Water | Electric,
	Fire:     Rock | Fire | Water | Dragon,
	Water:    Water | Grass | Dragon,
	Grass:    Flying | Poison | Bug | Steel | Fire | Grass | Dragon,
	Electric: Grass | Electric | Dragon,
	Psychic:  Steel | Psychic,
	Ice:      Steel | Fire | Water | Ice,
	Dragon:   Steel,
	Dark:     Fighting | Steel | Dark,
}

var doubleEffect = map[ElementalType]ElementalType{
	Fighting: Normal | Rock | Steel | Ice | Dark,
	Flying:   Fighting | Bug | Grass,
	Poison:   Grass,
	Ground:   Poison | Rock | Steel | Fire | Electric,
	Rock:     Flying | Bug | Fire | Ice,
	Bug:      Grass | Psychic | Dark,
	Ghost:    Ghost | Psychic,
	Steel:    Rock | Ice,
	Fire:     Bug | Steel | Grass | Ice,
	Water:    Ground | Rock | Fire,
	Grass:    Ground | Rock | Water,
	Electric: Flying | Water,
	Psychic:  Fighting | Poison,
	Ice:      Flying | Ground | Grass | Dragon,
	Dragon:   Dragon,
	Dark:     Ghost | Psychic,
}

// Get effectiveness of a given elemental type matchup.
func GetElementalEffect(move, def ElementalType) Effectiveness {
	if noEffect[move]&def > 0 {
		return NoEffect
	}

	reduce := bits.OnesCount32(uint32(halfEffect[move] & def))
	increase := bits.OnesCount32(uint32(doubleEffect[move] & def))
	effect := (increase - reduce) * 2
	if effect == 0 {
		return NormalEffect
	} else if effect > 0 {
		return Effectiveness(effect)
	} else {
		return Effectiveness(1 / math.Abs(float64(effect)))
	}
}

func (e ElementalType) String() string {
	result := ""
	for i := Normal; i <= Dark; i <<= 1 {
		if e&i > 0 {
			result += fmt.Sprintf("[%s]", elementalTypeStrings[i])
		}
	}
	return result
}

type Gender int

const (
	Genderless Gender = iota
	Female
	Male
)

func (g Gender) String() string {
	if g == Genderless {
		return ""
	} else if g == Female {
		return "♀"
	} else if g == Male {
		return "♂"
	} else {
		panic("Stringing gender reached unhandled condition")
	}
}

type Nature struct {
	StatUp   string
	StatDown string
}

type Ability struct {
	//TODO
}
