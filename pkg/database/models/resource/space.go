package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Space represents a sectio under a workspace.
type Space struct {
	models.Base
	Name        string       `json:"name"`
	AvatarURL   string       `json:"avatar_url"`
	CreatorID   uuid.UUID    `pg:"type:uuid" json:"creator_id"`
	WorkspaceID uuid.UUID    `pg:"type:uuid" json:"workspace_id"`
	Decks       []Deck       `pg:"rel:has-many" json:"decks"`
	Automations []Automation `pg:"rel:has-many" json:"automations"`
	Bases       []Base       `pg:"rel:has-many" json:"bases"`
}
