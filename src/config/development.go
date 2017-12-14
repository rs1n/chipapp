package config

import "upper.io/db.v3/postgresql"

var development = Config{
	IsDebug: true,
	Port:    3000,

	Postgres: &postgresql.ConnectionURL{
		Host:     "localhost",
		Database: "demodev",
		User:     "demodev",
		Password: "123",
	},
}
