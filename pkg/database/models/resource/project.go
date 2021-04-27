package resource

import (
	"github.com/gofrs/uuid"
	"github.com/gigamono/gigamono/pkg/database/models"
)

// Project represents a project under a workspace.
type Project struct {
	models.Base
	Name        string
	Avatar32URL string `gorm:"column:avatar_32_url"`
	CreatorID   uuid.UUID
	WorkspaceID   uuid.UUID
}
