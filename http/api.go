package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"sync"

	. "github.com/StevensSEC/pokemonbattlelib"
)

var battleLock sync.Mutex

func buildRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/pokedex/list/moves", HandleListMoves)
	mux.HandleFunc("/pokedex/generate", HandleGeneratePokemon)
	mux.HandleFunc("/battle/new", HandleCreateBattle)
	mux.HandleFunc("/battle/simulate", HandleBattleSimulate)
	mux.HandleFunc("/battle/context", HandleBattleContext)
	mux.HandleFunc("/battle/act", HandleBattleAct)
	mux.HandleFunc("/battle/results", HandleBattleResults)
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
		var moves []MoveId
		for _, id := range moveIds {
			moves = append(moves, MoveId(id))
		}
		args.Opts = append(args.Opts, WithMoves(moves...))
	} else {
		args.Opts = append(args.Opts, WithMoves(
			MoveId(rand.Intn(len(AllMoves)))+1,
			MoveId(rand.Intn(len(AllMoves)))+1,
			MoveId(rand.Intn(len(AllMoves)))+1,
			MoveId(rand.Intn(len(AllMoves)))+1))
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
	queuedTurns map[int]Turn
	turnLock    sync.Mutex
}

func newHttpBattle(b *Battle) *httpBattle {
	return &httpBattle{
		Battle:      NewBattle(),
		queuedTurns: make(map[int]Turn),
	}
}

func (hb *httpBattle) QueueNextTurn(targetId int, turn Turn) {
	hb.turnLock.Lock()
	defer hb.turnLock.Unlock()
	hb.queuedTurns[targetId] = turn
}

func (hb *httpBattle) FlushTurns() {
	for i, t := range hb.Battle.AllTargets() {
		p := hb.Battle.GetParty(t)
		switch a := (*p.Agent).(type) {
		case WaiterAgent:
			turn := hb.queuedTurns[i]
			c := a.Input()
			*c <- turn
			delete(hb.queuedTurns, i)
		}
	}
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

		hb := newHttpBattle(NewBattle())
		if len(args.Parties) > 0 {
			// deprecated
			for i := range args.Parties {
				// a := Agent(NewHttpCallbackAgent(args.CallbackUrls[i]))
				wa := NewWaiterAgent()
				a := Agent(wa)
				hb.AgentInputs = append(hb.AgentInputs, wa.Input())
				hb.Battle.AddParty(NewOccupiedParty(args.Parties[i]...), &a, i)
			}
		} else if len(args.Teams) > 0 {
			for t, team := range args.Teams {
				for _, party := range team.Parties {
					wa := NewWaiterAgent()
					a := Agent(wa)
					hb.AgentInputs = append(hb.AgentInputs, wa.Input())
					hb.Battle.AddParty(NewOccupiedParty(party.Pokemon...), &a, t)
				}
			}
		} else {
			w.WriteHeader(400)
			w.Write([]byte("Bad request: invalid body"))
		}
		battleLock.Lock()
		battles[nextBattleId] = hb
		battleLock.Unlock() // unlock immediately instead of defering

		err = hb.Battle.Start()
		if err != nil {
			log.Printf("Failed to start battle: %s", err)
			w.WriteHeader(500)
			w.Write([]byte("Internal server error: Failed to start battle"))
			return
		}

		log.Printf("Battle created: %v", hb.Battle)
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
	hb := battles[battleId]
	hb.FlushTurns()
	transactions, ended := hb.Battle.SimulateRound()

	type roundResults struct {
		Transactions []Transaction
		Ended        bool
		Parties      []*Party
	}

	results := roundResults{
		transactions,
		ended,
		hb.Battle.Parties(),
	}

	JSONResponse(w, results)
}

func HandleBattleContext(w http.ResponseWriter, r *http.Request) {
	battleId := parseNumberArg(r, "id")
	// partyId := parseNumberArg(r, "party")
	// slot := parseNumberArg(r, "slot")
	targetIdx := parseNumberArg(r, "target")

	b := battles[battleId].Battle
	target := b.AllTargets()[targetIdx]

	bytes, err := json.Marshal(b.GetBattleContext(target))
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
	targetIdx := parseNumberArg(r, "target")

	var hT HttpTurn
	err := json.NewDecoder(r.Body).Decode(&hT)
	if err != nil {
		respondSuccess(w, false)
		log.Printf("Failed to decode turn: %s", err)
		return
	}

	battles[battleId].QueueNextTurn(targetIdx, hT.GetTurn())

	w.WriteHeader(200)
	respondSuccess(w, true)
}

func HandleBattleResults(w http.ResponseWriter, r *http.Request) {
	battleId := parseNumberArg(r, "id")
	hb := battles[battleId]

	results := hb.Battle.GetResults()

	bytes, err := json.Marshal(results)
	if err != nil {
		log.Printf("Failed to marshal into JSON: %s", err)
		w.WriteHeader(500)
		w.Write([]byte(`Internal server error: failed to marshal battle results`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(bytes)
}

func HandleDumbAgent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"type": 0, "args": {"target": {"party":0, "slot": 0}, "move": 0}}`))
}

func respondSuccess(w http.ResponseWriter, success bool) {
	data, err := json.Marshal(responseSuccess{true})
	if err != nil {
		log.Panicf("Failed to marshal response.")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

type responseSuccess struct {
	Success bool `json:"success"`
}
