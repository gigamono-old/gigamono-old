package auth

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models"
)

// UserAccountCreds holds the necessary information to access a user's account.
type UserAccountCreds struct {
	models.Base
	Email        string `pg:",notnull,unique" json:"email"`
	PasswordHash string `json:"password_hash"`
}

// Create creates a user account credentials.
func (creds *UserAccountCreds) Create(db *database.DB) error {
	// Insert user account creds and return id.
	if _, err := db.Model(creds).Returning("*").Insert(); err != nil {
		return fmt.Errorf("creating user account creds in db: %v", err)
	}

	return nil
}

// GetByEmail gets user account credentials by email.
func (creds *UserAccountCreds) GetByEmail(db *database.DB) error {
	// Get the user by email.
	if err := db.Model(creds).Where("email = ?", creds.Email).Select(); err != nil {
		return err
	}

	return nil
}
