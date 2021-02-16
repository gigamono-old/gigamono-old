package auth

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
)

// Password represents a user's encrypted password.
type Password struct {
	models.Base
	OwnerID           uuid.UUID
	EncryptedPassword string
}
