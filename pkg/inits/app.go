package inits

import (
	"github.com/sageflow/sageflow/pkg/configs"
	"github.com/sageflow/sageflow/pkg/database"
	"github.com/sageflow/sageflow/pkg/logs"
	"github.com/sageflow/sageflow/pkg/secrets"
)

type App struct {
	Config  configs.SageflowConfig
	Secrets secrets.Manager
	DB      database.DB
	Kind    string
}

// NewApp is a common initialiser for Sageflow servers.
func NewApp(appKind string) (App, error) {
	// Set up log status file.
	logs.SetStatusLogFile()

	// Load sageflow config file.
	config, err := configs.LoadSageflowConfig()
	if err != nil {
		logs.FmtPrintln("Unable to load sageflow config file from env var `SAGEFLOW_CONFIG_FILE`:", err)
	}

	// Set up secret manager,
	secrets, err := secrets.NewManager(&config)
	if err != nil {
		logs.FmtPrintln("Unable to create secrets manager:", err)
	}

	// Connect to database.
	db, err := database.Connect(&config, secrets, appKind)
	if err != nil {
		return App{}, err
	}

	return App{Config: config, Secrets: secrets, DB: db, Kind: appKind}, nil
}
