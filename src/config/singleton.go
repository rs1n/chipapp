package config

import (
	"log"

	"github.com/rs1n/chipapp/src/config/environments"
)

var config *Config

func GetConfig() *Config {
	if config == nil {
		initConfig()
	}
	return config
}

func initConfig() {
	config = &Config{
		EnvConfig: environments.GetConfig(),
	}
	log.Printf("Using configuration: %+v", config)
}
