package server

import (
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	"github.com/chessnok/GoCalculator/orchestrator/internal/agents/manager"
	"github.com/chessnok/GoCalculator/orchestrator/internal/db"
)

// Config is a struct that contains the configuration for the server
type Config struct {
	Port             int
	CalculatorConfig *calculator.Config
	DB               *db.Postgres
	agentManager     *manager.AgentManager
}

// NewConfig creates a new configuration with the given orchestrator agentManager and returns it
func NewConfig(port int, config *calculator.Config, db *db.Postgres, agentManager *manager.AgentManager) *Config {
	return &Config{
		Port:             port,
		CalculatorConfig: config,
		DB:               db,
		agentManager:     agentManager,
	}
}
