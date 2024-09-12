package config

import (
	"log"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg  Config
	once sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		// Настраиваем Viper для работы с YAML-файлом
		viper.SetConfigName("config")    // Имя файла конфигурации (без расширения)
		viper.SetConfigType("yaml")      // Тип файла конфигурации
		viper.AddConfigPath("./config/") // Директория, где искать файл конфигурации

		viper.SetEnvPrefix("chater")
		replacer := strings.NewReplacer("", "_")
		viper.SetEnvKeyReplacer(replacer)

		// Загрузка переменных окружения (они могут переопределять значения из файла)
		viper.AutomaticEnv()

		// Загрузка конфигурации из файла
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("err reading config file: %v", err)
		}

		// Сопоставляем значения с нашей структурой конфигурации
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalf("err config decode: %v", err)
		}

		log.Println("config downloaded")
		log.Printf("config: %+v", cfg)
	})

	return &cfg
}
