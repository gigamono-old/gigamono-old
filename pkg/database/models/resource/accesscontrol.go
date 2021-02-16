package resource

import "github.com/sageflow/sageflow/pkg/database/models"

// AccessControl represents access control for managing access to resources.
type AccessControl struct {
	models.Base
	Name        string
	Description string
	XRole       []*Role `gorm:"many2many:access_controls_x_roles"`
}
