package application

import (
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit"
	"os"
	"strconv"
)

type Config struct {
	RabbitConfig     *rabbit.Config
	CalculatorConfig *calculator.Config
	Port             int
}

func NewConfig() *Config {
	port := 8080
	if len(os.Args) > 1 {
		p, err := strconv.Atoi(os.Args[1])
		if err != nil {
			p = 8080
		}
		port = p
	}
	return &Config{
		RabbitConfig:     rabbit.NewConfigFromEnv(),
		Port:             port,
		CalculatorConfig: calculator.NewConfigFromArgs(),
	}
}
