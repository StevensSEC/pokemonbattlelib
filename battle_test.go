package pokemonbattlelib

import (
	"fmt"
	"reflect"
	"testing"
)

// Blindly uses the first move on the first opponent pokemon.
type dumbAgent struct{}

func (dumbAgent) Act(b BattleInfo) Turn {
	opponent := b.Opponents()[0]
	return FightTurn{
		moveIdx:   0,
		targetIdx: opponent,
	}
}

func TestBattleSetup(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	b := Battle{}
	party1 := Party{
		Agent: &a1,
	}
	pkmn1 := GetPokemon(4)
	party1.AddPokemon(&pkmn1)
	party2 := Party{
		Agent: &a2,
	}
	pkmn2 := GetPokemon(7)
	party2.AddPokemon(&pkmn2)
	b.AddParty(&party1, &party2)
	b.SetTeams([][]int{{0}, {1}})
}

func TestBattleOneRound(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	b := Battle{}
	party1 := Party{
		Agent: &a1,
	}
	pound := GetMove(1)
	pkmn1 := GetPokemon(4)
	pkmn1.Stats = [6]uint{30, 10, 10, 10, 10, 10}
	pkmn1.CurrentHP = 30
	pkmn1.Moves[0] = &pound
	party1.AddPokemon(&pkmn1)
	party2 := Party{
		Agent: &a2,
	}
	pkmn2 := GetPokemon(7)
	pkmn2.Stats = [6]uint{30, 10, 10, 10, 10, 10}
	pkmn2.CurrentHP = 30
	pkmn2.Moves[0] = &pound
	party2.AddPokemon(&pkmn2)
	b.AddParty(&party1, &party2)
	b.SetTeams([][]int{{0}, {1}})

	b.Start()
	for p, party := range b.Parties {
		got := party.GetActive()
		if !reflect.DeepEqual(got, []int{0}) {
			t.Fatalf("Must send out first pokemon in each at the beginning of the battle. Party %d gave: %v", p, got)
		}
	}
	turns := b.SimulateRound()
	if len(turns) != 2 {
		t.Fatal("Expected only 2 turns to occur in a round")
	}
	logtest := []struct {
		turn Turn
		want string
	}{
		{
			turn: turns[0],
			want: "Charmander used Pound on Squirtle for 3 damage.",
		},
		{
			turn: turns[1],
			want: "Squirtle used Pound on Charmander for 3 damage.",
		},
	}
	for _, tt := range logtest {
		got := tt.turn.BattleLog()
		if got != tt.want {
			t.Errorf("Expected battle log to be %s, got %s", tt.want, got)
		}
	}
	// functionally arbitrary value, will need to be adjusted when damage calculation becomes more accurate
	expectedHp := uint(27)
	for _, party := range b.Parties {
		if party.Pokemon[0].CurrentHP != expectedHp {
			t.Errorf("Expected %s to have %d HP, got %d", party.Pokemon[0].GetName(), expectedHp, party.Pokemon[0].CurrentHP)
		}
	}
}

// Faster pokemon should go first.
func TestPokemonSpeed(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	b := Battle{}
	party1 := Party{
		Agent: &a1,
	}
	pound := GetMove(1)
	pkmn1 := GetPokemon(4)
	pkmn1.Moves[0] = &pound
	pkmn1.Stats = [6]uint{30, 10, 10, 10, 10, 10}
	party1.AddPokemon(&pkmn1)
	party2 := Party{
		Agent: &a2,
	}
	pkmn2 := GetPokemon(4)
	pkmn2.Moves[0] = &pound
	pkmn2.Stats = [6]uint{30, 10, 10, 10, 10, 12}
	party2.AddPokemon(&pkmn2)
	b.AddParty(&party1, &party2)
	b.SetTeams([][]int{{0}, {1}})

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
