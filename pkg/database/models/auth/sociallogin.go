package auth

import (
	"github.com/gofrs/uuid"
	"github.com/gigamono/gigamono/pkg/database/models"
)

// SocialLogin for social login information.
type SocialLogin struct {
	models.Base
	AppName       string
	UserAccountID uuid.UUID
}
