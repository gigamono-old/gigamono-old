package resource

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/gofrs/uuid"
)

// CreateUserIfNotExist creates a user if user does not already exist.
func CreateUserIfNotExist(
	db *database.DB,
	sessionUserID *uuid.UUID,
) (*uuid.UUID, error) {
	user := resource.User{
		Base: models.Base{
			ID: *sessionUserID,
		},
	}

	// Insert user if not exist.
	if _, err := db.Model(&user).OnConflict("(id) DO NOTHING").Insert(); err != nil {
		return &uuid.UUID{}, fmt.Errorf("creating user in db: %v", err)
	}

	return sessionUserID, nil
}
