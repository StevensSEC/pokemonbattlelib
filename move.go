package pokemonbattlelib

import "fmt"

type Category int

const (
	status Category = iota + 1
	physical
	special
)

func (c Category) String() string {
	switch c {
	case status:
		return "Status"
	case physical:
		return "Physical"
	case special:
		return "Special"
	default:
		panic("Unexpected value for move category")
	}
}

// Represents a Pokemon's move. Moves can deal damage, heal the user or allies, or cause status effects.
type Move struct {
	ID       int
	Name     string
	Type     int // update this to elementalType when PR #27 is closed
	Category Category
	Max_PP   int
	Priority int
	Power    int
	Accuracy int
}

func (m Move) String() string {
	return fmt.Sprintf("%v\nType: %v, Power: %v, Accuracy: %v\n", m.Name, m.Type, m.Power, m.Accuracy)
}
