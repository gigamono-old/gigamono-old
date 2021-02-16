package resource

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
)

// Profile stores information about the user.
type Profile struct {
	models.Base
	Username    string
	FirstName   string
	SecondName  string
	Email       string
	Avatar32URL string `gorm:"column:avatar_32_url"`
	UserID      uuid.UUID
}
