package pokemonbattlelib

type Nature struct {
	StatUp   int
	StatDown int
	name     string
}

func GetNatureTable() map[string]*Nature {
	//TODO: add all natures
	return map[string]*Nature{
		"adamant": {
			StatUp:   STAT_ATK,
			StatDown: STAT_SPATK,
			name:     "Adamant",
		},
		"hardy": {
			StatUp:   STAT_ATK,
			StatDown: STAT_ATK,
			name:     "Hardy",
		},
	}
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

func (n Nature) Equals(other *Nature) bool {
	return n.StatUp == other.StatUp &&
		n.StatDown == other.StatDown
}

// implement Stringer

func (n Nature) String() string {
	return n.name
}
