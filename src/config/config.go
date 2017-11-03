package config

import "github.com/rs1n/chipapp/src/config/environments"

type Config struct {
	environments.Config
}

var config = Config{
	Config: environments.GetConfig(),
}

func GetConfig() Config {
	return config
}
