package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

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

var _ = Describe("Move PP Consumption", func() {
	a := Agent(new(dumbAgent))
	It("should decrement move's PP by 1 when used", func() {
		b := New1v1Battle(
			GeneratePokemon(PkmnSquirtle, WithMoves(TestMoveNoDamage)), &a,
			GeneratePokemon(PkmnSquirtle, WithMoves(TestMoveNoDamage)), &a,
		)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransactionsInOrder(
			PPTransaction{
				Amount: -1,
			},
			PPTransaction{
				Amount: -1,
			},
		))
	})

	It("should decrease the opponent's last used move's PP by 4 when a pokemon uses Spite", func() {
		p1 := GeneratePokemon(PkmnCharmander, WithMoves(MoveSpite))
		p2 := GeneratePokemon(PkmnSquirtle, defaultMoveOpt)
		b := New1v1Battle(p1, &a, p2, &a)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		b.SimulateRound() // set Pokemon's last move
		p1.CurrentHP = p1.MaxHP()
		p2.CurrentHP = p2.MaxHP()
		p2.Moves[0].CurrentPP = 1
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransaction(PPTransaction{
			Move:   p2.Moves[0],
			Amount: -4,
		}))
		// Ensure that PP stays in bounds
		Expect(p2.Moves[0].CurrentPP).To(BeEquivalentTo(0))
	})
})
