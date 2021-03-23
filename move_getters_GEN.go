// Code generated - DO NOT EDIT.
// Regenerate with `go generate`.

package pokemonbattlelib

import "encoding/json"

func (m Move) Name() string           { return m.Data().Name }
func (m Move) Type() Type             { return m.Data().Type }
func (m Move) Category() MoveCategory { return m.Data().Category }
func (m Move) Targets() MoveTarget    { return m.Data().Targets }
func (m Move) Priority() int8         { return m.Data().Priority }
func (m Move) Power() uint            { return m.Data().Power }
func (m Move) Accuracy() uint         { return m.Data().Accuracy }
func (m Move) InitialMaxPP() uint8    { return m.Data().InitialMaxPP }

func (m *Move) MarshalJSON() ([]byte, error) {
	type alias Move
	return json.Marshal(&struct {
		Name         string
		Type         Type
		Category     MoveCategory
		Targets      MoveTarget
		Priority     int8
		Power        uint
		Accuracy     uint
		InitialMaxPP uint8
		*alias
	}{
		Name:         m.Name(),
		Type:         m.Type(),
		Category:     m.Category(),
		Targets:      m.Targets(),
		Priority:     m.Priority(),
		Power:        m.Power(),
		Accuracy:     m.Accuracy(),
		InitialMaxPP: m.InitialMaxPP(),
		alias:        (*alias)(m),
	})
}

func (m *Move) UnmarshalJSON(data []byte) error {
	type alias Move
	aux := &struct {
		*alias
	}{
		alias: (*alias)(m),
	}
	return json.Unmarshal(data, &aux)
}
