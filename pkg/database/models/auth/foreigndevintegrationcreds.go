package auth

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/gofrs/uuid"
)

// ForeignDevIntegrationCreds holds the necessary information for the app to make successful requests to third-party integrations.
type ForeignDevIntegrationCreds struct {
	models.Base
	Name          string    `json:"name"`
	IntegrationID uuid.UUID `json:"integration_id"`
	Specification string    `pg:"type:jsonb" json:"specification"`
}
