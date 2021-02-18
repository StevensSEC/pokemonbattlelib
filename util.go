package pokemonbattlelib

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func parseInt(s string) (n int) {
	if s == "" {
		return 0
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Panicln(err)
	}
	return n
}

func getCsvReader(path string) *csv.Reader {
	fmt.Printf("Reading csv: %s\n", path)
	file, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}
	reader := csv.NewReader(file)
	reader.ReuseRecord = true
	_, err = reader.Read() // skip header line
	if err != nil {
		log.Panicln(err)
	}
	return reader
}
