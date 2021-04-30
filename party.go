package pokemonbattlelib

import (
	"errors"
	"fmt"
	"log"
)

// Maximum number of Pokemon in a party
const MaxPartySize = 6

var ErrorPartyIndex = errors.New("invalid index for party")
var ErrorPartyFull = fmt.Errorf("party size cannot exceed max of %d Pokemon\n", MaxPartySize)

// A Party of Pokemon.
type Party struct {
	Pokemon      []*Pokemon
	PokemonRules PokemonValidationRules
}

func NewParty() *Party {
	return &Party{
		Pokemon:      make([]*Pokemon, 0),
		PokemonRules: PkmnRuleSetDefault,
	}
}

func NewOccupiedParty(pkmn ...*Pokemon) *Party {
	p := NewParty()
	err := p.AddPokemon(pkmn...)
	if err != nil {
		panic(err)
	}
	return p
}

// Adds 1 or more Pokemon to a Party
func (p *Party) AddPokemon(pkmn ...*Pokemon) error {
	if len(p.Pokemon)+len(pkmn) > MaxPartySize {
		return ErrorPartyFull
	}
	for i := range pkmn {
		if err := pkmn[i].Validate(p.PokemonRules); err != nil {
			return err
		}
	}
	p.Pokemon = append(p.Pokemon, pkmn...)
	return nil
}

// A Pokemon battleParty. Can hold up to 6 Pokemon. Also manages how many pokemon are out on the battlefield.
// Only used inside of Battles.
type battleParty struct {
	Party         *Party
	Agent         *Agent        // The agent that has control over this party
	activePokemon map[uint]uint // Map of slots for active Pokemon to actual party slots
	team          int           // The team that this party belongs to
}

func (p *battleParty) pokemon() []*Pokemon {
	return p.Party.Pokemon
}

// Adds 1 or more Pokemon to a Party
func (p *battleParty) AddPokemon(pkmn ...*Pokemon) error {
	return p.Party.AddPokemon(pkmn...)
}

// Sets a Pokemon to be active by its index in a party (0-5)
func (p *battleParty) SetActive(i int) {
	if p.IsActivePokemon(i) {
		log.Panicf("pokemon is already out on the battlefield")
	}
	p.activePokemon[i] = p.pokemon()[i]
}

// Sets a Pokemon to be inactive by its index in a party (0-5)
func (p *battleParty) SetInactive(i int) {
	if !p.IsActivePokemon(i) {
		log.Panicf("pokemon is not out on the battlefield")
	}
	delete(p.activePokemon, i)
}

// Checks if a Pokemon in a party is currently active
func (p *battleParty) IsActivePokemon(i int) bool {
	if i >= len(p.pokemon()) {
		log.Panicln(ErrorPartyIndex)
	}
	if _, ok := p.activePokemon[i]; ok {
		return true
	}
	return false
}

// Creates a map of party index to active Pokemon
func (p *battleParty) GetActivePokemon() map[int]Pokemon {
	allActive := make(map[int]Pokemon)
	for i, pokemon := range p.pokemon() {
		for _, active := range p.activePokemon {
			if pokemon == active {
				allActive[i] = *pokemon
			}
		}
	}
	return allActive
}
