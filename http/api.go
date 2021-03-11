package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	. "github.com/StevensSEC/pokemonbattlelib"
)

func buildRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/pokedex/list/moves", HandleListMoves)
	mux.HandleFunc("/pokedex/generate", HandleGeneratePokemon)
	mux.HandleFunc("/battle/new", HandleCreateBattle)
	mux.HandleFunc("/battle/simulate", HandleBattleSimulate)
	mux.HandleFunc("/battle/context", HandleBattleContext)
	mux.HandleFunc("/battle/act", HandleBattleAct)
	mux.HandleFunc("/agent/dumb", HandleDumbAgent)
	return mux
}

// List all available moves.
func HandleListMoves(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(AllMoves)
	if err != nil {
		log.Fatalf("Failed to marshal moves into JSON: %s", err)
	}
	w.WriteHeader(200)
	w.Write(bytes)
}

type pokemonGenerateArgs struct {
	NatDex int
	Opts   []GeneratePokemonOption
}

func (args pokemonGenerateArgs) Generate() *Pokemon {
	return GeneratePokemon(
		args.NatDex,
		args.Opts...,
	)
}

// Generate a pokemon.
func HandleGeneratePokemon(w http.ResponseWriter, r *http.Request) {
	args := pokemonGenerateArgs{}

	if len(r.FormValue("natdex")) != 0 {
		args.NatDex = parseNumberArg(r, "natdex")
	} else {
		args.NatDex = 1 + rand.Intn(493)
	}

	if len(r.FormValue("level")) != 0 {
		args.Opts = append(args.Opts, WithLevel(uint8(parseNumberArg(r, "level"))))
	} else {
		args.Opts = append(args.Opts, WithLevel(uint8(1+rand.Intn(99))))
	}

	if len(r.FormValue("moves")) != 0 {
		moveIds, err := parseIntSlice(r.FormValue("moves"))
		if err != nil {
			log.Fatalf("Invalid move ids %s: %s", r.FormValue("moves"), err)
		}
		var moves []*Move
		for _, id := range moveIds {
			moves = append(moves, GetMove(id))
		}
		args.Opts = append(args.Opts, WithMoves(moves...))
	} else {
		args.Opts = append(args.Opts, WithMoves(
			GetMove(1+rand.Intn(467)),
			GetMove(1+rand.Intn(467)),
			GetMove(1+rand.Intn(467)),
			GetMove(1+rand.Intn(467)),
		))
	}

	args.Opts = append(args.Opts,
		WithIVs(randValues(6)),
		WithEVs(randValues(42)),
	)

	pkmn := args.Generate()

	bytes, err := json.Marshal(pkmn)
	if err != nil {
		log.Fatalf("Failed to marshal generated pokemon into JSON: %s", err)
	}

	w.WriteHeader(200)
	w.Write(bytes)
}

var nextBattleId int
var battles = map[int]*httpBattle{}

type newBattleArgs struct {
	Parties      [][]*Pokemon
	CallbackUrls []string
}

type httpBattle struct {
	Battle      *Battle
	AgentInputs []*chan Turn
}

func HandleCreateBattle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var args newBattleArgs
		err := json.NewDecoder(r.Body).Decode(&args)
		if err != nil {
			log.Printf("Failed to decode newBattleArgs: %s", err)
			w.WriteHeader(400)
			w.Write([]byte("Bad request: Failed to parse arguments"))
			return
		}

		hb := httpBattle{
			Battle: NewBattle(),
		}
		for i := range args.Parties {
			// a := Agent(NewHttpCallbackAgent(args.CallbackUrls[i]))
			wa := NewWaiterAgent()
			a := Agent(wa)
			hb.AgentInputs = append(hb.AgentInputs, wa.Input())
			p := NewParty(&a, i)
			p.AddPokemon(args.Parties[i]...)
			hb.Battle.AddParty(p)
		}
		battles[nextBattleId] = &hb

		err = hb.Battle.Start()
		if err != nil {
			log.Printf("Failed to start battle: %s", err)
			w.WriteHeader(500)
			w.Write([]byte("Internal server error: Failed to start battle"))
			return
		}

		log.Printf("Battle created: %v", hb.Battle)
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("%d", nextBattleId)))
		nextBattleId++
	} else {
		w.WriteHeader(400)
		w.Write([]byte("Bad request: Wrong method"))
	}
}

func HandleBattleSimulate(w http.ResponseWriter, r *http.Request) {
	battleId := parseNumberArg(r, "id")
	log.Printf("Simulating round: id %d", battleId)
	b := battles[battleId].Battle
	transactions, ended := b.SimulateRound()

	type roundResults struct {
		Transactions []Transaction
		Ended        bool
	}

	results := roundResults{
		transactions,
		ended,
	}

	bytes, err := json.Marshal(results)
	if err != nil {
		log.Printf("Failed to marshal into JSON: %s", err)
		w.WriteHeader(500)
		w.Write([]byte(`Internal server error: failed to marshal round results`))
		return
	}

	w.WriteHeader(200)
	w.Write(bytes)
}

func HandleBattleContext(w http.ResponseWriter, r *http.Request) {
	battleId := parseNumberArg(r, "id")
	partyId := parseNumberArg(r, "party")
	slot := parseNumberArg(r, "slot")

	bytes, err := json.Marshal(battles[battleId].Battle.GetRoundContext(partyId, slot))
	if err != nil {
		log.Printf("Failed to marshal into JSON: %s", err)
		w.WriteHeader(500)
		w.Write([]byte(`Internal server error: failed to marshal battle context`))
		return
	}

	w.WriteHeader(200)
	w.Write(bytes)
}

func HandleBattleAct(w http.ResponseWriter, r *http.Request) {
	battleId := parseNumberArg(r, "id")
	agentId := parseNumberArg(r, "agent")

	var hT HttpTurn
	err := json.NewDecoder(r.Body).Decode(&hT)
	if err != nil {
		log.Panicf("Failed to decode turn: %s", err)
	}

	*battles[battleId].AgentInputs[agentId] <- hT.GetTurn()

	w.WriteHeader(200)
	w.Write([]byte("Turn has been queued."))
}

func HandleDumbAgent(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"type": 0, "args": {"target": {"party":0, "slot": 0}, "move": 0}}`))
}
