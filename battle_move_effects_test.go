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
			charmander := PkmnWithMoves(id)
			squirtle := PkmnWithMoves(MovePound)
			battle := New1v1Battle(charmander, &agent1, squirtle, &agent1)
			battle.rng = SimpleRNG()
			Expect(battle.Start()).To(Succeed())
			t, _ := battle.SimulateRound()
			Expect(t).To(HaveTransaction(ModifyStatTransaction{
				Target:        charmander,
				SelfInflicted: true,
				Stat:          stat,
				Stages:        stages,
			}))
			// Bound by min/max stat modifier
			charmander.StatModifiers[stat] = MaxStatModifier
			t, _ = battle.SimulateRound()
			Expect(t).To(HaveTransaction(ModifyStatTransaction{
				Target:        charmander,
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
			charmander := PkmnWithMoves(id)
			squirtle := PkmnWithMoves(MovePound)
			battle := New1v1Battle(charmander, &agent1, squirtle, &agent1)
			battle.rng = AlwaysRNG()
			Expect(battle.Start()).To(Succeed())
			t, _ := battle.SimulateRound()
			Expect(t).To(HaveTransaction(ModifyStatTransaction{
				Target: squirtle,
				Stat:   stat,
				Stages: stages,
			}))
		},
		Entry("Scary Face", MoveScaryFace, StatSpeed, -2),
		Entry("Leer", MoveLeer, StatDef, -1),
	)
})
