package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/StevensSEC/pokemonbattlelib"
)

func main() {
	port := 4000
	log.Printf("Listening on port %d\n", port)
	bl := log.New(os.Stdout, "battle", log.Ldate|log.Lshortfile)
	pokemonbattlelib.SetLogger(bl)
	router := buildRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
