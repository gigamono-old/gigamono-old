package seeds

import (
	"github.com/sageflow/sageflow/pkg/database/models"
)

// FakeUsers returns fake user details.
func FakeUsers(count int) ([]models.User, error) {
	uuids, err := generateUUIDs(count * 2)
	if err != nil {
		return []models.User{}, err
	}

	users := make([]models.User, count)
	for i := 0; i < count; i++ {
		users[i] = models.User{
			PasswordCredID: uuids[i],
		}

		users[i].ID = uuids[count+i]
	}

	return users, nil
}
