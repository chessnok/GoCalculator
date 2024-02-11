package application

import (
	"context"
	"github.com/chessnok/GoCalculator/orchestrator/http/server"
	db2 "github.com/chessnok/GoCalculator/orchestrator/internal/db"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit/queue"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
)

type Application struct {
	server   *server.Server
	context  context.Context
	conn     *amqp.Connection
	db       *db2.Postgres
	producer *queue.Producer
	consumer *queue.Consumer
}

func NewApplication(ctx context.Context) (*Application, error) {
	cfg := NewConfig()
	pg, err := db2.NewPostgres(db2.NewConfigFromEnv())
	if err != nil {
		return nil, err
	}
	pg.Init()
	conn, err := amqp.Dial(cfg.RabbitConfig.GetURI())
	if err != nil {
		return nil, err
	}
	producer := queue.NewProducer(conn, cfg.RabbitConfig.TaskQueueName, "text/plain")
	consumer := queue.NewConsumer(conn, cfg.RabbitConfig.ResultQueueName, pg.OnNewResult)
	return &Application{
		context:  ctx,
		server:   server.NewServer(server.NewConfig(8080, cfg.CalculatorConfig, pg)),
		db:       pg,
		conn:     conn,
		producer: producer,
		consumer: consumer,
	}, nil
}

func (a Application) Start() int {
	err := a.server.Start()
	if err != nil {
		return 1
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	return 0
}
