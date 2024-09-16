// cmd/chater/main.go
package main

import (
	"chater/internal/api"
	"chater/internal/config"
	"chater/internal/infrastructure/db"
	"chater/internal/infrastructure/repository"
	"chater/internal/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "chater/docs"

	"log"
)

// @title ChatGPT Backend API
// @version 1.0
// @description This is a sample server for a chat backend with JWT authentication.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:54321
// @BasePath /
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

	// Используем Gin для роутинга
	r := gin.Default()

	// Настройка CORS через данные из конфигурации
	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORS.AllowOrigins,
		AllowMethods:     cfg.CORS.AllowMethods,
		AllowHeaders:     cfg.CORS.AllowHeaders,
		AllowCredentials: cfg.CORS.AllowCredentials,
		MaxAge:           12 * time.Hour,
	}))

	// users
	userRepo := repository.NewGormUserRepository(database)
	authService := service.NewAuthService(userRepo, cfg.Auth.JWTKey, cfg.Auth.JWTKeyExpirationTimeH)
	authHandler := api.NewAuthController(authService)

	r.POST("auth/register", authHandler.Register)
	r.POST("auth/login", authHandler.Login)

	// chats
	chatRepo := repository.NewGormChatRepository(database)
	chatService := service.NewChatService(chatRepo)
	chatController := api.NewChatController(chatService)

	r.GET("chats/:user_id", chatController.GetChatsByUserIDHandler)

	// Настройка маршрута для Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск HTTP-сервера
	log.Printf("Server is running on port %s", cfg.App.Port)
	log.Fatal(r.Run(":" + cfg.App.Port))
}
