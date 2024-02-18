package application

import (
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit"
	"os"
	"strings"
)

type Config struct {
	Port             int
	CalculatorConfig *calculator.Config
	RabbitConfig     *rabbit.Config
	LoadDefautAgent  bool
}

func NewConfig() *Config {
	s := os.Getenv("LOAD_DEFAULT_AGENT")
	s = strings.ToLower(s)
	return &Config{
		Port:             8080,
		CalculatorConfig: calculator.NewConfigFromArgs(),
		RabbitConfig:     rabbit.NewConfigFromEnv(),
		LoadDefautAgent:  s == "true" || s == "1",
	}
}
