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
	return &ALL_ITEMS[itemID-1]
}

func (i *Item) UseItem(target *Pokemon, params interface{}) {
	switch i.Category {
	case ItemVitamins, ItemHealing, ItemRevival, ItemStatusCures:
		i.UseMedicine(target)
	}
}

func (i *Item) UseMedicine(target *Pokemon) {
	switch i.Name {
	case "berry-juice", "potion", "sweet-heart":
		target.RestoreHP(20)
	case "fresh-water", "super-potion":
		target.RestoreHP(50)
	case "antidote":
		target.CureStatusEffect(STATUS_POISON)
	}
}

func (i *Item) UseMoveItem(target *Pokemon, move *Move) {
	switch i.Name {
	case "elixir":
		for _, m := range target.Moves {
			m.RestorePP(10)
		}
	case "ether":
		move.RestorePP(10)
	case "max-elixir":
		for _, m := range target.Moves {
			m.RestorePP(m.MaxPP)
		}
	case "max-ether":
		move.RestorePP(move.MaxPP)
	}
}
