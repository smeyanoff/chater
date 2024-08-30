package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App struct {
		Port  string `env:"APP_PORT" 	env-default:"55677"`
		Host  string `env:"APP_HOST" 	env-default:"localhost"`
		Admin string `env:"APP_ADMIN" 	env-default:"admin"`
		Pwd   string `env:"APP_PWD" 	env-default:"admin"`
	}
	Database struct {
		Port int    `env:"POSTGRES_PORT"		env-default:"5432"`
		Host string `env:"POSTGRES_HOST"		env-default:"localhost"`
		DB   string `env:"POSTGRES_DB"			env-default:"chater"`
		User string `env:"POSTGRES_USER"		env-default:"user"`
		Pwd  string `env:"POSTGRES_PASSWORD"	env-default:"user"`
	}
	Auth struct {
		JWT struct {
			Key             string `env: JWT_KEY`
			ExpirationTimeH int    `env: JWT_EXP_TIME env-default:"72"`
		}
		OpenAIKey string `env: OPENAIKEY`
	}
}

var (
	cfg  Config
	once sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		err := cleanenv.ReadEnv(&cfg)
		if err != nil {
			log.Printf("config doesn't load: %s", err)
		}
	})
	log.Println("config has loaded")
	return &cfg
}
