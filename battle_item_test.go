package pokemonbattlelib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Misc. + Held Items", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))
	var holder *Pokemon

	setup := func(item Item, pkmn int) (*Battle, *Pokemon) {
		p1 := GeneratePokemon(PkmnSnorlax, WithLevel(25), WithMoves(MoveSplash))
		holder = GeneratePokemon(pkmn, WithLevel(25), WithMoves(MoveSplash))
		holder.HeldItem = item
		b := New1v1Battle(p1, &a1, holder, &a2)
		b.rng = SimpleRNG()
		Expect(b.Start()).To(Succeed())
		return b, holder
	}

	Context("Bad Held Items", func() {
		DescribeTable("Inflicts status after turn",
			func(item Item, status StatusCondition) {
				b, _ := setup(item, PkmnBulbasaur)
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(InflictStatusTransaction{
					Target:       target{1, 0},
					StatusEffect: status,
				}))
			},
			Entry("Flame Orb", ItemFlameOrb, StatusBurn),
			Entry("Toxic Orb", ItemToxicOrb, StatusBadlyPoison),
		)

		DescribeTable("Move last in priority bracket",
			func(item Item) {
				b, _ := setup(item, PkmnMagikarp)
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransactionsInOrder(
					UseMoveTransaction{
						User:   target{0, 0},
						Target: target{1, 0},
					},
					MoveFailTransaction{
						User:   target{0, 0},
						Reason: FailOther,
					},
					UseMoveTransaction{
						User:   target{1, 0},
						Target: target{0, 0},
					},
					MoveFailTransaction{
						User:   target{1, 0},
						Reason: FailOther,
					},
				))
			},
			Entry("Full Incense", ItemFullIncense),
			Entry("Lagging Tail", ItemLaggingTail),
		)

		Specify("Iron Ball", func() {
			b, holder := setup(ItemIronBall, PkmnPidgeot)
			attacker := b.getPokemon(target{0, 0})
			attacker.Moves[0] = GetMove(MoveEarthquake)
			// Flying immunity negated
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(UseMoveTransaction{
				User:   target{0, 0},
				Target: target{1, 0},
			}))
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{1, 0},
				Move:   GetMove(MoveEarthquake),
				Damage: 36,
			}))
			speed := holder.Speed()
			holder.HeldItem = ItemNone
			Expect(holder.Speed()).To(BeNumerically(">", speed))
		})

		Specify("Sticky Barb", func() {
			b, holder := setup(ItemStickyBarb, PkmnGrimer)
			attacker := b.getPokemon(target{0, 0})
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(UseMoveTransaction{
				User:   target{0, 0},
				Target: target{1, 0},
			}))
			// Holder takes 1/8 HP damage after every turn
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{1, 0},
				Damage: holder.MaxHP() / 8,
			}))
			// Contact moves damage attacker and pass the sticky barb to the attacker
			attacker.Moves[0] = GetMove(MoveTackle)
			t, _ = b.SimulateRound()
			Expect(t).To(HaveTransactionsInOrder(
				DamageTransaction{
					Target: target{0, 0},
					Damage: attacker.MaxHP() / 8,
				},
				GiveItemTransaction{
					Target: target{0, 0},
					Item:   ItemStickyBarb,
				},
				GiveItemTransaction{
					Target: target{1, 0},
					Item:   ItemNone,
				},
			))
			Expect(attacker.HeldItem).To(BeEquivalentTo(ItemStickyBarb))
		})
	})

	DescribeTable("Choice Items",
		func(item Item, stat func(*Pokemon) uint) {
			b, holder := setup(item, PkmnMachamp)
			b.SimulateRound()
			// Restricts user to one move
			Expect(holder.metadata[MetaLastMove]).To(BeEquivalentTo(holder.Moves[0]))
			// Stats are boosted by 50%
			effective := stat(holder)
			holder.HeldItem = ItemNone
			Expect(stat(holder)).To(BeNumerically("<", effective))
		},
		Entry("Choice Band", ItemChoiceBand, func(p *Pokemon) uint { return p.Attack() }),
		Entry("Choice Scarf", ItemChoiceScarf, func(p *Pokemon) uint { return p.Speed() }),
		Entry("Choice Specs", ItemChoiceSpecs, func(p *Pokemon) uint { return p.SpecialAttack() }),
	)

	Context("Miscellaneous Items", func() {
		Specify("Black Sludge", func() {
			// Heal poison types for 1/16 HP
			b, holder := setup(ItemBlackSludge, PkmnGrimer)
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(HealTransaction{
				Target: target{1, 0},
				Amount: holder.MaxHP() / 16,
			}))
			// Damage non-poison types for 1/8 HP
			b, holder = setup(ItemBlackSludge, PkmnAerodactyl)
			t, _ = b.SimulateRound()
			Expect(t).ToNot(HaveTransaction(HealTransaction{
				Target: target{1, 0},
				Amount: holder.MaxHP() / 16,
			}))
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{1, 0},
				Damage: holder.MaxHP() / 8,
			}))
		})

		Specify("Destiny Knot", func() {
			b, holder := setup(ItemDestinyKnot, PkmnMimeJr)
			attacker := b.getPokemon(target{0, 0})
			attacker.Moves[0] = GetMove(MoveAttract)
			attacker.Gender = GenderMale
			holder.Gender = GenderFemale
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(InflictStatusTransaction{
				Target:       target{0, 0},
				StatusEffect: StatusInfatuation,
			}))
		})

		Specify("Expert Belt", func() {
			b, holder := setup(ItemExpertBelt, PkmnMachamp)
			holder.Moves[0] = GetMove(MoveCloseCombat)
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(UseMoveTransaction{
				User:   target{1, 0},
				Target: target{0, 0},
			}))
			// Damage boosted by 20%
			Expect(t).To(HaveTransaction(
				DamageTransaction{
					Target: target{0, 0},
					Damage: 201,
				},
			))
		})

		Specify("Leftovers", func() {
			b, holder := setup(ItemLeftovers, PkmnSnorlax)
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(HealTransaction{
				Target: target{1, 0},
				Amount: holder.MaxHP() / 16,
			}))
		})

		Specify("Life Orb", func() {
			b, holder := setup(ItemLifeOrb, PkmnSnorlax)
			holder.Moves[0] = GetMove(MoveTackle)
			t, _ := b.SimulateRound()
			// Boost damage by 30%
			Expect(DamageDealt(t, holder)).To(Equal(32))
			// Take 10% of max HP
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{1, 0},
				Damage: holder.MaxHP() / 10,
			}))
		})

		Specify("Muscle Band", func() {
			b, holder := setup(ItemMuscleBand, PkmnSnorlax)
			holder.Moves[0] = GetMove(MoveTackle)
			t, _ := b.SimulateRound()
			// Boost physical move damage by 10%
			Expect(DamageDealt(t, holder)).To(Equal(27))
		})

		Specify("Shell Bell", func() {
			b, holder := setup(ItemShellBell, PkmnSnorlax)
			holder.Moves[0] = GetMove(MoveTackle)
			t, _ := b.SimulateRound()
			// Self-inflict 1/8 of dealt damage
			Expect(t).To(HaveTransaction(DamageTransaction{
				Target: target{1, 0},
				Damage: 3,
			}))
		})

		Specify("White Herb", func() {
			b, holder := setup(ItemWhiteHerb, PkmnSnorlax)
			holder.StatModifiers = [9]int{-1, -1, -1, -1, -1, -1, 0, -1, -1}
			b.SimulateRound()
			// Consumed after use
			Expect(holder.HeldItem).To(Equal(ItemNone))
			// Reset all lowered stat modifiers
			for _, stage := range holder.StatModifiers {
				if stage < 0 {
					Fail("Expected all lowered stats to be reset")
				}
			}
		})

		Specify("Wise Glasses", func() {
			b, holder := setup(ItemWiseGlasses, PkmnSnorlax)
			holder.Moves[0] = GetMove(MoveSurf)
			t, _ := b.SimulateRound()
			// Boost special move damage by 10%
			Expect(DamageDealt(t, holder)).To(Equal(16))
		})

		DescribeTable("Accuracy/evasion items",
			func(attacking Item, defending Item, op string) {
				b, holder := setup(ItemNone, PkmnSnorlax)
				opponent := b.getPokemon(target{0, 0})
				base := CalcAccuracy(b.Weather, holder, opponent, GetMove(MovePound))
				holder.HeldItem = attacking
				opponent.HeldItem = defending
				new := CalcAccuracy(b.Weather, holder, opponent, GetMove(MovePound))
				Expect(new).To(BeNumerically(op, base))
			},
			Entry("Wide Lens", ItemWideLens, ItemNone, ">"),
			Entry("Bright Powder", ItemNone, ItemBrightPowder, "<"),
			Entry("Lax Incense", ItemNone, ItemLaxIncense, "<"),
		)

		DescribeTable("Status curing held items",
			func(item Item, status StatusCondition) {
				b, holder := setup(item, PkmnSnorlax)
				holder.StatusEffects.apply(status)
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(CureStatusTransaction{
					Target:       target{1, 0},
					StatusEffect: status,
				}))
				// Item should be consumed after use
				Expect(holder.HeldItem).To(Equal(ItemNone))
			},
			Entry("Mental Herb", ItemMentalHerb, StatusInfatuation),
		)

		DescribeTable("Flinch inducing items",
			func(item Item) {
				b, holder := setup(ItemKingsRock, PkmnLucario)
				holder.Moves[0] = GetMove(MoveTackle)
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(InflictStatusTransaction{
					Target:       target{0, 0},
					StatusEffect: StatusFlinch,
				}))
			},
			Entry("King's Rock", ItemKingsRock),
			Entry("Razor Fang", ItemRazorFang),
		)

		DescribeTable("Weather duration boosting rocks",
			func(item Item, weather Weather, move MoveId) {
				b, holder := setup(item, PkmnCastform)
				holder.Moves[0] = GetMove(move)
				t, _ := b.SimulateRound()
				Expect(t).To(HaveTransaction(WeatherTransaction{
					Weather: weather,
					Turns:   8,
				}))
			},
			Entry("Damp Rock", ItemDampRock, WeatherRain, MoveRainDance),
			Entry("Heat Rock", ItemHeatRock, WeatherHarshSunlight, MoveSunnyDay),
			Entry("Icy Rock", ItemIcyRock, WeatherHail, MoveHail),
			Entry("Smooth Rock", ItemSmoothRock, WeatherSandstorm, MoveSandstorm),
		)
	})

	DescribeTable("Plates",
		func(item Item, expectedType Type) {
			b, holder := setup(ItemNone, PkmnArceus)
			receiver := b.getPokemon(target{0, 0})
			damage := CalcMoveDamage(b.Weather, holder, receiver, GetMove(MoveJudgment))
			holder.HeldItem = item
			heldDamage := CalcMoveDamage(b.Weather, holder, receiver, GetMove(MoveJudgment))
			Expect(holder.EffectiveType()).To(Equal(expectedType))
			if expectedType == TypeGhost { // Normal immune to ghost
				Expect(heldDamage).To(BeEquivalentTo(0))
			} else {
				Expect(heldDamage).To(BeNumerically(">", damage))
			}
		},
		Entry("Draco Plate", ItemDracoPlate, TypeDragon),
		Entry("Dread Plate", ItemDreadPlate, TypeDark),
		Entry("Earth Plate", ItemEarthPlate, TypeGround),
		Entry("Fist Plate", ItemFistPlate, TypeFighting),
		Entry("Flame Plate", ItemFlamePlate, TypeFire),
		Entry("Icicle Plate", ItemIciclePlate, TypeIce),
		Entry("Insect Plate", ItemInsectPlate, TypeBug),
		Entry("Iron Plate", ItemIronPlate, TypeSteel),
		Entry("Meadow Plate", ItemMeadowPlate, TypeGrass),
		Entry("Mind Plate", ItemMindPlate, TypePsychic),
		Entry("Sky Plate", ItemSkyPlate, TypeFlying),
		Entry("Splash Plate", ItemSplashPlate, TypeWater),
		Entry("Spooky Plate", ItemSpookyPlate, TypeGhost),
		Entry("Stone Plate", ItemStonePlate, TypeRock),
		Entry("Toxic Plate", ItemToxicPlate, TypePoison),
		Entry("Zap Plate", ItemZapPlate, TypeElectric),
	)

	DescribeTable("Type Enhancement",
		func(item Item, expectedType Type) {
			b, holder := setup(ItemNone, PkmnArceus)
			receiver := b.getPokemon(target{0, 0})
			m := GetMove(registerMoveWithType(expectedType))
			damage := CalcMoveDamage(b.Weather, holder, receiver, m)
			holder.HeldItem = item
			heldDamage := CalcMoveDamage(b.Weather, holder, receiver, m)
			if expectedType == TypeGhost { // Normal immune to ghost
				Expect(heldDamage).To(BeEquivalentTo(0))
			} else {
				Expect(heldDamage).To(BeNumerically(">", damage))
			}
		},
		Entry("BlackBelt", ItemBlackBelt, TypeFighting),
		Entry("BlackGlasses", ItemBlackGlasses, TypeDark),
		Entry("Charcoal", ItemCharcoal, TypeFire),
		Entry("DragonFang", ItemDragonFang, TypeDragon),
		Entry("HardStone", ItemHardStone, TypeRock),
		Entry("Magnet", ItemMagnet, TypeElectric),
		Entry("MetalCoat", ItemMetalCoat, TypeSteel),
		Entry("MiracleSeed", ItemMiracleSeed, TypeGrass),
		Entry("MysticWater", ItemMysticWater, TypeWater),
		Entry("NeverMeltIce", ItemNeverMeltIce, TypeIce),
		Entry("OddIncense", ItemOddIncense, TypePsychic),
		Entry("PoisonBarb", ItemPoisonBarb, TypePoison),
		Entry("RockIncense", ItemRockIncense, TypeRock),
		Entry("RoseIncense", ItemRoseIncense, TypeGrass),
		Entry("SeaIncense", ItemSeaIncense, TypeWater),
		Entry("SharpBeak", ItemSharpBeak, TypeFlying),
		Entry("SilkScarf", ItemSilkScarf, TypeNormal),
		Entry("SilverPowder", ItemSilverPowder, TypeBug),
		Entry("SoftSand", ItemSoftSand, TypeGround),
		Entry("SpellTag", ItemSpellTag, TypeGhost),
		Entry("TwistedSpoon", ItemTwistedSpoon, TypePsychic),
		Entry("WaveIncense", ItemWaveIncense, TypeWater),
	)
})

