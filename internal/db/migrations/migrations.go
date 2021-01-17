package migrations

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/sageflow/sageflow/pkg/database"
)

// PrepareMigrations prepares all migrations.
func PrepareMigrations(db *database.DB) *gormigrate.Gormigrate {
	// If table is missing on down migration, gormigrate returns an error.
	// So this is to make sure a migrations table is always present.
	createMigrationsTable(db)
	return gormigrate.New(db.DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		InitialTables1(),
	})
}

// RollbackAll rolls back all migrations. The function is not provided by gormigrate.
func RollbackAll(migrator *gormigrate.Gormigrate) error {
	var err error
	if err = migrator.RollbackTo("0"); err != nil {
		return err
	}
	err = migrator.RollbackLast()
	return err
}

// GetLastMigration get the previous migrations in the database.
func GetLastMigration(db *database.DB) string {
	var id string
	db.Table("migrations").
		Select("id").
		Order("id DESC").
		Limit(1).
		Row().
		Scan(&id)

	return id
}

// createMigrationsTable creates migrations table if one does not exist.
// This implementation is based on: https://github.com/go-gormigrate/gormigrate/blob/500825515543ab2fb0e9468eee2a88a303078f4c/gormigrate.go#L375
// Check gormigrate latest commit if anything breaks.
func createMigrationsTable(db *database.DB) error {
	type Migration struct {
		ID string `gorm:"type:varchar(255); primary_key; unique; column:id"`
	}

	migr := db.Migrator()

	if !migr.HasTable(&Migration{}) {
		return migr.CreateTable(&Migration{})
	}

	return nil
}
