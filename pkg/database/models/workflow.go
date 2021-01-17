package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
)

// Workflow represents a workflow.
type Workflow struct {
	Base
	Name             string
	Code             datatypes.JSON
	IsActive         bool
	IsDraft          bool
	FolderID         uuid.UUID
	CreatorID        uuid.UUID
	WorkflowInstance []WorkflowInstance
	XRole            []*Role   `gorm:"many2many:workflows_x_roles"`
	XEngine          []*Engine `gorm:"many2many:workflows_x_engines"`
}
