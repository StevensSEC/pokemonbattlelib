package pokemonbattlelib

func CalcMoveDamage(weather Weather, user, receiver *Pokemon, move *Move) (damage uint) {
	// Compute base damage
	levelEffect := uint((2 * user.Level / 5) + 2)
	movePower := move.Power()
	attack := user.Attack()
	defense := receiver.Defense()
	// Move modifiers
	if move.Category() == MoveCategorySpecial {
		attack = user.SpecialAttack()
		defense = receiver.SpecialDefense()
	}
	// Weather modifiers
	if weather == WeatherSandstorm {
		if receiver.EffectiveType()&TypeRock != 0 {
			defense = (defense * 150) / 100
		}
		if move.Id == MoveSolarBeam {
			movePower /= 2
		}
	}
	if weather == WeatherHail && move.Id == MoveSolarBeam {
		movePower /= 2
	}
	if weather == WeatherFog {
		if move.Id == MoveWeatherBall {
			movePower *= 2
		} else if move.Id == MoveSolarBeam {
			movePower /= 2
		}
	}
	damage = uint((((levelEffect * movePower * attack / defense) / 50) + 2))

	// Other modifiers
	elementalEffect := GetElementalEffect(move.Type(), receiver.EffectiveType(), receiver.HeldItem)
	// Account for ground type moves on grounded Pokemon
	if move.Type()&TypeGround > 0 && receiver.Type&TypeFlying > 0 && receiver.IsGrounded() {
		elementalEffect -= NoEffect
	}
	if elementalEffect > NormalEffect {
		damage <<= elementalEffect
	} else if elementalEffect < NormalEffect {
		damage >>= elementalEffect * -1 // bitshift operand must be positive
	}
	// Weather type modifier
	if rain, sun := weather == WeatherRain, weather == WeatherHarshSunlight; (rain && move.Type() == TypeWater) || (sun && move.Type() == TypeFire) {
		damage = (damage * 150) / 100
	} else if (rain && move.Type() == TypeFire) || (sun && move.Type() == TypeWater) {
		damage /= 2
	}
	// Stab modifier
	if move != nil && user.EffectiveType()&move.Type() != 0 {
		if user.Ability == AbilityAdaptability {
			damage *= 2
		} else {
			damage = (damage * 150) / 100
		}
	}
	// Item modifiers
	switch user.HeldItem {
	case ItemExpertBelt:
		if elementalEffect >= SuperEffective {
			damage = (damage * 120) / 100
		}
	case ItemChoiceBand, ItemChoiceScarf, ItemChoiceSpecs:
		if lastMove := user.metadata[MetaLastMove]; lastMove != nil && lastMove != move {
			blog.Panicf("cannot use move blocked by %s", user.HeldItem.Name())
		}
	case ItemLifeOrb:
		damage = (damage * 130) / 100
	case ItemMuscleBand:
		if move.Category() == MoveCategoryPhysical {
			damage = (damage * 110) / 100
		}
	case ItemWiseGlasses:
		if move.Category() == MoveCategorySpecial {
			damage = (damage * 110) / 100
		}
	}
	return damage
}

// Calculate the accuracy of a move, accounting for weather/evasion/held items/etc.
func CalcAccuracy(weather Weather, user, receiver *Pokemon, move *Move) uint {
	evasion := receiver.Evasion()
	accuracy := (move.Accuracy() * user.Accuracy() * evasion) / (100 * 100) // multiply by 100 twice because the user's accuracy and target's evasion are out of 100
	if weather == WeatherFog {
		accuracy = (accuracy * 3) / 5
	}
	switch user.HeldItem {
	case ItemWideLens:
		accuracy = (accuracy * 110) / 100
	}
	switch receiver.HeldItem {
	case ItemBrightPowder, ItemLaxIncense:
		accuracy -= accuracy / 10
	}
	return accuracy
}
