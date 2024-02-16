package application

import (
	"context"
	"github.com/chessnok/GoCalculator/orchestrator/http/server"
	"github.com/chessnok/GoCalculator/orchestrator/internal/agents/manager"
	db2 "github.com/chessnok/GoCalculator/orchestrator/internal/db"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit/queue"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"time"
)

type Application struct {
	server       *server.Server
	context      context.Context
	conn         *amqp.Connection
	db           *db2.Postgres
	producer     *queue.Producer
	consumer     *queue.Consumer
	agentManager *manager.AgentManager
}

func NewApplication(ctx context.Context) (*Application, error) {
	cfg := NewConfig()
	pg, err := db2.NewPostgres(db2.NewConfigFromEnv())
	if err != nil {
		return nil, err
	}
	err = pg.Init()
	if err != nil {
		return nil, err
	}
	conn, err := amqp.Dial(cfg.RabbitConfig.GetURI())
	if err != nil {
		return nil, err
	}
	producer := queue.NewProducer(conn, cfg.RabbitConfig.TaskQueueName, "text/plain")
	consumer := queue.NewConsumer(conn, cfg.RabbitConfig.ResultQueueName, pg.OnNewResult)
	agentManager := manager.NewAgentManager(pg.Agents, time.Second/4, cfg.CalculatorConfig)
	return &Application{
		context:      ctx,
		server:       server.NewServer(server.NewConfig(8080, cfg.CalculatorConfig, pg, producer, agentManager)),
		db:           pg,
		conn:         conn,
		producer:     producer,
		consumer:     consumer,
		agentManager: agentManager,
	}, nil
}

func (a Application) Start() int {
	go a.server.Start()
	a.agentManager.StartPing()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	return 0
}