var _ = Describe("Medicine Items", func() {
	var a1 Agent
	var a2 rcAgent
	var fullCure = StatusNonvolatileMask | StatusConfusion
	setup := func(item Item) (*Battle, *Pokemon) {
		a1 = Agent(new(dumbAgent))
		a2 = newRcAgent()
		_a2 := Agent(a2)
		user := GeneratePokemon(PkmnBulbasaur, WithLevel(100), WithMoves(MoveSplash))
		p2 := GeneratePokemon(PkmnCharmander, WithMoves(MoveSplash))
		b := New1v1Battle(user, &a1, p2, &_a2)
		Expect(b.Start()).To(Succeed())
		a2 <- ItemTurn{
			Item: item,
			Target: AgentTarget{
				target: target{0, 0},
			},
		}
		return b, user
	}

	DescribeTable("Healing (HP)",
		func(item Item, expectedHP int) {
			b, user := setup(item)
			fainted := PkmnDefault()
			if item.Category() == ItemCategoryRevival {
				user.CurrentHP = 0
				expectedHP = int(user.MaxHP())
				if item == ItemRevive {
					expectedHP /= 2
				}
			}
			if item == ItemSacredAsh {
				fainted.CurrentHP = 0
				b.parties[0].AddPokemon(fainted)
			}
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(HealTransaction{
				Target: target{0, 0},
				Amount: uint(expectedHP),
			}))
			if item == ItemSacredAsh {
				Expect(t).To(HaveTransaction(HealTransaction{
					Target: target{0, 1},
					Amount: fainted.MaxHP(),
				}))
			}
		},
		// Healing
		Entry("Berry Juice", ItemBerryJuice, 20),
		Entry("Energy Powder", ItemEnergyPowder, 50),
		Entry("Energy Root", ItemEnergyRoot, 200),
		Entry("Fresh Water", ItemFreshWater, 50),
		Entry("Full Restore", ItemFullRestore, 0),
		Entry("Hyper Potion", ItemHyperPotion, 200),
		Entry("Lemonade", ItemLemonade, 80),
		Entry("Max Potion", ItemMaxPotion, 0),
		Entry("Moomoo Milk", ItemMoomooMilk, 100),
		Entry("Potion", ItemPotion, 20),
		Entry("Soda Pop", ItemSodaPop, 60),
		Entry("Super Potion", ItemSuperPotion, 50),
		// Revival
		Entry("Max Revive", ItemMaxRevive, 0),
		Entry("Revival Herb", ItemRevivalHerb, 0),
		Entry("Revive", ItemRevive, 0),
		Entry("Sacred Ash", ItemSacredAsh, 0),
	)

	DescribeTable("Friendship",
		func(item Item, amount int) {
			b, user := setup(item)
			user.Friendship = MaxFriendship
			if item == ItemRevivalHerb {
				user.CurrentHP = 0
			}
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(FriendshipTransaction{
				Target: target{0, 0},
				Amount: amount,
			}))
		},
		Entry("Energy Powder", ItemEnergyPowder, -10),
		Entry("Energy Root", ItemEnergyRoot, -15),
		Entry("Revival Herb", ItemRevivalHerb, -20),
	)

	DescribeTable("PP Recovery",
		func(item Item, pp0, pp1 int) {
			b, user := setup(item)
			user.Moves[1] = GetMove(MoveTackle)
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(PPTransaction{
				Move:   user.Moves[0],
				Amount: int8(pp0),
			}))
			if pp1 != 0 {
				Expect(t).To(HaveTransaction(PPTransaction{
					Move:   user.Moves[1],
					Amount: int8(pp1),
				}))
			}
		},
		Entry("Elixir", ItemElixir, 10, 10),
		Entry("Ether", ItemEther, 10, 0),
		Entry("Max Elixir", ItemMaxElixir, 40, 35),
		Entry("Max Ether", ItemMaxEther, 40, 0),
	)

	DescribeTable("Status Cures",
		func(item Item, status StatusCondition) {
			b, user := setup(item)
			user.StatusEffects = status
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(CureStatusTransaction{
				Target:       target{0, 0},
				StatusEffect: status,
			}))
		},
		Entry("Full Restore", ItemFullRestore, fullCure),
		Entry("Antidote", ItemAntidote, StatusPoison),
		Entry("Awakening", ItemAwakening, StatusSleep),
		Entry("Burn Heal", ItemBurnHeal, StatusBurn),
		Entry("Full Heal", ItemFullHeal, fullCure),
		Entry("Heal Powder", ItemHealPowder, fullCure),
		Entry("Ice Heal", ItemIceHeal, StatusFreeze),
		Entry("Lava Cookie", ItemLavaCookie, fullCure),
		Entry("Old Gateau", ItemOldGateau, fullCure),
		Entry("Paralyze Heal", ItemParalyzeHeal, StatusParalyze),
	)
})

