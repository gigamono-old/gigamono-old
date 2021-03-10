package resource

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
				&Workspace{},
				&User{},
				&Profile{},
				&Group{},
				&AccessControl{},
				&Role{},
				&Folder{},
				&Engine{},
				&Workflow{},
				&Account{},
				&App{},
				&Theme{},
				&RESTHook{},
				&WorkflowInstance{},
				&Log{},
			)
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(
				"access_controls",
				"accounts",
				"apps",
				"engines",
				"folders",
				"groups",
				"profiles",
				"roles",
				"themes",
				"workflows",
				"workspaces",
				"users",
				"rest_hooks",
				"workflow_instances",
				"logs",
				"apps_x_accounts",
				"access_controls_x_roles",
				"themes_x_roles",
				"workflows_x_roles",
				"workflows_x_engines",
				"apps_x_roles",
				"users_x_groups",
				"users_x_roles",
				"users_x_workflows",
				"users_x_workspaces",
			)
		},
	}
}

// We need to use copies of models here because the ones in ./models can change during development.
// This prevents prevent inconsistent states between migrations.

// Base ...
type Base struct {
	ID        uuid.UUID `gorm:"primary_key; unique; type:uuid; column:id; default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

// AccessControl ...
type AccessControl struct {
	Base
	Name        string
	Description string
	XRole       []*Role `gorm:"many2many:access_controls_x_roles"`
}

// Account ...
type Account struct {
	Base
	UserID            uuid.UUID
	AuthAccessTokenID uuid.UUID `gorm:"unique; type:uuid"`
	XApp              []*App    `gorm:"many2many:apps_x_accounts"`
}

// App ...
type App struct {
	Base
	Name                string
	PublicID            uuid.UUID `gorm:"unique; type:uuid"`
	IsSecurityReviewed  bool
	IsOnAppEntityBehalf bool
	CreatorID           uuid.UUID
	AuthInfoID          uuid.UUID
	RESTHook            []RESTHook
	XRole               []*Role    `gorm:"many2many:apps_x_roles"`
	XAccount            []*Account `gorm:"many2many:apps_x_accounts"`
}

// Engine ...
type Engine struct {
	Base
	WorkflowInstance []WorkflowInstance
}

// Folder ...
type Folder struct {
	Base
	Name        string
	Description string
	Avatar32URL string `gorm:"column:avatar_32_url"`
	CreatorID   uuid.UUID
	Workflow    []Workflow
}

// Group ...
type Group struct {
	Base
	Name        string
	Avatar32URL string `gorm:"column:avatar_32_url"`
	CreatorID   uuid.UUID
	XUser       []*User `gorm:"many2many:users_x_groups"`
}

// Log ...
type Log struct {
	Base
	UserID             uuid.UUID
	EngineID           uuid.UUID
	WorkflowID         uuid.UUID
	WorkflowInstanceID uuid.UUID
	Message            string
	Level              string
}

// Profile ...
type Profile struct {
	Base
	Username    string
	FirstName   string
	LastName    string
	Email       string
	Avatar32URL string `gorm:"column:avatar_32_url"`
	UserID      uuid.UUID
}

// RESTHook ...
type RESTHook struct {
	Base
	UserID  uuid.UUID
	AppID   uuid.UUID
	HookURL string
}

// Role ...
type Role struct {
	Base
	Name           string
	Description    string
	XApp           []*App           `gorm:"many2many:apps_x_roles"`
	XUser          []*User          `gorm:"many2many:users_x_roles"`
	XAccessControl []*AccessControl `gorm:"many2many:access_controls_x_roles"`
	XTheme         []*Theme         `gorm:"many2many:themes_x_roles"`
	XWorkflow      []*Workflow      `gorm:"many2many:workflows_x_roles"`
}

// Theme ...
type Theme struct {
	Base
	Name      string
	Code      datatypes.JSON
	PublicID  uuid.UUID `gorm:"unique; type:uuid"`
	CreatorID uuid.UUID
	XRole     []*Role `gorm:"many2many:themes_x_roles"`
}

// User ...
type User struct {
	Base
	AuthUserID   *uuid.UUID `gorm:"unique; type:uuid"`
	Profile      Profile
	RefreshToken string // JWT.1.R
	Account      []Account
	RESTHook     []RESTHook
	AppID        []App        `gorm:"foreignKey:CreatorID"`
	Group        []Group      `gorm:"foreignKey:CreatorID"`
	Workflow     []Workflow   `gorm:"foreignKey:CreatorID"`
	Workspace    []Workspace  `gorm:"foreignKey:CreatorID"`
	Folder       []Folder     `gorm:"foreignKey:CreatorID"`
	Theme        []Theme      `gorm:"foreignKey:CreatorID"`
	XGroup       []*Group     `gorm:"many2many:users_x_groups"`
	XWorkspace   []*Workspace `gorm:"many2many:users_x_workspaces"`
	XRole        []*Role      `gorm:"many2many:users_x_roles"`
}

// Workflow ...
type Workflow struct {
	Base
	Name             string
	Code             datatypes.JSON
	IsActive         bool
	IsDraft          bool
	FolderID         *uuid.UUID
	CreatorID        *uuid.UUID
	WorkflowInstance []WorkflowInstance
	XRole            []*Role `gorm:"many2many:workflows_x_roles"`
}

// WorkflowInstance ...
type WorkflowInstance struct {
	Base
	CurrentTaskIndex int
	Dataflow         datatypes.JSON
	WorkflowID       uuid.UUID
	EngineID         uuid.UUID
}

// Workspace ...
type Workspace struct {
	Base
	Name        string
	Avatar32URL string `gorm:"column:avatar_32_url"`
	CreatorID   uuid.UUID
	XUser       []*User `gorm:"many2many:users_x_workspaces"`
}
