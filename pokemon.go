package pokemonbattlelib

type Pokemon struct {
    NatDex uint16                    // National Pokedex Number
    Nickname string                  // player-given name for the Pokemon
    Level Level                      // value from 1-100 influencing stats
    Ability Ability                  // name of this Pokemon's ability
    TotalExperience uint             // the total amount of exp this Pokemon has gained, influencing its level
    IVs [6]IndividualValue           // values from 0-31 that represents a Pokemon's 'genetic' potential
    EVs [6]EffortValue               // values from 0-255 that represents a Pokemon's training in a particular stat
    Nature Nature                    // represents a Pokemon's disposition and affects stats
    Stats [6]Stat                    // the actual stats of a Pokemon determined from the above data
    CurrentHP uint                   // the remaining HP of this Pokemon 
    HeldItem Item                    // the item a Pokemon is holding
    Moves [4]Move                    // the moves the Pokemon currenly knows
    Friendship uint8                 // how close this Pokemon is to its Trainer
    OriginalTrainerName string       // the name of the first Trainer who caught this Pokemon
    OriginalTrainerID uint16         // a number associated with the first Trainer who caught this Pokemon
}

func (p *Pokemon) GetName() string {
	return PokemonNames[p.NatDex]
}

func (p *Pokemon) GetExpGroup() ExpGroup {
    //TODO
}

func (p *Pokemon) GetElementalTypes() (ElementalType, ElementalType) {
    // TODO
}

func (p *Pokemon) GetBaseStats() BaseStats[6] {
    // TODO
} 

func (p *Pokemon) GetMoveList() []Move {
    //TODO
}

func (p *Pokemon) GetLearnSet() map[Level][]Move {
    //TODO
}

func (p *Pokemon) GetBaseExperienceYield() uint8 {
    //TODO
} 

func (p *Pokemon) GetEffortValueYield() [6]uint8 {
    //TODO
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
    // and set the stats to the computed amounts
}

type Level uint8

func (l Level) IsInValidRange() bool {
    return l > Level(1) && l <= Level(100)
}

type BaseStat uint8

type IndividualValue uint8

func (iv IndividualValue) IsInValidRange() bool {
    return iv > IndividualValue(0) && iv <= IndividualValue(31)
}

type EffortValue uint8

type Stat uint

type Item string

