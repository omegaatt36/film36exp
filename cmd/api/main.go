package main

import (
	"context"

	"github.com/omegaatt36/film36exp/app"
	"github.com/omegaatt36/film36exp/app/api"
	"github.com/omegaatt36/film36exp/logging"
	"github.com/urfave/cli/v2"
)

// Main is the entry point of the application.
func Main(ctx context.Context) {
	logging.Init(!api.IsLocal())

	stopped := api.NewServer().Start(ctx)

	<-stopped
	<-ctx.Done()
	logging.Info("Shutting down")
}

func main() {
	app := app.App{
		Main:  Main,
		Flags: []cli.Flag{},
	}

	app.Run()
}
