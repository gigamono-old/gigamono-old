package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
)

// User stores information about the user.
type User struct {
	models.Base
	Profile     Profile     `pg:"rel:belongs-to"`
	Workspaces  []Workspace `pg:"rel:has-many, join_fk:creator_id"`
	Projects    []Project   `pg:"rel:has-many, join_fk:creator_id"`
	XWorkspaces []Workspace `pg:"many2many:x_users_workspaces" json:"-"`
}
