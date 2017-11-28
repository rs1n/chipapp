package config

import (
	"os"

	"github.com/globalsign/mgo"
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

	Mongo *mgo.DialInfo
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
