package pokemonbattlelib

type Gender int

const (
	Genderless Gender = iota
	Female
	Male
)

// implement Stringer

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
