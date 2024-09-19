package entities

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Content  string
	SenderID uint  `gorm:"index"` // Индекс для SenderID
	ChatID   uint  `gorm:"index"` // Индекс для ChatID
	Sender   *User `gorm:"foreignKey:SenderID;references:ID"`
	Chat     *Chat `gorm:"foreignKey:ChatID;references:ID"`
}
