// cmd/chater/main.go
package main

import (
	"chater/internal/api"
	"chater/internal/config"
	"chater/internal/domain/service"
	"chater/internal/infrastructure/db"
	"chater/internal/infrastructure/repository"

	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to DB
	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Автоматическая миграция
	if err := db.AutoMigrate(database); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	userRepo := repository.NewGormUserRepository(database)
	authService := service.NewAuthService(userRepo, cfg.Auth.JWT.Key, cfg.Auth.JWT.ExpirationTimeH)
	authHandler := api.NewAuthHandler(authService)

	http.HandleFunc("/register", authHandler.Register)
	http.HandleFunc("/login", authHandler.Login)

	log.Printf("App is running on port %s", cfg.App.Port)
	log.Fatal(http.ListenAndServe(cfg.App.Host+":"+cfg.App.Port, nil))
}
