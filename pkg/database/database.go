package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/gigamono/gigamono/pkg/configs"
	"github.com/gigamono/gigamono/pkg/secrets"
	"github.com/go-pg/pg/v10"
)

// DB contains a db connection.
type DB struct {
	*pg.DB
}

type queryPrinter struct{}

func (logger queryPrinter) BeforeQuery(ctx context.Context, _ *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (logger queryPrinter) AfterQuery(ctx context.Context, queryEvent *pg.QueryEvent) error {
	query, _ := queryEvent.FormattedQuery()
	fmt.Printf(">> %v \n\n", string(query))
	return nil
}

// Connect connects to specified postgres database.
func Connect(secrets secrets.Manager, databaseKind string, config configs.GigamonoConfig) (DB, error) {
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

	// Print queries in dev environment.
	if config.Environment == configs.Development {
		db.AddQueryHook(queryPrinter{})
	}

	return DB{DB: db}, nil
}
