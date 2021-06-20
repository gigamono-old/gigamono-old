package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gofrs/uuid"
)

// XWorkspaceIntegration is represents users membership to a workspace.
type XWorkspaceIntegration struct {
	models.BaseNoID
	WorkspaceID   uuid.UUID `pg:",pk,type:uuid" json:"workspace_id"`
	IntegrationID uuid.UUID `pg:",pk,type:uuid" json:"integration_id"`
}

func init() {
	orm.RegisterTable(XUserWorkspaceMembership{})
}
