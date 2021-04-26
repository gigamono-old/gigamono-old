package resource

import (
	"github.com/gofrs/uuid"
	"github.com/gigamono/gigamono/pkg/database/models"
)

// Account represents credentials needed for an app to authorize access.
type Account struct {
	models.Base
	UserID            uuid.UUID
	AuthAccessTokenID uuid.UUID `gorm:"unique; type:uuid"`
	XApp              []*App    `gorm:"many2many:apps_x_accounts"`
}
