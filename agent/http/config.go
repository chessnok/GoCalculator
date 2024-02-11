package server

import (
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
)

// Config is a struct that contains the configuration for the server
type Config struct {
	Port       int
	Calculator *calculator.Calculator
}

// NewConfig creates a new configuration with the given orchestrator manager and returns it
func NewConfig(port int, calc *calculator.Calculator) *Config {
	return &Config{
		Port:       port,
		Calculator: calc,
	}
}
