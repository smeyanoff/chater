package entities

import (
	"chater/internal/domain/valueobject"

	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Name        valueobject.ChatName `gorm:"not null"`
	OwnerID     uint                 `gorm:"not null"`                                                          // Явно указываем внешний ключ (поле ID владельца)
	Owner       *User                `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Связь с пользователем
	ChatMembers []*User              `gorm:"many2many:chat_users;"`                                             // Многие-ко-многим
	Messages    []*Message           `gorm:"constraint:OnDelete:CASCADE;"`                                      // Каскадное удаление сообщений при удалении чата
}
