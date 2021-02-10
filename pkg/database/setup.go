package database

import (
	"errors"
	"github.com/sageflow/sageflow/pkg/secrets"
	"log"
	"os"
	"time"

	"github.com/sageflow/sageflow/pkg/logs"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect connects to the appropriate database.
func Connect(secrets secrets.Manager) (DB, error) {
	connectionURI, err := secrets.Get("RESOURCE_DB_CONNECTION_URI", make(map[string]string))
	if err != nil {
		return DB{}, err
	}

	newLogger := createStatusLogger()

	kind, err := ToDBKind(os.Getenv("RESOURCE_DB_TYPE"))
	if err != nil {
		return DB{}, err
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
		return DB{}, errors.New("Unsupported database type")
	}

	log.Printf(
		"Database successfully connected: %v: %v\n",
		kind,
		db.Migrator().CurrentDatabase(),
	)

	return DB{db, kind}, nil
}

func createStatusLogger() *logger.Interface {
	file, err := logs.OpenOrCreateLogFile("status.log")
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
