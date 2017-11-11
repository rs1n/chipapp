package config

import (
	"time"

	"gopkg.in/mgo.v2"
)

var development = Config{
	IsDebug: true,
	Port:    3000,

	Mongo: &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "demodev",
		Username: "demodev",
		Password: "123",
		Timeout:  5 * time.Second,
	},
}
