package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Workflow stores information about a workflow.
// TODO: js_code and wasm_code columns.
type Workflow struct {
	models.Base
	Name             string     `json:"name"`
	IsActive         bool       `json:"is_active"`
	WorkflowPath     string     `json:"workflow_path"`
	ServerlessJSPath string     `json:"serverless_js_path"`
	CreatorID        *uuid.UUID `json:"creator_id"`
}
