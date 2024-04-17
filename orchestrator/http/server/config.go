package server

import (
	"github.com/chessnok/GoCalculator/orchestrator/internal/db"
	agent_proto "github.com/chessnok/GoCalculator/proto"
)

// Config is a struct that contains the configuration for the server
type Config struct {
	Port             int
	CalculatorConfig *agent_proto.Config
	DB               *db.Postgres
}

// NewConfig creates a new configuration with the given orchestrator agentManager and returns it
func NewConfig(port int, config *agent_proto.Config, db *db.Postgres) *Config {
	return &Config{
		Port:             port,
		CalculatorConfig: config,
		DB:               db,
	}
}
