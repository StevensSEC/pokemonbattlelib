package pokemonbattlelib

// import (
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/ginkgo/extensions/table"
// 	. "github.com/onsi/gomega"
// )

// TODO: Move tests involving moves here

// Creates a new move and returns the ID of the move created
func NewMove(power uint, category MoveCategory, moveType Type) MoveId {
	m := MoveData{
		Power:        power,
		Category:     category,
		Type:         moveType,
		Accuracy:     100,
		InitialMaxPP: ^uint8(0),
	}
	AllMoves = append(AllMoves, m)
	return MoveId(len(AllMoves))
}

var TestMoveDefault = NewMove(10, MoveCategoryPhysical, TypeNormal)
var TestMoveNoDamage = NewMove(0, MoveCategoryPhysical, TypeNormal)
