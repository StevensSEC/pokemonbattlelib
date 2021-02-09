package pokemonbattlelib

import "fmt"

// Represents a Pokemon's move. Moves can deal damage, heal the user or allies, or cause status effects.
type Move struct {
	Name     string
	Type     string
	Category string
	Max_PP   int
	Priority int
	Power    int
	Accuracy int
}

func (m Move) String() string {
	return fmt.Sprintf("%v\nType: %v, Power: %v, Accuracy: %v\n", m.Name, m.Type, m.Power, m.Accuracy)
}
