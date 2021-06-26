package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gofrs/uuid"
)

// XWorkspaceInstalledIntegration is represents users membership to a workspace.
type XWorkspaceInstalledIntegration struct {
	models.BaseNoID
	WorkspaceID   uuid.UUID `pg:",pk,type:uuid" json:"workspace_id"`
	IntegrationID uuid.UUID `pg:",pk,type:uuid" json:"integration_id"`
}

func init() {
	orm.RegisterTable((*XWorkspaceInstalledIntegration)(nil))
}
