package config

import "os"

const (
	envProduction = "production"
	envTest       = "test"

	envVarName = "CHIP_ENV"
)

type Config struct {
	IsDebug bool
	Port    int
}

func NewConfig() *Config {
	cfg := development // Default configuration.
	switch appEnvironment() {
	case envProduction:
		cfg = production
	case envTest:
		cfg = test
	}
	return &cfg
}

func appEnvironment() string {
	return os.Getenv(envVarName)
}
