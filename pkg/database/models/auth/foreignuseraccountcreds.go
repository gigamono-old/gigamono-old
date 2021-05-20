package auth

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// ForeignUserAccountCreds holds the necessary information to access user's third-party resources.
type ForeignUserAccountCreds struct {
	models.Base
	UserID          uuid.UUID `json:"user_id"`
	IntegrationID   uuid.UUID `json:"integration_id"`
	EncAccessToken  string    `json:"enc_access_token"`
	EncRefreshToken string    `json:"enc_refresh_token"`
}
