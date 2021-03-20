package pokemonbattlelib

import (
	"errors"
	"fmt"
	"log"
)

// A Pokemon party. Can hold up to 6 Pokemon. Also manages how many pokemon are out on the battlefield.
type party struct {
	Agent         *Agent           // The agent that has control over this party
	pokemon       []*Pokemon       // The Pokemon in the party
	activePokemon map[int]*Pokemon // Map containing slots and references to active Pokemon on the battlefield
	team          int              // The team that this party belongs to
	PokemonRules  PokemonValidationRules
}

// Maximum number of Pokemon in a party
const MaxPartySize = 6

var PartyIndexError = errors.New("invalid index for party")
var ErrorPartyFull = fmt.Errorf("party size cannot exceed max of %d Pokemon\n", MaxPartySize)

// Creates a new party to store Pokemon and assigns them to a team
func NewParty(agent *Agent, team int) *party {
	return &party{
		Agent:         agent,
		pokemon:       make([]*Pokemon, 0),
		activePokemon: make(map[int]*Pokemon),
		team:          team,
		PokemonRules:  PkmnRuleSetDefault,
	}
}

// Creates a new party and fills it out with the passed Pokemon
func NewOccupiedParty(agent *Agent, team int, pkmn ...*Pokemon) *party {
	party := NewParty(agent, team)
	err := party.AddPokemon(pkmn...)
	if err != nil {
		panic(err)
	}
	return party
}

// Adds 1 or more Pokemon to a Party
func (p *party) AddPokemon(pkmn ...*Pokemon) error {
	if len(p.pokemon)+len(pkmn) > MaxPartySize {
		return ErrorPartyFull
	}
	for i := range pkmn {
		if err := pkmn[i].Validate(p.PokemonRules); err != nil {
			return err
		}
	}
	p.pokemon = append(p.pokemon, pkmn...)
	return nil
}

// Sets a Pokemon to be active by its index in a party (0-5)
func (p *party) SetActive(i int) {
	if p.IsActivePokemon(i) {
		log.Panicf("pokemon is already out on the battlefield")
	}
	p.activePokemon[i] = p.pokemon[i]
}

// Sets a Pokemon to be inactive by its index in a party (0-5)
func (p *party) SetInactive(i int) {
	if !p.IsActivePokemon(i) {
		log.Panicf("pokemon is not out on the battlefield")
	}
	delete(p.activePokemon, i)
}

// Checks if a Pokemon in a party is currently active
func (p *party) IsActivePokemon(i int) bool {
	if i >= len(p.pokemon) {
		log.Panicln(PartyIndexError)
	}
	if _, ok := p.activePokemon[i]; ok {
		return true
	}
	return false
}

// Creates a map of party index to active Pokemon
func (p *party) GetActivePokemon() map[int]Pokemon {
	allActive := make(map[int]Pokemon)
	for i, pokemon := range p.pokemon {
		for _, active := range p.activePokemon {
			if pokemon == active {
				allActive[i] = *pokemon
			}
		}
	}
	return allActive
}
