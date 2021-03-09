package secrets

import (
	"errors"
	"os"

	"github.com/sageflow/sageflow/pkg/configs"

	"github.com/joho/godotenv"
)

// EnvManager manages secrets stored in .env files
type EnvManager struct{}

// NewEnvManager creates a new EnvManager.
func NewEnvManager(_ *configs.SageflowConfig) (EnvManager, error) {
	// Load .env file.
	return EnvManager{}, godotenv.Load()
}

// Get get a secret by its key.
func (mgr *EnvManager) Get(key string, _ ...interface{}) (string, error) {
	secret := os.Getenv(key)
	if secret == "" {
		return secret, errors.New("Env Secrets Manager: " + key + " env variable is missing or empty")
	}

	return secret, nil
}

// Set does nothing.
func (mgr *EnvManager) Set(_ string, _ string, _ ...interface{}) error {
	return nil
}
