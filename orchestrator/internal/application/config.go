package application

import (
	"github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit"
	agentproto "github.com/chessnok/GoCalculator/proto"
)

type Config struct {
	Port             int
	CalculatorConfig *agentproto.Config
	RabbitConfig     *rabbit.Config
}

func NewConfig() *Config {
	defaultTime := int64(1)
	return &Config{
		Port: 8080,
		CalculatorConfig: &agentproto.Config{
			AddExecutionTime: defaultTime,
			SubExecutionTime: defaultTime,
			MulExecutionTime: defaultTime,
			DivExecutionTime: defaultTime,
		},
		RabbitConfig: rabbit.NewConfigFromEnv(),
	}
}
