package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	Addr string `yaml:"addr" env-default:":8080"`
	DB   Dbconfig
}

type Dbconfig struct {
	Name     string `yaml:"name" env-required:"true"`
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-required:"true"`
	Username string `yaml:"username" env-default:"postgres"`
}

var cfg Config

func InitConfig(configpath string) Config {
	if err := cleanenv.ReadConfig(configpath, &cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
