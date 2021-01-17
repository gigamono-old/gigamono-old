package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
)

// WorkflowInstance represents a running, paused or stopped workflow instance.
type WorkflowInstance struct {
	Base
	CurrentTaskIndex int
	Dataflow         datatypes.JSON
	WorkflowID       uuid.UUID
	EngineID         uuid.UUID
}

