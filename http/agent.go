package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	. "github.com/StevensSEC/pokemonbattlelib"
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
	switch hT.Type {
	case 0:
		var t FightTurn
		// HACK: required to unmarshal private fields
		bytes, err := json.Marshal(hT.Args)
		if err != nil {
			log.Panicf("Failed to decode turn arguments: %s", err)
		}
		err = json.Unmarshal(bytes, &t)
		if err != nil {
			log.Panicf("Failed to decode turn arguments: %s", err)
		}
		turn = Turn(t)
	case 1:
		var t ItemTurn
		// HACK: required to unmarshal private fields
		bytes, err := json.Marshal(hT.Args)
		if err != nil {
			log.Panicf("Failed to decode turn arguments: %s", err)
		}
		err = json.Unmarshal(bytes, &t)
		if err != nil {
			log.Panicf("Failed to decode turn arguments: %s", err)
		}
		turn = Turn(t)
	default:
		log.Panic("unknown turn number")
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
