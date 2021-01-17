package database

import (
	"github.com/gofrs/uuid"
	"github.com/sageflow/sageflow/pkg/database/models"
	"gorm.io/datatypes"
)

/// Users

/// Folders

// CreateFolder creates a folder.
func (db *DB) CreateFolder() *models.Folder {
	folder := models.Folder{}
	db.Create(&folder)
	return &folder
}

/// Workflows

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

/// Engines

// CreateEngine creates an engine.
func (db *DB) CreateEngine() *models.Engine {
	engine := models.Engine{}
	db.Create(&engine)
	return &engine
}

/// Logs

// WriteLog writes a log to the database.
func (db *DB) WriteLog() *models.Log {
	log := models.Log{}
	db.Create(&log)
	return &log
}

/// Workflows X Engines

// RegisterEngineWorkflow ...
func (db *DB) RegisterEngineWorkflow(workflow *models.Workflow, engine *models.Engine) error {
	return db.Model(workflow).
		Association("XEngine").
		Append(engine)
}
