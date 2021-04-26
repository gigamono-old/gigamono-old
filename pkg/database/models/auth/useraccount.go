package auth

import (
	"github.com/gofrs/uuid"
	"github.com/gigamono/gigamono/pkg/database/models"
)

// UserAccount stores information for authenticating a user.
type UserAccount struct {
	models.Base
	ResourceUserID    uuid.UUID
	EncryptedPassword string
	Username          string
	Email             string
	RefreshToken      string
	SocialLogin       SocialLogin
}
