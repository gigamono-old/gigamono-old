package resource

import (
	"github.com/sageflow/sagedb/pkg/database"
	"github.com/sageflow/sageflow/pkg/database/models/resource"
)

// WriteLog writes a log to the database.
func WriteLog(db *database.DB) (resource.Log, error) {
	log := resource.Log{}
	if err := db.Create(&log); err != nil {
		return resource.Log{}, nil
	}
	return log, nil
}
