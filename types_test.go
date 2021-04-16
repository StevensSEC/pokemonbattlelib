package pokemonbattlelib

import (
	"encoding/json"
	"fmt"
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Elemental types", func() {
	DescribeTable("GetElementalEffect()",
		func(move, def Type, want Effectiveness) {
			Expect(GetElementalEffect(move, def, ItemNone)).To(Equal(want))
		},
		Entry("Water -> Fire", TypeWater, TypeFire, SuperEffective),
		Entry("Grass -> Water", TypeGrass, TypeWater, SuperEffective),
		Entry("Fire -> Water", TypeFire, TypeWater, Ineffective),
		Entry("Ghost -> Normal", TypeGhost, TypeNormal, NoEffect),
		Entry("Steel -> Ground", TypeSteel, TypeGround, NormalEffect),
		Entry("Fire -> Fire and Water", TypeFire, TypeFire|TypeWater, VeryIneffective),
		Entry("Water -> Ground and Rock", TypeWater, TypeGround|TypeRock, VerySuperEffective),
	)

	DescribeTable("Stringer",
		func(value Type, want string) {
			Expect(fmt.Sprintf("%s", value)).To(Equal(want))
		},
		Entry("Normal", TypeNormal, "[Normal]"),
		Entry("Water", TypeWater, "[Water]"),
		Entry("Fire", TypeFire, "[Fire]"),
		Entry("Fire and Grass", TypeFire|TypeGrass, "[Fire][Grass]"),
	)
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
			{
				name:  "Should not clear status that does not exist",
				from:  StatusNone,
				clear: StatusBurn,
				want:  StatusNone,
			},
			{
				name:  "Should not clear freeze",
				from:  StatusParalyze,
				clear: StatusFreeze,
				want:  StatusParalyze,
			},
			{
				name:  "Should clear only existing effects",
				from:  StatusParalyze | StatusBound | StatusCantEscape,
				clear: StatusConfusion | StatusBound | StatusCantEscape,
				want:  StatusParalyze,
			},
			{
				name:  "Should clear all non-volatile effects",
				from:  StatusBurn | StatusConfusion,
				clear: StatusNonvolatileMask,
				want:  StatusConfusion,
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

var _ = Describe("Natures", func() {
	Context("Auto-generated functions", func() {
		var entries []TableEntry
		for i := NatureAdamant; i <= NatureTimid; i++ {
			entries = append(entries, Entry(fmt.Sprintf("Nature #%d", i), i))
		}

		DescribeTable("GetStatModifiers() doesn't output StatHP",
			func(n Nature) {
				gotUp, gotDown := n.GetStatModifiers()
				// Testing the exact results doesn't really make sense, since the function is auto-generated.
				// So I'm just to make sure the generated code didn't didn't fuck up.
				Expect(gotUp).ToNot(Equal(StatHP))
				Expect(gotDown).ToNot(Equal(StatHP))
			},
			entries...,
		)

		DescribeTable("Stringer doesn't output \"Nature\"",
			func(n Nature) {
				got := fmt.Sprintf("%s", n)
				// Testing the exact results doesn't really make sense, since the function is auto-generated.
				// So I'm just to make sure the generated code doesn't produce unnecessary text.
				Expect(got).ToNot(ContainSubstring("Nature"))
			},
			entries...,
		)

		It("should panic if given an invalid Nature", func() {
			Expect(func() { _ = Nature(255).String() }).To(Panic())
			Expect(func() { _, _ = Nature(255).GetStatModifiers() }).To(Panic())
		})
	})
})

var _ = Describe("Abilities", func() {
	var entries []TableEntry
	for i := AbilityAdaptability; i <= AbilityWonderGuard; i++ {
		entries = append(entries, Entry(fmt.Sprintf("Ability #%d", i), i))
	}

	DescribeTable("Stringer doesn't output \"Ability\"",
		func(a Ability) {
			got := fmt.Sprintf("%s", a)
			// Testing the exact results doesn't really make sense, since the function is auto-generated.
			// So I'm just to make sure the generated code doesn't produce unnecessary text.
			Expect(got).ToNot(ContainSubstring("Ability"))
		},
		entries...,
	)

	It("should panic if given an invalid Ability", func() {
		Expect(func() { _ = Ability(math.MaxUint16).String() }).To(Panic())
	})
})

var _ = Describe("target", func() {
	It("should unmarshal json", func() {
		bytes := []byte("{\"Party\": 1, \"Slot\": 2}")
		var t target
		Expect(json.Unmarshal(bytes, &t)).To(Succeed())
		Expect(t.party).To(Equal(1))
		Expect(t.partySlot).To(Equal(2))
	})

	It("should unmarshal json from marshalled json", func() {
		bytes, err := json.Marshal(target{
			party:     1,
			partySlot: 2,
		})
		Expect(err).To(Succeed())
		Expect(string(bytes)).To(ContainSubstring("2"))
		var t target
		Expect(json.Unmarshal(bytes, &t)).To(Succeed())
		Expect(t.party).To(Equal(1))
		Expect(t.partySlot).To(Equal(2))
	})
})
