package pokemonbattlelib

type RNG interface {
	SetSeed(seed uint)
	Get(min, max int) int
}

type LCRNG uint32

func (g *LCRNG) SetSeed(seed uint) {
	*g = LCRNG(seed)
}

func (g *LCRNG) Get(min, max int) int {
	*g = LCRNG(0x41C64E6D)*(*g) + LCRNG(0x00006073)
	diff := max - min
	return int(*g)%diff + min
}

// Rolls for a 1/n chance, returning whether it was hit
func (g *LCRNG) Roll(n int) bool {
	return g.Get(1, n) == 1
}
