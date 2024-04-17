package application

import (
	"github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit"
	agent_proto "github.com/chessnok/GoCalculator/proto"
)

type Config struct {
	RabbitConfig     *rabbit.Config
	CalculatorConfig *agent_proto.Config
}

func NewConfig() *Config {
	return &Config{
		RabbitConfig: rabbit.NewConfigFromEnv(),
		CalculatorConfig: &agent_proto.Config{
			AddExecutionTime: 0,
			SubExecutionTime: 0,
			MulExecutionTime: 0,
			DivExecutionTime: 0,
		},
	}
}