var _ = Describe("Battle Items", func() {
	var a1 rcAgent
	var a2 Agent
	setup := func(item Item) (*Battle, *Pokemon) {
		a1 = newRcAgent()
		_a1 := Agent(a1)
		a2 = Agent(new(dumbAgent))
		user := GeneratePokemon(PkmnBulbasaur, WithLevel(100), WithMoves(MoveSplash))
		p2 := GeneratePokemon(PkmnCharmander, WithMoves(MoveSplash))
		b := New1v1Battle(user, &_a1, p2, &a2)
		Expect(b.Start()).To(Succeed())
		a1 <- ItemTurn{
			Item: item,
			Target: AgentTarget{
				target: target{0, 0},
			},
		}
		return b, user
	}

	DescribeTable("Flutes",
		func(item Item, status StatusCondition) {
			b, user := setup(item)
			user.StatusEffects = status
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(CureStatusTransaction{
				Target:       target{0, 0},
				StatusEffect: status,
			}))
		},
		Entry("Blue Flute", ItemBlueFlute, StatusSleep),
		Entry("Red Flute", ItemRedFlute, StatusInfatuation),
		Entry("Yellow Flute", ItemYellowFlute, StatusConfusion),
	)

	DescribeTable("Stat Boosts",
		func(item Item, stat, stages int) {
			b, _ := setup(item)
			t, _ := b.SimulateRound()
			Expect(t).To(HaveTransaction(ModifyStatTransaction{
				Target: target{0, 0},
				Stat:   stat,
				Stages: stages,
			}))
			Expect(t).To(HaveTransaction(FriendshipTransaction{
				Target: target{0, 0},
				Amount: 1,
			}))
		},
		Entry("Dire Hit", ItemDireHit, StatCritChance, 2),
		Entry("X Accuracy", ItemXAccuracy, StatAccuracy, 1),
		Entry("X Attack", ItemXAttack, StatAtk, 1),
		Entry("X Defense", ItemXDefense, StatDef, 1),
		Entry("X Sp. Atk", ItemXSpAtk, StatSpAtk, 1),
		Entry("X Sp. Def", ItemXSpDef, StatSpDef, 1),
		Entry("X Speed", ItemXSpeed, StatSpeed, 1),
	)

	Specify("Guard Spec", func() {
		b, user := setup(ItemGuardSpec)
		b.SimulateRound()
		Expect(user.metadata).To(HaveKeyWithValue(MetaStatChangeImmune, 4))
		for stat := 0; stat < len(user.StatModifiers); stat += 1 {
			b.QueueTransaction(ModifyStatTransaction{
				Target: target{0, 0},
				Stat:   stat,
				Stages: -1,
			})
		}
		b.QueueTransaction(ModifyStatTransaction{
			Target:        target{0, 0},
			SelfInflicted: true,
			Stat:          StatAtk,
			Stages:        -2,
		})
		// No-op for RC Agent
		a1 <- ItemTurn{
			Item: ItemNone,
			Target: AgentTarget{
				target: target{0, 0},
			},
		}
		b.SimulateRound()
		Expect(user.StatModifiers).ToNot(ContainElement(-1))
		// Self-inflicted stat debuffs are valid
		Expect(user.StatModifiers[StatAtk]).To(Equal(-2))
		// Decreases every round
		Expect(user.metadata).To(HaveKeyWithValue(MetaStatChangeImmune, 3))
	})
})

