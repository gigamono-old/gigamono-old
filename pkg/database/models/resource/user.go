package resource

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models"
)

// User stores information about a user.
type User struct {
	models.Base
	Profile                   Profile       `pg:"rel:belongs-to" json:"profile"`
	Workspaces                []Workspace   `pg:"rel:has-many,join_fk:creator_id" json:"workspaces"`
	Spaces                    []Space       `pg:"rel:has-many,join_fk:creator_id" json:"spaces"`
	Decks                     []Deck        `pg:"rel:has-many,join_fk:creator_id" json:"decks"`
	Automations               []Automation  `pg:"rel:has-many,join_fk:creator_id" json:"automations"`
	Bases                     []Base        `pg:"rel:has-many,join_fk:creator_id" json:"bases"`
	Boards                    []Board       `pg:"rel:has-many,join_fk:creator_id" json:"boards"`
	Workflows                 []Workflow    `pg:"rel:has-many,join_fk:creator_id" json:"workflows"`
	Tables                    []Table       `pg:"rel:has-many,join_fk:creator_id" json:"tables"`
	Integrations              []Integration `pg:"rel:has-many,join_fk:creator_id" json:"integrations"`
	XUserWorkspaceMemberships []Workspace   `pg:"many2many:x_user_workspace_memberships" json:"-"`
}

// Exists user with specified id exists.
func (user *User) Exists(db *database.DB) (bool, error) {
	// TODO: Sec: Permission.
	// Check if user exists.
	val, err := db.Model(user).WherePK().Exists()
	if err != nil {
		return val, fmt.Errorf("fetching user from db: %v", err)
	}

	return val, nil
}

// Create creates a user.
func (user *User) Create(db *database.DB) error {
	// Insert user.
	if _, err := db.Model(user).Insert(); err != nil {
		return fmt.Errorf("creating user in db: %v", err)
	}

	return nil
}

// GetByID gets a user by id.
func (user *User) GetByID(db *database.DB) error {
	// TODO: Sec: Permission.
	// Select the user with the specified user ID.
	if err := db.Model(user).WherePK().Relation("Profile").Select(); err != nil {
		return fmt.Errorf("fetching user from db: %v", err)
	}

	return nil
}

// GetPreferencesByID gets a user's preferences.
func (user *User) GetPreferencesByID(db *database.DB) (*Preferences, error) {
	// TODO: Sec: Permission.
	// Select the user's preferences.
	preferences := Preferences{}
	if err := db.Model(&preferences).Join("JOIN users ON users.id = preferences.user_id").First(); err != nil {
		return &preferences, fmt.Errorf("fetching user's preferences from db: %v", err)
	}

	return &preferences, nil
}
