package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Addr string `yaml:"addr" env-required:"true"`
	Db   Db
}

type Db struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	Username string `yaml:"username" env-required:"true"`
	Name     string `yaml:"name" env-required:"true"`
}

var cfg Config

func InitConfig(configpath string) Config {
	if err := cleanenv.ReadConfig(configpath, &cfg); err != nil {
		log.Fatal("config_read_err")
	}
	return cfg
}
