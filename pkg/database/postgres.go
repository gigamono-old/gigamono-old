package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectPostgresDB connects to Postgres database.
func ConnectPostgresDB(connectionString string, newLogger *logger.Interface) *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(connectionString),
		&gorm.Config{
			Logger: *newLogger,
		},
	)

	if err != nil {
		log.Panic("Cannot continue without a database.\n")
	}

	return db
}
