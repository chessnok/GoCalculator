package application

import (
	"context"
	"github.com/chessnok/GoCalculator/http/server"
	"os"
	"os/signal"
)

type Application struct {
	server  *server.Server
	context context.Context
}

func NewApplication(ctx context.Context) *Application {
	return &Application{
		context: ctx,
		server:  server.NewServer(server.NewConfig()),
	}
}

func (a Application) Start() int {
	err := a.server.Start()
	if err != nil {
		return 1
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	return 0
}
