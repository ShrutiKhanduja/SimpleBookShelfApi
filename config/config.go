package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Server serverConf
}

type serverConf struct {
	Port         uint16        `env:"SERVER_PORT,default=8080"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE"`
	Debug        bool          `env:"DEBUG,default=true"`
}

func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
