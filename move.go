package pokemonbattlelib

import "fmt"

type MoveCategory int

const (
	Status MoveCategory = iota + 1
	Physical
	Special
)

func (c MoveCategory) String() string {
	switch c {
	case Status:
		return "Status"
	case Physical:
		return "Physical"
	case Special:
		return "Special"
	default:
		panic("Unexpected value for move category")
	}
}

// Represents a Pokemon's move. Moves can deal damage, heal the user or allies, or cause status effects.
type Move struct {
	ID       int
	Name     string
	Type     int // update this to elementalType when PR #27 is complete
	Category MoveCategory
	Max_PP   int
	Priority int
	Power    int
	Accuracy int
}

func (m Move) String() string {
	return fmt.Sprintf("%v\nType: %v, Power: %v, Accuracy: %v\n", m.Name, m.Type, m.Power, m.Accuracy)
}
