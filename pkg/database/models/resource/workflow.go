package resource

import (
	"github.com/gigamono/gigamono/pkg/configs"
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Workflow stores information about a workflow.
type Workflow struct {
	models.Base
	Name          string
	Specification configs.WorkflowConfig `pg:"type:jsonb"`
	IsActive      bool                   `json:"is_active"`
	CreatorID     *uuid.UUID             `json:"creator_id"`
}
