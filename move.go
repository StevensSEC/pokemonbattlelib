package pokemonbattlelib

import "fmt"

// Represents a Pokemon's move. Moves can deal damage, heal the user or allies, or cause status effects.
type Move struct{
	Name string
	Type string
	Category string
	Max_PP int
	Priority int
	Power int
	Accuracy int
	//Side effects not implemented yet
}

//Stringer implementation
func (m Move) String() string{
	return fmt.Sprintf("Name: %v\nType: %v\nCategory: %v\nMax PP: %v\nPriority: %v\nPower: %v\nAccuracy: %v\n", m.Name, m.Type, m.Category, m.Max_PP, m.Priority, m.Power, m.Accuracy);
}
