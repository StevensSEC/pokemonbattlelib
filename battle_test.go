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
	pkmn1 := GetPokemon(4)
	party1.AddPokemon(&pkmn1)
	party2 := Party{
		Agent: &a2,
	}
	pkmn2 := GetPokemon(7)
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
	b.SimulateRound()
	// output:
	// TODO: Implement fight {0 1}
	// TODO: Implement fight {0 0}
}

// Faster pokemon should go first.
func TestPokemonSpeed(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
    b := Battle{}
	party1 := Party{
		Agent: &a1,
	}
	pkmn1 := GetPokemon(4)
    pkmn1.setSpeed(10)
	party1.AddPokemon(&pkmn1)
	party2 := Party{
		Agent: &a2,
	}
	pkmn2 := GetPokemon(7)
    pkmn2.setSpeed(12)
	party2.AddPokemon(&pkmn2)
	b.AddParty(&party1, &party2)
	b.SetTeams([][]int{{0}, {1}})

	b.Start()
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
