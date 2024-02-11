package queue

import (
	"github.com/streadway/amqp"
)

// Consumer - struct for receiving messages from rabbitmq
type Consumer struct {
	// conn - connection to rabbitmq
	conn *amqp.Connection
	// queueName - name of queue
	queueName string
	// onMessage - function for processing messages
	onMessage func(*amqp.Delivery)
}

// NewConsumer - create new consumer
func NewConsumer(conn *amqp.Connection, queueName string, onMessage func(*amqp.Delivery)) *Consumer {
	return &Consumer{
		conn:      conn,
		queueName: queueName,
		onMessage: onMessage,
	}
}

// Consume - start consuming messages. This method is blocking
func (c *Consumer) Consume() error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
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
	if err != nil {
		return err
	}
	for msg := range msgs {
		c.onMessage(&msg)
	}
	return nil
}
