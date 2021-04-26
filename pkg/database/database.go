package database

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/gigamono/gigamono/pkg/configs"
	"github.com/gigamono/gigamono/pkg/secrets"

	"github.com/gigamono/gigamono/pkg/logs"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

// Connect connects to the appropriate database.
func Connect(config *configs.SageflowConfig, secrets secrets.Manager, appKind string) (DB, error) {
	newLogger := createStatusLogger() // Create a new logger.
	kind := DBKind(-1)
	var connectionURI string
	var err error
	var db *gorm.DB

	// Choose appropriate URI to use.
	switch strings.ToUpper(appKind) {
	case "RESOURCE", "RES":
		connectionURI, err = secrets.Get("RESOURCE_DB_CONNECTION_URI")
		if err != nil {
			return DB{}, err
		}
	case "AUTH":
		connectionURI, err = secrets.Get("AUTH_DB_CONNECTION_URI")
		if err != nil {
			return DB{}, err
		}
	default:
		return DB{}, errors.New("Unsupported application type for connecting to database. Resource and Auth app types supported at the moment")
	}

	// Get the db kind from connection URI.
	if index := strings.IndexByte(connectionURI, ':'); index >= 0 {
		kind, err = ToDBKind(connectionURI[:index])
		if err != nil {
			return DB{}, err
		}
	} else {
		return DB{}, errors.New("Invalid connection URI: No semi-colon in URI")
	}

	// Connect using the appropriate driver.
	switch kind {
	case Postgres:
		db = ConnectPostgresDB(connectionURI, newLogger)
	case MySQL:
		db = ConnectMySQLDB(connectionURI, newLogger)
	case SQLite3:
		db = ConnectSQLite3DB(connectionURI, newLogger)
	default:
		return DB{}, errors.New("Unsupported database type")
	}

	log.Printf(
		"Database successfully connected: %v: %v\n",
		kind,
		db.Migrator().CurrentDatabase(),
	)

	return DB{
		DB:   db,
		Kind: kind,
	}, nil
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
