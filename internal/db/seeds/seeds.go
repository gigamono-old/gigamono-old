package seeds

import (
	"errors"
	"fmt"

	"github.com/gofrs/uuid"

	"github.com/sageflow/sageflow/pkg/database"
	"github.com/sageflow/sageflow/pkg/database/models"
)

// Seed represents a seed in the database.
type Seed struct {
	models.Base
	TableName string
	SeedID    uuid.UUID `gorm:"type:uuid"`
}

// AddAll adds all seed data to the DB.
func AddAll(db *database.DB) error {
	var err error

	// Creates seeds table if one does not exist.
	if err = createSeedsTable(db); err != nil {
		return err
	}

	// Load user seeds to the database.
	if err = loadUsers(db); err != nil {
		return err
	}

	return nil
}

// RemoveAll removes all seed data in the DB.
func RemoveAll(db *database.DB) error {
	tableNames, err := getSeedTableNames(db)
	if err != nil {
		return err
	}

	fmt.Println("tables: ", tableNames)

	for tableName := range tableNames {
		// Delete seeds in a specified table.
		// Sec: TODO: Need to remove concat. Gorm currently doesn't handle table name escaping well.
		db.Exec(
			"DELETE FROM "+tableName+" WHERE id IN (SELECT seed_id FROM seeds WHERE table_name = ?)",
			tableName,
		)

		// Delete rows associated with a specified table in seeds table.
		db.Exec("DELETE FROM seeds WHERE table_name = ?", tableName)
	}

	return nil
}

// Add adds seed data for a table to the DB.
func Add(db *database.DB, tableName string) error {
	var err error

	// Creates seeds table if one does not exist.
	if err = createSeedsTable(db); err != nil {
		return err
	}

	// Load specified table with seeds.
	switch tableName {
	case "users":
		if err = loadUsers(db); err != nil {
			return err
		}

	default:
		return errors.New("Table not seeded (Perhaps table name not updated)")
	}

	return nil
}

// Remove removes seed data for a table in the DB.
func Remove(db *database.DB, tableName string) error {
	tableNames, err := getSeedTableNames(db)
	if err != nil {
		return err
	}

	// Check if table exists.
	if _, ok := tableNames[tableName]; ok {
		fmt.Println("table: ", tableName)

		// Delete seeds in a specified table.
		// Sec: TODO: Need to remove concat. Gorm currently doesn't handle table name escaping well.
		db.Exec(
			"DELETE FROM "+tableName+" WHERE id IN (SELECT seed_id FROM seeds WHERE table_name = ?)",
			tableName,
		)

		// Delete rows associated with a specified table in seeds table.
		db.Exec("DELETE FROM seeds WHERE table_name = ?", tableName)

	}

	return nil
}

// createSeedsTable creates seeds table if one does not exist.
func createSeedsTable(db *database.DB) error {
	migr := db.Migrator()

	if !migr.HasTable(&Seed{}) {
		return migr.CreateTable(&Seed{})
	}

	return nil
}

// generateUUIDs generates random UUIDs.
func generateUUIDs(count int) ([]uuid.UUID, error) {
	var err error
	uuids := make([]uuid.UUID, count)

	for i := 0; i < count; i++ {
		uuids[i], err = uuid.NewV4()
		if err != nil {
			return []uuid.UUID{}, err
		}
	}

	return uuids, nil
}

// getSeedTableNames gets all tables with seeds.
func getSeedTableNames(db *database.DB) (map[string]struct{}, error) {
	rows, err := db.Table("seeds").
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

// loadUsers loads seeds into the users table.
func loadUsers(db *database.DB) error {
	users, err := FakeUsers(10)
	if err != nil {
		return err
	}

	for _, user := range users {
		db.Create(&user)
		db.Create(&Seed{
			TableName: db.GetTableName(user),
			SeedID:    user.ID,
		})
	}

	return nil
}
