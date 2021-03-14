package database

import (
	"errors"
	"strings"
)

// DBKind represents the type of the database.
type DBKind int

// ...
const (
	Postgres DBKind = iota
	MySQL
	SQLite3
)

func (kind DBKind) String() string {
	var res string

	switch kind {
	case Postgres:
		res = "postgres"
	case MySQL:
		res = "mysql"
	case SQLite3:
		res = "sqlite"
	}

	return res
}

// ToDBKind converts string representation to
func ToDBKind(ty string) (DBKind, error) {
	switch strings.ToLower(ty) {
	case "postgres", "postgresql", "psql":
		return Postgres, nil
	case "mysql", "mysqldb":
		return MySQL, nil
	case "sqlite":
		return SQLite3, nil
	default:
		return 0, errors.New("Unsupported database type")
	}
}
