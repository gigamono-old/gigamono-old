package resource

import (
	"github.com/gofrs/uuid"
	"github.com/gigamono/gigamono/pkg/database/models"
	"gorm.io/datatypes"
)

// Workflow represents a workflow.
type Workflow struct {
	models.Base
	Name             string
	Code             datatypes.JSON
	IsActive         bool
	IsDraft          bool
	FolderID         *uuid.UUID
	CreatorID        *uuid.UUID
	WorkflowInstance []WorkflowInstance
	XRole            []*Role   `gorm:"many2many:workflows_x_roles"`
	XEngine          []*Engine `gorm:"many2many:workflows_x_engines"`
}
