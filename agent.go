package pokemonbattlelib

// Party agnostic interface to get turns from a human or bot. Makes turns like a Trainer or Wild Pokemon would.
type AgentType int

const (
	AGENT_TRAINER AgentType = iota
	AGENT_NPC
	AGENT_WILD_POKEMON
)

type Agent interface {
	Act(b BattleInfo) Turn   // Called for each pokemon that the agent has control over. Allowed to block for an unlimited amount of time.
	GetAgentType() AgentType // To determine if an agent is a trainer, NPC, or wild Pokemon
}
