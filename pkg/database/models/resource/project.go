package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Project represents a project under a workspace.
type Project struct {
	models.Base
	Name        string
	Avatar32URL string    `pg:"avatar_32_url" json:"avatar_32_url"`
	CreatorID   uuid.UUID `pg:"type:uuid" json:"creator_id"`
}
