package resource

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database"
	"github.com/sageflow/sageflow/pkg/database/models/resource"
	"gorm.io/gorm"
)

// CreateUserWithProfile creates a user without password credentials.
// Arguments are used to create Profile for user.
func CreateUserWithProfile(
	db *database.DB,
	username string,
	firstName string,
	lastName string,
	email string,
	avatar32URL string,
	authUserID uuid.UUID,
	refreshToken string,
) (resource.User, resource.Profile, error) {
	user := resource.User{
		AuthUserID:   &authUserID,
		RefreshToken: refreshToken,
	}

	profile := resource.Profile{
		Username:    username,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		Avatar32URL: avatar32URL,
		UserID:      user.ID,
	}

	// Add user and profile entries. Reverse if there is an error.
	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		return tx.Create(&profile).Error
	}); err != nil {
		return resource.User{}, resource.Profile{}, err
	}

	return user, profile, nil
}
