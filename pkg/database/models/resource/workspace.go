package resource

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
)

// Workspace represents a workspace.
type Workspace struct {
	models.Base
	Name        string
	Avatar32URL string `gorm:"column:avatar_32_url"`
	CreatorID   uuid.UUID
	XUser       []*User `gorm:"many2many:users_x_workspaces"`
}
