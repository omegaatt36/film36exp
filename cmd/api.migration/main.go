package main

import (
	"context"
	"log"

	"github.com/omegaatt36/film36exp/app"
	"github.com/omegaatt36/film36exp/logging"
	"github.com/omegaatt36/film36exp/rdb/database"
	"github.com/omegaatt36/film36exp/rdb/database/migration"
	v1 "github.com/omegaatt36/film36exp/rdb/database/migration/v1"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/urfave/cli/v2"
)

type config struct {
	rollback bool

	appEnvironment string
}

var cfg config

// Main starts process in cli.
func Main(ctx context.Context) {
	logging.Init(cfg.appEnvironment != "local")

	db := database.GetDB(database.Default).Debug()
	mg := migration.NewMigrator(db,
		[]*gormigrate.Migration{
			&v1.CreateUserAndFilmLogAndPhoto,
		},
	)

	if cfg.rollback {
		err := mg.Rollback()
		if err != nil {
			log.Fatalf("Could not RollbackLast: %v", err)
		}

		return
	}

	err := mg.Upgrade()
	if err != nil {
		log.Fatalf("Could not Upgrade: %v", err)
	}
}

func main() {
	app := app.App{
		Main: Main,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "rollback-last",
				EnvVars:     []string{"ROLLBACK_LAST"},
				Value:       false,
				Destination: &cfg.rollback,
			},
			&cli.StringFlag{
				Name:        "app-env",
				EnvVars:     []string{"APP_ENV"},
				Destination: &cfg.appEnvironment,
				Required:    false,
				DefaultText: "local",
				Value:       "local",
			},
		},
	}

	app.Run()
}
