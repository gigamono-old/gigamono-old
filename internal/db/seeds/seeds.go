package seeds

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"

	"github.com/sageflow/sageflow/internal/db/seeds/common"
	"github.com/sageflow/sageflow/internal/db/seeds/resource"
	"github.com/sageflow/sageflow/pkg/database"
)

// Loader represent a pair of fake data generating function and its generated ids.
type Loader = struct {
	lambda func(*database.DB, int) ([]uuid.UUID, error)
	ids    []uuid.UUID
}

// Seeder represents a seeder instance.
type Seeder struct {
	DB          *database.DB
	appKind     string
	tableLoader map[string]Loader
}

// NewSeeder creates a new seeder.
func NewSeeder(db *database.DB, appKind string) Seeder {
	tableLoader := make(map[string]Loader)

	switch strings.ToUpper(appKind) {
	case "AUTH":
	default:
		tableLoader["users"] = Loader{
			lambda: resource.LoadFakeUsers,
			ids:    []uuid.UUID{},
		}
		tableLoader["profiles"] = Loader{
			lambda: resource.LoadFakeProfiles,
			ids:    []uuid.UUID{},
		}
	}
	return Seeder{DB: db, appKind: appKind, tableLoader: tableLoader}
}

// AddAll adds all seed data to the DB.
func (seeder *Seeder) AddAll() error {
	var err error

	// Creates seeds table if one does not exist.
	if err = seeder.createSeedsTable(); err != nil {
		return err
	}

	for tableName := range seeder.tableLoader {
		if err := seeder.Add(tableName); err != nil {
			return err
		}
	}

	return nil
}

// RemoveAll removes all seed data in the DB.
func (seeder *Seeder) RemoveAll() error {
	tableNames, err := seeder.getSeedTableNames()
	if err != nil {
		return err
	}

	fmt.Println("tables: ", tableNames)

	for tableName := range tableNames {
		// Delete seeds in a specified table.
		// Sec: TODO: Need to remove concat. Gorm currently doesn't handle table name escaping well.
		seeder.DB.Exec(
			"DELETE FROM "+tableName+" WHERE id IN (SELECT seed_id FROM seeds WHERE table_name = ?)",
			tableName,
		)

		// Delete rows associated with a specified table in seeds table.
		seeder.DB.Exec("DELETE FROM seeds WHERE table_name = ?", tableName)
	}

	return nil
}

// Add adds seed data for a table to the DB.
func (seeder *Seeder) Add(tableName string) error {
	var err error

	// Creates seeds table if one does not exist.
	if err = seeder.createSeedsTable(); err != nil {
		return err
	}

	// Get fake data loader from map.
	loader := seeder.tableLoader[tableName]
	if loader.lambda == nil {
		return errors.New("Specified table name does not exist in seed list")
	}

	// Generate fake data.
	loader.ids, err = loader.lambda(seeder.DB, 10)

	return err
}

// Remove removes seed data for a table in the DB.
func (seeder *Seeder) Remove(tableName string) error {
	tableNames, err := seeder.getSeedTableNames()
	if err != nil {
		return err
	}

	// Check if table exists.
	if _, ok := tableNames[tableName]; ok {
		fmt.Println("table: ", tableName)

		// Delete seeds in a specified table.
		// Sec: TODO: Need to remove concat. Gorm currently doesn't handle table name escaping well.
		seeder.DB.Exec(
			"DELETE FROM "+tableName+" WHERE id IN (SELECT seed_id FROM seeds WHERE table_name = ?)",
			tableName,
		)

		// Delete rows associated with a specified table in seeds table.
		seeder.DB.Exec("DELETE FROM seeds WHERE table_name = ?", tableName)

	}

	return nil
}

// createSeedsTable creates seeds table if one does not exist.
func (seeder *Seeder) createSeedsTable() error {
	migr := seeder.DB.Migrator()

	if !migr.HasTable(&common.Seed{}) {
		return migr.CreateTable(&common.Seed{})
	}

	return nil
}

// getSeedTableNames gets all tables with seeds.
func (seeder *Seeder) getSeedTableNames() (map[string]struct{}, error) {
	rows, err := seeder.DB.Table("seeds").
		Distinct("table_name").
		Select("table_name").
		Rows()

	if err != nil {
		return map[string]struct{}{}, err
	}

	// Using maps instead of array because it make checking for "contains" easy.
	tableNames := map[string]struct{}{}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		tableNames[tableName] = struct{}{}
	}

	return tableNames, nil
}
