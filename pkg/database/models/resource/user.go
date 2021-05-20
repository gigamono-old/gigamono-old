package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
)

// User stores information about a user.
type User struct {
	models.Base
	Profile     Profile     `pg:"rel:belongs-to" json:"profile"`
	Workflow    Workflow    `pg:"rel:belongs-to,join_fk:creator_id" json:"workflow"`
	Workspaces  []Workspace `pg:"rel:has-many,join_fk:creator_id" json:"workdpaces"`
	Projects    []Project   `pg:"rel:has-many,join_fk:creator_id" json:"projects"`
	XWorkspaces []Workspace `pg:"many2many:x_users_workspaces" json:"-"`
}
