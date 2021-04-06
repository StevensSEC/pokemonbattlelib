package pokemonbattlelib

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transaction Marshalling", func() {
	Specify("all Transactions must marshall to have different values for \"type\"", func() {
		transactions := []Transaction{
			DamageTransaction{},
			FriendshipTransaction{},
			EVTransaction{},
			ItemTransaction{},
			PPTransaction{},
			HealTransaction{},
			InflictStatusTransaction{},
			CureStatusTransaction{},
			FaintTransaction{},
			SendOutTransaction{},
			WeatherTransaction{},
			EndBattleTransaction{},
			ImmobilizeTransaction{},
			MoveFailTransaction{},
			ModifyStatTransaction{},
		}
		seen := []int{}
		for _, t := range transactions {
			bytes, _ := json.Marshal(t)
			var c struct {
				Type int `json:"type"`
			}
			json.Unmarshal(bytes, &c)

			Expect(seen).ToNot(ContainElement(c.Type))
			seen = append(seen, c.Type)
		}
	})
})
