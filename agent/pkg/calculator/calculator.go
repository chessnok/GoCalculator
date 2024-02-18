package calculator

import (
	"errors"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/result"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/task"
	"github.com/streadway/amqp"
	"sync"
)

var (
	ErrDivisionByZero   = errors.New("division by zero")
	ErrInvalidOperation = errors.New("invalid operation")
)

type Calculator struct {
	LastOperationID string
	Cnt             int
	mu              sync.RWMutex
	Config          *Config
	Tasks           chan interface{}
	Results         chan interface{}
}

func NewCalculator(config *Config, tasks, results chan interface{}) *Calculator {
	return &Calculator{
		Config:  config,
		Tasks:   tasks,
		Results: results,
	}
}

func (c *Calculator) Start() {
	for i := 0; i < c.Config.ParallelWorkers; i++ {
		go func() {
			for {
				tsk := <-c.Tasks
				task := task.TaskFromDelivery(tsk.(*amqp.Delivery))
				c.mu.Lock()
				c.LastOperationID = task.Id
				c.mu.Unlock()
				resp, err := c.calc(task.Operation, task.A, task.B)
				var res *result.Result
				if err != nil {
					res = result.NewResult(task.Id, 0, true, err.Error())
				} else {
					res = result.NewResult(task.Id, resp, false, "")
				}
				c.Results <- res
			}
		}()
	}
}

func (c *Calculator) Stop() {
	close(c.Tasks)
	close(c.Results)
}
func (c *Calculator) TaskReceived(delivery *amqp.Delivery) {
	task := task.TaskFromDelivery(delivery)
	c.Tasks <- *task
}
