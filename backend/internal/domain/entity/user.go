package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	Organization string
	Department   string
	Name         string
	MiddleName   string
	Surname      string
	OwnChats     []*Chat `gorm:"foreignKey:OwnerID"`    // Явная связь с полем OwnerID в Chat
	InChats      []*Chat `gorm:"many2many:chat_users;"` // Многие-ко-многим
}
