package pokemonbattlelib

import (
	"fmt"
	"math"
	"math/bits"
)

type ElementalType uint32

const (
	Normal ElementalType = 1 << iota
	Fight
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
	Fight:    "Fight",
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

type Effectiveness float64

const (
	VeryIneffective    Effectiveness = 0.25
	Ineffective        Effectiveness = 0.5
	NoEffect           Effectiveness = 0
	NormalEffect       Effectiveness = 1
	SuperEffective     Effectiveness = 2
	VerySuperEffective Effectiveness = 4
)

var noEffect = map[ElementalType]ElementalType{
	Normal:   Ghost,
	Fight:    Ghost,
	Poison:   Steel,
	Ground:   Flying,
	Ghost:    Normal,
	Electric: Ground,
	Psychic:  Dark,
}
var halfEffect = map[ElementalType]ElementalType{
	Normal:   Rock | Steel,
	Fight:    Flying | Poison | Bug | Psychic,
	Flying:   Rock | Steel | Electric,
	Poison:   Poison | Ground | Rock | Ghost,
	Ground:   Bug | Grass,
	Rock:     Fight | Ground | Steel,
	Bug:      Fight | Flying | Poison | Ghost | Steel | Fire,
	Ghost:    Steel | Dark,
	Steel:    Steel | Fire | Water | Electric,
	Fire:     Rock | Fire | Water | Dragon,
	Water:    Water | Grass | Dragon,
	Grass:    Flying | Poison | Bug | Steel | Fire | Grass | Dragon,
	Electric: Grass | Electric | Dragon,
	Psychic:  Steel | Psychic,
	Ice:      Steel | Fire | Water | Ice,
	Dragon:   Steel,
	Dark:     Fight | Steel | Dark,
}
var doubleEffect = map[ElementalType]ElementalType{
	Fight:    Normal | Rock | Steel | Ice | Dark,
	Flying:   Fight | Bug | Grass,
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
	Psychic:  Fight | Poison,
	Ice:      Flying | Ground | Grass | Dragon,
	Dragon:   Dragon,
	Dark:     Ghost | Psychic,
}

func GetEffect(move, def ElementalType) Effectiveness {
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
