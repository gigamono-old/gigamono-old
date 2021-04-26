package auth

import (
	"github.com/gofrs/uuid"
	"github.com/gigamono/gigamono/pkg/database/models"
)

// ClientAccount stores information for authenticating a client.
type ClientAccount struct {
	models.Base
	ClientID              uuid.UUID // Public-facing ID // SecureRandom.hex(32)
	EncryptedClientSecret string    // For verifying client // UUID
	Kind                  string    // Confidential or Public
	RedirectURI           string
	IsFirstParty          bool
}
