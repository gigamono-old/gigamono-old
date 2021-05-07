package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Workflow stores information about a workflow.
// TODO: js_code and wasm_code columns.
type Workflow struct {
	models.Base
	Name          string     `json:"name"`
	Specification string     `pg:"type:jsonb" json:"specification"`
	IsActive      bool       `json:"is_active"`
	CreatorID     *uuid.UUID `json:"creator_id"`
}