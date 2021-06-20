package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Base stores information about a base.
// No specification file is stored. Instead it is generated from model.
type Base struct {
	models.Base
	Name      string    `json:"name"`
	CreatorID uuid.UUID `json:"creator_id"`
	SpaceID   uuid.UUID `pg:"type:uuid" json:"space_id"`
	Tables    []Table   `pg:"rel:has-many" json:"tables"`
}
