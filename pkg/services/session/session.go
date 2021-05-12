package session

import (
	"github.com/gofrs/uuid"
)

// Tokens are used to identify session.
type Tokens struct {
	AccessToken  string
	RefreshToken string
}

// GetSessionUser authenticates and returns session's user id from the access token.
func GetSessionUser(tokens Tokens) (uuid.UUID, error) {
	// TODO: Sec: Auth.
	// Validate session (not pre-session) csrf token.
	// Validate session JWT tokens.
	// Return JWT claim subject. (not uuid).
	return uuid.FromString("4b523a9f-1be2-45ca-99fc-005f12581467")
}

// // ...
// func CheckOrCreatePreSession(tokens Tokens) {
// 	// TODO: ...
// 	// Check if pre-session csrf token exists
// 	// Create one if it doesn't.
// }
