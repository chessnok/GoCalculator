package queue

import "github.com/streadway/amqp"

// Producer - struct for sending messages to rabbitmq
type Producer struct {
	// conn - connection to rabbitmq
	conn *amqp.Connection
	// queueName - name of queue
	queueName string
	// t - type of message
	t string
}

func NewProducer(conn *amqp.Connection, queueName, t string) *Producer {
	return &Producer{
		conn:      conn,
		queueName: queueName,
		t:         t,
	}
}

// Send sends a message. If the queue doesn't exist, it will be created.
func (p *Producer) Send(msg []byte) error {
	ch, err := p.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Check if queue exists, if not, declare it
	_, err = ch.QueueInspect(p.queueName)
	if err != nil {
		ch, _ = p.conn.Channel()
		_, err = ch.QueueDeclare(
			p.queueName,
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

	return ch.Publish(
		"",
		p.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: p.t,
			Body:        msg,
		},
	)
}
