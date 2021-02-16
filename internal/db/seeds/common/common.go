package common

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
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
