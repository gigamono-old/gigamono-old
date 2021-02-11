package controllers

import "github.com/sageflow/sageflow/pkg/database/models"

// RegisterEngineWorkflow ...
func (db *DB) RegisterEngineWorkflow(workflow *models.Workflow, engine *models.Engine) error {
	return db.Model(workflow).
		Association("XEngine").
		Append(engine)
}
