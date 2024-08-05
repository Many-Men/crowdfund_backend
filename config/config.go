package config

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload"
	"sync"
)

type EnvMode string

const (
	Dev  EnvMode = "dev"
	Prod EnvMode = "prod"
)

type Config struct {
	Server  Server
	MongoDB MongoDB
	API     API
}

type (
	MongoDB struct {
		Host     string `env:"DB_HOST"`
		Port     string `env:"DB_PORT"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Database string `env:"DB_NAME"`
	}

	Server struct {
		Port string `env:"SERVER_PORT"`
	}

	API struct {
		EnrichAPI string `env:"ENRICH_API_URL"`
	}
)

var (
	instance Config
	once     sync.Once
)

func Load() *Config {
	once.Do(func() {
		if err := env.Parse(&instance); err != nil {
			panic(err)
		}
	})

	return &instance
}
