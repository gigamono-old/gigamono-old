package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Deck stores information about a deck.
// No specification file is stored. Instead it is generated from model.
type Deck struct {
	models.Base
	Name                  string    `json:"name"`
	OutputBundleFolderURL string    `json:"output_bundle_folder_url"`
	CreatorID             uuid.UUID `pg:"type:uuid" json:"creator_id"`
	SpaceID               uuid.UUID `pg:"type:uuid" json:"space_id"`
	Boards                []Board   `pg:"rel:has-many" json:"boards"`
}
