package pokemonbattlelib

// A Party can use Items in a battle for different effects. A Pokemon can hold one Item.
type Item uint16

// Represents data associated with an `Item`.
type ItemData struct {
	Name        string
	Category    ItemCategory
	FlingPower  int
	FlingEffect FlingEffect
	Flags       ItemFlags
}

type ItemCategory uint8

// Fix: remove unnecessary items
const (
	ItemCategoryStatBoost ItemCategory = iota + 1
	ItemCategoryEffortDrop
	ItemCategoryMedicine
	ItemCategoryOther
	ItemCategoryInAPinch
	ItemCategoryPickyHealing
	ItemCategoryTypeProtection
	ItemCategoryBakingOnly
	ItemCategoryCollectibles
	ItemCategoryEvolution
	ItemCategorySpelunking
	ItemCategoryHeldItems
	ItemCategoryChoice
	ItemCategoryEffortTraining
	ItemCategoryBadHeldItems
	ItemCategoryTraining
	ItemCategoryPlates
	ItemCategorySpeciesSpecific
	ItemCategoryTypeEnhancement
	ItemCategoryEventItems
	ItemCategoryGameplay
	ItemCategoryPlotAdvancement
	ItemCategoryUnused
	ItemCategoryLoot
	ItemCategoryAllMail
	// Medicine
	ItemCategoryVitamins
	ItemCategoryHealing
	ItemCategoryPPRecovery
	ItemCategoryRevival
	ItemCategoryStatusCures

	ItemCategoryMulch
	ItemCategorySpecialBalls
	ItemCategoryStandardBalls
	ItemCategoryDexCompletion
	ItemCategoryScarves
	ItemCategoryAllMachines
	ItemCategoryFlutes
	ItemCategoryApricornBalls
	ItemCategoryApricornBox
	ItemCategoryDataCards
	ItemCategoryJewels
	ItemCategoryMiracleShooter
	ItemCategoryMegaStones
	ItemCategoryMemories
	ItemCategoryZCrystals
)

type FlingEffect uint8

const (
	FlingBadlyPoison FlingEffect = iota + 1
	FlingBurn
	FlingActivateBerry
	FlingActivateHerb
	FlingParalyze
	FlingPoison
	FlingFlinch
)

// Properties that an item can have.
type ItemFlags uint8

const (
	// The item can be consumed
	FlagConsumable ItemFlags = 1 << iota
	// The item can be held by a pokemon
	FlagHoldable
	// The item has passive effects when held.
	FlagHoldablePassive
	// The item can be used by a pokemon when held.
	FlagHoldableActive
	// The item can be used in battle
	FlagUsableInBattle
)

// Retrieves an Item's data.
// For item effects, see https://github.com/veekun/pokedex/blob/master/pokedex/data/csv/item_prose.csv
func (i Item) Data() *ItemData {
	if i == ItemNone {
		return &ItemData{}
	}
	return &AllItems[i-1]
}

// Dispatches the correct item handler based on its category
func (p *Pokemon) UseItem(i Item) []Transaction {
	switch i.Data().Category {
	case ItemCategoryHealing, ItemCategoryRevival, ItemCategoryStatusCures:
		return p.UseMedicine(i)
	case ItemCategoryInAPinch:
		return p.UseBerryInAPinch(i)
	}
	return make([]Transaction, 0)
}

// Uses a medicine item which affects HP and status effects
func (p *Pokemon) UseMedicine(i Item) (t []Transaction) {
	switch i {
	case ItemPotion:
		t = append(t, p.RestoreHP(20))
	}
	return t
}

func (p *Pokemon) UseBerryInAPinch(i Item) (t []Transaction) {
	switch i {
	case ItemApicotBerry:
		t = append(t, ModifyStatTransaction{
			Target: p,
			Stat:   StatSpDef,
			Stages: 1,
		})
	case ItemCustapBerry:
		// TODO: Force pokemon to go first
	case ItemGanlonBerry:
		t = append(t, ModifyStatTransaction{
			Target: p,
			Stat:   StatDef,
			Stages: 1,
		})
	case ItemLansatBerry:
		t = append(t, ModifyStatTransaction{
			Target: p,
			Stat:   StatCritChance,
			Stages: 2,
		})
	case ItemLiechiBerry:
		t = append(t, ModifyStatTransaction{
			Target: p,
			Stat:   StatAtk,
			Stages: 1,
		})
	case ItemMicleBerry:
		// TODO: Perfect accuracy for next move
	case ItemPetayaBerry:
		t = append(t, ModifyStatTransaction{
			Target: p,
			Stat:   StatSpAtk,
			Stages: 1,
		})
	case ItemSalacBerry:
		t = append(t, ModifyStatTransaction{
			Target: p,
			Stat:   StatSpeed,
			Stages: 1,
		})
	case ItemStarfBerry:
		// TODO: boost random stat, requires battle RNG to be available.
	}
	return t
}
