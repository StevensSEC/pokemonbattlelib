package pokemonbattlelib

func calcMoveDamage(weather Weather, user, receiver *Pokemon, move *Move) (damage uint) {
	weatherMod := 1.0
	if rain, sun := weather == WeatherRain, weather == WeatherHarshSunlight; (rain && move.Type() == TypeWater) || (sun && move.Type() == TypeFire) {
		weatherMod = 1.5
	} else if (rain && move.Type() == TypeFire) || (sun && move.Type() == TypeWater) {
		weatherMod = 0.5
	}

	stab := 1.0
	if move != nil && user.EffectiveType()&move.Type() != 0 {
		stab = 1.5
		if user.Ability == AbilityAdaptability {
			stab = 2.0
		}
	}

	modifier := weatherMod * stab
	levelEffect := float64((2 * user.Level / 5) + 2)
	movePower := float64(move.Power())
	attack := float64(user.Attack())
	defense := float64(receiver.Defense())
	// Move modifiers
	if move.Category() == MoveCategorySpecial {
		attack = float64(user.SpecialAttack())
		defense = float64(receiver.SpecialDefense())
	}
	// Weather modifiers
	if weather == WeatherSandstorm {
		if receiver.EffectiveType()&TypeRock != 0 {
			defense *= 1.5
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
	damage = uint((((levelEffect * movePower * attack / defense) / 50) + 2) * modifier)
	elementalEffect := GetElementalEffect(move.Type(), receiver.EffectiveType())
	if elementalEffect > NormalEffect {
		damage <<= elementalEffect
	} else if elementalEffect < NormalEffect {
		damage >>= elementalEffect * -1 // bitshift operand must be positive
	}
	return damage
}
