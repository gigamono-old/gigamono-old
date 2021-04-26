package resource

import (
	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
)

// RegisterEngineWorkflow ...
func RegisterEngineWorkflow(db *database.DB, workflow *resource.Workflow, engine *resource.Engine) error {
	return db.Model(workflow).
		Association("XEngine").
		Append(engine)
}
