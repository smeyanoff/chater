package entities

import (
	"chater/internal/domain/valueobject"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Name        *valueobject.ChatName `gorm:"not null"`
	ChatOwner   User                  `gorm:"foreignKey:ChatOwnerId"`
	Users       []*User               `gorm:"many2many:chat_users;"`
	ChatOwnerId uint
	Messages    []*Message
}
