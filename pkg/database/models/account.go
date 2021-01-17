package models

import (
	"github.com/gofrs/uuid"
)

// Account represents credentials needed for an app to authorize access.
type Account struct {
	Base
	UserID            uuid.UUID
	AccessTokenCredID uuid.UUID `gorm:"unique; type:uuid"`
	XApp              []*App    `gorm:"many2many:apps_x_accounts"`
}
