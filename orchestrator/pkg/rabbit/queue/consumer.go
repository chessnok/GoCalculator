package queue

import (
	"fmt"
	"github.com/streadway/amqp"
)

// Consumer - struct for receiving messages from rabbitmq
type Consumer struct {
	// conn - connection to rabbitmq
	conn *amqp.Connection
	// queueName - name of queue
	queueName string
	Stop      chan struct{}
}

// NewConsumer - create new consumer
func NewConsumer(conn *amqp.Connection, queueName string) *Consumer {
	return &Consumer{
		conn:      conn,
		queueName: queueName,
	}
}

// Consume - start consuming messages. This method is blocking
func (c *Consumer) Consume(onMessage func(*amqp.Delivery)) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	_, err = ch.QueueInspect(c.queueName)
	if err != nil {
		ch, _ = c.conn.Channel()
		_, err = ch.QueueDeclare(
			c.queueName,
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return err
		}
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
