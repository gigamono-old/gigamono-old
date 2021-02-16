package auth

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
)

// AccessToken represents the user's access token gotten from OAuth flow.
type AccessToken struct {
	models.Base
	AppID                uuid.UUID
	OwnerID              uuid.UUID
	EncryptedAccessToken string // Sec: Encrypted with AUTH_SECRET_KEY
}
