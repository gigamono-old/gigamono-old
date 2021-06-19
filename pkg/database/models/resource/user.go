package resource

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models"
)

// User stores information about a user.
type User struct {
	models.Base
	Profile      Profile       `pg:"rel:belongs-to" json:"profile"`
	Workspaces   []Workspace   `pg:"rel:has-many,join_fk:creator_id" json:"workspaces"`
	Spaces       []Space       `pg:"rel:has-many,join_fk:creator_id" json:"spaces"`
	Integrations []Integration `pg:"rel:has-many,join_fk:creator_id" json:"integrations"`
	XWorkspaces  []Workspace   `pg:"many2many:x_users_workspaces" json:"-"`
}

// CreateIfNotExist creates a user if user does not already exist.
func (user *User) CreateIfNotExist(db *database.DB) error {
	// Insert user if not exist.
	if _, err := db.Model(user).OnConflict("(id) DO NOTHING").Insert(); err != nil {
		return fmt.Errorf("creating user in db: %v", err)
	}

	return nil
}
