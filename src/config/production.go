package config

import (
	"os"
	"time"

	"github.com/globalsign/mgo"
)

var production = Config{
	IsDebug:   false,
	Port:      3000,
	SecretKey: []byte(os.Getenv(envVarSecretKey)),

	Mongo: &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "demoprod",
		Username: "demoprod",
		Password: os.Getenv(envVarDbPassword),
		Timeout:  5 * time.Second,
	},
}
