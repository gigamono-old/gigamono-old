package resource

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/configs"
	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/gofrs/uuid"
)

// CreateWorkflow creates a workflow.
func CreateWorkflow(
	db *database.DB,
	sessionUserID *uuid.UUID,
	workflowName string,
	workflowSpecification configs.WorkflowConfig,
) (string, error) {
	// TODO: Sec: Permission.
	workflow := resource.Workflow{
		Name:          workflowName,
		Specification: workflowSpecification,
		CreatorID:     sessionUserID,
	}

	// Insert workflow and return id.
	if _, err := db.Model(&workflow).Returning("id").Insert(); err != nil {
		return "", fmt.Errorf("unable to create workflow: %v", err)
	}

	return workflow.ID.String(), nil
}

// ActivateWorkflow updates the `is_active` field of a workflow.
func ActivateWorkflow(db *database.DB, _ *uuid.UUID, workflowID *uuid.UUID) (string, error) {
	// TODO: Sec: Permission.
	workflow := &resource.Workflow{}
	workflow.ID = *workflowID
	workflow.IsActive = true

	// Update details specified workflow.
	if _, err := db.Model(&workflow).Update(); err != nil {
		return "", fmt.Errorf("unable to activate workflow: %v", err)
	}

	return workflow.ID.String(), nil
}

// GetWorkflow gets a workflow by id.
func GetWorkflow(db *database.DB, _ *uuid.UUID, workflowID *uuid.UUID) (*resource.Workflow, error) {
	// TODO: Sec: Permission.
	workflow := new(resource.Workflow)

	// Select the workflow with the specified workflowID
	if err := db.Model(&workflow).Where("id = ?", workflowID).Select(); err != nil {
		return &resource.Workflow{}, fmt.Errorf("unable to get workflow: %v", err)
	}

	return workflow, nil
}
