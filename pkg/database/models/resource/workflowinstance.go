package resource

import (
	"github.com/gofrs/uuid"
	"github.com/gigamono/gigamono/pkg/database/models"
	"gorm.io/datatypes"
)

// WorkflowInstance represents a running, paused or stopped workflow instance.
type WorkflowInstance struct {
	models.Base
	CurrentTaskIndex int
	Dataflow         datatypes.JSON
	WorkflowID       uuid.UUID
	EngineID         uuid.UUID
}
