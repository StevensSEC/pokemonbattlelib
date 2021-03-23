package main

import (
	"encoding/json"
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
	w.Header().Set("Content-Type", "application/json")
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
			moves = append(moves, GetMove(MoveId(id)))
		}
		args.Opts = append(args.Opts, WithMoves(moves...))
	} else {
		args.Opts = append(args.Opts, WithMoves(
			GetMove(MoveId(rand.Intn(len(AllMoves)))+1),
			GetMove(MoveId(rand.Intn(len(AllMoves)))+1),
			GetMove(MoveId(rand.Intn(len(AllMoves)))+1),
			GetMove(MoveId(rand.Intn(len(AllMoves)))+1),
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bytes)
}

var nextBattleId int
var battles = map[int]*httpBattle{}

type newBattleArgs struct {
	Teams []struct {
		Parties []struct {
			Pokemon []*Pokemon `json:"pokemon"`
		} `json:"parties"`
	} `json:"teams"`
	Parties [][]*Pokemon `json:"parties"` // deprecated
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
		if len(args.Parties) > 0 {
			// deprecated
			for i := range args.Parties {
				// a := Agent(NewHttpCallbackAgent(args.CallbackUrls[i]))
				wa := NewWaiterAgent()
				a := Agent(wa)
				hb.AgentInputs = append(hb.AgentInputs, wa.Input())
				p := NewParty(&a, i)
				p.AddPokemon(args.Parties[i]...)
				hb.Battle.AddParty(p)
			}
		} else if len(args.Teams) > 0 {
			for t, team := range args.Teams {
				for _, party := range team.Parties {
					wa := NewWaiterAgent()
					a := Agent(wa)
					hb.AgentInputs = append(hb.AgentInputs, wa.Input())
					p := NewParty(&a, t)
					p.AddPokemon(party.Pokemon...)
					hb.Battle.AddParty(p)
				}
			}
		} else {
			w.WriteHeader(400)
			w.Write([]byte("Bad request: invalid body"))
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
		data, err := json.Marshal(struct {
			BattleId      int
			ActivePokemon int
		}{
			BattleId:      nextBattleId,
			ActivePokemon: 2, // TODO: change for double battles
		})
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Internal server error: Failed to marshal response"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
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

	w.Header().Set("Content-Type", "application/json")
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

	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"type": 0, "args": {"target": {"party":0, "slot": 0}, "move": 0}}`))
}
