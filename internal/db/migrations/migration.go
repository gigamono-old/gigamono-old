package migrations

import (
	"errors"
	"fmt"
	"time"

	"github.com/sageflow/sageflow/pkg/logs"

	"gorm.io/gorm"
)

const initialTablesName = "InitialTables"

// Migration represnts the migration table.
type Migration struct {
	ID        string `gorm:"type:varchar(255); primary_key; unique; column:id;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

// Migratable contains information for applying and rolling back a migration.
type Migratable struct {
	ID   string
	Up   func(db *gorm.DB) error
	Down func(db *gorm.DB) error
}

// Migrator encapsulates the migration list and operations.
type Migrator struct {
	DB                *gorm.DB
	Migrations        []Migratable
	InitialTables     []interface{}
	InitialJoinTables []string
}

// NewMigrator creates a new migrator.
func NewMigrator(db *gorm.DB) Migrator {
	createMigrationsTable(db)
	return Migrator{DB: db}
}

// AddMigration adds a new migration to migrator's list.
func (migrator *Migrator) AddMigration(migration Migratable) error {
	// Check if ID is not reserved ID
	if migration.ID == initialTablesName {
		return errors.New("adding migration: `" + initialTablesName + "` is a reserved id")
	}

	migrator.Migrations = append(migrator.Migrations, migration)

	return nil
}

// AddMigrations adds multiple migrations to migrator's list.
func (migrator *Migrator) AddMigrations(migrations []Migratable) error {
	return nil
}

// AddInitialTables adds initial tables to the list.
func (migrator *Migrator) AddInitialTables(models ...interface{}) {
	migrator.InitialTables = models
}

// AddInitialJoinTables add initial join tables. This is only useful during initial table roll back.
func (migrator *Migrator) AddInitialJoinTables(models ...string) {
	migrator.InitialJoinTables = models
}

// Up applies the all migrations starting from the latest migration.
func (migrator *Migrator) Up() error {
	return migrator.UpByOne()
}

// Down rolls back the all migrations starting from the latest migration.
func (migrator *Migrator) Down() error {
	return migrator.DownByOne()
}

// UpTo applies migrations up to the specified id starting from the latest migration.
func (migrator *Migrator) UpTo(id string) error {
	return nil
}

// DownTo applies migrations down to the specified id starting from the latest migration.
func (migrator *Migrator) DownTo(id string) error {
	return nil
}

// UpByOne applies the next migration.
func (migrator *Migrator) UpByOne() error {
	lastMigration := Migration{}
	err := migrator.DB.First(&lastMigration).Order("id DESC").Limit(1).Error

	var migrationName string

	if errors.Is(err, gorm.ErrRecordNotFound) { // No migration record.
		// TODO: Abstract
		if migrator.hasIntialTables() {
			migrationName = initialTablesName

			if err := migrator.migrateInitialTables(); err != nil { // TODO: Idempotent
				logs.FmtPrintf("ERR\t\t%s\n", initialTablesName)
				return err
			}
		} else if migrator.hasMigrations() {
			migration := migrator.Migrations[0]
			migrationName = migration.ID

			if err := migrator.migrate(migration); err != nil { // TODO: Idempotent
				return err
			}
		} else {
			return fmt.Errorf("no migration available")
		}
	} else { // Migration record exists.
		// Find migration in stored migrations.
		var foundMigrationIndex = -1
		for idx, migr := range migrator.Migrations {
			if migr.ID == lastMigration.ID {
				foundMigrationIndex = idx
				break
			}
		}

		if foundMigrationIndex == -1 && lastMigration.ID != initialTablesName { // Migration not found in list
			return fmt.Errorf("cannot find the last migration in migration list")
		}

		if foundMigrationIndex == len(migrator.Migrations)-1 { // foundMigrationIndex is the last migration
			return fmt.Errorf("no new migration")
		}

		migration := migrator.Migrations[foundMigrationIndex+1]
		migrationName = migration.ID

		if err := migrator.migrate(migration); err != nil { // TODO: Idempotent
			return err
		}
	}

	logs.FmtPrintf("OK\t\t%s", migrationName)

	return nil
}

// DownByOne rolls back the next migration.
func (migrator *Migrator) DownByOne() error {
	id, err := migrator.getDownMigration()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("no previous migration to rollback to")
		}
		return err
	}

	switch id {
	case initialTablesName:
		if err := migrator.demigrateInitialTables(); err != nil { // TODO: Idempotent
			return err
		}
	default:
		// Find migration in stored migrations.
		var foundMigrationIndex = -1
		for idx, migr := range migrator.Migrations {
			if migr.ID == id {
				foundMigrationIndex = idx
				break
			}
		}

		if foundMigrationIndex == -1 { // Not found
			return fmt.Errorf("cannot find the last migration in migration list")
		}

		if err := migrator.demigrate(migrator.Migrations[foundMigrationIndex+1]); err != nil { // TODO: Idempotent
			return err
		}
	}

	logs.FmtPrintf("OK\t\t%s", id)

	return nil
}

// migrate applies a migration.
func (migrator *Migrator) migrate(migration Migratable) error {
	// Apply up migration and rollback if unsuccessful.
	if err := migrator.DB.Transaction(migration.Up); err != nil {
		return fmt.Errorf("migrating `%s`: %s", migration.ID, err)
	}

	return nil
}

// demigrate rolls back a migration.
func (migrator *Migrator) demigrate(migration Migratable) error {
	// Apply down demigration and rollback if unsuccessful.
	if err := migrator.DB.Transaction(migration.Down); err != nil {
		return fmt.Errorf("migrating `%s`: %s", migration.ID, err)
	}
	return nil
}

// migrateInitialTables migrates initial tables and assoiated join tables.
func (migrator *Migrator) migrateInitialTables() error {
	// Migrate tables and add . Drop tables if there is an error.
	if err := migrator.DB.Transaction(func(tx *gorm.DB) error {
		if err := migrator.DB.AutoMigrate(migrator.InitialTables...); err != nil {
			return err
		}
		return migrator.DB.Create(&Migration{ID: initialTablesName}).Error
	}); err != nil {
		return fmt.Errorf("migrating initial tables: %s", err)
	}

	return nil
}

// demigrateInitialTables demigrates initial tables and assoiated join tables.
func (migrator *Migrator) demigrateInitialTables() error {
	// Drop tables.
	if err := migrator.DB.Transaction(func(tx *gorm.DB) error {
		for _, model := range migrator.InitialTables {
			if err := migrator.DB.Migrator().DropTable(model); err != nil {
				return err
			}
		}
		for _, model := range migrator.InitialJoinTables {
			if err := migrator.DB.Migrator().DropTable(model); err != nil {
				return err
			}
		}
		return migrator.DB.Delete(&Migration{ID: initialTablesName}).Error
	}); err != nil {
		return fmt.Errorf("demigrating initial tables: %s", err)
	}

	return nil
}

// getDownMigration gets the last migration.
func (migrator *Migrator) getDownMigration() (string, error) {
	lastMigration := Migration{}

	if err := migrator.DB.First(&lastMigration).Order("id DESC").Limit(1).Error; err != nil {
		return "", err
	}

	return lastMigration.ID, nil
}

// hasIntialTables checks if migrator has initial tables and assoiated join tables.
func (migrator *Migrator) hasIntialTables() bool {
	return len(migrator.InitialTables) > 0 || len(migrator.InitialJoinTables) > 0
}

// hasMigrations checks if migrator has other migrations.
func (migrator *Migrator) hasMigrations() bool {
	return len(migrator.Migrations) > 0
}

// createMigrationsTable creates migrations table if one does not exist.
func createMigrationsTable(db *gorm.DB) error {
	migr := db.Migrator()

	if !migr.HasTable(&Migration{}) {
		return migr.CreateTable(&Migration{})
	}

	return nil
}
