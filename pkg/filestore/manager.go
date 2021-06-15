package filestore

// Manager abstracts how certain files are managed.
type Manager interface {
	WriteToFile(filename string, content []byte, opts ...interface{}) (string, error)
	ReadFile(filename string, opts ...interface{}) (string, error)
	GetActualPath(filename string) string
	GetPublicPath(filename string) string
}

// NewManager creates a new file manager based on settings in your gigamono.yaml file.
func NewManager(actualRootPath string, publicRootPath string) (Manager, error) {
	// TODO: Currently only supports local directory manager.
	manager, err := NewLocalManager(actualRootPath, publicRootPath)
	return &manager, err
}
