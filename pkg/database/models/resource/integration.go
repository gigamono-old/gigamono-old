package resource

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Integration stores information about an integration.
type Integration struct {
	models.Base
	Name                            string      `json:"name"`
	SpecificationFileURL            string      `json:"specification_file_url"`
	CreatorID                       uuid.UUID   `json:"creator_id"`
	XWorkspaceInstalledIntegrations []Workspace `pg:"many2many:x_workspace_installed_integrations" json:"-"`
}

// Create creates an integration.
func (integration *Integration) Create(db *database.DB) error {
	// TODO: Sec: Permission.
	// Insert integration in db.
	if _, err := db.Model(integration).Insert(); err != nil {
		return fmt.Errorf("creating integration in db: %v", err)
	}

	return nil
}

// GetByID gets an integration by id.
func (integration *Integration) GetByID(db *database.DB) error {
	// TODO: Sec: Permission.
	// Select the integration with the specified integration ID
	if err := db.Model(integration).WherePK().Select(); err != nil {
		return fmt.Errorf("fetching integration from db: %v", err)
	}

	return nil
}
