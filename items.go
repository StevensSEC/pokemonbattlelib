package pokemonbattlelib

import "fmt"

type Item struct {
	ID          int
	Name        string
	Category    ItemCategory
	FlingPower  int
	FlingEffect FlingEffect
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

// For item effects, see https://github.com/veekun/pokedex/blob/master/pokedex/data/csv/item_prose.csv
// Retrieves an item using its ID
// Can also use constants like ITEM_POTION or ITEM_REVIVE
func GetItem(itemID int) Item {
	for _, item := range ALL_ITEMS {
		if item.ID == itemID {
			return item
		}
	}
	panic(fmt.Sprintf("unknown item with ID %v\n", itemID))
}

// Dispatches the correct item handler based on its category
func (p *Pokemon) UseItem(i *Item) {
	switch i.Category {
	case ItemHealing, ItemRevival, ItemStatusCures:
		p.UseMedicine(i)
	case ItemVitamins:
		p.UseVitamins(i)
	}
}

// Activates an held item's effect in battle
func (p *Pokemon) UseHeldItem() {
	if p.HeldItem == nil {
		return
	}
}

// Uses a medicine item which affects HP and status effects
func (p *Pokemon) UseMedicine(i *Item) {
	switch i.ID {
	case ITEM_BERRY_JUICE, ITEM_POTION:
		p.RestoreHP(20)
	case ITEM_FRESH_WATER, ITEM_SUPER_POTION:
		p.RestoreHP(50)
	case ITEM_ANTIDOTE:
		p.CureStatusEffect(STATUS_POISON)
	}
}

// Uses a vitamin which affects EVs and friendship
func (p *Pokemon) UseVitamins(i *Item) {
	// https://bulbapedia.bulbagarden.net/wiki/Friendship#In_Generation_IV
	friendAmounts := []uint8{5, 3, 2}
	switch i.ID {
	case ITEM_CALCIUM:
		if p.EVs[STAT_ATK] < 100 {
			p.AddEVs(STAT_ATK, 10)
		}
	}
	p.AddFriendship(friendAmounts[p.Friendship/100])
}

// Uses a move item which affects PP
func (p *Pokemon) UseMoveItem(i *Item, move *Move) {
	switch i.ID {
	case ITEM_ELIXIR:
		for _, m := range p.Moves {
			m.RestorePP(10)
		}
	case ITEM_ETHER:
		move.RestorePP(10)
	case ITEM_MAX_ELIXIR:
		for _, m := range p.Moves {
			m.RestorePP(m.MaxPP)
		}
	case ITEM_MAX_ETHER:
		move.RestorePP(move.MaxPP)
	}
}

// Gets the damage multiplier for items during a turn
func (p *Pokemon) GetDamageMultiplier(i *Item, move *Move) float64 {
	multiplier := 1.0
	switch i.ID {
	// Plates
	case ITEM_DRACO_PLATE:
		if move.Type == Dragon {
			multiplier *= 1.20
		}
	// Species-specific
	case ITEM_ADAMANT_ORB: // Affects Dialga
		if p.NatDex == 483 && (move.Type == Steel || move.Type == Dragon) {
			multiplier *= 1.20
		}
	}
	return multiplier
}
