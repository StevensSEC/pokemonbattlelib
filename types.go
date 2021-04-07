package pokemonbattlelib

import (
	"fmt"
	"math"
	"math/bits"
	"strings"
)

// Constants for looking up Pokemon stats
const (
	// Base stats
	StatHP = iota
	StatAtk
	StatDef
	StatSpAtk
	StatSpDef
	StatSpeed
	// Fighting stats
	StatCritChance
	StatAccuracy
	StatEvasion
)

// Table of critical hit chances (denominator of 1/X)
var CritChances = [5]int{16, 8, 4, 3, 2}

// Table for accuracy/evasion multiplier
var AccuracyEvasionMultiplier = map[int]uint{
	-6: 300,
	-5: 266,
	-4: 250,
	-3: 200,
	-2: 166,
	-1: 133,
	0:  100,
	1:  75,
	2:  60,
	3:  50,
	4:  43,
	5:  36,
	6:  33,
}

// The min and max number of stages that a stat can be modified is [-6, 6]
const (
	MaxStatModifier = 6
	MinStatModifier = -6
)

// Represents a bit mask of all elemental types
type Type uint32

const (
	TypeNormal Type = 1 << iota
	TypeFighting
	TypeFlying
	TypePoison
	TypeGround
	TypeRock
	TypeBug
	TypeGhost
	TypeSteel
	TypeFire
	TypeWater
	TypeGrass
	TypeElectric
	TypePsychic
	TypeIce
	TypeDragon
	TypeDark
)

var typeStrings = map[Type]string{
	TypeNormal:   "Normal",
	TypeFighting: "Fighting",
	TypeFlying:   "Flying",
	TypePoison:   "Poison",
	TypeGround:   "Ground",
	TypeRock:     "Rock",
	TypeBug:      "Bug",
	TypeGhost:    "Ghost",
	TypeSteel:    "Steel",
	TypeFire:     "Fire",
	TypeWater:    "Water",
	TypeGrass:    "Grass",
	TypeElectric: "Electric",
	TypePsychic:  "Psychic",
	TypeIce:      "Ice",
	TypeDragon:   "Dragon",
	TypeDark:     "Dark",
}

// Represents effectiveness of an elemental type matchup.
type Effectiveness int8

const NoEffect Effectiveness = math.MinInt8 + 1 // separate because it fucks up the iota
const (
	VeryIneffective Effectiveness = iota - 2
	Ineffective
	NormalEffect
	SuperEffective
	VerySuperEffective
)

var noEffect = map[Type]Type{
	TypeNormal:   TypeGhost,
	TypeFighting: TypeGhost,
	TypePoison:   TypeSteel,
	TypeGround:   TypeFlying,
	TypeGhost:    TypeNormal,
	TypeElectric: TypeGround,
	TypePsychic:  TypeDark,
}

var halfEffect = map[Type]Type{
	TypeNormal:   TypeRock | TypeSteel,
	TypeFighting: TypeFlying | TypePoison | TypeBug | TypePsychic,
	TypeFlying:   TypeRock | TypeSteel | TypeElectric,
	TypePoison:   TypePoison | TypeGround | TypeRock | TypeGhost,
	TypeGround:   TypeBug | TypeGrass,
	TypeRock:     TypeFighting | TypeGround | TypeSteel,
	TypeBug:      TypeFighting | TypeFlying | TypePoison | TypeGhost | TypeSteel | TypeFire,
	TypeGhost:    TypeSteel | TypeDark,
	TypeSteel:    TypeSteel | TypeFire | TypeWater | TypeElectric,
	TypeFire:     TypeRock | TypeFire | TypeWater | TypeDragon,
	TypeWater:    TypeWater | TypeGrass | TypeDragon,
	TypeGrass:    TypeFlying | TypePoison | TypeBug | TypeSteel | TypeFire | TypeGrass | TypeDragon,
	TypeElectric: TypeGrass | TypeElectric | TypeDragon,
	TypePsychic:  TypeSteel | TypePsychic,
	TypeIce:      TypeSteel | TypeFire | TypeWater | TypeIce,
	TypeDragon:   TypeSteel,
	TypeDark:     TypeFighting | TypeSteel | TypeDark,
}

