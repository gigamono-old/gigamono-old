package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectSQLite3DB connects to SQLite database.
func ConnectSQLite3DB(connectionString string, newLogger *logger.Interface) *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open(connectionString),
		&gorm.Config{
			Logger: *newLogger,
		},
	)

	if err != nil {
		log.Panic("Cannot continue without a database.\n")
	}

	return db
}
