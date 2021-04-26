package resource

import (
	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
)

// CreateEngine creates an engine.
func CreateEngine(db *database.DB) (resource.Engine, error) {
	engine := resource.Engine{}
	if err := db.Create(&engine); err != nil {
		return resource.Engine{}, nil
	}
	return engine, nil
}
