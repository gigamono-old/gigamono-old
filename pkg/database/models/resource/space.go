package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Space represents a sectio under a workspace.
type Space struct {
	models.Base
	Name        string    `json:"name"`
	Avatar32URL string    `pg:"avatar_32_url" json:"avatar_32_url"`
	CreatorID   uuid.UUID `pg:"type:uuid" json:"creator_id"`
}
