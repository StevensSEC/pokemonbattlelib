package pokemonbattlelib

type Nature struct {
	StatUp   int
	StatDown int
}

var (
    Adamant *Nature = &Nature{
        StatUp: STAT_ATK,
        StatDown: STAT_DEF,
    }

    Hardy = &Nature{
        StatUp: STAT_ATK,
        StatDown: STAT_ATK,
    }
)
