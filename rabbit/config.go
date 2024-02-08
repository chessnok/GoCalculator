package rabbit

import (
	"fmt"
	"os"
)

type Config struct {
	Host            string
	Port            int
	Username        string
	Password        string
	TaskQueueName   string
	ResultQueueName string
}

func NewConfigFromEnv() *Config {
	return &Config{
		Host:            os.Getenv("RABBITMQ_HOST"),
		Port:            5672,
		Username:        os.Getenv("RABBITMQ_DEFAULT_USER"),
		Password:        os.Getenv("RABBITMQ_DEFAULT_PASS"),
		TaskQueueName:   os.Getenv("RABBITMQ_TASK_QUEUE"),
		ResultQueueName: os.Getenv("RABBITMQ_RESULT_QUEUE"),
	}
}

func (c *Config) GetURI() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Username, c.Password, c.Host, c.Port)
}
