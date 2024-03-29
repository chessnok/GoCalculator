package application

import (
	"context"
	"fmt"
	server "github.com/chessnok/GoCalculator/agent/http"
	"github.com/chessnok/GoCalculator/agent/pkg/calculator"
	queue2 "github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit/queue"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
)

type Application struct {
	context       context.Context
	config        *Config
	server        *server.Server
	connection    *amqp.Connection
	consumer      *queue2.Consumer
	producer      *queue2.Producer
	calculator    *calculator.Calculator
	resultChannel chan interface{}
}

func NewApplication(ctx context.Context) *Application {
	godotenv.Load("env/.env.go", "env/.env.pg", "env/.env.rmq")
	cfg := NewConfig()
	conn, err := amqp.Dial(cfg.RabbitConfig.GetURI())
	if err != nil {
		log.Default().Println(err)
	}
	tasks := make(chan interface{}, 10)
	results := make(chan interface{}, 10)
	producer := queue2.NewProducer(conn, cfg.RabbitConfig.ResultQueueName, "text/plain")
	calc := calculator.NewCalculator(cfg.CalculatorConfig, tasks, results)
	consumer := queue2.NewConsumer(conn, cfg.RabbitConfig.TaskQueueName)
	return &Application{
		context:       ctx,
		config:        cfg,
		connection:    conn,
		consumer:      consumer,
		producer:      producer,
		server:        server.NewServer(server.NewConfig(cfg.Port, calc)),
		calculator:    calc,
		resultChannel: results,
	}
}

func (a Application) Start() int {
	go func() {
		err := a.server.Start()
		if err != nil {
			return
		}
	}()
	if a.connection == nil || a.consumer == nil || a.producer == nil {
		return 1
	}
	defer a.connection.Close()
	go func() {
		a.consumer.Consume(func(delivery *amqp.Delivery) {
			fmt.Println(a.calculator.Cnt)
			a.calculator.Tasks <- delivery
		})
	}()
	go func() {
		for m := range a.resultChannel {
			err := a.producer.SendJson(m)
			if err != nil {
				log.Default().Println(err)
			}
		}
	}()
	go func() {
		a.calculator.Start()
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	a.calculator.Stop()
	return 0
}
