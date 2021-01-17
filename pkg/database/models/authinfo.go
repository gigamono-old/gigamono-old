package models

import (
	"gorm.io/datatypes"
)

// AuthInfo represents information an application needs for authentication to take place.
type AuthInfo struct {
	Base
	Name string
	Code datatypes.JSON
	App  App
}
