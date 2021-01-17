package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectMySQLDB connects to MySQL database.
func ConnectMySQLDB(connectionString string, newLogger *logger.Interface) *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(connectionString),
		&gorm.Config{
			Logger: *newLogger,
		},
	)

	if err != nil {
		log.Panic("Cannot continue without a database.\n")
	}

	return db
}
