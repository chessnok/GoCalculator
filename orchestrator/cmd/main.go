package main

import (
	"context"
	"github.com/chessnok/GoCalculator/orchestrator/internal/application"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	os.Exit(MainWithCode(ctx))
}

func MainWithCode(ctx context.Context) int {
	app, err := application.NewApplication(ctx)
	if err != nil {
		log.Default().Println(err)
		return 1
	}
	return app.Start()
}
