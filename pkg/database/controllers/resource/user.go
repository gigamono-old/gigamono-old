package resource

import (
	"github.com/sageflow/sagedb/pkg/database"
	"github.com/sageflow/sageflow/pkg/database/models/resource"
)

// CreateUnprotectedUserWithProfile creates a user without password credentials.
// Arguments are used to create Profile for user.
// SEC: Only used for dev.
func CreateUnprotectedUserWithProfile(
	db *database.DB,
	username string,
	firstName string,
	secondName string,
	email string,
	avatar32URL string,
) (resource.User, resource.Profile, error) {
	user := resource.User{}

	if err := db.Create(&user).Error; err != nil {
		return resource.User{}, resource.Profile{}, err
	}

	profile := resource.Profile{
		Username:    username,
		FirstName:   firstName,
		SecondName:  secondName,
		Email:       email,
		Avatar32URL: avatar32URL,
		UserID:      user.ID,
	}

	if err := db.Create(&profile).Error; err != nil {
		return resource.User{}, resource.Profile{}, err
	}

	return user, profile, nil
}
