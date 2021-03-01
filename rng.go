package pokemonbattlelib

// Provides an interface for generating pseudo random numbers.
type RNG interface {
	SetSeed(seed uint)
	Get(min, max int) int // Get a random number.
	Roll(x, y int) bool   // Generate a random boolean with a probability of x/y. Returns true x out of y times.
}

// Implements the LCRNG algorithm used by the original Pokemon Diamond/Pearl,
// and all other mainline Pokemon games that released on the GBA/DS.
// Note that this type must only be used in pointer form.
type LCRNG uint32

func (g *LCRNG) SetSeed(seed uint) {
	*g = LCRNG(seed)
}

func (g *LCRNG) Get(min, max int) int {
	*g = LCRNG(0x41C64E6D)*(*g) + LCRNG(0x00006073)
	diff := max - min
	return int(*g)%diff + min
}

// Rolls for a x/y chance, returning whether it was hit
func (g *LCRNG) Roll(x, y int) bool {
	return g.Get(1, y) <= x
}
