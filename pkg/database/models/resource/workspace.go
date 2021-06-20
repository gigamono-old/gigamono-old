package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Workspace represents a workspace.
type Workspace struct {
	models.Base
	Name      string    `json:"name"`
	AvatarURL string    `json:"avatar_url"`
	CreatorID uuid.UUID `pg:"type:uuid" json:"creator_id"`
	Spaces    []Space   `pg:"rel:has-many" json:"spaces"`
	XUsers    []User    `pg:"many2many:x_users_workspaces" json:"-"`
}
