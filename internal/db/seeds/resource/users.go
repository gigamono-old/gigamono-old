package resource

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/internal/db/seeds/common"
	"github.com/sageflow/sageflow/pkg/database"
	"github.com/sageflow/sageflow/pkg/database/models/resource"
)

// LoadFakeUsers returns fake user details.
func LoadFakeUsers(db *database.DB, count int) ([]uuid.UUID, error) {
	uuids, err := common.GenerateUUIDs(count * 2)
	if err != nil {
		return []uuid.UUID{}, err
	}

	userIDs := make([]uuid.UUID, count)

	for i := 0; i < count; i++ {
		user := resource.User{
			PasswordCredID: uuids[i],
		}

		user.ID = uuids[count+i]
		userIDs[i] = uuids[count+i]

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
