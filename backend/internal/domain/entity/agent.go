package entities

import "gorm.io/gorm"

type Agent struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Description string
	BaseURL     string        `gorm:"not null"`                                        // Базовый URL агента
	Routes      []*AgentRoute `gorm:"foreignKey:AgentID;constraint:OnDelete:CASCADE;"` // Связанные маршруты
}

type AgentRoute struct {
	gorm.Model
	AgentID    uint   `gorm:"not null"` // Ссылка на агента
	Name       string `gorm:"not null"` // Имя эндпоинта, например, "respond" или "billing"
	Path       string `gorm:"not null"` // Путь, например, "/respond" или "/billing"
	HttpMethod string `gorm:"not null"` // Метод HTTP, например, "POST", "GET"
}
