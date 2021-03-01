package main

import (
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func randValues(max int) [6]uint8 {
	return [6]uint8{
		uint8(rand.Intn(max)),
		uint8(rand.Intn(max)),
		uint8(rand.Intn(max)),
		uint8(rand.Intn(max)),
		uint8(rand.Intn(max)),
		uint8(rand.Intn(max)),
	}
}

// helper to parse number args from http requests
func parseNumberArg(r *http.Request, param string) int {
	n, err := strconv.Atoi(r.FormValue(param))
	if err != nil {
		log.Fatalf("Invalid number for %s: %s", param, err)
	}
	return n
}

// takes a comma separated string of numbers, and turns it into a slice of numbers
func parseIntSlice(s string) ([]int, error) {
	var nums []int
	for _, sn := range strings.Split(s, ",") {
		n, err := strconv.Atoi(sn)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}
