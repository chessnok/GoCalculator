package queue

import (
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

// Consumer - struct for receiving messages from rabbitmq
type Consumer struct {
	// conn - connection to rabbitmq
	conn *amqp091.Connection
	// queueName - name of queue
	queueName string
	Stop      chan struct{}
}

// NewConsumer - create new consumer
func NewConsumer(conn *amqp091.Connection, queueName string) *Consumer {
	c := &Consumer{
		conn:      conn,
		queueName: queueName,
	}
	ch, err := conn.Channel()
	if err != nil {
		return c
	}
	defer ch.Close()
	declareQueue(ch, queueName)
	return &Consumer{
		conn:      conn,
		queueName: queueName,
	}
}

// Consume - start consuming messages. This method is blocking
func (c *Consumer) Consume(onMessage func(*amqp091.Delivery)) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		c.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	go func() {
		defer ch.Close()
		for {
			select {
			case msg, ok := <-msgs:
				if ok {
					fmt.Println("Received a message. Body: " + string(msg.Body))
					onMessage(&msg)
				} else {
					return
				}
			case <-c.Stop:
				return
			}
		}
	}()
	return nil
}
