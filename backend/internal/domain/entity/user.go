package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	Email      string `gorm:"unique;not null"`
	Name       string
	MiddleName string
	Surname    string
	OwnChats   []*Chat  `gorm:"foreignKey:OwnerID"`
	InChats    []*Chat  `gorm:"many2many:chat_users;"`
	InGroups   []*Group `gorm:"many2many:group_users;"`
}
