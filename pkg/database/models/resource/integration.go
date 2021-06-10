package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// Integration stores information about an integration.
type Integration struct {
	models.Base
	Name            string     `json:"name"`
	IntegrationPath string     `json:"integration_path"`
	CreatorID       *uuid.UUID `json:"creator_id"`
}
