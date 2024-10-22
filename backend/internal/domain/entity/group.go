package entities

import (
	"chater/internal/domain/valueobject"

	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Name         valueobject.GroupName `gorm:"not null"`
	OwnerID      uint                  `gorm:"not null"`
	GroupOwner   *User                 `gorm:"foreignKey:OwnerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	GroupMembers []*User               `gorm:"many2many:group_users;"`
}
