package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Board stores information about a board.
type Board struct {
	models.Base
	Name                 string    `json:"name"`
	SpecificationFileURL string    `json:"specification_file_url"`
	CreatorID            uuid.UUID `json:"creator_id"`
	DeckID               uuid.UUID `json:"deck_id"`
}
