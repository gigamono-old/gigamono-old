package resource

import (
	"github.com/gigamono/gigamono/pkg/database"
	"github.com/gigamono/gigamono/pkg/database/models/resource"
)

// CreateFolder creates a folder.
func CreateFolder(db *database.DB) (resource.Folder, error) {
	folder := resource.Folder{}
	if err := db.Create(&folder); err != nil {
		return resource.Folder{}, nil
	}
	return folder, nil
}
