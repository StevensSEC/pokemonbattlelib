package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stat-changing moves", func() {
	agent1 := Agent(new(dumbAgent))

	DescribeTable("Move Stat-change data",
		func(id MoveId, stat int) {
			move := GetMove(id)
			Expect(move.AffectedStat()).To(Equal(uint8(stat)))
		},
		Entry("Howl", MoveHowl, StatAtk),
		Entry("Double Team", MoveDoubleTeam, StatEvasion),
		Entry("Harden", MoveHarden, StatDef),
	)

	DescribeTable("Should modify the User's stats",
		func(id MoveId, stat, stages int) {
			charmander := GeneratePokemon(PkmnCharmander, WithMoves(id))
			squirtle := GeneratePokemon(PkmnSquirtle, WithMoves(MovePound))
			battle := New1v1Battle(charmander, &agent1, squirtle, &agent1)
			battle.rng = SimpleRNG()
			Expect(battle.Start()).To(Succeed())
			t, _ := battle.SimulateRound()
			Expect(t).To(HaveTransaction(ModifyStatTransaction{
				Target:        target{0, 0},
				SelfInflicted: true,
				Stat:          stat,
				Stages:        stages,
			}))
			// Bound by min/max stat modifier
			charmander.StatModifiers[stat] = MaxStatModifier
			t, _ = battle.SimulateRound()
			Expect(t).To(HaveTransaction(ModifyStatTransaction{
				Target:        target{0, 0},
				SelfInflicted: true,
				Stat:          stat,
				Stages:        stages,
			}))
			Expect(charmander.StatModifiers[stat]).To(BeEquivalentTo(MaxStatModifier))
		},
		Entry("Howl", MoveHowl, StatAtk, +1),
		Entry("Double Team", MoveDoubleTeam, StatEvasion, +1),
		Entry("Harden", MoveHarden, StatDef, 1),
	)

	DescribeTable("Should modify the Opponents's stats",
		func(id MoveId, stat, stages int) {
			charmander := GeneratePokemon(PkmnCharmander, WithMoves(id))
			squirtle := GeneratePokemon(PkmnSquirtle, WithMoves(MovePound))
			battle := New1v1Battle(charmander, &agent1, squirtle, &agent1)
			battle.rng = AlwaysRNG()
			Expect(battle.Start()).To(Succeed())
			t, _ := battle.SimulateRound()
			Expect(t).To(HaveTransaction(ModifyStatTransaction{
				Target: target{1, 0},
				Stat:   stat,
				Stages: stages,
			}))
		},
		Entry("Scary Face", MoveScaryFace, StatSpeed, -2),
		Entry("Leer", MoveLeer, StatDef, -1),
	)
})

var _ = Describe("Move Effects", func() {
	a1 := Agent(new(dumbAgent))

	It("should cause flinching", func() {
		b := New1v1Battle(
			GeneratePokemon(PkmnMightyena, WithLevel(20), WithIVs([6]uint8{0, 0, 0, 0, 0, 31}), WithEVs([6]uint8{0, 0, 0, 0, 0, 252}), WithMoves(MoveBite)), &a1,
			GeneratePokemon(PkmnPonyta, WithLevel(20), WithIVs([6]uint8{31, 0, 31, 0, 31, 0}), WithEVs([6]uint8{252, 0, 0, 0, 0, 0}), WithMoves(MoveTackle)), &a1,
		)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransaction(DamageTransaction{
			Target: target{1, 0},
		}))
		Expect(t).To(HaveTransactionsInOrder(
			InflictStatusTransaction{
				Target:       target{1, 0},
				StatusEffect: StatusFlinch,
			},
			ImmobilizeTransaction{
				Target:       target{1, 0},
				StatusEffect: StatusFlinch,
			},
		))
	})

	It("should raise speed on flinch for pokemon with steadfast ability", func() {
		b := New1v1Battle(
			GeneratePokemon(PkmnMightyena, WithLevel(20), WithIVs([6]uint8{0, 0, 0, 0, 0, 31}), WithEVs([6]uint8{0, 0, 0, 0, 0, 252}), WithMoves(MoveBite)), &a1,
			GeneratePokemon(PkmnPonyta, WithLevel(20), WithIVs([6]uint8{31, 0, 31, 0, 31, 0}), WithEVs([6]uint8{252, 0, 0, 0, 0, 0}), WithMoves(MoveTackle), WithAbility(AbilitySteadfast)), &a1,
		)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransaction(DamageTransaction{
			Target: target{1, 0},
		}))
		Expect(t).To(HaveTransactionsInOrder(
			InflictStatusTransaction{
				Target:       target{1, 0},
				StatusEffect: StatusFlinch,
			},
			ModifyStatTransaction{
				Target: target{1, 0},
				Stat:   StatSpeed,
				Stages: 1,
			},
			ImmobilizeTransaction{
				Target:       target{1, 0},
				StatusEffect: StatusFlinch,
			},
		))
	})

	PIt("should deal damage and modify the user's stat (MoveMetaCategoryDamageRaise)", func() {
		b := New1v1Battle(
			GeneratePokemon(PkmnMightyena, WithLevel(20), WithMoves(TestMoveDamageAndStatChangeSelf)), &a1,
			GeneratePokemon(PkmnPonyta, WithLevel(20), WithMoves(TestMoveNoDamage)), &a1,
		)
		b.rng = AlwaysRNG()
		Expect(b.Start()).To(Succeed())
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransactionsInOrder(
			DamageTransaction{
				Target: target{1, 0},
			},
			ModifyStatTransaction{
				Target:        target{0, 0},
				SelfInflicted: true,
				Stat:          int(b.GetPokemon(target{0, 0}).Moves[0].AffectedStat()),
				Stages:        int(b.GetPokemon(target{0, 0}).Moves[0].StatStages()),
			},
		))
	})
})

