package config

import (
	"github.com/bygui86/go-cli-dashboard/example-app/logging"
	"github.com/bygui86/go-cli-dashboard/example-app/utils"
)

const (
	shutdownTimeoutEnvVar  = "SHUTDOWN_TIMEOUT"
	shutdownTimeoutDefault = 10
)

type Config struct {
	ShutdownTimeout int
}

func LoadConfig() *Config {
	logging.Log.Debug("Load HTTP server general configurations")
	return &Config{
		ShutdownTimeout: utils.GetIntEnv(shutdownTimeoutEnvVar, shutdownTimeoutDefault),
	}
}
