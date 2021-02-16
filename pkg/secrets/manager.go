package secrets

import (
	"github.com/sageflow/sageflow/pkg/configs"
)

// Manager abstract how application secrets are gotten.
// Secrets may be from a .env file or some external like Hashicorp Vault.
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
