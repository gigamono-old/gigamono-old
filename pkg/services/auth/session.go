package auth

import (
	"github.com/gofrs/uuid"
)

// Tokens are used to identify session.
type Tokens struct {
	AccessToken  string
	RefreshToken string
}

// GetSessionUserID authenticates and returns session's user id from the access token.
func GetSessionUserID(tokens Tokens) (uuid.UUID, error) {
	// TODO: Sec: Auth.
	return uuid.FromString("4b523a9f-1be2-45ca-99fc-005f12581467")
}
