package application

import (
	"github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit"
	agent_proto "github.com/chessnok/GoCalculator/proto"
	"os"
	"strings"
	"time"
)

type Config struct {
	Port             int
	CalculatorConfig *agent_proto.Config
	RabbitConfig     *rabbit.Config
	LoadDefautAgent  bool
}

func NewConfig() *Config {
	s := os.Getenv("LOAD_DEFAULT_AGENT")
	s = strings.ToLower(s)
	defaultTime := int64(time.Millisecond * 1)
	return &Config{
		Port: 8080,
		CalculatorConfig: &agent_proto.Config{
			AddExecutionTime: defaultTime,
			SubExecutionTime: defaultTime,
			MulExecutionTime: defaultTime,
			DivExecutionTime: defaultTime,
		},
		RabbitConfig:    rabbit.NewConfigFromEnv(),
		LoadDefautAgent: s == "true" || s == "1",
	}
}
