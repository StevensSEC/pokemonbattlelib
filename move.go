package pokemonbattlelib

type MoveCategory uint8

const (
	MoveCategoryStatus MoveCategory = iota
	MoveCategoryPhysical
	MoveCategorySpecial
)

func (c MoveCategory) String() string {
	switch c {
	case MoveCategoryStatus:
		return "Status"
	case MoveCategoryPhysical:
		return "Physical"
	case MoveCategorySpecial:
		return "Special"
	default:
		panic("Unexpected value for move category")
	}
}

// Sets the bounds on move priority to [-7, 5]
const (
	MovePriorityMin = 5
	MovePriorityMax = -7
)

// Targets that the move can specify
type MoveTarget uint8

const (
	MoveTargetSpecificMove MoveTarget = iota + 1
	MoveTargetSelectedMeFirst
	MoveTargetAlly
	MoveTargetUsersField
	MoveTargetUserOrAlly
	MoveTargetOpponentsField
	MoveTargetUser
	MoveTargetRandomOpponent
	MoveTargetAllOthers
	MoveTargetSelected
	MoveTargetAllOpponents
	MoveTargetEntireField
	MoveTargetUserAndAllies
	MoveTargetAll
)

// Helps define the bahavior of moves.
type MoveMetaCategory uint8

const (
	MoveMetaCategoryDamage           MoveMetaCategory = iota // Deal damage to target pokemon.
	MoveMetaCategoryAilment                                  // Apply a status condition to the target of the move.
	MoveMetaCategoryNetGoodStats                             // Modify the target pokemon's stats.
	MoveMetaCategoryHeal                                     // Heals the user or a target ally.
	MoveMetaCategoryDamageAilment                            // Deal damage to target pokemon, and have a chance to apply a status condition to the target of the move.
	MoveMetaCategorySwagger                                  // Raise the target's stats, and apply confusion.
	MoveMetaCategoryDamageLower                              // Deal damage to target pokemon, and have a chance to apply a stat modifier to the target of the move.
	MoveMetaCategoryDamageRaise                              // Deal damage to target pokemon, and have a chance to apply a stat modifier to the user of the move.
	MoveMetaCategoryDamageHeal                               // Deal damage to target pokemon, and restore some of the user's HP. However, this is effectively a noop, because this functionality is only present on moves that have `Drain != 0`.
	MoveMetaCategoryOhko                                     // One hit knock out. If the move hits the target pokemon, it is generally guaranteed to die (there are exceptions). Always fails if the target pokemon's level is greater than the user's level.
	MoveMetaCategoryWholeFieldEffect                         // Applies a special effect to the entire battlefield (usually weather).
	MoveMetaCategoryFieldEffect                              // Applies a special effect to the user's side of the battlefield.
	MoveMetaCategoryForceSwitch                              // Forces the target pokemon to switch out.
	MoveMetaCategoryUnique                                   // For moves that have specific, one-off rules.
)

type MoveFlags uint32

const (
	// This move ignores the target's Substitute.
	FlagAuthentic MoveFlags = 1 << iota
	// This move is blocked by Bulletproof.
	FlagBallistics
	// This move has 1.5× its usual power when used by a Pokémon with Strong Jaw.
	FlagBite
	// This move has a charging turn that can be skipped with a Power Herb.
	FlagCharge
	// User touches the target. This triggers some abilities (e.g., Static) and items (e.g., Sticky Barb).
	FlagContact
	// This move triggers the ability Dancer.
	FlagDance
	// This move can be used while frozen to force the Pokémon to defrost.
	FlagDefrost
	// In triple battles, this move can be used on either side to target the farthest away opposing Pokémon. (non gen 4)
	FlagDistance
	// This move cannot be used in high Gravity.
	FlagGravity
	// This move is blocked by Heal Block.
	FlagHeal
	// This move is blocked by Aroma Veil and cured by Mental Herb.
	FlagMental
	// Copied by Mirror Move: A Pokémon targeted by this move can use Mirror Move to copy it.
	FlagMirror
	// Not usable in sky battles (non gen 4)
	FlagNonSkyBattle
	// Pokémon with Overcoat and Grass-type Pokémon are immune to this move.
	FlagPowder
	// This move will not work if the target has used Detect or Protect this turn.
	FlagProtect
	// This move has 1.5× its usual power when used by a Pokémon with Mega Launcher. (non gen 4)
	FlagPulse
	// This move has 1.2× its usual power when used by a Pokémon with Iron Fist.
	FlagPunch
	// The turn after this move is used, the Pokémon's action is skipped so it can recharge.
	FlagRecharge
	// This move may be reflected back at the user with Magic Coat or Magic Bounce.
	FlagReflectable
	// This move will be stolen if another Pokémon has used Snatch this turn.
	FlagSnatch
	// Pokémon with Soundproof are immune to this move.
	FlagSound
)

type MoveId uint16

type MoveData struct {
	Name         string
	Type         Type
	Category     MoveCategory
	Targets      MoveTarget
	Priority     int8
	Power        uint
	Accuracy     uint
	InitialMaxPP uint8
	// Metadata
	MinHits       int
	MaxHits       int
	MinTurns      int
	MaxTurns      int
	Drain         int // Percentage (out of 100) of the damage dealt by this move is given back to the user as HP. If the value is negative, it's recoil damage.
	Healing       int
	CritRate      int
	AilmentChance int
	FlinchChance  int
	StatChance    int
	Flags         MoveFlags
	AffectedStat  uint8 // if this move modifies stats, this is the stat it modifies
	StatStages    int8  // if 0, then this move does not modify stats
	Ailment       StatusCondition
	MetaCategory  MoveMetaCategory
}

// Represents a Pokemon's move. Moves can deal damage, heal the user or allies, or cause status effects.
type Move struct {
	Id        MoveId
	CurrentPP uint8
	MaxPP     uint8
}

//go:generate go run ./scripts/getters/gen_getters.go -for Move -data MoveData

// Retrieves a Pokemon move given its move ID
func GetMove(id MoveId) *Move {
	m := Move{
		Id: id,
	}
	m.CurrentPP = m.InitialMaxPP()
	m.MaxPP = m.InitialMaxPP()
	return &m
}

// Grabs move's constant data
func (m *Move) Data() *MoveData {
	if m.Id > MoveId(len(AllMoves)) {
		blog.Panicf("Move (id: %d) is an invalid move", m.Id)
	}
	if m.Id == MoveNone {
		return &MoveData{}
	}
	return &AllMoves[m.Id-1]
}

func (m Move) String() string {
	return m.Name()
}
