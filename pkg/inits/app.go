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
	Config              configs.GigamonoConfig
	Secrets             secrets.Manager
	WorkflowFilestore   filestore.Manager
	ServerlessFilestore filestore.Manager
	DB                  database.DB
	Kind                ServiceKind
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

	workflowManager := (filestore.Manager)(nil)
	serverlessManager := (filestore.Manager)(nil)
	switch serviceKind {
	case API:
		// Set how workflow files are stored.
		workflowManager, err = filestore.NewManager(config.Filestore.Workflow.Path)
		if err != nil {
			err := fmt.Errorf("initialising app: unable to create a workflows filestore manager: %v", err)
			logs.FmtPrintln(err)
			return App{}, err
		}
	case WorkflowEngineMainServer, WorkflowEngineWebhookService, WorkflowEngineRunnableSupervisor:
		// Set how serverless files are stored.
		serverlessManager, err = filestore.NewManager(config.Filestore.Serverless.Path)
		if err != nil {
			err := fmt.Errorf("initialising app: unable to create a workflows filestore manager: %v", err)
			logs.FmtPrintln(err)
			return App{}, err
		}
	}

	// Connect to database.
	db, err := database.Connect(secretsManager, serviceKind.DatabaseKind())
	if err != nil {
		err := fmt.Errorf("initialising app: unable to connect to db: %v", err)
		logs.FmtPrintln(err)
		return App{}, err
	}

	return App{
		Config:              config,
		Secrets:             secretsManager,
		WorkflowFilestore:   workflowManager,
		ServerlessFilestore: serverlessManager,
		DB:                  db,
		Kind:                serviceKind,
	}, nil
}
