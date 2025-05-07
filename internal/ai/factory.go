package ai

import (
	"fmt"

	"github.com/GabrielChaves1/easycommit/internal/config"
)

type AgentType string

const (
	AgentTypeOpenAI AgentType = "openai"
	AgentTypeGemini AgentType = "gemini"
)

type AgentFactory func(*config.Config) Agent

var agentFactory = map[AgentType]AgentFactory{
	AgentTypeOpenAI: func(cfg *config.Config) Agent {
		return NewOpenAIClient(cfg.APIKey)
	},
	AgentTypeGemini: func(cfg *config.Config) Agent {
		return NewGeminiClient(cfg.APIKey)
	},
}

func NewAgent(agentType string, opts ...Option) (Agent, error) {
	aType := AgentType(agentType)
	if aType == "" {
		return nil, fmt.Errorf("agent type cannot be empty")
	}
	factory, ok := agentFactory[aType]
	if !ok {
		return nil, fmt.Errorf("unknown agent type: %s", agentType)
	}

	cfg := &config.Config{}
	for _, opt := range opts {
		opt(cfg)
	}

	return factory(cfg), nil
}

func (a AgentType) String() string {
	return string(a)
}

func (a AgentType) IsValid() bool {
	_, ok := agentFactory[a]
	return ok
}

func ListAgentTypes() []string {
	types := make([]string, 0, len(agentFactory))
	for at := range agentFactory {
		types = append(types, string(at))
	}
	return types
}
