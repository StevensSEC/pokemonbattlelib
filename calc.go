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
			defense *= (defense * 150) / 100
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

	// apply modifiers
	elementalEffect := GetElementalEffect(move.Type(), receiver.EffectiveType())
	if elementalEffect > NormalEffect {
		damage <<= elementalEffect
	} else if elementalEffect < NormalEffect {
		damage >>= elementalEffect * -1 // bitshift operand must be positive
	}

	if rain, sun := weather == WeatherRain, weather == WeatherHarshSunlight; (rain && move.Type() == TypeWater) || (sun && move.Type() == TypeFire) {
		damage = (damage * 150) / 100
	} else if (rain && move.Type() == TypeFire) || (sun && move.Type() == TypeWater) {
		damage /= 2
	}

	if move != nil && user.EffectiveType()&move.Type() != 0 {
		if user.Ability == AbilityAdaptability {
			damage *= 2
		} else {
			damage = (damage * 150) / 100
		}
	}

	// Item modifiers
	switch user.HeldItem {
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
