package pokemonbattlelib

type ExpGroup int

const (
    InvalidGroup ExpGroup = iota
    Erratic
    Fast
    MediumFast
    MediumSlow
    Slow
    Fluctuating
)

