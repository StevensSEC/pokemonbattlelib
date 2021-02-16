package pokemonbattlelib

import (
	"fmt"
	"reflect"
	"testing"
)

type dumbAgent struct{}

// Blindly uses the first move on the first opponent pokemon.
func (a dumbAgent) Act(ctx *BattleContext) Turn {
	// You can use `a` (reference to self) for self-targeting turns
	for _, target := range ctx.Targets {
		return FightTurn{
			Move:   0,
			Target: target,
		}
	}
	panic("no opponents found")
}

func TestBattleSetup(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := NewPokemon(4)
	party1.AddPokemon(&pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := NewPokemon(7)
	party2.AddPokemon(&pkmn2)
	b := NewBattle()
	b.AddParty(party1, party2)
}

func TestBattleOneRound(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := NewPokemon(4)
	pkmn1.Stats = [6]uint{30, 10, 10, 10, 10, 10}
	pkmn1.CurrentHP = 30
	pound := NewMove(1)
	pkmn1.Moves[0] = &pound
	party1.AddPokemon(&pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := NewPokemon(7)
	pkmn2.Stats = [6]uint{30, 10, 10, 10, 10, 10}
	pkmn2.CurrentHP = 30
	pkmn2.Moves[0] = &pound
	party2.AddPokemon(&pkmn2)
	b := NewBattle()
	b.AddParty(party1, party2)
	err := b.Start()
	if err != nil {
		t.Fatal("failed to start battle")
	}
	for p, party := range b.parties {
		got := party.GetActivePokemon()
		if !reflect.DeepEqual(got, map[int]Pokemon{0: *party.pokemon[0]}) {
			t.Fatalf("Must send out first pokemon in each at the beginning of the battle. Party %d gave: %v", p, got)
		}
	}
	b.SimulateRound()
	// functionally arbitrary value, will need to be adjusted when damage calculation becomes more accurate
	expectedHp := uint(27)
	for _, party := range b.parties {
		if party.pokemon[0].CurrentHP != expectedHp {
			t.Errorf("Expected %s to have %d HP, got %d", party.pokemon[0].GetName(), expectedHp, party.pokemon[0].CurrentHP)
		}
	}

	// output:
	// TODO: Implement fight {0 1}
	// TODO: Implement fight {0 0}
}

// Tests if active Pokemon are set correctly
func TestActivePokemon(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := NewPokemon(4)
	party1.AddPokemon(&pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := NewPokemon(7)
	pkmn3 := NewPokemon(9)
	party2.AddPokemon(&pkmn2, &pkmn3)
	b := NewBattle()
	b.AddParty(party1, party2)
	party1.SetActive(0)
	if len(b.parties[0].activePokemon) != 1 {
		t.Error("expected party 1 to have 1 active Pokemon")
	}
	party1.SetInactive(0)
	if len(b.parties[0].activePokemon) != 0 {
		t.Error("expected party 1 to have no active Pokemon")
	}
	party2.SetActive(1)
	if n := b.parties[1].activePokemon[1].NatDex; n != 9 {
		t.Errorf("expected party 2 to have an active Pokemon with dex number 9, received %v\n", n)
	}
}

func TestGetAllies(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := NewPokemon(4)
	party1.AddPokemon(&pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := NewPokemon(7)
	pkmn3 := NewPokemon(9)
	party2.AddPokemon(&pkmn2, &pkmn3)
	b := NewBattle()
	b.AddParty(party1, party2)
	err := b.Start()
	if err != nil {
		t.Fatal("failed to start battle")
	}
	if n := len(b.GetAllies(party1)); n != 1 {
		t.Errorf("expected party 1 to have 1 active ally, received %v\n", n)
	}
	if n := len(b.GetAllies(party2)); n != 1 {
		t.Errorf("expected party 2 to have 1 active ally, received %v\n", n)
	}
}

func TestGetOpponents(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := NewPokemon(4)
	party1.AddPokemon(&pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := NewPokemon(7)
	pkmn3 := NewPokemon(9)
	party2.AddPokemon(&pkmn2, &pkmn3)
	b := NewBattle()
	b.AddParty(party1, party2)
	err := b.Start()
	if err != nil {
		t.Fatal("failed to start battle")
	}
	if n := len(b.GetOpponents(party1)); n != 1 {
		t.Errorf("expected party 1 to have 1 active opponents, received %v\n", n)
	}
	if n := len(b.GetOpponents(party2)); n != 1 {
		t.Errorf("expected party 2 to have 1 active opponent, received %v\n", n)
	}
}

// Faster pokemon should go first.
func TestPokemonSpeed(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := NewPokemon(4)
	pound := NewMove(1)
	pkmn1.Moves[0] = &pound
	pkmn1.Stats = [6]uint{30, 10, 10, 10, 10, 10}
	party1.AddPokemon(&pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := NewPokemon(4)
	pkmn2.Moves[0] = &pound
	pkmn2.Stats = [6]uint{30, 10, 10, 10, 10, 12}
	party2.AddPokemon(&pkmn2)
	b := NewBattle()
	b.AddParty(party1, party2)
	err := b.Start()
	if err != nil {
		t.Fatal("failed to start battle")
	}
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
