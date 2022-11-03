package plexnet

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type AppConfig struct {
	Port     string `yaml:"port" env:"PORT"`
	AppName  string `yaml:"appName" env:"APP_NAME"`
	BasePath string `yaml:"basePath"`
}

func LoadConfig(config interface{}) {
	err := cleanenv.ReadConfig("config.yaml", &config)
	if err != nil {
		log.Fatalf("error reading config file %v", err)
	}
}
