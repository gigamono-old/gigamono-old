package resource

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Profile stores information about the user.
type Profile struct {
	models.Base
	Username  *string   `json:"username"`
	FirstName *string   `json:"first_name"`
	LastName  *string   `json:"last_name"`
	Email     *string   `json:"email"`
	AvatarURL *string   `json:"avatar_url"`
	UserID    uuid.UUID `pg:"type:uuid,notnull" json:"user_id"`
}

// Create creates a profile.
func (profile *Profile) Create(db *database.DB) error {
	// TODO: Sec: Permission.
	// Insert profile in db.
	if _, err := db.Model(profile).Insert(); err != nil {
		return fmt.Errorf("creating profile in db: %v", err)
	}

	return nil
}

// UpdateByID updates all the fields of a profile by id.
func (profile *Profile) UpdateByID(db *database.DB) error {
	// TODO: Sec: Permission.
	// Update profile in db.
	if _, err := db.
		Model(profile).
		Column( // Sec: columns allowed to be updated.
			"username",
			"first_name",
			"last_name",
			"email",
			"avatar_url",
		).
		WherePK().
		Returning("*").
		Update(); err != nil {
		return fmt.Errorf("updating profile in db: %v", err)
	}

	return nil
}

// GetByID gets an profile by id.
func (profile *Profile) GetByID(db *database.DB) error {
	// TODO: Sec: Permission.
	// Select the profile with the specified profile ID.
	if err := db.Model(profile).WherePK().Select(); err != nil {
		return fmt.Errorf("fetching profile from db: %v", err)
	}

	return nil
}