var doubleEffect = map[Type]Type{
	TypeFighting: TypeNormal | TypeRock | TypeSteel | TypeIce | TypeDark,
	TypeFlying:   TypeFighting | TypeBug | TypeGrass,
	TypePoison:   TypeGrass,
	TypeGround:   TypePoison | TypeRock | TypeSteel | TypeFire | TypeElectric,
	TypeRock:     TypeFlying | TypeBug | TypeFire | TypeIce,
	TypeBug:      TypeGrass | TypePsychic | TypeDark,
	TypeGhost:    TypeGhost | TypePsychic,
	TypeSteel:    TypeRock | TypeIce,
	TypeFire:     TypeBug | TypeSteel | TypeGrass | TypeIce,
	TypeWater:    TypeGround | TypeRock | TypeFire,
	TypeGrass:    TypeGround | TypeRock | TypeWater,
	TypeElectric: TypeFlying | TypeWater,
	TypePsychic:  TypeFighting | TypePoison,
	TypeIce:      TypeFlying | TypeGround | TypeGrass | TypeDragon,
	TypeDragon:   TypeDragon,
	TypeDark:     TypeGhost | TypePsychic,
}

// Get effectiveness of a given elemental type matchup.
func GetElementalEffect(move, def Type, item Item) Effectiveness {
	reduce := bits.OnesCount32(uint32(halfEffect[move] & def))
	increase := bits.OnesCount32(uint32(doubleEffect[move] & def))
	effect := increase - reduce
	if noEffect[move]&def > 0 {
		return NoEffect
	}
	return Effectiveness(effect)
}

func (t Type) String() string {
	result := ""
	for i := TypeNormal; i <= TypeDark; i <<= 1 {
		if t&i > 0 {
			result += fmt.Sprintf("[%s]", typeStrings[i])
		}
	}
	return result
}

type Gender int

const (
	GenderGenderless Gender = iota
	GenderFemale
	GenderMale
)

func (g Gender) String() string {
	if g == GenderGenderless {
		return ""
	} else if g == GenderFemale {
		return "♀"
	} else if g == GenderMale {
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
	StatusNonvolatileMask StatusCondition = 0b111
	StatusVolatileMask    StatusCondition = math.MaxUint32 ^ StatusNonvolatileMask
)

func (s *StatusCondition) check(c StatusCondition) bool {
	vCheck := (*s&StatusVolatileMask)&(c&StatusVolatileMask) == c&StatusVolatileMask
	if c&StatusNonvolatileMask > 0 {
		nvCheck := (*s&StatusNonvolatileMask)^(c&StatusNonvolatileMask) == 0
		return nvCheck && vCheck
	} else {
		return vCheck
	}
}

// Applies the given status conditions to this status condition. Non-volatile status conditions are overwritten.
func (s *StatusCondition) apply(c StatusCondition) {
	nv := c & StatusNonvolatileMask
	if nv == 0 {
		nv = *s & StatusNonvolatileMask
	}
	v := StatusVolatileMask & (*s | c)
	*s = nv | v
}

// Clears the given status conditions from this status condition.
func (s *StatusCondition) clear(c StatusCondition) {
	if c != StatusNonvolatileMask && *s&StatusNonvolatileMask != c&StatusNonvolatileMask {
		c &= StatusVolatileMask
	}
	*s &= ^c
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

type Ability uint16

// Natures can affect a Pokemon's stats, increasing one and decreasing another.
// *Constants, GetStatModifiers(), and String() are auto-generated.*
type Nature uint8

// Deprecated: getNatureModifers is deprecated. Use Nature.GetStatModifiers() instead.
func (n Nature) getNatureModifers() [6]float64 {
	up, down := n.GetStatModifiers()
	natureModifiers := [6]float64{-1, 1, 1, 1, 1, 1} // hp is not affected by nature

	// tried to multiply natureModifiers by both 1.1 and 0.9, caused rounding errors
	if up != down {
		natureModifiers[up] = 1.1
		natureModifiers[down] = 0.9
	}

	return natureModifiers
}

// Weather effects
type Weather int

const (
	WeatherClearSkies Weather = iota
	WeatherHarshSunlight
	WeatherRain
	WeatherSandstorm
	WeatherHail
	WeatherFog
)

type MoveFailReason uint8

const (
	FailOther MoveFailReason = iota
	FailMiss
	FailDodge
)
