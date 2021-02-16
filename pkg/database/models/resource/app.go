package resource

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
)

// App represents the application a workflow step runs.
type App struct {
	models.Base
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
