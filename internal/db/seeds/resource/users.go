package resource

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/internal/db/seeds/common"
	"github.com/sageflow/sageflow/pkg/database"
	"github.com/sageflow/sageflow/pkg/database/models/resource"
)

// LoadFakeUsers loads fake user details.
func LoadFakeUsers(db *database.DB, count int) ([]uuid.UUID, error) {
	// Generate UUIDs.
	uuids, err := common.GenerateUUIDs(count * 2)
	if err != nil {
		return []uuid.UUID{}, err
	}

	userIDs := make([]uuid.UUID, count)
	passwordCredIDs := make([]*uuid.UUID, count) // TODO. Add actual passwords.

	for i := 0; i < count; i++ {
		user := resource.User{
			PasswordCredID: passwordCredIDs[i],
		}

		user.ID = uuids[i]
		userIDs[i] = user.ID

		// Load item.
		if err := db.Create(&user).Error; err != nil {
			return []uuid.UUID{}, err
		}

		// Load seed.
		if err := db.Create(&common.Seed{
			TableName: db.GetTableName(user),
			SeedID:    user.ID,
		}).Error; err != nil {
			return []uuid.UUID{}, err
		}
	}

	return userIDs, nil
}
