package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/StevensSEC/pokemonbattlelib"
)

var seedFlag = flag.Int64("seed", 0, "Set the seed of the random number generator. 0 means the seed is not specified, so the seed will be set to something effectively unpredictable.")

const defaultPort = 4000

func main() {
	flag.Parse()

	var seed int64
	if *seedFlag != 0 {
		seed = *seedFlag
	} else {
		seed = time.Now().UnixNano()
	}
	log.Printf("Seed: %d", seed)
	rand.Seed(seed)

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
