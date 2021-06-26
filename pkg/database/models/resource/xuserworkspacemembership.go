package resource

import (
	"github.com/gigamono/gigamono/pkg/database/models"
	"github.com/go-pg/pg/v10/orm"
	"github.com/gofrs/uuid"
)

// XUserWorkspaceMembership is represents users membership to a workspace.
type XUserWorkspaceMembership struct {
	models.BaseNoID
	UserID      uuid.UUID `pg:",pk,type:uuid" json:"user_id"`
	WorkspaceID uuid.UUID `pg:",pk,type:uuid" json:"workspace_id"`
}

func init() {
	orm.RegisterTable((*XUserWorkspaceMembership)(nil))
}
