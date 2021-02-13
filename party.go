package pokemonbattlelib

import (
	"errors"
	"log"
)

// A Pokemon party. Can hold up to 6 Pokemon. Also manages how many pokemon are out on the battlefield.
type party struct {
	Agent         *Agent     // The agent that has control over this party
	pokemon       []*Pokemon // The Pokemon in the party
	activePokemon []*Pokemon // The pokemon in the party that are out on the battlefield
}

const MAX_PARTY_SIZE = 6

var PartyIndexError = errors.New("invalid index for party")

// Creates a new party to store Pokemon
func NewParty(agent *Agent) *party {
	return &party{
		Agent:         agent,
		pokemon:       make([]*Pokemon, 0),
		activePokemon: make([]*Pokemon, 0),
	}
}

// Adds 1 or more Pokemon to a Party
func (p *party) AddPokemon(pkmn ...*Pokemon) {
	if len(p.pokemon)+len(pkmn) > MAX_PARTY_SIZE {
		log.Panicf("party size cannot exceed max of %v Pokemon\n", MAX_PARTY_SIZE)
	}
	p.pokemon = append(p.pokemon, pkmn...)
}

// Sets a Pokemon to be active by its index in a party (0-5)
func (p *party) SetActive(i int) {
	if p.IsActivePokemon(i) {
		log.Panicf("pokemon is already out on the battlefield")
	}
	p.activePokemon = append(p.activePokemon, p.pokemon[i])
}

// Sets a Pokemon to be inactive by its index in a party (0-5)
func (p *party) SetInactive(i int) {
	if !p.IsActivePokemon(i) {
		log.Panicf("pokemon is not out on the battlefield")
	}
	newActive := make([]*Pokemon, 0)
	for _, active := range p.activePokemon {
		if active != p.pokemon[i] {
			newActive = append(newActive, active)
		}
	}
	p.activePokemon = newActive
}

// Checks if a Pokemon in a party is currently active
func (p *party) IsActivePokemon(i int) bool {
	if i >= len(p.pokemon) {
		log.Panicln(PartyIndexError)
	}
	for _, active := range p.activePokemon {
		if active == p.pokemon[i] {
			return true
		}
	}
	return false
}
