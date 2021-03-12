package pokemonbattlelib

import (
	"io/ioutil"
	"log"
)

// internal logger
var blog *log.Logger

func init() {
	blog = log.New(ioutil.Discard, "", 0)
}

// Set the internal logger.
func SetLogger(l *log.Logger) {
	blog = l
}

// Get the internal logger.
func GetLogger() *log.Logger {
	return blog
}
