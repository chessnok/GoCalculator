package application

import (
	"context"
	server "github.com/chessnok/GoCalculator/agent/http"
	"github.com/chessnok/GoCalculator/agent/internal/calculator"
	"github.com/chessnok/GoCalculator/rabbit/queue"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
)

type Application struct {
	context    context.Context
	config     *Config
	server     *server.Server
	connection *amqp.Connection
	consumer   *queue.Consumer
	producer   *queue.Producer
	calculator *calculator.Calculator
}

func NewApplication(ctx context.Context) *Application {
	cfg := NewConfig()
	conn, err := amqp.Dial(cfg.RabbitConfig.GetURI())
	if err != nil {
		log.Default().Println(err)
		return nil
	}
	producer := queue.NewProducer(conn, cfg.RabbitConfig.ResultQueueName, "text/plain")
	calc := calculator.NewCalculator(cfg.CalculatorConfig, producer)
	consumer := queue.NewConsumer(conn, cfg.RabbitConfig.TaskQueueName, calc.TaskReceived)
	return &Application{
		context:    ctx,
		config:     cfg,
		connection: conn,
		consumer:   consumer,
		producer:   producer,
		server:     server.NewServer(server.NewConfig(cfg.Port, cfg.CalculatorConfig)),
		calculator: calc,
	}
}

func (a Application) Start() int {
	if a.connection == nil || a.consumer == nil || a.producer == nil {
		return 1
	}
	defer a.connection.Close()
	go func() {
		if err := a.consumer.Consume(); err != nil {
			log.Default().Println(err)
		}
	}()
	go func() {
		err := a.server.Start()
		if err != nil {
			return
		}
	}()
	go func() {
		a.calculator.Start()
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	return 0
}
