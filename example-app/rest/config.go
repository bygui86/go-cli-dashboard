package rest

import (
	"github.com/bygui86/go-cli-dashboard/example-app/logging"
	"github.com/bygui86/go-cli-dashboard/example-app/utils"
)

const (
	restHostEnvVar = "REST_HOST"
	restPortEnvVar = "REST_PORT"

	restHostDefault = "localhost"
	restPortDefault = 8080
)

type Config struct {
	RestHost string
	RestPort int
}

func loadConfig() *Config {
	logging.Log.Debug("Load REST configurations")
	return &Config{
		RestHost: utils.GetStringEnv(restHostEnvVar, restHostDefault),
		RestPort: utils.GetIntEnv(restPortEnvVar, restPortDefault),
	}
}
