package main

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/StevensSEC/pokemonbattlelib"
)

func buildRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/pokedex/list/moves", HandleListMoves)
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
