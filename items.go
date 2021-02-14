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
