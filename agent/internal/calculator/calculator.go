package calculator

import (
	"errors"
	"github.com/chessnok/GoCalculator/agent/internal/message"
	"github.com/chessnok/GoCalculator/rabbit/queue"
	"github.com/streadway/amqp"
)

var (
	ErrDivisionByZero   = errors.New("division by zero")
	ErrInvalidOperation = errors.New("invalid operation")
)

type Calculator struct {
	config   Config
	Tasks    chan message.Task
	producer *queue.Producer
}

func NewCalculator(config *Config, producer *queue.Producer) *Calculator {
	return &Calculator{
		config:   *config,
		Tasks:    make(chan message.Task),
		producer: producer,
	}
}

func (c *Calculator) Start() {
	for i := 0; i < c.config.ParallelWorkers; i++ {
		go func() {
			for task := range c.Tasks {
				result, err := c.calc(task.Operation, task.A, task.B)
				var res *message.Result
				if err != nil {
					res = message.NewResult(task.Id, 0, true, "Error while running the operation")
				} else {
					res = message.NewResult(task.Id, result, false, "")
				}
				err = c.producer.Send(res.ToJSON())
			}
		}()
	}
}
func (c *Calculator) TaskReceived(delivery *amqp.Delivery) {
	task := message.TaskFromDelivery(delivery)
	c.Tasks <- *task
}
