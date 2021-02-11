package controllers

import "github.com/sageflow/sageflow/pkg/database/models"

// CreateFolder creates a folder.
func (db *DB) CreateFolder() *models.Folder {
	folder := models.Folder{}
	db.Create(&folder)
	return &folder
}
