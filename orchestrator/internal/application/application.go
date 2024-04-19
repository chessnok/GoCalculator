package application

import (
	"context"
	"github.com/chessnok/GoCalculator/orchestrator/http/server"
	db2 "github.com/chessnok/GoCalculator/orchestrator/internal/db"
	manager2 "github.com/chessnok/GoCalculator/orchestrator/internal/expressions/manager"
	"github.com/chessnok/GoCalculator/orchestrator/internal/grpc"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit/queue"
	"github.com/rabbitmq/amqp091-go"
	"os"
	"os/signal"
)

type Application struct {
	server             *server.Server
	context            context.Context
	conn               *amqp091.Connection
	db                 *db2.Postgres
	producer           *queue.Producer
	consumer           *queue.Consumer
	expressionsManager *manager2.TasksManager
	grpcServer         *grpc.Server
	cfg                *Config
}

func NewApplication(ctx context.Context) (*Application, error) {
	cfg := NewConfig()
	pdConfig := db2.NewConfigFromEnv()
	pg, err := db2.NewPostgres(pdConfig)
	if err != nil {
		return nil, err
	}
	err = pg.Init()
	if err != nil {
		return nil, err
	}
	conn, err := amqp091.Dial(cfg.RabbitConfig.GetURI())
	if err != nil {
		return nil, err
	}
	producer := queue.NewProducer(conn, cfg.RabbitConfig.TaskQueueName, "text/plain")
	consumer := queue.NewConsumer(conn, cfg.RabbitConfig.ResultQueueName)
	expressionsManager := manager2.NewTasksManager(pg, producer, consumer, cfg.CalculatorConfig)
	grpcServer := grpc.NewServer(cfg.CalculatorConfig, pg)
	return &Application{
		cfg:                cfg,
		context:            ctx,
		server:             server.NewServer(server.NewConfig(8080, cfg.CalculatorConfig, pg)),
		db:                 pg,
		conn:               conn,
		producer:           producer,
		consumer:           consumer,
		expressionsManager: expressionsManager,
		grpcServer:         grpcServer,
	}, nil
}

func (a *Application) Start() int {
	go grpc.Start(a.cfg.CalculatorConfig, a.db)
	go a.server.Start()
	go a.expressionsManager.Start()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	return 0
}
