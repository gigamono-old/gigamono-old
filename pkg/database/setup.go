package database

import (
	"log"
	"os"
	"time"

	"github.com/sageflow/sageflow/pkg/logs"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect connects to the appropriate database.
func Connect() *DB {
	connectionURI := os.Getenv("RESOURCE_DB_CONNECTION_URI")
	newLogger := createStatusLogger()

	kind, err := ToDBKind(os.Getenv("RESOURCE_DB_TYPE"))
	if err != nil {
		log.Panicf("%v\n", err)
	}

	var db *gorm.DB

	switch kind {
	case POSTGRES:
		db = ConnectPostgresDB(connectionURI, newLogger)
	case MYSQL:
		db = ConnectMySQLDB(connectionURI, newLogger)
	case SQLITE3:
		db = ConnectSQLite3DB(connectionURI, newLogger)
	default:
		log.Panic("Unsupported database type\n")
	}

	log.Printf(
		"Database successfully connected: %v: %v\n",
		kind,
		db.Migrator().CurrentDatabase(),
	)

	return &DB{db, kind}
}

func createStatusLogger() *logger.Interface {
	file, err := logs.OpenLogFile("status.log")
	if err != nil {
		log.Printf("Cannot open or create 'logs/status.log' file: %v\nFalling back to stdout/stderr\n", err)
	}

	newLogger := logger.New(
		log.New(file, "\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Warn, // Log level
			Colorful:      false,       // Disable color
		},
	)

	return &newLogger
}
