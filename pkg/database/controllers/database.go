package controllers

import (
	"gorm.io/gorm"
)

// DB contains a db connection.
type DB struct {
	*gorm.DB
	Kind DBKind
}

// GetTableName gets a model's real table name.
func (db *DB) GetTableName(model interface{}) string {
	stmt := &gorm.Statement{DB: db.DB}
	stmt.Parse(model)
	return stmt.Schema.Table
}
