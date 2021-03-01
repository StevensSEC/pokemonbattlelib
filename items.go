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
	ItemStatBoost ItemCategory = iota + 1
	ItemEffortDrop
	ItemMedicine
	ItemOther
	ItemInAPinch
	ItemPickyHealing
	ItemTypeProtection
	ItemBakingOnly
	ItemCollectibles
	ItemEvolution
	ItemSpelunking
	ItemHeldItems
	ItemChoice
	ItemEffortTraining
	ItemBadHeldItems
	ItemTraining
	ItemPlates
	ItemSpeciesSpecific
	ItemTypeEnhancement
	ItemEventItems
	ItemGameplay
	ItemPlotAdvancement
	ItemUnused
	ItemLoot
	ItemAllMail
	// Medicine
	ItemVitamins
	ItemHealing
	ItemPPRecovery
	ItemRevival
	ItemStatusCures

	ItemMulch
	ItemSpecialBalls
	ItemStandardBalls
	ItemDexCompletion
	ItemScarves
	ItemAllMachines
	ItemFlutes
	ItemApricornBalls
	ItemApricornBox
	ItemDataCards
	ItemJewels
	ItemMiracleShooter
	ItemMegaStones
	ItemMemories
	ItemZCrystals
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
// Can also use constants like ITEM_POTION or ITEM_REVIVE
// For item effects, see https://github.com/veekun/pokedex/blob/master/pokedex/data/csv/item_prose.csv
func GetItem(itemID int) Item {
	for _, item := range ALL_ITEMS {
		if item.ID == itemID {
			return item
		}
	}
	panic(fmt.Sprintf("unknown item with ID %v\n", itemID))
}

// Dispatches the correct item handler based on its category
func (p *Pokemon) UseItem(i *Item) []Transaction {
	switch i.Category {
	case ItemHealing, ItemRevival, ItemStatusCures:
		return p.UseMedicine(i)
	}
	return make([]Transaction, 0)
}

// Uses a medicine item which affects HP and status effects
func (p *Pokemon) UseMedicine(i *Item) (t []Transaction) {
	switch i.ID {
	case ITEM_POTION:
		t = append(t, p.RestoreHP(20))
	}
	return t
}
