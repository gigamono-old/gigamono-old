package filestore

// Manager abstracts how certain files are managed.
type Manager interface {
	CreateFile(filename string, content []byte, opts ...interface{}) (string, error)
	WriteToFile(filename string, content []byte, opts ...interface{}) error
}

// NewManager creates a new file manager based on settings in your gigamono.yaml file.
func NewManager(rootPath string) (Manager, error) {
	// TODO: Currently only supports local directory manager.
	manager, err := NewLocalManager(rootPath)
	return &manager, err
}
