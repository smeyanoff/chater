package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	Organization string
	Department   string
	Chats        []*Chat `gorm:"many2many:chat_users;"`
}
