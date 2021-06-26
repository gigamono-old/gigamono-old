package inits

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/configs"
	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/filestore"
	"github.com/gigamono/gigamono/pkg/logs"
	"github.com/gigamono/gigamono/pkg/secrets"
)

// App represents states common to every Gigamono service.
type App struct {
	Config  configs.GigamonoConfig
	Secrets secrets.Manager
	Filestore
	DB   database.DB
	Kind ServiceKind
}

// Filestore holds the different filestore managers.
type Filestore struct {
	Project   filestore.Manager
	Extension filestore.Manager
	Image     filestore.Manager
}

// NewApp is a common initialiser for Gigamono services.
func NewApp(serviceKind ServiceKind) (App, error) {
	// Set log status file.
	logs.SetStatusLogFile() // TODO: Abstract

	// Load gigamono config file.
	config, err := configs.LoadGigamonoConfig()
	if err != nil {
		err := fmt.Errorf("initialising app: unable to load gigamono config file from env var `GIGAMONO_CONFIG_FILE`: %v", err)
		logs.FmtPrintln(err)
		return App{}, err
	}

	// Set how secret manager.
	secretsManager, err := secrets.NewManager(&config)
	if err != nil {
		err := fmt.Errorf("initialising app: unable to create a secrets manager: %v", err)
		logs.FmtPrintln(err)
		return App{}, err
	}

	// Set how project files are stored.
	projectManager, err := filestore.NewManager(
		config.Filestore.Project.Paths.Public,
		config.Filestore.Project.Paths.Private,
	)
	if err != nil {
		err := fmt.Errorf("initialising app: unable to create a project filestore manager: %v", err)
		logs.FmtPrintln(err)
		return App{}, err
	}

	// Set how extension files are stored.
	extensionManager, err := filestore.NewManager(
		config.Filestore.Extension.Paths.Public,
		config.Filestore.Extension.Paths.Private,
	)
	if err != nil {
		err := fmt.Errorf("initialising app: unable to create a extension filestore manager: %v", err)
		logs.FmtPrintln(err)
		return App{}, err
	}

	// Set how image files are stored.
	imageManager, err := filestore.NewManager(
		config.Filestore.Image.Paths.Public,
		config.Filestore.Image.Paths.Private,
	)
	if err != nil {
		err := fmt.Errorf("initialising app: unable to create a images filestore manager: %v", err)
		logs.FmtPrintln(err)
		return App{}, err
	}

	// Connect to database.
	db, err := database.
		Connect(
			secretsManager,
			serviceKind.DatabaseKind(),
			config,
		)
	if err != nil {
		err := fmt.Errorf("initialising app: unable to connect to db: %v", err)
		logs.FmtPrintln(err)
		return App{}, err
	}

	return App{
		Config:  config,
		Secrets: secretsManager,
		Filestore: Filestore{
			Project:   projectManager,
			Extension: extensionManager,
			Image:     imageManager,
		},
		DB:   db,
		Kind: serviceKind,
	}, nil
}
