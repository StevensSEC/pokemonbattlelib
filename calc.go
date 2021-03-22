package pokemonbattlelib

func calcMoveDamage(b *Battle, user, receiver *Pokemon, move *Move) (damage float64) {
	weather := 1.0
	if rain, sun := b.Weather == WeatherRain, b.Weather == WeatherHarshSunlight; (rain && move.Type() == TypeWater) || (sun && move.Type() == TypeFire) {
		weather = 1.5
	} else if (rain && move.Type() == TypeFire) || (sun && move.Type() == TypeWater) {
		weather = 0.5
	}
	crit := 1.0
	if b.rng.Roll(1, user.CritChance()) {
		crit = 2.0
	}
	stab := 1.0
	if move != nil && user.EffectiveType()&move.Type() != 0 {
		stab = 1.5
		if user.Ability == AbilityAdaptability {
			stab = 2.0
		}
	}
	// Compute damage modifier
	modifier := weather * crit * stab
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
	if b.Weather == WeatherSandstorm {
		if receiver.EffectiveType()&TypeRock != 0 {
			defense *= 1.5
		}
		if move.Id == MoveSolarBeam {
			movePower /= 2
		}
	}
	if b.Weather == WeatherHail && move.Id == MoveSolarBeam {
		movePower /= 2
	}
	if b.Weather == WeatherFog {
		if move.Id == MoveWeatherBall {
			movePower *= 2
		} else if move.Id == MoveSolarBeam {
			movePower /= 2
		}
	}
	// Item modifiers
	switch user.HeldItem {
	case ItemLifeOrb:
		modifier *= 1.30
	case ItemMuscleBand:
		if move.Category() == MoveCategoryPhysical {
			modifier *= 1.10
		}
	case ItemWiseGlasses:
		if move.Category() == MoveCategorySpecial {
			modifier *= 1.10
		}
	}
	damage = (((levelEffect * movePower * attack / defense) / 50) + 2) * modifier
	return damage
}
