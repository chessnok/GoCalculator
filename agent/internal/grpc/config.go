package grpc

import "fmt"

type Config struct {
	Port int
	Host string
}

func NewConfig() *Config {
	return &Config{
		Port: 5000,
		Host: "localhost",
	}
}

func (c *Config) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
