package task

import (
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type Task struct {
	Id        string  `json:"id"`
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
}

func TaskFromDelivery(delivery *amqp091.Delivery) *Task {
	logger := log.Default()
	task := Task{}
	err := json.Unmarshal(delivery.Body, &task)
	if err != nil {
		logger.Println(err)
		return nil
	}
	return &task
}
