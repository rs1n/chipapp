package config

import (
	"os"

	"upper.io/db.v3/postgresql"
)

const (
	envProduction = "production"
	envTest       = "test"

	envVarName       = "CHIP_ENV"
	envVarDbPassword = "CHIPAPP_DATABASE_PASSWORD"
)

type Config struct {
	IsDebug bool
	Port    int

	Postgres *postgresql.ConnectionURL
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
