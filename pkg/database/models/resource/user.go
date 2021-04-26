package resource

import (
	"github.com/gofrs/uuid"
	"github.com/gigamono/gigamono/pkg/database/models"
)

// User stores information about the user.
type User struct {
	models.Base
	AuthUserID   *uuid.UUID `gorm:"unique; type:uuid"`
	Profile      Profile
	RefreshToken string // JWT.1.R
	Account      []Account
	RESTHook     []RESTHook
	AppID        []App        `gorm:"foreignKey:CreatorID"`
	Group        []Group      `gorm:"foreignKey:CreatorID"`
	Workflow     []Workflow   `gorm:"foreignKey:CreatorID"`
	Workspace    []Workspace  `gorm:"foreignKey:CreatorID"`
	Folder       []Folder     `gorm:"foreignKey:CreatorID"`
	Theme        []Theme      `gorm:"foreignKey:CreatorID"`
	XGroup       []*Group     `gorm:"many2many:users_x_groups"`
	XWorkspace   []*Workspace `gorm:"many2many:users_x_workspaces"`
	XRole        []*Role      `gorm:"many2many:users_x_roles"`
}
