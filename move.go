package pokemonbattlelib

import "fmt"

type MoveCategory uint8

const (
	MoveCategoryStatus MoveCategory = iota
	MoveCategoryPhysical
	MoveCategorySpecial
)

func (c MoveCategory) String() string {
	switch c {
	case MoveCategoryStatus:
		return "Status"
	case MoveCategoryPhysical:
		return "Physical"
	case MoveCategorySpecial:
		return "Special"
	default:
		panic("Unexpected value for move category")
	}
}

// Sets the bounds on move priority to [-7, 5]
const (
	MovePriorityMin = 5
	MovePriorityMax = -7
)

// Targets that the move can specify
type MoveTarget int

const (
	TARGET_SPECIFIC_MOVE MoveTarget = iota + 1
	TARGET_SELECTED_ME_FIRST
	TARGET_ALLY
	TARGET_USERS_FIELD
	TARGET_USER_OR_ALLY
	TARGET_OPPONENTS_FIELD
	TARGET_USER
	TARGET_RANDOM_OPPONENT
	TARGET_ALL_OTHERS
	TARGET_SELECTED
	TARGET_ALL_OPPONENTS
	TARGET_ENTIRE_FIELD
	TARGET_USER_AND_ALLIES
	TARGET_ALL
)

// Represents a Pokemon's move. Moves can deal damage, heal the user or allies, or cause status effects.
type Move struct {
	ID        int
	Name      string
	Type      ElementalType
	Category  MoveCategory
	Targets   MoveTarget
	CurrentPP int
	MaxPP     int
	Priority  int
	Power     int
	Accuracy  int
}

// Retrieves a Pokemon move given its move ID
func GetMove(id int) *Move {
	for _, m := range ALL_MOVES {
		if m.ID == id {
			return &m
		}
	}
	panic("move not found")
}

func (m Move) String() string {
	return fmt.Sprintf("%v\nType: %v, Power: %v, Accuracy: %v\n", m.Name, m.Type, m.Power, m.Accuracy)
}
