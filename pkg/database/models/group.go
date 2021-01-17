package models

import (
	"github.com/gofrs/uuid"
)

// Group represents a group.
type Group struct {
	Base
	Name        string
	Avatar32URL string `gorm:"column:avatar_32_url"`
	CreatorID   uuid.UUID
	XUser       []*User `gorm:"many2many:users_x_groups"`
}
