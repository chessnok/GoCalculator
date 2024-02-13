package calculator

import (
	"errors"
	"github.com/chessnok/GoCalculator/agent/internal/message"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/rabbit/queue"
	"github.com/streadway/amqp"
	"sync"
)

var (
	ErrDivisionByZero   = errors.New("division by zero")
	ErrInvalidOperation = errors.New("invalid operation")
)

type Calculator struct {
	LastOperationID string
	mu              sync.RWMutex
	Config          *Config
	Tasks           chan message.Task
	producer        *queue.Producer
}

func NewCalculator(config *Config, producer *queue.Producer) *Calculator {
	return &Calculator{
		Config:   config,
		Tasks:    make(chan message.Task),
		producer: producer,
	}
}

func (c *Calculator) Start() {
	for i := 0; i < c.Config.ParallelWorkers; i++ {
		go func() {
			for task := range c.Tasks {
				c.mu.Lock()
				c.LastOperationID = task.Id
				c.mu.Unlock()
				result, err := c.calc(task.Operation, task.A, task.B)
				var res *message.Result
				if err != nil {
					res = message.NewResult(task.Id, 0, true, "Error while running the operation")
				} else {
					res = message.NewResult(task.Id, result, false, "")
				}
				err = c.producer.SendJson(res)
			}
		}()
	}
}

func (c *Calculator) Stop() {
	close(c.Tasks)
}
func (c *Calculator) TaskReceived(delivery *amqp.Delivery) {
	task := message.TaskFromDelivery(delivery)
	c.Tasks <- *task
}
