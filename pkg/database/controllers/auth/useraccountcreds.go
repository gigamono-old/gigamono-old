package auth

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models/auth"
	"github.com/gofrs/uuid"
)

// CreateUserAccountCreds creates a user account credentials.
func CreateUserAccountCreds(db *database.DB, email string, passwordHash string) (*uuid.UUID, error) {
	credentials := auth.UserAccountCreds{
		Email:        email,
		PasswordHash: passwordHash,
	}

	// Insert user account creds and return id.
	if _, err := db.Model(&credentials).Returning("id").Insert(); err != nil {
		return &uuid.UUID{}, fmt.Errorf("creating user account creds in db: %v", err)
	}

	return &credentials.ID, nil
}

// GetUserAccountCreds gets user account credentials by email.
func GetUserAccountCreds(db *database.DB, email string) (*auth.UserAccountCreds, error) {
	credentials := auth.UserAccountCreds{}

	// Get the user by email.
	if err := db.Model(&credentials).Where("email = ?", email).Select(); err != nil {
		return &auth.UserAccountCreds{}, err
	}

	return &credentials, nil
}
