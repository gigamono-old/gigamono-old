package resource

import (
	"github.com/gofrs/uuid"
	"github.com/gigamono/gigamono/pkg/database/models"
)

// Group represents a group.
type Group struct {
	models.Base
	Name        string
	Avatar32URL string `gorm:"column:avatar_32_url"`
	CreatorID   uuid.UUID
	XUser       []*User `gorm:"many2many:users_x_groups"`
}
