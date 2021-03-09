package resource

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database"
	"github.com/sageflow/sageflow/pkg/database/models/resource"
	"gorm.io/datatypes"
)

// CreateWorkflow creates a workflow.
func CreateWorkflow(
	db *database.DB,
	workflowName string,
	workflowString string,
	isActive bool,
	isDraft bool,
	folderID *uuid.UUID,
	creatorID *uuid.UUID,
) (resource.Workflow, error) {
	// Workflow with a nil folder UUID implies default folder.
	// Workflow with a nil creator UUID implies unregistered user.

	workflow := resource.Workflow{
		Name:      workflowName,
		Code:      datatypes.JSON(workflowString),
		IsActive:  isActive,
		IsDraft:   isDraft,
		FolderID:  folderID,
		CreatorID: creatorID,
	}

	if err := db.Create(&workflow).Error; err != nil {
		return resource.Workflow{}, err
	}

	return workflow, nil
}
