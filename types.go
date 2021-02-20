package pokemonbattlelib

import (
	"fmt"
	"math"
	"math/bits"
	"strings"
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

// Represents volatile and non-volatile status conditions on a pokemon.
type StatusCondition uint32

const (
	StatusNone StatusCondition = 0

	// Non volatile - first 3 bits
	StatusBurn StatusCondition = iota
	StatusFreeze
	StatusParalyze
	StatusPoison
	StatusBadlyPoison
	StatusSleep
)
const (
	// volatile - the rest of the bits
	StatusBound StatusCondition = 1 << (iota + 3)
	StatusCantEscape
	StatusConfusion
	StatusCursed
	StatusEmbargo
	StatusFlinch
	StatusHealBlock
	StatusIdentified
	StatusInfatuation
	StatusLeechSeed
	StatusNightmare
	StatusPerishSong
	StatusTaunt
	StatusTorment
)

// Status effect strings in past tense. Mostly for use in battle log.
var statusStrings = map[StatusCondition]string{
	StatusNone:        "",
	StatusBurn:        "burned",
	StatusFreeze:      "frozen",
	StatusParalyze:    "paralyzed",
	StatusPoison:      "poisoned",
	StatusBadlyPoison: "badly poisoned",
	StatusSleep:       "asleep",
	StatusBound:       "bound",
	StatusCantEscape:  "can't escape",
	StatusConfusion:   "confused",
	StatusCursed:      "cursed",
	StatusEmbargo:     "embargoed",
	StatusFlinch:      "flinched",
	StatusHealBlock:   "heal blocked",
	StatusInfatuation: "infatuated",
	StatusIdentified:  "identified",
	StatusLeechSeed:   "leeched",
	StatusNightmare:   "nightmared",
	StatusPerishSong:  "perished",
	StatusTaunt:       "taunted",
	StatusTorment:     "tormented",
}

const (
	NONVOLATILE_STATUS_MASK StatusCondition = 0b111
	VOLATILE_STATUS_MASK    StatusCondition = math.MaxUint32 ^ NONVOLATILE_STATUS_MASK
)

func (s *StatusCondition) check(c StatusCondition) bool {
	vCheck := (*s&VOLATILE_STATUS_MASK)&(c&VOLATILE_STATUS_MASK) == c&VOLATILE_STATUS_MASK
	if c&NONVOLATILE_STATUS_MASK > 0 {
		nvCheck := (*s&NONVOLATILE_STATUS_MASK)^(c&NONVOLATILE_STATUS_MASK) == 0
		return nvCheck && vCheck
	} else {
		return vCheck
	}
}

func (s *StatusCondition) apply(c StatusCondition) {
	nv := c & NONVOLATILE_STATUS_MASK
	if nv == 0 {
		nv = *s & NONVOLATILE_STATUS_MASK
	}
	v := VOLATILE_STATUS_MASK & (*s | c)
	*s = nv | v
}

func (s *StatusCondition) clear(c StatusCondition) {
	*s ^= c
}

func (s StatusCondition) String() string {
	// fast path
	if val, ok := statusStrings[s]; ok {
		return val
	}

	// slower path
	result := []string{}
	for c := StatusBurn; c <= StatusSleep; c++ {
		if s.check(c) {
			result = append(result, statusStrings[c])
		}
	}
	for c := StatusBound; c <= StatusTorment; c <<= 1 {
		if s.check(c) {
			result = append(result, statusStrings[c])
		}
	}
	return strings.Join(result, ", ")
}

type Ability struct {
	//TODO
}

type Nature struct {
	StatUp   int
	StatDown int
	name     string
}

// Constants for looking up natures
const (
	HARDY = iota + 1
	LONELY
	ADAMANT
	NAUGHTY
	BRAVE
	BOLD
	DOCILE
	IMPISH
	LAX
	RELAXED
	MODEST
	MILD
	BASHFUL
	RASH
	QUIET
	CALM
	GENTLE
	CAREFUL
	QUIRKY
	SASSY
	TIMID
	HASTY
	JOLLY
	NAIVE
	SERIOUS
)

func GetNature(nature int) *Nature {
	natures := map[int]*Nature{
		//TODO: add all natures
		HARDY: {
			StatUp:   STAT_ATK,
			StatDown: STAT_ATK,
			name:     "Hardy",
		},
		ADAMANT: {
			StatUp:   STAT_ATK,
			StatDown: STAT_SPATK,
			name:     "Adamant",
		},
	}
	return natures[nature]
}

func (n Nature) getNatureModifers() [6]float64 {
	natureModifiers := [6]float64{-1, 1, 1, 1, 1, 1} // hp is not affected by nature
	natureModifiers[n.StatUp] = 1.1
	natureModifiers[n.StatDown] = 0.9

	// tried to multiply natureModifiers by both 1.1 and 0.9, caused rounding errors
	if n.StatUp == n.StatDown {
		natureModifiers[n.StatUp] = 1
	}

	return natureModifiers
}
