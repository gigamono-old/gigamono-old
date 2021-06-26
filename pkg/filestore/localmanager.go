package filestore

import (
	"io/ioutil"
	"os"

	"github.com/gigamono/gigamono/pkg/files"
)

// LocalManager manages code
type LocalManager struct {
	PublicRootPath  string
	PrivateRootPath string
}

// NewLocalManager creates a new local directory manager.
func NewLocalManager(publicRootPath string, privateRootPath string) (LocalManager, error) {
	// Create or open file directory.
	if err := files.OpenOrCreateFolder(privateRootPath); err != nil {
		return LocalManager{}, err
	}

	return LocalManager{
		PublicRootPath:  publicRootPath,
		PrivateRootPath: privateRootPath,
	}, nil
}

// WriteToFile writes to a file. Creates file if it does not exist.
//
// Returns private file path.
func (mgr *LocalManager) WriteToFile(filename string, content []byte, opts ...interface{}) (string, error) {
	filePath := mgr.GetPrivatePath(filename)

	// Open or create file and write content to it.
	if _, err := files.WriteToFile(filePath, content); err != nil {
		return "", err
	}

	return filePath, nil
}

// ReadFile reads a file.
//
// Returns the file content.
func (mgr *LocalManager) ReadFile(filename string, opts ...interface{}) (string, error) {
	filePath := mgr.GetPrivatePath(filename)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// DeleteFile deletes a file permanently.
//
// Returns private file path.
func (mgr *LocalManager) DeleteFile(filename string, opts ...interface{}) (string, error) {
	filePath := mgr.GetPrivatePath(filename)

	if err := os.Remove(filePath); err != nil {
		return "", err
	}

	return filePath, nil
}

// GetPrivatePath gets the private path.
func (mgr *LocalManager) GetPrivatePath(filename string) string {
	return mgr.PrivateRootPath + "/" + filename
}

// GetPublicPath gets the public path.
func (mgr *LocalManager) GetPublicPath(filename string) string {
	return mgr.PublicRootPath + "/" + filename
}
