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
	for _, target := range ctx.Opponents {
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
	pkmn1 := GetPokemon(4)
	party1.AddPokemon(&pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := GetPokemon(7)
	party2.AddPokemon(&pkmn2)
	b := NewBattle()
	b.AddParty(party1, party2)
}

func TestBattleOneRound(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := GetPokemon(4)
	pkmn1.Stats = [6]uint{30, 10, 10, 10, 10, 10}
	pkmn1.CurrentHP = 30
	pound := GetMove(1)
	pkmn1.Moves[0] = &pound
	party1.AddPokemon(&pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := GetPokemon(7)
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
	transactions, _ := b.SimulateRound()
	if len(transactions) != 2 {
		t.Fatal("Expected only 2 transactions to occur in a round")
	}
	logtest := []struct {
		turn Transaction
		want string
	}{
		{
			turn: transactions[0],
			want: "Charmander used Pound on Squirtle for 3 damage.",
		},
		{
			turn: transactions[1],
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
	for _, party := range b.parties {
		if party.pokemon[0].CurrentHP != expectedHp {
			t.Errorf("Expected %s to have %d HP, got %d", party.pokemon[0].GetName(), expectedHp, party.pokemon[0].CurrentHP)
		}
	}
}

// Tests if active Pokemon are set correctly
func TestActivePokemon(t *testing.T) {
	a := Agent(dumbAgent{})
	party := NewParty(&a, 0)
	pkmn1 := GetPokemon(7)
	pkmn2 := GetPokemon(9)
	party.AddPokemon(&pkmn1, &pkmn2)
	party.SetActive(0)
	if len(party.activePokemon) != 1 {
		t.Error("expected party to have 1 active Pokemon")
	}
	party.SetInactive(0)
	if len(party.activePokemon) != 0 {
		t.Error("expected party to have no active Pokemon")
	}
	party.SetActive(1)
	if n := party.activePokemon[1].NatDex; n != 9 {
		t.Errorf("expected party to have an active Pokemon with dex number 9, received %v\n", n)
	}
}

func TestGetPartyPokemon(t *testing.T) {
	a1 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := GetPokemon(4)
	pkmn2 := GetPokemon(7)
	pkmn3 := GetPokemon(11)
	party1.AddPokemon(&pkmn1, &pkmn2, &pkmn3)
	b := NewBattle()
	b.AddParty(party1)
	p := b.getPokemon(0, 1)
	if p.NatDex != 7 {
		t.Errorf("expected Pokemon with dex number 7, received %v\n", p.NatDex)
	}
}

func TestGetAllies(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := GetPokemon(4)
	party1.AddPokemon(&pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := GetPokemon(7)
	pkmn3 := GetPokemon(9)
	party2.AddPokemon(&pkmn2, &pkmn3)
	b := NewBattle()
	b.AddParty(party1, party2)
	err := b.Start()
	if err != nil {
		t.Fatal("failed to start battle")
	}

	for _, party := range []*party{party1, party2} {
		allies := b.GetAllies(party)
		if n := len(allies); n != 1 {
			t.Errorf("expected party to have 1 active ally, received %v\n", n)
		}
		for _, ally := range allies {
			if party.team != ally.Team {
				t.Errorf("expected party allies to match team, but %d != %d", party.team, ally.Team)
			}
		}
	}
}

func TestGetOpponents(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := GetPokemon(4)
	party1.AddPokemon(&pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := GetPokemon(7)
	pkmn3 := GetPokemon(9)
	party2.AddPokemon(&pkmn2, &pkmn3)
	b := NewBattle()
	b.AddParty(party1, party2)
	err := b.Start()
	if err != nil {
		t.Fatal("failed to start battle")
	}
	for _, party := range []*party{party1, party2} {
		opponents := b.GetOpponents(party)
		if n := len(opponents); n != 1 {
			t.Errorf("expected party to have 1 active opponent, received %v\n", n)
		}
		for _, opponent := range opponents {
			if party.team == opponent.Team {
				t.Errorf("expected party opponents to not match team, but %d == %d", party.team, opponent.Team)
			}
		}
	}
}

// Faster pokemon should go first.
func TestPokemonSpeed(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := GetPokemon(4)
	pound := GetMove(1)
	pkmn1.Moves[0] = &pound
	pkmn1.Stats = [6]uint{30, 10, 10, 10, 10, 10}
	pkmn1.CurrentHP = 30
	party1.AddPokemon(&pkmn1)
	party2 := NewParty(&a2, 1)
	pkmn2 := GetPokemon(7)
	pkmn2.Moves[0] = &pound
	pkmn2.Stats = [6]uint{30, 10, 10, 10, 10, 12}
	pkmn2.CurrentHP = 30
	party2.AddPokemon(&pkmn2)
	b := NewBattle()
	b.AddParty(party1, party2)
	err := b.Start()
	if err != nil {
		t.Fatal("failed to start battle")
	}

	transactions, _ := b.SimulateRound()
	if len(transactions) != 2 {
		t.Fatal("Expected only 2 transactions to occur in a round")
	}
	logtest := []struct {
		turn Transaction
		want string
	}{
		{
			turn: transactions[0],
			want: "Squirtle used Pound on Charmander for 3 damage.",
		},
		{
			turn: transactions[1],
			want: "Charmander used Pound on Squirtle for 3 damage.",
		},
	}
	for _, tt := range logtest {
		got := tt.turn.BattleLog()
		if got != tt.want {
			t.Errorf("Expected battle log to be %s, got %s", tt.want, got)
		}
	}
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

func TestBattleFaintAutoSwitch(t *testing.T) {
	a1 := Agent(dumbAgent{})
	a2 := Agent(dumbAgent{})
	party1 := NewParty(&a1, 0)
	pkmn1 := GetPokemon(4)
	pkmn1.Stats = [6]uint{30, 1, 1, 10, 10, 200}
	pkmn1.CurrentHP = 1
	pkmn3 := GetPokemon(387)
	pkmn3.Stats = [6]uint{30, 10, 10, 10, 10, 10}
	pkmn3.CurrentHP = 30
	pound := GetMove(1)
	pkmn1.Moves[0] = &pound
	pkmn3.Moves[0] = &pound
	party1.AddPokemon(&pkmn1, &pkmn3)
	party2 := NewParty(&a2, 1)
	pkmn2 := GetPokemon(7)
	pkmn2.Stats = [6]uint{30, 800, 800, 10, 10, 10}
	pkmn2.CurrentHP = 30
	pkmn2.Moves[0] = &pound
	party2.AddPokemon(&pkmn2)
	b := NewBattle()
	b.AddParty(party1, party2)
	err := b.Start()
	if err != nil {
		t.Fatal("failed to start battle")
	}

	transactions, _ := b.SimulateRound()
	if n := len(transactions); n != 4 {
		t.Errorf("Expected 4 transactions to occur, got %d", n)
	}
	logtest := []struct {
		turn Transaction
		want string
	}{
		{
			turn: transactions[0],
			want: "Charmander used Pound on Squirtle for 2 damage.",
			// Charmander smashed his nubby little fist into Squirtle as
			// hard as he could. Spectators gasped and winced when the
			// impact created a very audible crack. But it was not
			// Squirtle's shell that broke, it was Charmanders knuckles.
			// The Squirtle was unfazed.
		},
		{
			turn: transactions[1],
			want: "Squirtle used Pound on Charmander for 1282 damage.",
			// Ash watched in horror as his Charmander was obliterated from the
			// battlefield. "Critical hit!" echoed the automated announcer. The
			// Squirtle snarled, now covered in the entrails of his previous
			// opponent. "OH GOD, WHAT THE FUCK!?" sobbed Ash, "Is my friend
			// really gone forever? Please tell me I'm dreaming, this can't be real!"
		},
		{
			turn: transactions[2],
			want: "Charmander fainted.",
		},
		{
			turn: transactions[3],
			want: "Turtwig was sent out.",
		},
	}
	for _, tt := range logtest {
		got := tt.turn.BattleLog()
		if got != tt.want {
			t.Errorf("Expected battle log to be %s, got %s", tt.want, got)
		}
	}
}
