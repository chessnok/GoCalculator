package main

import (
	"context"
	"github.com/chessnok/GoCalculator/agent/internal/application"
	"os"
)

func main() {
	ctx := context.Background()
	os.Exit(MainWithCode(ctx))
}

func MainWithCode(ctx context.Context) int {
	app := application.NewApplication(ctx)
	return app.Start()
}
