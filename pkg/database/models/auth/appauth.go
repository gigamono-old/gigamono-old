package auth

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
	"gorm.io/datatypes"
)

// AppAuth represents information an application needs for administrative authentication.
type AppAuth struct {
	models.Base
	Name  string
	Code  datatypes.JSON
	AppID uuid.UUID
}

