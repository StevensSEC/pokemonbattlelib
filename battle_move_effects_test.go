package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stat-changing moves", func() {
	agent1 := Agent(new(dumbAgent))

	DescribeTable("Move Stat-change data",
		func(id MoveId) {
			// move := GetMove(id)
		},
		Entry("Howl", MoveHowl),
		Entry("Harden", MoveHarden),
	)

	DescribeTable("Changing Pokemon stat modifiers",
		func(id MoveId, stat, stages int) {
			charmander := GeneratePokemon(PkmnCharmander, defaultMoveOpt)
			squirtle := GeneratePokemon(PkmnSquirtle, defaultMoveOpt)
			battle := New1v1Battle(charmander, &agent1, squirtle, &agent1)
			battle.rng = SimpleRNG()
			charmander.Moves[0] = GetMove(id)
			Expect(battle.Start()).To(Succeed())
			t, _ := battle.SimulateRound()
			Expect(t).To(HaveTransaction(ModifyStatTransaction{
				Target: charmander,
				Stat:   stat,
				Stages: stages,
			}))
			// Bound by min/max stat modifier
			charmander.StatModifiers[stat] = MaxStatModifier
			t, _ = battle.SimulateRound()
			Expect(t).To(HaveTransaction(ModifyStatTransaction{
				Target: charmander,
				Stat:   stat,
				Stages: stages,
			}))
			Expect(charmander.StatModifiers[stat]).To(BeEquivalentTo(MaxStatModifier))
		},
		Entry("Howl", MoveHowl, StatAtk, +1),
		Entry("Double Team", MoveDoubleTeam, StatEvasion, +1),
	)
})
