package app

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	PortAddr string `yaml:"port_addr" env:"port_addr"`
	Env      string `yaml:"env" env-default:"development"`
	ConfigBD `yaml:"storage"`
}

type ConfigBD struct {
	Host     string `yaml:"host" env:"host"`
	Port     string `yaml:"port" env:"port"`
	DBName   string `yaml:"dbname" env:"dbname" `
	User     string `yaml:"user" env:"user"`
	Password string `yaml:"password" env:"password"`
	SSLMode  string `yaml:"sslmode" env:"sslmode"`
}

var cfg *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{}
		if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
			log.Fatalln(err)
		}
	})
	return cfg
}
