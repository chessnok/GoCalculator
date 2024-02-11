package db

import "os"

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
}

func NewConfigFromEnv() *Config {
	return &Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     5432,
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DB:       os.Getenv("POSTGRES_DB"),
	}
}
