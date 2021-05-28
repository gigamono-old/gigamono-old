package filestore

import (
	"github.com/gigamono/gigamono/pkg/files"
)

// LocalManager manages code
type LocalManager struct {
	RootPath string
}

// NewLocalManager creates a new local directory manager.
func NewLocalManager(rootPath string) (LocalManager, error) {
	// Create or open file directory.
	files.OpenOrCreateFolder(rootPath)

	return LocalManager{
		RootPath: rootPath,
	}, nil
}

// CreateFile creates a file.
func (mgr *LocalManager) CreateFile(filename string, content []byte, opts ...interface{}) (string, error) {
	//
	return "", nil
}

// WriteToFile writes to a file.
func (mgr *LocalManager) WriteToFile(filename string, content []byte, opts ...interface{}) error {
	//
	return nil
}
