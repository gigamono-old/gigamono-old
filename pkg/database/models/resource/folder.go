package resource

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
)

// Folder represents the folder workflows and the likes can be kept.
type Folder struct {
	models.Base
	Name        string
	Description string
	Avatar32URL string `gorm:"column:avatar_32_url"`
	CreatorID   uuid.UUID
	Workflow    []Workflow
}
