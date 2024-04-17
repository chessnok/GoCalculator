package queue

import "github.com/rabbitmq/amqp091-go"

func declareQueue(ch *amqp091.Channel, name string) error {
	_, err := ch.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}
