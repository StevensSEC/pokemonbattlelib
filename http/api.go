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
	return mux
}

// List all available moves.
func HandleListMoves(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(ALL_MOVES)
	if err != nil {
		log.Fatalf("Failed to marshal moves into JSON: %s", err)
	}
	w.WriteHeader(200)
	w.Write(bytes)
}

type httpPokemon struct {
	Pokemon
	Name string
}

// Convert a Pokemon into an httpPokemon, suitable for JSON marshalling for HTTP responses that contain pokemon.
func convertPokemon(p *Pokemon) *httpPokemon {
	return &httpPokemon{
		Pokemon: *p,
		Name:    p.GetName(),
	}
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

	bytes, err := json.Marshal(convertPokemon(pkmn))
	if err != nil {
		log.Fatalf("Failed to marshal generated pokemon into JSON: %s", err)
	}

	w.WriteHeader(200)
	w.Write(bytes)
}
