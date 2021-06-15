package filestore

import (
	"io/ioutil"

	"github.com/gigamono/gigamono/pkg/files"
)

// LocalManager manages code
type LocalManager struct {
	ActualRootPath string
	PublicRootPath string
}

// NewLocalManager creates a new local directory manager.
func NewLocalManager(actualRootPath string, publicRootPath string) (LocalManager, error) {
	// Create or open file directory.
	if err := files.OpenOrCreateFolder(actualRootPath); err != nil {
		return LocalManager{}, err
	}

	return LocalManager{
		ActualRootPath: actualRootPath,
		PublicRootPath: publicRootPath,
	}, nil
}

// WriteToFile writes to a file. Creates file if it does not exist.
func (mgr *LocalManager) WriteToFile(filename string, content []byte, opts ...interface{}) (string, error) {
	filePath := mgr.GetActualPath(filename)

	// Open or create file and write content to it.
	if _, err := files.WriteToFile(filePath, content); err != nil {
		return "", err
	}

	return filePath, nil
}

// ReadFile reads a file.
func (mgr *LocalManager) ReadFile(filename string, opts ...interface{}) (string, error) {
	filePath := mgr.GetActualPath(filename)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// GetActualPath gets the actual path.
func (mgr *LocalManager) GetActualPath(filename string) string {
	return mgr.ActualRootPath + "/" + filename
}

// GetPublicPath gets the public path.
func (mgr *LocalManager) GetPublicPath(filename string) string {
	return mgr.PublicRootPath + "/" + filename
}
