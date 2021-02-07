package pokemonbattlelib

type Pokemon struct {
    NatDex uint16                    // National Pokedex Number
    SpeciesName string               // the name of the species (e.g. Bulbasaur)
    Nickname string                  // player-given name for the Pokemon
    Level Level                      // value from 1-100 influencing stats
    Ability Ability                  // name of this Pokemon's ability
    TotalExperience uint             // the total amount of exp this Pokemon has gained, influencing its level
    ExpGroup ExpGroup                // the experience group this Pokemon belongs to
    Types [2]ElementalType           // this Pokemon's type(s)
    BaseStats [6]BaseStat            // values from 1-255 that represent a Pokemon potential based on their species
    IVs [6]IndividualValue           // values from 0-31 that represents a Pokemon's 'genetic' potential
    EVs [6]EffortValue               // values from 0-255 that represents a Pokemon's training in a particular stat
    Nature Nature                    // represents a Pokemon's disposition and affects stats
    Stats [6]Stat                    // the actual stats of a Pokemon determined from the above data
    HeldItem Item                    // the item a Pokemon is holding
    Moves [4]Move                    // the moves the Pokemon currenly knows
    MoveList []Move                  // all moves this Pokemon is eligible to learn
    LearnSet map[Level][]Move        // the moves that this Pokemon learns at each new level
    Friendship uint8                 // how close this Pokemon is to its Trainer
    BaseExperienceYield uint8        // a factor influencing how much experience this Pokemon will provide when defeated
    EffortValueYield [6]uint8        // the effort values this Pokemon will provide when defeated
    OriginalTrainerName string       // the name of the first Trainer who caught this Pokemon
    OriginalTrainerID uint16         // a number associated with the first Trainer who caught this Pokemon
}

func (p *Pokemon) GetName() string {
	return PokemonNames[p.NatDex]
}

func (p *Pokemon) ComputeLevel() {
    // TODO: Implement function to compute
    // pokemon's level based on its
    // total exp and exp group
    // and set the level to that computed amount
} 

func (p *Pokemon) ComputeStats() {
    // TODO: Implement function to compute
    // Pokemon's actual stats based on the Pokemon's
    // base stats, IVs, EVs, level, and Nature
}

type Level uint8

func (l Level) IsInValidRange() bool {
    return l > Level(1) && l <= Level(100)
}

type ExpGroup int

const (
    InvalidGroup ExpGroup = iota
    Erratic
    Fast
    MediumFast
    MediumSlow
    Slow
    Fluctuating
)

type BaseStat uint8

type IndividualValue uint8

func (iv IndividualValue) IsInValidRange() bool {
    return iv > IndividualValue(0) && iv <= IndividualValue(31)
}

type EffortValue uint8

type Stat uint

type Item string

