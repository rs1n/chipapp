package config

import (
	"time"

	"github.com/globalsign/mgo"
)

var development = Config{
	IsDebug:   true,
	Port:      3000,
	SecretKey: []byte("MSjLRL86sgY9UTxe"),

	Mongo: &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "demodev",
		Username: "demodev",
		Password: "123",
		Timeout:  5 * time.Second,
	},
}
