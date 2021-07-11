package config

import (
	"log"
	"time"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Debug  bool `env:"DEBUG,default=true"`
	Db     dbConf
	Server serverConf
}
type dbConf struct {
	Host     string `env:"DB_HOST,default=db"`
	Port     int    `env:"DB_PORT,default=3306"`
	Username string `env:"DB_USER,default=myapp_user"`
	Password string `env:"DB_PASS,myapp_pass"`
	DbName   string `env:"DB_NAME,myapp_db"`
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
