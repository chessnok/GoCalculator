package queue

import (
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
)

// Producer - struct for sending messages to rabbitmq
type Producer struct {
	// conn - connection to rabbitmq
	conn *amqp091.Connection
	// queueName - name of queue
	queueName string
	// t - type of message
	t string
}

func NewProducer(conn *amqp091.Connection, queueName, t string) *Producer {
	p := &Producer{
		conn:      conn,
		queueName: queueName,
		t:         t,
	}
	ch, err := p.conn.Channel()
	if err != nil {
		return nil
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
			return nil
		}
	}
	return p
}

// send sends a message. If the queue doesn't exist, it will be created.
func (p *Producer) send(msg []byte) error {
	ch, err := p.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	return ch.Publish(
		"",
		p.queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: p.t,
			Body:        msg,
		},
	)
}

func (p *Producer) SendJson(msg interface{}) error {
	fmt.Println("Sending message. Body: " + fmt.Sprintf("%v", msg))
	t, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return p.send(t)
}
