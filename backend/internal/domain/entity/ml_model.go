package entities

import "gorm.io/gorm"

type MLModel struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Description string
	BaseURL     string `gorm:"not null"`
}
