package auth

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// InitialTables1 returns the migration for creating the initial table.
func InitialTables1() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1_initial_tables",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(
				&AccessToken{},
				&AppAuth{},
				&Password{},
			)
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(
				"access_tokens",
				"app_auths",
				"passwords",
			)
		},
	}
}

// Base ...
type Base struct {
	ID        uuid.UUID `gorm:"primary_key; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

// AccessToken ...
type AccessToken struct {
	Base
	AppID                uuid.UUID
	OwnerID              uuid.UUID
	EncryptedAccessToken string // Sec: Encrypted with AUTH_SECRET_KEY
}

// AppAuth ...
type AppAuth struct {
	Base
	Name  string
	Code  datatypes.JSON
	AppID uuid.UUID
}

// Password ...
type Password struct {
	Base
	OwnerID           uuid.UUID
	EncryptedPassword string
}

