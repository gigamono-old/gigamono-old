package auth

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/gofrs/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// InitialTables1 returns the migration for creating the initial table.
// TODO: Not idempotent.
func InitialTables1() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1_initial_tables",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(
				&UserAccount{},
				&AccessToken{},
				&AppCredentials{},
				&ClientAccount{},
				&SocialLogin{},
			)
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(
				"access_tokens",
				"app_credentials",
				"client_accounts",
				"social_logins",
				"user_accounts",
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

// AppCredentials ...
type AppCredentials struct {
	Base
	Name  string
	Code  datatypes.JSON
	AppID uuid.UUID
}

// ClientAccount ...
type ClientAccount struct {
	Base
	ClientID              uuid.UUID // Public-facing ID // SecureRandom.hex(32)
	EncryptedClientSecret string    // For verifying client // UUID
	Kind                  string    // Confidential or Public
	RedirectURI           string
	IsFirstParty          bool
}

// SocialLogin ...
type SocialLogin struct {
	Base
	AppName       string
	UserAccountID uuid.UUID
}

// UserAccount ...
type UserAccount struct {
	Base
	ResourceUserID    uuid.UUID
	EncryptedPassword string
	Username          string
	Email             string
	RefreshToken      string
	SocialLogin       SocialLogin
}
