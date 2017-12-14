package config

import "upper.io/db.v3/postgresql"

var test = Config{
	IsDebug: true,
	Port:    3000,

	Postgres: &postgresql.ConnectionURL{
		Host:     "localhost",
		Database: "demotest",
		User:     "demotest",
		Password: "123",
	},
}
