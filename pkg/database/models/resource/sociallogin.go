package resource

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
)

// SocialLogin for social login information.
type SocialLogin struct {
	models.Base
	AppName string
	UserID  uuid.UUID
}
