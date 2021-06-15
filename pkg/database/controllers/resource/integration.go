package resource

import (
	"fmt"

	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
	"github.com/gofrs/uuid"
)

// CreateIntegration creates an integration.
func CreateIntegration(
	db *database.DB,
	sessionUserID *uuid.UUID,
	integrationName string,
	specificationPath string,
) (*resource.Integration, error) {
	// TODO: Sec: Permission.
	integration := resource.Integration{
		Name:              integrationName,
		CreatorID:         sessionUserID,
		SpecificationPath: specificationPath,
	}

	// Insert integration and return id.
	if _, err := db.Model(&integration).Insert(); err != nil {
		return &integration, fmt.Errorf("creating integration in db: %v", err)
	}

	return &integration, nil
}

// GetIntegration gets an integration by id.
func GetIntegration(db *database.DB, integrationID *uuid.UUID) (*resource.Integration, error) {
	// TODO: Sec: Permission.
	integration := resource.Integration{}

	// Select the integration with the specified integrationID
	if err := db.Model(&integration).Where("id = ?", integrationID).Select(); err != nil {
		return &integration, fmt.Errorf("fetching integrationn from db: %v", err)
	}

	return &integration, nil
}
