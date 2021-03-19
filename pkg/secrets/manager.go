package secrets

import (
	"github.com/sageflow/sageflow/pkg/configs"
)

// Manager abstracts how application secrets are managed.
// Secrets may be from a .env file or some external service like Hashicorp Vault.
type Manager interface {
	Get(key string, opts ...interface{}) (string, error)
	Set(key string, value string, opts ...interface{}) error
}

// NewManager creates a new secrets manager based on settings in your sageflow.yaml file.
func NewManager(config *configs.SageflowConfig) (Manager, error) {
	// TODO: Currently only supports EnvManager
	manager, err := NewEnvManager(config)
	return &manager, err
}
