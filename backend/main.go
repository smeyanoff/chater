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
	_ "chater/internal/domain/valueobject"

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

// @securityDefinitions.apikey BearerAuth
// @in cookie
// @name token

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

	// Repos
	chatRepo := repository.NewGormChatRepository(database)
	userRepo := repository.NewGormUserRepository(database)
	messageRepo := repository.NewGormMessageRepository(database)
	groupRepo := repository.NewGormGroupRepository(database)

	// Serices
	authService := service.NewAuthService(userRepo, cfg.Auth.JWTKey, cfg.Auth.JWTKeyExpirationTimeH)
	chatService := service.NewChatService(chatRepo, userRepo, groupRepo)
	messageService := service.NewMessageService(messageRepo)
	groupService := service.NewGroupService(groupRepo, userRepo)

	// Controllers
	authHandler := api.NewAuthController(authService)
	chatController := api.NewChatController(chatService)
	messageController := api.NewMessageController(messageService)
	groupController := api.NewGroupController(groupService)

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

	apiV1 := r.Group("/v1")

	// Настройка маршрута для Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// auth

	apiV1.POST("/auth/register", authHandler.Register)
	apiV1.POST("/auth/login", authHandler.Login)

	// middelware secured JWT endpoints
	jwtSecret := cfg.Auth.JWTKey
	auth := api.JWTAuthMiddleware(jwtSecret)
	apiV1.Use(auth) // Connect middleware

	// chats

	apiV1.GET("/chats", chatController.GetChatsForUser)
	apiV1.POST("/chats", chatController.CreateChat)

	// messages

	apiV1.GET("/chats/:chat_id/messages/ws", messageController.MessageWebSocketController)
	apiV1.GET("/chats/:chat_id/messages", messageController.GetMessages)

	// groups

	apiV1.POST("/groups", groupController.CreateGroup)
	apiV1.GET("/groups", groupController.GetAllUserGroups)
	apiV1.DELETE("/groups/:group_id", groupController.DeleteGroup)
	apiV1.POST("/groups/:group_id/users", groupController.AddUserToGroup)
	apiV1.DELETE("/groups/:group_id/users", groupController.DeleteUserFromGroup)

	// Запуск HTTP-сервера
	log.Printf("Server is running on port %s", cfg.App.Port)
	log.Fatal(r.Run(":" + cfg.App.Port))
}
