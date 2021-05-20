package auth

import "github.com/gigamono/gigamono/pkg/database/models"

// UserAccountCreds holds the necessary information to access a user's account.
type UserAccountCreds struct {
	models.Base
	Email        string `pg:",notnull,unique" json:"email"`
	PasswordHash string `json:"password_hash"`
}
