package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Auth     AuthConfig
	Cors     CORSConfig
}

type AppConfig struct {
	Port  string `env:"APP_PORT" 	env-default:"55677" env-upd`
	Host  string `env:"APP_HOST" 	env-default:"localhost"`
	Admin string `env:"APP_ADMIN" 	env-default:"admin"`
	Pwd   string `env:"APP_PWD" 	env-default:"admin"`
}

type DatabaseConfig struct {
	Port int    `env:"DATABASE_PORT"		env-default:"5432"`
	Host string `env:"DATABASE_HOST"		env-default:"localhost"`
	DB   string `env:"DATABASE_DB"			env-default:"chater"`
	User string `env:"DATABASE_USER"		env-default:"user"`
	Pwd  string `env:"DATABASE_PASSWORD"	env-default:"user"`
}

type AuthConfig struct {
	JWT       JWTConfig
	OpenAIKey string `env: OPENAIKEY`
}

type JWTConfig struct {
	Key             string `env: JWT_KEY`
	ExpirationTimeH int    `env: JWT_EXP_TIME env-default:"72"`
}

type CORSConfig struct {
	AllowOrigins     []string `yaml: "allowOrigins"`
	AllowMethods     []string `yaml: "allowMethods"`
	AllowHeaders     []string `yaml: "allowHeaders"`
	AllowCredentials bool     `yaml: "allowCredentials"`
}

var (
	cfg  Config
	once sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		err := cleanenv.ReadEnv(&cfg)
		if err != nil {
			log.Printf("err read env cfg: %s", err)
		}
		err = cleanenv.ReadConfig("./config/cors-config.yaml", cfg)
		if err != nil {
			log.Printf("err read cors cfg: %s", err)
		}
	})
	log.Println("config has been loaded")
	return &cfg
}
