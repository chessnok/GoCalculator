package server

import (
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	"github.com/chessnok/GoCalculator/orchestrator/internal/db"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit/queue"
)

// Config is a struct that contains the configuration for the server
type Config struct {
	Port             int
	CalculatorConfig *calculator.Config
	DB               *db.Postgres
	Producer         *queue.Producer
}

// NewConfig creates a new configuration with the given orchestrator manager and returns it
func NewConfig(port int, config *calculator.Config, db *db.Postgres, producer *queue.Producer) *Config {
	return &Config{
		Port:             port,
		CalculatorConfig: config,
		DB:               db,
		Producer:         producer,
	}
}
