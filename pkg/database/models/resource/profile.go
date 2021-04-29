package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Profile stores information about the user.
type Profile struct {
	models.Base
	Username    string
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string
	Avatar32URL string    `pg:"avatar_32_url" json:"avatar_32_url"`
	UserID      uuid.UUID `pg:"type:uuid, notnull" json:"user_id"`
}
