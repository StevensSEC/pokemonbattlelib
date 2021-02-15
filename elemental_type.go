package pokemonbattlelib

type ElementalType uint32

var effect int

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

func GetEffect(move, def ElementalType) int {
	if noEffect[move]&def == def {
		effect = 0
	} else if halfEffect[move]&def == def {
		effect = 50
	} else if doubleEffect[move]&def == def {
		effect = 200
	} else {
		effect = 100
	}

	return effect
}
