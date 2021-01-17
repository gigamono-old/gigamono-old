package models

import (
	"github.com/gofrs/uuid"
)

// Profile stores information about the user.
type Profile struct {
	Base
	Username    string
	FirstName   string
	SecondName  string
	Email       string
	Avatar32URL string `gorm:"column:avatar_32_url"`
	UserID      uuid.UUID
}
