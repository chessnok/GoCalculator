package calculator

import (
	"errors"
	"fmt"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/result"
	"github.com/chessnok/GoCalculator/orchestrator/pkg/task"
	agentproto "github.com/chessnok/GoCalculator/proto"
	"github.com/rabbitmq/amqp091-go"
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
	Config          *agentproto.Config
	Tasks           chan interface{}
	Results         chan interface{}
	Workers         int
}

func NewCalculator(config *agentproto.Config, tasks, results chan interface{}) *Calculator {
	return &Calculator{
		Config:  config,
		Tasks:   tasks,
		Results: results,
	}
}

func (c *Calculator) Start() {
	for i := 0; i < GetWorkersCount(); i++ {
		go func() {
			for {
				tsk := <-c.Tasks
				task2 := task.TaskFromDelivery(tsk.(*amqp091.Delivery))
				c.mu.Lock()
				c.LastOperationID = task2.Id
				c.mu.Unlock()
				resp, err := c.calc(task2.Operation, task2.A, task2.B)
				var res *result.Result
				if err != nil {
					res = result.NewResult(task2.Id, 0, true, err.Error())
				} else {
					res = result.NewResult(task2.Id, resp, false, "")
				}
				fmt.Println("New result: ", res)
				c.Results <- res
			}
		}()
	}
}

func (c *Calculator) Stop() {
	close(c.Tasks)
	close(c.Results)
}
func (c *Calculator) TaskReceived(delivery *amqp091.Delivery) {
	tsk := task.TaskFromDelivery(delivery)
	c.Tasks <- *tsk
}
