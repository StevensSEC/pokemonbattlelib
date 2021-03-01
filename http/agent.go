package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	. "github.com/StevensSEC/pokemonbattlelib"
	"github.com/mitchellh/mapstructure"
)

type HttpCallbackAgent struct {
	CallbackUrl string
}

func (a HttpCallbackAgent) Act(ctx *BattleContext) Turn {
	ctxBytes, err := json.Marshal(ctx)
	if err != nil {
		log.Fatalf("Failed to marshal into JSON: %s", err)
	}
	r := bytes.NewReader(ctxBytes)
	resp, err := http.Post(a.CallbackUrl, "application/json", r)

	var hT HttpTurn
	err = json.NewDecoder(resp.Body).Decode(&hT)
	if err != nil {
		log.Fatalf("Failed to decode response: %s", err)
	}

	return hT.GetTurn()
}

func NewHttpCallbackAgent(url string) HttpCallbackAgent {
	return HttpCallbackAgent{
		CallbackUrl: url,
	}
}

type HttpTurn struct {
	Type int         `json:"type"`
	Args interface{} `json:"args"`
}

func (hT *HttpTurn) GetTurn() Turn {
	var turn Turn
	var err error
	switch hT.Type {
	case 0:
		var t FightTurn
		err = mapstructure.Decode(hT.Args, &t)
		turn = Turn(t)
	case 1:
		var t ItemTurn
		err = mapstructure.Decode(hT.Args, &t)
		turn = Turn(t)
	default:
		panic("unknown turn number")
	}
	if err != nil {
		panic(err)
	}
	return turn
}