var _ = Describe("Draining moves", func() {
	a1 := Agent(new(dumbAgent))
	var b *Battle

	BeforeEach(func() {
		b = NewSingleBattle(
			NewOccupiedParty(
				GeneratePokemon(PkmnRoselia,
					WithLevel(25),
					WithMoves(MoveGigaDrain),
				),
			),
			&a1,
			NewOccupiedParty(
				GeneratePokemon(PkmnBidoof,
					WithLevel(25),
					WithMoves(MoveSplash),
				),
			),
			&a1,
		)
		b.rng = SimpleRNG()
		Expect(b.Start()).To(Succeed())
	})

	It("should damage the target and heal the user", func() {
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransactionsInOrder(
			UseMoveTransaction{
				User:   target{0, 0},
				Target: target{1, 0},
			},
			DamageTransaction{
				Target: target{1, 0},
				Damage: 61,
			},
			HealTransaction{
				Target: target{0, 0},
				Amount: 30,
			},
		))
	})

	It("should heal more when the user is holding a big root", func() {
		b.GetPokemon(target{0, 0}).HeldItem = ItemBigRoot
		t, _ := b.SimulateRound()

		Expect(t).To(HaveTransactionsInOrder(
			UseMoveTransaction{
				User:   target{0, 0},
				Target: target{1, 0},
			},
			DamageTransaction{
				Target: target{1, 0},
				Damage: 61,
			},
			HealTransaction{
				Target: target{0, 0},
				Amount: 39,
			},
		))
	})
})

var _ = Describe("Recoil moves", func() {
	a1 := Agent(new(dumbAgent))
	var b *Battle

	BeforeEach(func() {
		b = NewSingleBattle(
			NewOccupiedParty(
				GeneratePokemon(PkmnPikachu,
					WithLevel(25),
					WithMoves(MoveVoltTackle),
				),
			),
			&a1,
			NewOccupiedParty(
				GeneratePokemon(PkmnBidoof,
					WithLevel(25),
					WithMoves(MoveSplash),
				),
			),
			&a1,
		)
		b.rng = SimpleRNG()
		Expect(b.Start()).To(Succeed())
	})

	It("should damage the target and damage the user", func() {
		t, _ := b.SimulateRound()
		Expect(t).To(HaveTransactionsInOrder(
			DamageTransaction{
				Target: target{1, 0},
				Damage: 57,
			},
			DamageTransaction{
				Target: target{0, 0},
				Damage: 18,
			},
		))
	})
})

var _ = Describe("Fixed damage moves", func() {
	a1 := Agent(new(dumbAgent))

	It("should reduce the target's HP to half that amount when Super Fang is used", func() {
		hoppip := GeneratePokemon(PkmnHoppip,
			WithLevel(25),
			WithMoves(MoveSplash),
		)
		const HOPPIP_HP = 41
		hoppip.Stats[StatHP] = HOPPIP_HP
		hoppip.CurrentHP = HOPPIP_HP

		b := NewSingleBattle(
			NewOccupiedParty(
				GeneratePokemon(PkmnRaticate,
					WithLevel(25),
					WithMoves(MoveSuperFang),
				),
			),
			&a1,
			NewOccupiedParty(
				hoppip,
			),
			&a1,
		)
		b.rng = SimpleRNG()
		Expect(b.Start()).To(Succeed())

		t, _ := b.SimulateRound()
		Expect(DamageDealt(t, target{0, 0})).To(Equal(HOPPIP_HP / 2))
		t, _ = b.SimulateRound()
		Expect(DamageDealt(t, target{0, 0})).To(Equal((HOPPIP_HP / 2) / 2))
	})
})
