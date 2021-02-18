package pokemonbattlelib

import (
	"errors"
	"log"
)

// A Pokemon party. Can hold up to 6 Pokemon. Also manages how many pokemon are out on the battlefield.
type party struct {
	Agent         *Agent           // The agent that has control over this party
	pokemon       []*Pokemon       // The Pokemon in the party
	activePokemon map[int]*Pokemon // Map containing slots and references to active Pokemon on the battlefield
	team          int              // The team that this party belongs to
}

const MAX_PARTY_SIZE = 6

var PartyIndexError = errors.New("invalid index for party")

// Creates a new party to store Pokemon and assigns them to a team
func NewParty(agent *Agent, team int) *party {
	return &party{
		Agent:         agent,
		pokemon:       make([]*Pokemon, 0),
		activePokemon: make(map[int]*Pokemon),
		team:          team,
	}
}

// Creates a new party and fills it out with the passed Pokemon
func NewOccupiedParty(agent *Agent, team int, pkmn ...*Pokemon) *party {
	party := NewParty(agent, team)
	party.AddPokemon(pkmn...)
	return party
}

// Creates a new party and fills it out with Pokemon corresponding the the dex number array passed
// at the specified level
func NewPartyFromDexNumbers(agent *Agent, team int, dexNums []int, level int) *party {
	if len(dexNums) > MAX_PARTY_SIZE {
		log.Panicf("party size cannot exceed max of %v Pokemon\n", MAX_PARTY_SIZE)
	}
	party := NewParty(agent, team)
	for _, dexNum := range dexNums {
		pkmn := GeneratePokemon(uint16(dexNum), uint8(level))
		party.AddPokemon(&pkmn)
	}
	return party
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