var _ = Describe("In-a-pinch Berries", func() {
	a1 := Agent(new(dumbAgent))
	a2 := Agent(new(dumbAgent))
	var holder *Pokemon

	setup := func(item Item) *Battle {
		holder = GeneratePokemon(
			PkmnGrotle,
			WithLevel(25),
			WithMoves(MoveSplash),
			WithIVs([6]uint8{1, 1, 1, 20, 1, 1}),
		)
		holder.HeldItem = item
		holder.CurrentHP = holder.MaxHP() / 4
		b := New1v1Battle(
			GeneratePokemon(PkmnCombusken, WithLevel(25), WithMoves(MoveSplash)), &a1,
			holder, &a2,
		)
		b.rng = SimpleRNG()
		Expect(b.Start()).To(Succeed())
		return b
	}

	DescribeTable("Stat changing in-a-pinch berries",
		func(item Item, stat, stages int) {
			b := setup(item)
			t, _ := b.SimulateRound()

			Expect(t).To(HaveTransaction(ItemTransaction{
				Target: target{1, 0},
				IsHeld: true,
				Item:   holder.HeldItem,
			}))
			Expect(t).To(HaveTransaction(ModifyStatTransaction{
				Target: target{1, 0},
				Stat:   stat,
				Stages: stages,
			}))
			Expect(b.getPokemon(target{0, 0}).HeldItem).To(Equal(ItemNone))
			Expect(b.getPokemon(target{1, 0}).HeldItem).To(Equal(ItemNone))
		},
		Entry("Apicot Berry", ItemApicotBerry, StatSpDef, 1),
		Entry("Ganlon Berry", ItemGanlonBerry, StatDef, 1),
		Entry("Lansat Berry", ItemLansatBerry, StatCritChance, 2),
		Entry("Liechi Berry", ItemLiechiBerry, StatAtk, 1),
		Entry("Micle Berry", ItemMicleBerry, StatAccuracy, 1),
		Entry("Petaya Berry", ItemPetayaBerry, StatSpAtk, 1),
		Entry("Salac Berry", ItemSalacBerry, StatSpeed, 1),
		Entry("Starf Berry", ItemStarfBerry, StatSpeed, 2), //Don't know what to do with the StatSpeed (don't know how to implement rng into this)
	)
})
