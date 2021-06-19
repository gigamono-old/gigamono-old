package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Profile stores information about the user.
type Profile struct {
	models.Base
	Username    string      `json:"username"`
	FirstName   string      `json:"first_name"`
	LastName    string      `json:"last_name"`
	Email       string      `json:"email"`
	AvatarURL   string      `json:"avatar_url"`
	UserID      uuid.UUID   `pg:"type:uuid,notnull" json:"user_id"`
}
