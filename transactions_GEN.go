// Code generated - DO NOT EDIT.
// Regenerate with `go generate`.

package pokemonbattlelib

import "encoding/json"

func (t UseMoveTransaction) MarshalJSON() ([]byte, error) {
	type alias UseMoveTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 0,
		Name: "UseMoveTransaction",
		Args: (*alias)(&t),
	})
}

func (t DamageTransaction) MarshalJSON() ([]byte, error) {
	type alias DamageTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 1,
		Name: "DamageTransaction",
		Args: (*alias)(&t),
	})
}

func (t FriendshipTransaction) MarshalJSON() ([]byte, error) {
	type alias FriendshipTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 2,
		Name: "FriendshipTransaction",
		Args: (*alias)(&t),
	})
}

func (t EVTransaction) MarshalJSON() ([]byte, error) {
	type alias EVTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 3,
		Name: "EVTransaction",
		Args: (*alias)(&t),
	})
}

func (t ItemTransaction) MarshalJSON() ([]byte, error) {
	type alias ItemTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 4,
		Name: "ItemTransaction",
		Args: (*alias)(&t),
	})
}

func (t PPTransaction) MarshalJSON() ([]byte, error) {
	type alias PPTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 5,
		Name: "PPTransaction",
		Args: (*alias)(&t),
	})
}

func (t GiveItemTransaction) MarshalJSON() ([]byte, error) {
	type alias GiveItemTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 6,
		Name: "GiveItemTransaction",
		Args: (*alias)(&t),
	})
}

func (t HealTransaction) MarshalJSON() ([]byte, error) {
	type alias HealTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 7,
		Name: "HealTransaction",
		Args: (*alias)(&t),
	})
}

func (t InflictStatusTransaction) MarshalJSON() ([]byte, error) {
	type alias InflictStatusTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 8,
		Name: "InflictStatusTransaction",
		Args: (*alias)(&t),
	})
}

func (t CureStatusTransaction) MarshalJSON() ([]byte, error) {
	type alias CureStatusTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 9,
		Name: "CureStatusTransaction",
		Args: (*alias)(&t),
	})
}

func (t FaintTransaction) MarshalJSON() ([]byte, error) {
	type alias FaintTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 10,
		Name: "FaintTransaction",
		Args: (*alias)(&t),
	})
}

func (t SendOutTransaction) MarshalJSON() ([]byte, error) {
	type alias SendOutTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 11,
		Name: "SendOutTransaction",
		Args: (*alias)(&t),
	})
}

func (t WeatherTransaction) MarshalJSON() ([]byte, error) {
	type alias WeatherTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 12,
		Name: "WeatherTransaction",
		Args: (*alias)(&t),
	})
}

func (t EndBattleTransaction) MarshalJSON() ([]byte, error) {
	type alias EndBattleTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 13,
		Name: "EndBattleTransaction",
		Args: (*alias)(&t),
	})
}

func (t ImmobilizeTransaction) MarshalJSON() ([]byte, error) {
	type alias ImmobilizeTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 14,
		Name: "ImmobilizeTransaction",
		Args: (*alias)(&t),
	})
}

func (t MoveFailTransaction) MarshalJSON() ([]byte, error) {
	type alias MoveFailTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 15,
		Name: "MoveFailTransaction",
		Args: (*alias)(&t),
	})
}

func (t ModifyStatTransaction) MarshalJSON() ([]byte, error) {
	type alias ModifyStatTransaction
	return json.Marshal(&struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Args *alias `json:"args"`
	}{
		Type: 16,
		Name: "ModifyStatTransaction",
		Args: (*alias)(&t),
	})
}
