package controllers

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
	"gorm.io/datatypes"
)

// CreateWorkflow creates a workflow.
func (db *DB) CreateWorkflow(
	workflowName string,
	workflowString string,
	isActive bool,
	isDraft bool,
	folderID uuid.UUID,
	creatorID uuid.UUID,
) *models.Workflow {
	workflow := models.Workflow{
		Name:      workflowName,
		Code:      datatypes.JSON(workflowString),
		IsActive:  isActive,
		IsDraft:   isDraft,
		FolderID:  folderID,
		CreatorID: creatorID,
	}
	db.Create(&workflow)
	return &workflow
}
