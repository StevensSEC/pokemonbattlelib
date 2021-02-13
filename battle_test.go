package pokemonbattlelib

import (
	"fmt"
	"reflect"
	"testing"
)

type dumbAgent struct{}

// Blindly uses the first move on the first opponent pokemon.
func (a dumbAgent) Act(ctx *Context) Turn {
	// You can use `a` (reference to self) for self-targeting turns
	for agent, party := range ctx.Opponents {
		for i := range party {
			return FightTurn{
				agent:  agent,
				move:   0,
				target: i,
			}
		}
	}
	panic("no opponents found")
}

func TestBattleSetup(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := NewPokemon(4)
	party1.AddPokemon(pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := NewPokemon(7)
	party2.AddPokemon(pkmn2)
	b := NewBattle()
	b.AddParty(party1, party2)
}

func TestBattleOneRound(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := NewPokemon(4)
	party1.AddPokemon(pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := NewPokemon(7)
	party2.AddPokemon(pkmn2)
	b := NewBattle()
	b.AddParty(party1, party2)
	b.Start()
	for p, party := range b.parties {
		got := party.GetActivePokemon()
		if !reflect.DeepEqual(got, map[int]Pokemon{0: *party.pokemon[0]}) {
			t.Fatalf("Must send out first pokemon in each at the beginning of the battle. Party %d gave: %v", p, got)
		}
	}
	b.SimulateRound()
	// output:
	// TODO: Implement fight {0 1}
	// TODO: Implement fight {0 0}
}

// Faster pokemon should go first.
func TestPokemonSpeed(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := NewPokemon(4)
	pkmn1.Stats[5] = 10
	party1.AddPokemon(pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := NewPokemon(4)
	pkmn2.Stats[5] = 12
	party2.AddPokemon(pkmn2)
	b := NewBattle()
	b.AddParty(party1, party2)
	b.Start()
	b.SimulateRound()
	b.SimulateRound()
	// FIXME: ideally should check battle log/history
	// FIXME: For some reason, the output is not actually checked correctly, see #49
	// Output:
	// TODO: Implement fight {0 0}
	// TODO: Implement fight {0 1}
}

func TestTurnPriority(t *testing.T) {
	tests := []struct {
		turn Turn
		want int
	}{
		{
			turn: FightTurn{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%T priority", tt.turn), func(t *testing.T) {
			t.Parallel()
			got := tt.turn.Priority()
			if got != tt.want {
				t.Errorf("TurnPriority(%T) got %v, want %v", tt.turn, got, tt.want)
			}
		})
	}
}
