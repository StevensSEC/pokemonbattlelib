package pokemonbattlelib

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

// implement Stringer

func (n Nature) String() string {
	return n.name
}
