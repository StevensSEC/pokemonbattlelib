package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/StevensSEC/pokemonbattlelib"
)

const defaultPort = 4000

func main() {
	var port int
	if p := os.Getenv("PORT"); len(p) > 0 {
		var err error
		port, err = strconv.Atoi(p)
		if err != nil {
			log.Printf("Invalid port number: %s using default", p)
			port = defaultPort
		}
	} else {
		port = defaultPort
	}
	log.Printf("Listening on port %d\n", port)
	bl := log.New(os.Stdout, "[battle] ", log.Lshortfile)
	pokemonbattlelib.SetLogger(bl)
	router := buildRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
