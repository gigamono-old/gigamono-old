package models

import (
	"github.com/gofrs/uuid"
)

// RESTHook represents a webhook used by an app on behalf of a user.
type RESTHook struct {
	Base
	UserID  uuid.UUID
	AppID   uuid.UUID
	HookURL string
}
