package pokemonbattlelib

// Party agnostic interface to get turns from a human or bot. Makes turns like a Trainer or Wild Pokemon would.
type Agent interface {
	Act(ctx *BattleContext) Turn // Called for each pokemon that the agent has control over. Allowed to block for an unlimited amount of time.
}
