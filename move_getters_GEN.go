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
func (m Move) MinHits() int           { return m.Data().MinHits }
func (m Move) MaxHits() int           { return m.Data().MaxHits }
func (m Move) MinTurns() int          { return m.Data().MinTurns }
func (m Move) MaxTurns() int          { return m.Data().MaxTurns }
func (m Move) Drain() int             { return m.Data().Drain }
func (m Move) Healing() int           { return m.Data().Healing }
func (m Move) CritRate() int          { return m.Data().CritRate }
func (m Move) AilmentChance() int     { return m.Data().AilmentChance }
func (m Move) FlinchChance() int      { return m.Data().FlinchChance }
func (m Move) StatChance() int        { return m.Data().StatChance }
func (m Move) Flags() MoveFlags       { return m.Data().Flags }

func (m *Move) MarshalJSON() ([]byte, error) {
	type alias Move
	return json.Marshal(&struct {
		*alias
		Name          string
		Type          Type
		Category      MoveCategory
		Targets       MoveTarget
		Priority      int8
		Power         uint
		Accuracy      uint
		InitialMaxPP  uint8
		MinHits       int
		MaxHits       int
		MinTurns      int
		MaxTurns      int
		Drain         int
		Healing       int
		CritRate      int
		AilmentChance int
		FlinchChance  int
		StatChance    int
		Flags         MoveFlags
	}{
		alias:         (*alias)(m),
		Name:          m.Name(),
		Type:          m.Type(),
		Category:      m.Category(),
		Targets:       m.Targets(),
		Priority:      m.Priority(),
		Power:         m.Power(),
		Accuracy:      m.Accuracy(),
		InitialMaxPP:  m.InitialMaxPP(),
		MinHits:       m.MinHits(),
		MaxHits:       m.MaxHits(),
		MinTurns:      m.MinTurns(),
		MaxTurns:      m.MaxTurns(),
		Drain:         m.Drain(),
		Healing:       m.Healing(),
		CritRate:      m.CritRate(),
		AilmentChance: m.AilmentChance(),
		FlinchChance:  m.FlinchChance(),
		StatChance:    m.StatChance(),
		Flags:         m.Flags(),
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
