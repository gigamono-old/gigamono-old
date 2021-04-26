package resource

import (
	"errors"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofrs/uuid"
	"github.com/gigamono/gigamono/internal/db/seeds/common"
	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
)

// LoadFakeProfiles loads fake user profiles
func LoadFakeProfiles(db *database.DB, count int) ([]uuid.UUID, error) {
	// Generate UUIDs.
	uuids, err := common.GenerateUUIDs(count * 2)
	if err != nil {
		return []uuid.UUID{}, err
	}

	// Get details.
	users, err := generateUsers(db, count)
	if err != nil {
		return []uuid.UUID{}, err
	}

	profileIDs := make([]uuid.UUID, count)
	faker := gofakeit.NewCrypto()

	for i := 0; i < count; i++ {
		profile := resource.Profile{
			Username:    faker.Username(),
			FirstName:   faker.FirstName(),
			LastName:    faker.LastName(),
			Email:       faker.Email(),
			Avatar32URL: faker.URL(),
			UserID:      users[i].ID,
		}

		profile.ID = uuids[i]
		profileIDs[i] = profile.ID

		// Load item.
		if err := db.Create(&profile).Error; err != nil {
			return []uuid.UUID{}, err
		}

		// Load seed.
		if err := db.Create(&common.Seed{
			TableName: db.GetTableName(profile),
			SeedID:    profile.ID,
		}).Error; err != nil {
			return []uuid.UUID{}, err
		}
	}

	return profileIDs, nil
}

func generateUsers(db *database.DB, count int) ([]resource.User, error) {
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
