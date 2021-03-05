package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 4000
	log.Printf("Listening on port %d\n", port)
	router := buildRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 4000), router))
}
