package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TODO: Move tests involving moves here

func registerMoveWithType(t Type) MoveId {
	return RegisterMove(MoveData{
		Power:        10,
		Type:         t,
		Category:     MoveCategoryPhysical,
		InitialMaxPP: 100,
	})
}

var TestMoveDefault = RegisterMove(MoveData{Name: "Default", Category: MoveCategoryPhysical, Power: 10, Accuracy: 100, InitialMaxPP: 100})
var TestMoveNoDamage = RegisterMove(MoveData{Name: "No Damage", InitialMaxPP: 100})
var TestMoveInflictBurn = RegisterMove(MoveData{Name: "Inflict Burn", Category: MoveCategoryStatus, InitialMaxPP: 100, Ailment: StatusBurn})
var TestMoveDamageAndInflictBurn = RegisterMove(MoveData{Name: "Damage and Inflict Burn", Category: MoveCategoryPhysical, InitialMaxPP: 100, AilmentChance: 100, Ailment: StatusBurn, Power: 10})

var _ = Describe("Move Status Inflict", func() {
	a := Agent(new(dumbAgent))
	It("should inflict burn", func() {
		b := New1v1Battle(
			PkmnWithMoves(TestMoveInflictBurn), &a,
			PkmnWithMoves(TestMoveNoDamage), &a,
		)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransaction(
			InflictStatusTransaction{
				Target:       target{1, 0},
				StatusEffect: StatusBurn,
			},
		))
	})

	It("should deal damage and inflict burn", func() {
		// see: https://github.com/StevensSEC/pokemonbattlelib/pull/366#discussion_r624527986
		b := New1v1Battle(
			PkmnWithMoves(TestMoveDamageAndInflictBurn), &a,
			PkmnWithMoves(TestMoveNoDamage), &a,
		)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransactionsInOrder(
			DamageTransaction{
				Target: target{1, 0},
				Move:   b.getPokemon(target{0, 0}).Moves[0],
				Damage: 4,
			},
			InflictStatusTransaction{
				Target:       target{1, 0},
				StatusEffect: StatusBurn,
			},
		))
	})

	It("should inflict paralysis from MoveStunSpore", func() {
		pkmn1 := PkmnWithMoves(TestMoveNoDamage)
		pkmn2 := PkmnWithMoves(MoveStunSpore)
		b := New1v1Battle(pkmn1, &a, pkmn2, &a)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransaction(
			InflictStatusTransaction{
				Target:       target{0, 0},
				StatusEffect: StatusParalyze,
			},
		))
	})

	It("should inflict poison from MovePoisonSting", func() {
		pkmn1 := PkmnWithMoves(TestMoveNoDamage)
		pkmn2 := PkmnWithMoves(MovePoisonSting)
		b := New1v1Battle(pkmn1, &a, pkmn2, &a)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransaction(
			InflictStatusTransaction{
				Target:       target{0, 0},
				StatusEffect: StatusPoison,
			},
		))
	})
})

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
		p1 := PkmnDefault()
		p2 := PkmnDefault()
		p1.Moves[0] = GetMove(MoveSpite)
		b := New1v1Battle(p1, &a, p2, &a)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		p2.metadata[MetaLastMove] = p2.Moves[0]
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

var _ = Describe("Move Damage", func() {
	a := Agent(new(dumbAgent))

	It("should always do 20 damage when using sonic boom", func() {
		b := New1v1Battle(
			PkmnWithMoves(MoveSonicBoom), &a,
			PkmnWithMoves(TestMoveNoDamage), &a,
		)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransactionsInOrder(
			UseMoveTransaction{
				User:   target{0, 0},
				Target: target{1, 0},
				Move:   b.getPokemon(target{0, 0}).Moves[0],
			},
			DamageTransaction{
				Target: target{1, 0},
				Move:   b.getPokemon(target{0, 0}).Moves[0],
				Damage: 20,
			},
		))
	})
})
