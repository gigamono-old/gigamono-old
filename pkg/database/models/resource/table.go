package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Table stores information about a table in a base.
type Table struct {
	models.Base
	Name                 string    `json:"name"`
	SpecificationFileURL string    `json:"specification_file_url"`
	CreatorID            uuid.UUID `json:"creator_id"`
	BaseID               uuid.UUID `json:"base_id"`
}
