package controllers

import "github.com/sageflow/sageflow/pkg/database/models"

// CreateEngine creates an engine.
func (db *DB) CreateEngine() *models.Engine {
	engine := models.Engine{}
	db.Create(&engine)
	return &engine
}
