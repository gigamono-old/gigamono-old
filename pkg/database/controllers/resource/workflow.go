package resource

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/gofrs/uuid"
)

// CreateWorkflow creates a workflow.
func CreateWorkflow(
	db *database.DB,
	sessionUserID *uuid.UUID,
	workflowName string,
	specificationFileURL string,
) (*resource.Workflow, error) {
	// TODO: Sec: Permission.
	workflow := resource.Workflow{
		Name:                 workflowName,
		CreatorID:            sessionUserID,
		SpecificationFileURL: specificationFileURL,
	}

	// Insert workflow and return id.
	if _, err := db.Model(&workflow).Insert(); err != nil {
		return &workflow, fmt.Errorf("creating workflow in db: %v", err)
	}

	return &workflow, nil
}

// ActivateWorkflow updates the `is_active` field of a workflow.
func ActivateWorkflow(db *database.DB, _ *uuid.UUID, workflowID *uuid.UUID) (*resource.Workflow, error) {
	// TODO: Sec: Permission.
	workflow := resource.Workflow{}
	workflow.IsActive = true

	// Update details specified workflow.
	if _, err := db.Model(&workflow).Set("is_active = ?is_active").Where("id = ?", workflowID).Update(); err != nil {
		return &workflow, fmt.Errorf("activating workflow in db: %v", err)
	}

	return &workflow, nil
}

// GetWorkflow gets a workflow by id.
func GetWorkflow(db *database.DB, _ *uuid.UUID, workflowID *uuid.UUID) (*resource.Workflow, error) {
	// TODO: Sec: Permission.
	workflow := resource.Workflow{}

	// Select the workflow with the specified workflowID
	if err := db.Model(&workflow).Where("id = ?", workflowID).Select(); err != nil {
		return &workflow, fmt.Errorf("fetching workflow from db: %v", err)
	}

	return &workflow, nil
}
