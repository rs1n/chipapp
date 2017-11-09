package config

import "log"

var config *Config

func GetConfig() *Config {
	if config == nil {
		initConfig()
	}
	return config
}

func initConfig() {
	config = NewConfig()
	log.Printf("Using configuration: %+v", config)
}
