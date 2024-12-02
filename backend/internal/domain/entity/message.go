package entities

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Content  string `gorm:"not null"`
	SenderID uint   `gorm:"index; not null"`
	ChatID   uint   `gorm:"index; not null"`
	Sender   *User  `gorm:"foreignKey:SenderID;references:ID"`
	Chat     *Chat  `gorm:"foreignKey:ChatID;references:ID"`
}
