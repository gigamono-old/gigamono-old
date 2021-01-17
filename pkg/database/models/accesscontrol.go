package models

// AccessControl represents access control for managing access to resources.
type AccessControl struct {
	Base
	Name        string
	Description string
	XRole       []*Role `gorm:"many2many:access_controls_x_roles"`
}
