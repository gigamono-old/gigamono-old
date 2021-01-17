package models

// Engine represents an engine.
type Engine struct {
	Base
	WorkflowInstance []WorkflowInstance
	XWorkflow        []*Workflow      `gorm:"many2many:workflows_x_engines"`
}

