package config

import (
	"os"

	"upper.io/db.v3/postgresql"
)

var production = Config{
	IsDebug: false,
	Port:    3000,

	Postgres: &postgresql.ConnectionURL{
		Host:     "localhost",
		Database: "demoprod",
		User:     "demoprod",
		Password: os.Getenv(envVarDbPassword),
	},
}
