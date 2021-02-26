package pokemonbattlelib

import (
	"fmt"
	"math"
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
	Immune             Effectiveness = 0
	VeryIneffective    Effectiveness = 0.25
	Ineffective        Effectiveness = 0.5
	NeutralEffect      Effectiveness = 1
	SuperEffective     Effectiveness = 2
	VerySuperEffective Effectiveness = 4
)

// Maps attacker, defender(s) to effectiveness
// If key is not found, defaults to NormalEffect
var typeChart = map[ElementalType]map[ElementalType]Effectiveness{
	Normal: {
		Rock | Steel: Ineffective,
		Ghost:        Immune,
	},
	Fighting: {
		Normal | Rock | Steel | Ice | Dark: SuperEffective,
		Flying | Poison | Bug | Psychic:    Ineffective,
		Ghost:                              Immune,
	},
	Flying: {
		Fighting | Bug | Grass:  SuperEffective,
		Rock | Steel | Electric: Ineffective,
	},
	Poison: {
		Grass:                          SuperEffective,
		Poison | Ground | Rock | Ghost: Ineffective,
		Steel:                          Immune,
	},
	Ground: {
		Poison | Rock | Steel | Fire | Electric: SuperEffective,
		Bug | Grass:                             Ineffective,
		Flying:                                  Immune,
	},
	Rock: {
		Flying | Bug | Fire | Ice: SuperEffective,
		Fighting | Ground | Steel: Ineffective,
	},
	Bug: {
		Grass | Psychic | Dark:                            SuperEffective,
		Fighting | Flying | Poison | Ghost | Steel | Fire: Ineffective,
	},
	Ghost: {
		Ghost | Psychic: SuperEffective,
		Steel | Dark:    Ineffective,
		Normal:          Immune,
	},
	Steel: {
		Rock | Ice:                      SuperEffective,
		Steel | Fire | Water | Electric: Ineffective,
	},
	Fire: {
		Bug | Steel | Ice:            SuperEffective,
		Rock | Fire | Water | Dragon: Ineffective,
	},
	Water: {
		Ground | Rock | Fire:   SuperEffective,
		Water | Grass | Dragon: Ineffective,
	},
	Grass: {
		Ground | Rock | Water: SuperEffective,
		Flying | Poison | Bug | Steel | Fire | Grass | Dragon: Ineffective,
	},
	Electric: {
		Flying | Water:            SuperEffective,
		Grass | Electric | Dragon: Ineffective,
		Ground:                    Immune,
	},
	Psychic: {
		Fighting | Poison: SuperEffective,
		Steel | Psychic:   Ineffective,
		Dark:              Immune,
	},
	Ice: {
		Flying | Ground | Grass | Dragon: SuperEffective,
		Steel | Fire | Water | Ice:       Ineffective,
	},
	Dragon: {
		Dragon: SuperEffective,
		Steel:  Ineffective,
	},
	Dark: {
		Ghost | Psychic:         SuperEffective,
		Fighting | Steel | Dark: Ineffective,
	},
}

// Get effectiveness of a given elemental type matchup.
func GetElementalEffect(attacking, defending ElementalType) Effectiveness {
	multiType := false
	total := NeutralEffect
	for elem := Normal; elem <= Dark; elem <<= 1 {
		if elem&defending == 0 {
			continue
		}
		for keys, effect := range typeChart[attacking] {
			if keys&defending != 0 {
				if !multiType {
					multiType = true
					total = effect
				} else {
					total *= effect
				}
			}
		}
	}
	return total
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
	ID int // The ID of the ability
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
