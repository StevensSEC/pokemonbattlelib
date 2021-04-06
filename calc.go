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

	// Other modifiers
	elementalEffect := GetElementalEffect(move.Type(), receiver.EffectiveType())
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
	case ItemIronBall:
		// TODO: make flying not immune to ground
	case ItemChoiceBand, ItemChoiceScarf, ItemChoiceSpecs:
		// TODO: boost attack by 50% (in p.GetAttack)
		// TODO: boost special attak by 50% (in p.GetSpAtk)
		if lastMove := user.metadata[MetaLastMove]; lastMove != nil && lastMove != move {
			blog.Panicf("cannot use move blocked by %s", user.HeldItem.Name)
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
