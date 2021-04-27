package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// User stores information about the user.
type User struct {
	models.Base
	AuthUserID *uuid.UUID `gorm:"unique; type:uuid"`
	Profile    Profile
	Workspace  []Workspace  `gorm:"foreignKey:CreatorID"`
	Project    []Project    `gorm:"foreignKey:CreatorID"`
	XWorkspace []*Workspace `gorm:"many2many:users_x_workspaces"`
}
