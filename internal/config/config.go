package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		Port  string `env:"SERVER_PORT" 	env-default:"55677"`
		Host  string `env:"SERVER_HOST" 	env-default:"localhost"`
		Admin string `env:"SERVER_ADMIN" 	env-default:"admin"`
		Pwd   string `env:"SERVER_PWD" 	env-default:"admin"`
	}
	Database struct {
		Port string `env:"DB_PORT"		env-default:"5432"`
		Host string `env:"DB_HOST"		env-default:"localhost"`
		User string `env:"DB_USER"		env-default:"user"`
		Pwd  string `env:"DB_PWD"		env-default:"user"`
	}
	APIKeys struct {
		OpenAIKey string `env: OPENAIKEY`
	}
}

var (
	cfg    Config
	once   sync.Once
	logger log.Logger
)

func LoadConfig() *Config {
	once.Do(func() {
		err := cleanenv.ReadEnv(&cfg)
		if err != nil {
			logger.Printf("config doesn't load: %s", err)
		}
	})
	logger.Println("config has loaded")
	return &cfg
}
