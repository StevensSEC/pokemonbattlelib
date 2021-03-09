package pokemonbattlelib

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
	MoveTargetSpecificMove MoveTarget = iota + 1
	MoveTargetSelectedMeFirst
	MoveTargetAlly
	MoveTargetUsersField
	MoveTargetUserOrAlly
	MoveTargetOpponentsField
	MoveTargetUser
	MoveTargetRandomOpponent
	MoveTargetAllOthers
	MoveTargetSelected
	MoveTargetAllOpponents
	MoveTargetEntireField
	MoveTargetUserAndAllies
	MoveTargetAll
)

// Represents a Pokemon's move. Moves can deal damage, heal the user or allies, or cause status effects.
type Move struct {
	ID        int
	Name      string
	Type      Type
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
	for _, m := range AllMoves {
		if m.ID == id {
			return &m
		}
	}
	panic("move not found")
}

func (m Move) String() string {
	return m.Name
}
