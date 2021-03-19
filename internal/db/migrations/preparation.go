package migrations

import (
	"strings"

	"github.com/sageflow/sageflow/internal/db/migrations/resource"
)

//
func (migrator *Migrator) PrepareMigrations(appKind string) {
	switch strings.ToLower(appKind) {
	case "auth":
		prepareAuthMigrations(migrator)
	default:
		prepareResourceMigrations(migrator)
	}
}

func prepareAuthMigrations(migrator *Migrator) {
	migrator.AddInitialTables()
	migrator.AddMigrations([]Migratable{})
}

func prepareResourceMigrations(migrator *Migrator) {
	migrator.AddInitialTables(resource.GetInitialTables()...)
	migrator.AddInitialJoinTables(resource.GetInitialJoinTables()...)
	migrator.AddMigrations([]Migratable{})
}
