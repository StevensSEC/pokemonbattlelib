package pokemonbattlelib

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
// Creates a new item from its ID
func NewItem(itemID int) *Item {
	item := ALL_ITEMS[itemID-1]
	return &item
}

// Dispatches the correct item handler based on its category
func (i *Item) UseItem(target *Pokemon) {
	switch i.Category {
	case ItemHealing, ItemRevival, ItemStatusCures:
		i.UseMedicine(target)
	case ItemVitamins:
		i.UseVitamins(target)
	}
}

// Uses a medicine item which affects HP and status effects
func (i *Item) UseMedicine(target *Pokemon) {
	switch i.ID {
	case ITEM_BERRY_JUICE, ITEM_POTION:
		target.RestoreHP(20)
	case ITEM_FRESH_WATER, ITEM_SUPER_POTION:
		target.RestoreHP(50)
	case ITEM_ANTIDOTE:
		target.CureStatusEffect(STATUS_POISON)
	}
}

// Uses a vitamin which affects EVs and friendship
func (i *Item) UseVitamins(target *Pokemon) {
	// https://bulbapedia.bulbagarden.net/wiki/Friendship#In_Generation_IV
	friendAmounts := []uint8{5, 3, 2}
	switch i.ID {
	case ITEM_CALCIUM:
		if target.EVs[STAT_ATK] < 100 {
			target.AddEVs(STAT_ATK, 10)
		}
	}
	target.AddFriendship(friendAmounts[target.Friendship/100])
}

// Uses a move item which affects PP
func (i *Item) UseMoveItem(target *Pokemon, move *Move) {
	switch i.ID {
	case ITEM_ELIXIR:
		for _, m := range target.Moves {
			m.RestorePP(10)
		}
	case ITEM_ETHER:
		move.RestorePP(10)
	case ITEM_MAX_ELIXIR:
		for _, m := range target.Moves {
			m.RestorePP(m.MaxPP)
		}
	case ITEM_MAX_ETHER:
		move.RestorePP(move.MaxPP)
	}
}
