package resource

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
	"gorm.io/datatypes"
)

// Theme represents a theme used to change the look of the application.
type Theme struct {
	models.Base
	Name      string
	Code      datatypes.JSON
	PublicID  uuid.UUID `gorm:"unique; type:uuid"`
	CreatorID uuid.UUID
	XRole     []*Role `gorm:"many2many:themes_x_roles"`
}
