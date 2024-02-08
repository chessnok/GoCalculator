package server

import "github.com/chessnok/GoCalculator/agent/internal/calculator"

// Config is a struct that contains the configuration for the server
type Config struct {
	Port             int
	CalculatorConfig *calculator.Config
}

// NewConfig creates a new configuration with the given orchestrator manager and returns it
func NewConfig(port int, cc *calculator.Config) *Config {
	return &Config{
		Port:             port,
		CalculatorConfig: cc,
	}
}
