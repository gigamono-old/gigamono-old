package resource

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gigamono/gigamono/internal/db/seeds/common"
	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/gofrs/uuid"
)

// LoadFakeProfiles loads fake user profiles
func LoadFakeProfiles(db *database.DB, count int) ([]uuid.UUID, error) {
	// Generate UUIDs.
	uuids, err := common.GenerateUUIDs(count * 2)
	if err != nil {
		return []uuid.UUID{}, err
	}

	// Get users.
	users, err := common.GetUsers(db, count)
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
