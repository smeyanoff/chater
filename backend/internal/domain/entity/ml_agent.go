package entities

import "gorm.io/gorm"

type Agent struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Description string
	ModelID     uint     `gorm:"not null;index"`
	MLModel     *MLModel `gorm:"foreignKey:ModelID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Context     string
}
