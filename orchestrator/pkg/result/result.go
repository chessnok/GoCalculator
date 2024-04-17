package result

import (
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"log"
)

type Result struct {
	Id     string  `json:"id"`
	Result float64 `json:"result"`
	IsErr  bool    `json:"is_err"`
	Error  string  `json:"error"`
}

func NewResult(id string, result float64, isErr bool, error string) *Result {
	return &Result{
		Id:     id,
		Result: result,
		IsErr:  isErr,
		Error:  error,
	}
}

func ResultFromDelivery(delivery *amqp091.Delivery) (*Result, error) {
	logger := log.Default()
	res := Result{}
	err := json.Unmarshal(delivery.Body, &res)
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	return &res, nil
}
