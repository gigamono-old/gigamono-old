package email

import "github.com/gigamono/gigamono/pkg/configs"

// Driver abstracts how the application sends emails.
type Driver interface {
	Send(from string, to []string) error
}

// NewDriver creates a new email driver based on settings in your gigamono.yaml file.
func NewDriver(config *configs.SageflowConfig) (Driver, error) {
	// TODO: Currently only supports EnvManager
	// manager, err := NewSendGridDriver(config)
	// return &manager, err
	return nil, nil
}
