package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Workflow stores information about a workflow.
type Workflow struct {
	models.Base
	Name              string     `json:"name"`
	IsActive          bool       `json:"is_active"`
	SpecificationPath string     `pg:"specification_path" json:"specification_path"`
	ServerlessJSPath  string     `pg:"serverless_js_path" json:"serverless_js_path"`
	CreatorID         *uuid.UUID `json:"creator_id"`
}
