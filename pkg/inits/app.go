package inits

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/configs"
	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/filestore"
	"github.com/gigamono/gigamono/pkg/logs"
	"github.com/gigamono/gigamono/pkg/secrets"
)

// App holds common important states of a service.
type App struct {
	Config  configs.GigamonoConfig
	Secrets secrets.Manager
	DB      database.DB
	Kind    string
}

// NewApp is a common initialiser for Gigamono servers.
func NewApp(appKind string) (App, error) {
	// Set log status file.
	logs.SetStatusLogFile() // TODO: Abstract

	// Set filestore avatars location.
	filestore.SetAvatarsLocation() // TODO: Abstract

	// Load gigamono config file.
	config, err := configs.LoadGigamonoConfig()
	if err != nil {
		err := fmt.Errorf("initialising app: unable to load gigamono config file from env var `GIGAMONO_CONFIG_FILE`: %v", err)
		logs.FmtPrintln(err)
		return App{}, err
	}

	// Set up secret manager,
	secrets, err := secrets.NewManager(&config)
	if err != nil {
		err := fmt.Errorf("initialising app: unable to create a secrets manager: %v", err)
		logs.FmtPrintln(err)
		return App{}, err
	}

	// Connect to database.
	db, err := database.Connect(&config, secrets, appKind)
	if err != nil {
		err := fmt.Errorf("initialising app: unable to connect to db: %v", err)
		logs.FmtPrintln(err)
		return App{}, err
	}

	return App{Config: config, Secrets: secrets, DB: db, Kind: appKind}, nil
}
