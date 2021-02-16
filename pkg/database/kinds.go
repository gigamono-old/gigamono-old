package database

import (
	"errors"
	"strings"
)

// DBKind represents the type of the database.
type DBKind int

// ...
const (
	POSTGRES DBKind = iota
	MYSQL
	SQLITE3
)

func (kind DBKind) String() string {
	var res string

	switch kind {
	case POSTGRES:
		res = "postgres"
	case MYSQL:
		res = "mysql"
	case SQLITE3:
		res = "sqlite"
	}

	return res
}

// ToDBKind converts string representation to
func ToDBKind(ty string) (DBKind, error) {
	switch strings.ToUpper(ty) {
	case "POSTGRES", "POSTGRESQL", "PSQL":
		return POSTGRES, nil
	case "MYSQL", "MYSQLDB":
		return MYSQL, nil
	case "SQLITE":
		return SQLITE3, nil
	default:
		return 0, errors.New("Unsupported database type")
	}
}
