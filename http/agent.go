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
		log.Panicf("Failed to marshal into JSON: %s", err)
	}
	r := bytes.NewReader(ctxBytes)
	log.Println("Requesting turn from agent callback")
	resp, err := http.Post(a.CallbackUrl, "application/json", r)
	if err != nil {
		log.Panicf("Failed to post battle context to callback URL: %s", err)
	}

	var hT HttpTurn
	err = json.NewDecoder(resp.Body).Decode(&hT)
	if err != nil {
		log.Panicf("Failed to decode response: %s", err)
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
		log.Panic("unknown turn number")
	}
	if err != nil {
		log.Panicf("Failed to decode map into turn struct: %s", err)
	}
	return turn
}

// Blocks until it receives a turn on a channel.
type WaiterAgent struct {
	in chan Turn
}

func NewWaiterAgent() WaiterAgent {
	return WaiterAgent{
		in: make(chan Turn, 1),
	}
}

func (a *WaiterAgent) Input() *chan Turn {
	return &a.in
}

func (a WaiterAgent) Act(ctx *BattleContext) Turn {
	return <-a.in
}
