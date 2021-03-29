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
type MoveTarget uint8

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

type MoveId uint16

type MoveData struct {
	Name         string
	Type         Type
	Category     MoveCategory
	Targets      MoveTarget
	Priority     int8
	Power        uint
	Accuracy     uint
	InitialMaxPP uint8
}

// Represents a Pokemon's move. Moves can deal damage, heal the user or allies, or cause status effects.
type Move struct {
	Id        MoveId
	CurrentPP uint8
	MaxPP     uint8
}

//go:generate go run ./scripts/gen_getters.go -for Move -data MoveData

// Retrieves a Pokemon move given its move ID
func GetMove(id MoveId) *Move {
	m := Move{
		Id: id,
	}
	m.CurrentPP = m.InitialMaxPP()
	m.MaxPP = m.InitialMaxPP()
	return &m
}

// Grabs move's constant data
func (m *Move) Data() *MoveData {
	if m.Id > MoveId(len(AllMoves)) {
		blog.Panicf("Move (id: %d) is an invalid move", m.Id)
	}
	if m.Id == MoveNone {
		return &MoveData{}
	}
	return &AllMoves[m.Id-1]
}

func (m Move) String() string {
	return m.Name()
}
