package resource

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Workflow stores information about a workflow.
type Workflow struct {
	models.Base
	Name                      string     `json:"name"`
	IsActive                  bool       `json:"is_active"`
	SpecificationFileURL      string     `json:"specification_file_url"`
	OutputServerlessJSFileURL string     `json:"output_serverless_js_url"`
	CreatorID                 *uuid.UUID `json:"creator_id"`
}

// Create creates a workflow.
func (workflow *Workflow) Create(db *database.DB) error {
	// TODO: Sec: Permission.
	// Insert workflow and return id.
	if _, err := db.Model(workflow).Insert(); err != nil {
		return fmt.Errorf("creating workflow in db: %v", err)
	}

	return nil
}

// ActivateByID updates the `is_active` field of a workflow selected by id.
func (workflow *Workflow) ActivateByID(db *database.DB) error {
	// TODO: Sec: Permission.
	workflow.IsActive = true

	// Update details specified workflow.
	if _, err := db.Model(workflow).Set("is_active = ?is_active").Where("id = ?", workflow.ID).Update(); err != nil {
		return fmt.Errorf("activating workflow in db: %v", err)
	}

	return nil
}

// GetByID gets a workflow by id.
func (workflow *Workflow) GetByID(db *database.DB) error {
	// TODO: Sec: Permission.
	// Select the workflow with the specified workflow ID.
	if err := db.Model(workflow).Where("id = ?", workflow.ID).Select(); err != nil {
		return fmt.Errorf("fetching workflow from db: %v", err)
	}

	return nil
}
