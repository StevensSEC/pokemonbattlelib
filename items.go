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

func (i Item) Name() string {
	return i.Data().Name
}

func (i Item) Category() ItemCategory {
	return i.Data().Category
}

func (i Item) FlingPower() int {
	return i.Data().FlingPower
}

func (i Item) FlingEffect() FlingEffect {
	return i.Data().FlingEffect
}

func (i Item) Flags() ItemFlags {
	return i.Data().Flags
}
