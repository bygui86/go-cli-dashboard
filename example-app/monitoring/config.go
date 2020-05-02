package monitoring

import (
	"github.com/bygui86/go-cli-dashboard/example-app/logging"
	"github.com/bygui86/go-cli-dashboard/example-app/utils"
)

const (
	monitorHostEnvVar = "MONITOR_HOST"
	monitorPortEnvVar = "MONITOR_PORT"

	monitorHostDefault = "localhost"
	monitorPortDefault = 9090
)

type Config struct {
	MonitorHost string
	MonitorPort int
}

func loadConfig() *Config {
	logging.Log.Debug("Load Monitoring configurations")
	return &Config{
		MonitorHost: utils.GetStringEnv(monitorHostEnvVar, monitorHostDefault),
		MonitorPort: utils.GetIntEnv(monitorPortEnvVar, monitorPortDefault),
	}
}
