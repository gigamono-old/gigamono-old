package filestore

// Manager abstracts how certain files are managed.
type Manager interface {
	// WriteToFile writes to a file. Creates file if it does not exist.
	WriteToFile(filename string, content []byte, opts ...interface{}) (string, error)

	// ReadFile reads a file.
	ReadFile(filename string, opts ...interface{}) (string, error)

	// DeleteFile deletes a file. May be permanent.
	DeleteFile(filename string, opts ...interface{}) (string, error)

	// GetPublicPath gets the public path.
	GetPublicPath(filename string) string

	// GetPrivatePath gets the private path.
	GetPrivatePath(filename string) string
}

// NewManager creates a new file manager based on settings in your gigamono.yaml file.
func NewManager(publicRootPath string, privateRootPath string) (Manager, error) {
	// TODO: Currently only supports local directory manager.
	manager, err := NewLocalManager(publicRootPath, privateRootPath)
	return &manager, err
}
