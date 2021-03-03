package pokemonbattlelib

import "fmt"

// A Party can use Items in a battle for different effects. A Pokemon can hold one Item.
type Item struct {
	ID          int
	Name        string
	Category    ItemCategory
	FlingPower  int
	FlingEffect FlingEffect
	Flags       ItemFlags
}
type ItemCategory int

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

type FlingEffect int

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

// Retrieves an item using its ID
// Can also use constants like ItemPotion or ItemRevive
// For item effects, see https://github.com/veekun/pokedex/blob/master/pokedex/data/csv/item_prose.csv
func GetItem(itemID int) Item {
	for _, item := range AllItems {
		if item.ID == itemID {
			return item
		}
	}
	panic(fmt.Sprintf("unknown item with ID %v\n", itemID))
}

// Dispatches the correct item handler based on its category
func (p *Pokemon) UseItem(i *Item) []Transaction {
	switch i.Category {
	case ItemCategoryHealing, ItemCategoryRevival, ItemCategoryStatusCures:
		return p.UseMedicine(i)
	case ItemCategoryInAPinch:
		return p.UseBerryInAPinch(i)
	}
	return make([]Transaction, 0)
}

// Uses a medicine item which affects HP and status effects
func (p *Pokemon) UseMedicine(i *Item) (t []Transaction) {
	switch i.ID {
	case ItemPotion:
		t = append(t, p.RestoreHP(20))
	}
	return t
}

func (p *Pokemon) UseBerryInAPinch(i *Item) (t []Transaction) {
	switch i.ID {
	case ItemApicotBerry:
	case ItemCustapBerry:
	case ItemGanlonBerry:
	case ItemLansatBerry:
	case ItemLiechiBerry:
	case ItemMicleBerry:
	case ItemPetayaBerry:
	case ItemSalacBerry:
	case ItemStarfBerry:
	}
	return t
}
