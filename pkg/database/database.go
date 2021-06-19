package database

import (
	"errors"

	"github.com/gigamono/gigamono/pkg/secrets"
	"github.com/go-pg/pg/v10"
)

// DB contains a db connection.
type DB struct {
	*pg.DB
}

// Connect connects to specified postgres database.
func Connect(secrets secrets.Manager, databaseKind string) (DB, error) {
	var connectionURI string
	var err error

	switch databaseKind {
	case "resource":
		connectionURI, err = secrets.Get("RESOURCE_DB_CONNECTION_URI")
		if err != nil {
			return DB{}, err
		}
	case "auth":
		connectionURI, err = secrets.Get("AUTH_DB_CONNECTION_URI")
		if err != nil {
			return DB{}, err
		}
	default:
		return DB{}, errors.New("Unsupported database type. Resource or auth database types supported")
	}

	opt, err := pg.ParseURL(connectionURI)
	if err != nil {
		return DB{}, errors.New("Bad database connection URI")
	}

	db := pg.Connect(opt)

	return DB{DB: db}, nil
}
