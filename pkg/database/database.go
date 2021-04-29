package database

import (
	"errors"
	"strings"

	"github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/go-pg/pg/v10/orm"

	"github.com/gigamono/gigamono/pkg/secrets"
	"github.com/go-pg/pg/v10"
)

// DB contains a db connection.
type DB struct {
	*pg.DB
}

func init() {
	// Register all junction tables.
	orm.RegisterTable(resource.XUsersWorkspaces{})
}

// Connect connects to specified postgres database.
func Connect(secrets secrets.Manager, serviceKind string) (DB, error) {
	var connectionURI string
	var err error

	switch strings.ToLower(serviceKind) {
	case "res", "resource":
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
		return DB{}, errors.New("Unsupported service type for connecting to database. Resource or auth supported")
	}

	opt, err := pg.ParseURL(connectionURI)
	if err != nil {
		return DB{}, errors.New("Bad database connection URI")
	}

	db := pg.Connect(opt)

	return DB{DB: db}, nil
}
