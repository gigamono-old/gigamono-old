package auth

import "github.com/gigamono/gigamono/pkg/database/models"

// UserAccountAccess holds the necessary information to access a user's account.
type UserAccountAccess struct {
	models.Base
	Email       string `json:"email"`
	EncPassword string `json:"enc_password"`
}
