package application

import (
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit"
)

type Config struct {
	Port             int
	CalculatorConfig *calculator.Config
	RabbitConfig     *rabbit.Config
}

func NewConfig() *Config {
	return &Config{
		Port:             8080,
		CalculatorConfig: calculator.NewConfig(),
		RabbitConfig:     rabbit.NewConfigFromEnv(),
	}
}
