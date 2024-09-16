package entities

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Content  string
	AuthorID uint
	Author   User
	ChatID   uint
}
