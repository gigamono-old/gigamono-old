package resource

import (
	"github.com/gofrs/uuid"
	"github.com/gigamono/gigamono/pkg/database/models"
)

// Profile stores information about the user.
type Profile struct {
	models.Base
	Username    string
	FirstName   string
	LastName    string
	Email       string
	Avatar32URL string `gorm:"column:avatar_32_url"`
	UserID      uuid.UUID
}
