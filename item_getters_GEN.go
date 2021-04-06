// Code generated - DO NOT EDIT.
// Regenerate with `go generate`.

package pokemonbattlelib

import "encoding/json"

func (i Item) Name() string             { return i.Data().Name }
func (i Item) Category() ItemCategory   { return i.Data().Category }
func (i Item) FlingPower() int          { return i.Data().FlingPower }
func (i Item) FlingEffect() FlingEffect { return i.Data().FlingEffect }
func (i Item) Flags() ItemFlags         { return i.Data().Flags }

func (i Item) MarshalJSON() ([]byte, error) {
	type alias Item
	return json.Marshal(&struct {
		Id          alias
		Name        string
		Category    ItemCategory
		FlingPower  int
		FlingEffect FlingEffect
		Flags       ItemFlags
	}{
		Id:          alias(i),
		Name:        i.Name(),
		Category:    i.Category(),
		FlingPower:  i.FlingPower(),
		FlingEffect: i.FlingEffect(),
		Flags:       i.Flags(),
	})
}

func (i *Item) UnmarshalJSON(data []byte) error {
	type alias Item
	type dataalias struct{ Id uint16 }
	if data[0] == '{' {
		var aux dataalias
		if err := json.Unmarshal(data, &aux); err != nil {
			return err
		}
		*i = Item(aux.Id)
		return nil
	} else {
		var aux alias
		if err := json.Unmarshal(data, &aux); err != nil {
			return err
		}
		*i = Item(aux)
		return nil
	}
}
