package resource

import (
	"github.com/gofrs/uuid"
	"github.com/gigamono/gigamono/pkg/database/models"
)

// RESTHook represents a webhook used by an app on behalf of a user.
type RESTHook struct {
	models.Base
	UserID  uuid.UUID
	AppID   uuid.UUID
	HookURL string
}
