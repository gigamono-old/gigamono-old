package common

import (
	"errors"

	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/gofrs/uuid"
)

// Seed represents a seed in the database.
type Seed struct {
	models.Base
	TableName string
	SeedID    uuid.UUID `gorm:"type:uuid"`
}

// GenerateUUIDs generates random UUIDs.
func GenerateUUIDs(count int) ([]uuid.UUID, error) {
	var err error
	uuids := make([]uuid.UUID, count)

	for i := 0; i < count; i++ {
		uuids[i], err = uuid.NewV4()
		if err != nil {
			return []uuid.UUID{}, err
		}
	}

	return uuids, nil
}

// GetUsers gets users from the database.
func GetUsers(db *database.DB, count int) ([]resource.User, error) {
	var users []resource.User

	// Get users from the db.
	if err := db.Limit(count).Find(&users).Error; err != nil {
		return []resource.User{}, err
	}

	// Check if users from the db match count.
	if len(users) < count {
		return []resource.User{}, errors.New("seeding: users generated from the db is not up to expected `count`")
	}

	return users, nil
}
