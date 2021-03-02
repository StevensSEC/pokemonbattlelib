package pokemonbattlelib

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Type effectiveness", func() {
	tests := []struct {
		move Type
		def  Type
		want Effectiveness
	}{
		{
			move: TypeWater,
			def:  TypeFire,
			want: SuperEffective,
		},
		{
			move: TypeGrass,
			def:  TypeWater,
			want: SuperEffective,
		},
		{
			move: TypeFire,
			def:  TypeWater,
			want: Ineffective,
		},
		{
			move: TypeGhost,
			def:  TypeNormal,
			want: NoEffect,
		},
		{
			move: TypeSteel,
			def:  TypeGround,
			want: NormalEffect,
		},
		{
			move: TypeFire,
			def:  TypeFire | TypeWater,
			want: VeryIneffective,
		},
		{
			move: TypeWater,
			def:  TypeGround | TypeRock,
			want: VerySuperEffective,
		},
	}
	for _, tt := range tests {
		Expect(GetElementalEffect(tt.move, tt.def)).To(Equal(tt.want))
	}
})

var _ = Describe("Elemental types", func() {
	tests := []struct {
		value Type
		want  string
	}{
		{
			value: TypeWater,
			want:  "[Water]",
		},
		{
			value: TypeNormal,
			want:  "[Normal]",
		},
		{
			value: TypeFire,
			want:  "[Fire]",
		},
		{
			value: TypeFire | TypeGrass,
			want:  "[Fire][Grass]",
		},
	}
	for _, tt := range tests {
		Expect(fmt.Sprintf("%s", tt.value)).To(Equal(tt.want))
	}
})

var _ = Describe("Gender", func() {
	It("should show correct string for gender", func() {
		Expect(GenderGenderless.String()).To(Equal(""))
		Expect(GenderFemale.String()).To(Equal("♀"))
		Expect(GenderMale.String()).To(Equal("♂"))
		Expect(func() {
			var _ = Gender(-1).String()
		}).To(Panic())
	})
})

var _ = Describe("Status conditions", func() {
	It("should check if status conditions are present", func() {
		tests := []struct {
			status StatusCondition
			check  StatusCondition
			want   bool
		}{
			{
				status: StatusNone,
				check:  StatusNone,
				want:   true,
			},
			{
				status: StatusBurn,
				check:  StatusBurn,
				want:   true,
			},
			{
				status: StatusSleep | StatusCursed,
				check:  StatusSleep,
				want:   true,
			},
			{
				status: StatusSleep | StatusCursed,
				check:  StatusCursed,
				want:   true,
			},
			{
				status: StatusParalyze,
				check:  StatusBurn,
				want:   false,
			},
			{
				status: StatusParalyze | StatusConfusion,
				check:  StatusBurn,
				want:   false,
			},
		}
		for _, tt := range tests {
			Expect(tt.status.check(tt.check)).To(Equal(tt.want))
		}
	})

	It("should apply status conditions", func() {
		tests := []struct {
			name  string
			from  StatusCondition
			apply StatusCondition
			want  StatusCondition
		}{
			{
				name:  "Should apply burn",
				from:  StatusNone,
				apply: StatusBurn,
				want:  StatusBurn,
			},
			{
				name:  "Should apply burn and remove poison",
				from:  StatusPoison,
				apply: StatusBurn,
				want:  StatusBurn,
			},
			{
				name:  "Should apply freeze and remove sleep",
				from:  StatusSleep,
				apply: StatusFreeze,
				want:  StatusFreeze,
			},
			{
				name:  "Should preserve bound, replace sleep with poison",
				from:  StatusSleep | StatusBound,
				apply: StatusPoison,
				want:  StatusPoison | StatusBound,
			},
			{
				name:  "Should preserve sleep, apply bound",
				from:  StatusSleep,
				apply: StatusBound,
				want:  StatusSleep | StatusBound,
			},
			{
				name:  "Should preserve sleep, leech seed, apply bound",
				from:  StatusSleep | StatusLeechSeed,
				apply: StatusBound,
				want:  StatusSleep | StatusLeechSeed | StatusBound,
			},
		}
		for _, tt := range tests {
			tt.from.apply(tt.apply)
			Expect(tt.from).To(Equal(tt.want))
		}
	})

	It("should clear status effects if they are present", func() {
		tests := []struct {
			name  string
			from  StatusCondition
			clear StatusCondition
			want  StatusCondition
		}{
			{
				name:  "Should clear burn",
				from:  StatusBurn,
				clear: StatusBurn,
				want:  StatusNone,
			},
			{
				name:  "Should preserve bound, clear sleep",
				from:  StatusSleep | StatusBound,
				clear: StatusSleep,
				want:  StatusBound,
			},
			{
				name:  "Should preserve bound, apply bound",
				from:  StatusSleep | StatusBound,
				clear: StatusBound,
				want:  StatusSleep,
			},
			{
				name:  "Should preserve sleep, leech seed, clear bound",
				from:  StatusSleep | StatusLeechSeed | StatusBound,
				clear: StatusBound,
				want:  StatusSleep | StatusLeechSeed,
			},
		}
		for _, tt := range tests {
			tt.from.clear(tt.clear)
			Expect(tt.from).To(Equal(tt.want))
		}
	})

	It("should render the correct string for status conditions", func() {
		tests := []struct {
			name string
			cond StatusCondition
			want string
		}{
			{
				name: "None",
				cond: StatusNone,
				want: "",
			},
			{
				name: "Burn",
				cond: StatusBurn,
				want: "burned",
			},
			{
				name: "Poison",
				cond: StatusPoison,
				want: "poisoned",
			},
			{
				name: "Sleep",
				cond: StatusSleep,
				want: "asleep",
			},
			{
				name: "Freeze, Bound",
				cond: StatusFreeze | StatusBound,
				want: "frozen, bound",
			},
			{
				name: "Paralyzed, Bound, Confusion, Torment",
				cond: StatusParalyze | StatusBound | StatusConfusion | StatusTorment,
				want: "paralyzed, bound, confused, tormented",
			},
		}
		for _, tt := range tests {
			Expect(fmt.Sprintf("%s", tt.cond)).To(Equal(tt.want))
		}
	})
})
