package pokemonbattlelib

var fullCure = StatusNonvolatileMask | StatusConfusion
var medicineData = map[Item]struct {
	Heal       uint
	Cure       StatusCondition
	Friendship [3]int
}{
	// Healing
	ItemBerryJuice:   {Heal: 20},
	ItemEnergyPowder: {Heal: 50, Friendship: [3]int{-5, -5, -10}},
	ItemEnergyRoot:   {Heal: 200, Friendship: [3]int{-10, -10, -15}},
	ItemFreshWater:   {Heal: 50},
	ItemFullRestore:  {Cure: fullCure},
	ItemHyperPotion:  {Heal: 200},
	ItemLemonade:     {Heal: 80},
	ItemMoomooMilk:   {Heal: 100},
	ItemPotion:       {Heal: 20},
	ItemSodaPop:      {Heal: 60},
	ItemSuperPotion:  {Heal: 50},
	// Revival
	ItemRevivalHerb: {Friendship: [3]int{-15, -15, -20}},
	// Status Cures
	ItemAntidote:     {Cure: StatusPoison},
	ItemAwakening:    {Cure: StatusSleep},
	ItemBurnHeal:     {Cure: StatusBurn},
	ItemFullHeal:     {Cure: fullCure},
	ItemHealPowder:   {Cure: fullCure, Friendship: [3]int{-5, -5, -10}},
	ItemIceHeal:      {Cure: StatusFreeze},
	ItemLavaCookie:   {Cure: fullCure},
	ItemOldGateau:    {Cure: fullCure},
	ItemParalyzeHeal: {Cure: StatusParalyze},
}
var typeItemData = map[Item]Type{
	// Plates
	ItemDracoPlate:  TypeDragon,
	ItemDreadPlate:  TypeDark,
	ItemEarthPlate:  TypeGround,
	ItemFistPlate:   TypeFighting,
	ItemFlamePlate:  TypeFire,
	ItemIciclePlate: TypeIce,
	ItemInsectPlate: TypeBug,
	ItemIronPlate:   TypeSteel,
	ItemMeadowPlate: TypeGrass,
	ItemMindPlate:   TypePsychic,
	ItemSkyPlate:    TypeFlying,
	ItemSplashPlate: TypeWater,
	ItemSpookyPlate: TypeGhost,
	ItemStonePlate:  TypeRock,
	ItemToxicPlate:  TypePoison,
	ItemZapPlate:    TypeElectric,
	// Type Enhancement
	ItemBlackBelt:    TypeFighting,
	ItemBlackGlasses: TypeDark,
	ItemCharcoal:     TypeFire,
	ItemDragonFang:   TypeDragon,
	ItemHardStone:    TypeRock,
	ItemMagnet:       TypeElectric,
	ItemMetalCoat:    TypeSteel,
	ItemMiracleSeed:  TypeGrass,
	ItemMysticWater:  TypeWater,
	ItemNeverMeltIce: TypeIce,
	ItemOddIncense:   TypePsychic,
	ItemPoisonBarb:   TypePoison,
	ItemRockIncense:  TypeRock,
	ItemRoseIncense:  TypeGrass,
	ItemSeaIncense:   TypeWater,
	ItemSharpBeak:    TypeFlying,
	ItemSilkScarf:    TypeNormal,
	ItemSilverPowder: TypeBug,
	ItemSoftSand:     TypeGround,
	ItemSpellTag:     TypeGhost,
	ItemTwistedSpoon: TypePsychic,
	ItemWaveIncense:  TypeWater,
}

//go:generate go run data/gen.go
