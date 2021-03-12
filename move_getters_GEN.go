// Code generated - DO NOT EDIT.
// Regenerate with `go generate`.

package pokemonbattlelib

func (n Move) Name() string           { return n.Data().Name }
func (n Move) Type() Type             { return n.Data().Type }
func (n Move) Category() MoveCategory { return n.Data().Category }
func (n Move) Targets() MoveTarget    { return n.Data().Targets }
func (n Move) Priority() int8         { return n.Data().Priority }
func (n Move) Power() uint            { return n.Data().Power }
func (n Move) Accuracy() uint         { return n.Data().Accuracy }
func (n Move) InitialMaxPP() uint8    { return n.Data().InitialMaxPP }
