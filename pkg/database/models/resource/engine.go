package resource

import "github.com/sageflow/sageflow/pkg/database/models"

// Engine represents an engine.
type Engine struct {
	models.Base
	WorkflowInstance []WorkflowInstance
	XWorkflow        []*Workflow `gorm:"many2many:workflows_x_engines"`
}
