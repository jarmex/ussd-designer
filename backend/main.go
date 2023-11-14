package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/jarmex/ussd-designer/src/application"
)

func main() {
	app := application.New(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := app.Start(ctx); err != nil {
		fmt.Println("failed to start app:", err)
		os.Exit(1)
	}
}
