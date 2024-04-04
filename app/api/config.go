//go:generate go-enum -f=$GOFILE

package api

import (
	"github.com/urfave/cli/v2"

	"github.com/omegaatt36/film36exp/cliflag"
)

// Env represents the environment of the application.
// ENUM(
// local
// development
// production
// )
type Env string

type config struct {
	appEnvironment string
	listenAddr     string
}

var defaultConfig config

func init() {
	cliflag.Register(&defaultConfig)
}

// CliFlags returns cli flags to setup cache package.
func (cfg *config) CliFlags() []cli.Flag {
	var flags []cli.Flag

	flags = append(flags, &cli.StringFlag{
		Name:        "app-env",
		EnvVars:     []string{"APP_ENV"},
		Destination: &cfg.appEnvironment,
		Required:    false,
		DefaultText: EnvLocal.String(),
		Value:       EnvLocal.String(),
	}, &cli.StringFlag{
		Name:        "listen-addr",
		EnvVars:     []string{"LISTEN_ADDR"},
		Destination: &cfg.listenAddr,
		Required:    false,
		DefaultText: ":8070",
		Value:       ":8070",
	})

	return flags
}

// IsLocal will return true if the APP_ENV is equals to local.
func IsLocal() bool {
	return defaultConfig.appEnvironment == EnvLocal.String()
}

// IsDevelopment will return true if the APP_ENV is equals to dev.
func IsDevelopment() bool {
	return defaultConfig.appEnvironment == EnvDevelopment.String()
}

// IsProduction will return true if the APP_ENV is equals to prod.
func IsProduction() bool {
	return defaultConfig.appEnvironment == EnvProduction.String()
}

// GetAppEnv returns the current app environment.
func GetAppEnv() string {
	return defaultConfig.appEnvironment
}
