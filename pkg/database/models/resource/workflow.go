package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Workflow stores information about a workflow.
type Workflow struct {
	models.Base
	Name                      string     `json:"name"`
	IsActive                  bool       `json:"is_active"`
	SpecificationFileURL      string     `json:"specification_file_url"`
	OutputServerlessJSFileURL string     `json:"output_serverless_js_url"`
	CreatorID                 *uuid.UUID `json:"creator_id"`
}
