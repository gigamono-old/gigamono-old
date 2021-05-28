package secrets

import (
	"github.com/gigamono/gigamono/pkg/configs"
)

// Manager abstracts how application secrets are managed.
//
// Secrets may be from a .env file or some external service like Hashicorp Vault.
type Manager interface {
	Get(key string, opts ...interface{}) (string, error)
	Set(key string, value string, opts ...interface{}) error
}

// NewManager creates a new secrets manager based on settings in your gigamono.yaml file.
func NewManager(config *configs.GigamonoConfig) (Manager, error) {
	// TODO: Currently only supports EnvManager
	manager, err := NewEnvManager(config)
	return &manager, err
}
