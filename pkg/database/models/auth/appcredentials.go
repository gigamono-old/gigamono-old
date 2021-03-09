package auth

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
	"gorm.io/datatypes"
)

// AppCredentials represents information an application needs for administrative authentication.
type AppCredentials struct {
	models.Base
	Name  string
	Code  datatypes.JSON
	AppID uuid.UUID
}
