package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Automation stores information about an automation.
// No specification file is stored. Instead it is generated from model.
type Automation struct {
	models.Base
	Name                      string     `json:"name"`
	OutputServerlessJSFileURL string     `json:"output_serverless_js_file_url"`
	CreatorID                 uuid.UUID  `json:"creator_id"`
	SpaceID                   uuid.UUID  `pg:"type:uuid" json:"space_id"`
	Workflows                 []Workflow `pg:"rel:has-many" json:"workflows"`
}
