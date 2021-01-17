package models

import (
	"github.com/gofrs/uuid"
)

// App represents the application a workflow step runs.
type App struct {
	Base
	Name                string
	PublicID            uuid.UUID `gorm:"unique; type:uuid"`
	IsSecurityReviewed  bool
	IsOnAppEntityBehalf bool
	CreatorID           uuid.UUID
	AuthInfoID          uuid.UUID
	RESTHook            []RESTHook
	XRole               []*Role    `gorm:"many2many:apps_x_roles"`
	XAccount            []*Account `gorm:"many2many:apps_x_accounts"`
}
