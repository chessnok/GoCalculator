package application

import (
	"context"
	"fmt"
	"github.com/chessnok/GoCalculator/agent/internal/calculator"
	"github.com/chessnok/GoCalculator/agent/internal/grpc"
	queue2 "github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit/queue"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"os/signal"
)

type Application struct {
	config        *Config
	connection    *amqp091.Connection
	consumer      *queue2.Consumer
	producer      *queue2.Producer
	calculator    *calculator.Calculator
	grpc          *grpc.GRPC
	resultChannel chan interface{}
}

func NewApplication(ctx context.Context) *Application {
	cfg := NewConfig()
	conn, err := amqp091.Dial(cfg.RabbitConfig.GetURI())
	if err != nil {
		log.Default().Println(err)
	}
	tasks := make(chan interface{}, 10)
	results := make(chan interface{}, 10)
	producer := queue2.NewProducer(conn, cfg.RabbitConfig.ResultQueueName, "text/plain")
	calc := calculator.NewCalculator(cfg.CalculatorConfig, tasks, results)
	consumer := queue2.NewConsumer(conn, cfg.RabbitConfig.TaskQueueName)
	rpc := grpc.NewGRPC(grpc.NewConfig())
	return &Application{
		config:        cfg,
		connection:    conn,
		consumer:      consumer,
		grpc:          rpc,
		producer:      producer,
		calculator:    calc,
		resultChannel: results,
	}
}

func (a Application) Start() int {
	if a.connection == nil || a.consumer == nil || a.producer == nil || a.calculator == nil || a.grpc == nil {
		return 1
	}
	go func() {
		a.grpc.Connect()
	}()
	go func() {
		a.consumer.Consume(func(delivery *amqp091.Delivery) {
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
	a.grpc.Close()
	a.connection.Close()
	return 0
}
