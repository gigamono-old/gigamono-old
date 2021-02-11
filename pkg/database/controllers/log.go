package controllers

import "github.com/sageflow/sageflow/pkg/database/models"

// WriteLog writes a log to the database.
func (db *DB) WriteLog() *models.Log {
	log := models.Log{}
	db.Create(&log)
	return &log
}
