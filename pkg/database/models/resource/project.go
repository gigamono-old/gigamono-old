package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Project represents a project under a space.
type Project struct {
	models.Base
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatar_url"`
	CreatorID uuid.UUID `pg:"type:uuid" json:"creator_id"`
}
