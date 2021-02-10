package database

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/sageflow/sageflow/pkg/configs"
	"github.com/sageflow/sageflow/pkg/secrets"

	"github.com/sageflow/sageflow/pkg/logs"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect connects to the appropriate database.
func Connect(config *configs.SageflowConfig, secrets secrets.Manager, appKind string) (DB, error) {
	newLogger := createStatusLogger() // Create a new logger.
	connectionURI := ""
	kind := DBKind(-1)
	var err error

	switch strings.ToUpper(appKind) {
	case "RESOURCE":
		connectionURI, err = secrets.Get("RESOURCE_DB_CONNECTION_URI", make(map[string]string))
		if err != nil {
			return DB{}, err
		}

		kind, err = ToDBKind(config.Database.Resource.Kind)
		if err != nil {
			return DB{}, err
		}
	case "AUTH":
		connectionURI, err = secrets.Get("AUTH_DB_CONNECTION_URI", make(map[string]string))
		if err != nil {
			return DB{}, err
		}

		kind, err = ToDBKind(config.Database.Auth.Kind)
		if err != nil {
			return DB{}, err
		}
	default:
		return DB{}, errors.New("Unsupported application type for connecting to database. Resource and Auth app types supported at the moment")
	}

	var db *gorm.DB

	// Connect using teh appropriate driver.
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
