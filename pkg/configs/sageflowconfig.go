package configs

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// SageflowConfig holds Sageflow configurations.
type SageflowConfig struct {
	Version  uint
	Kind     ConfigKind
	Metadata struct {
		Authors []Author
	}
	Execution struct {
		UseSubprocess bool `mapstructure:"use_subprocess"`
	}
	Server struct {
		API struct {
			Port int
		}
		Engine struct {
			Port int
		}
		Auth struct {
			Port int
		}
	}
	Database struct {
		Resource struct {
			URI  string
			Kind string
		}
		Auth struct {
			URI  string
			Kind string
		}
	}
	SecretsManager struct {
		kind string
	} `mapstructure:"secrets_manager"`
}

// NewSageflowConfig creates a SageflowConfig from string. Supports JSON, TOML and YAML string format.
func NewSageflowConfig(sageflowString string, format ConfigFormat) (SageflowConfig, error) {
	config := SageflowConfig{}
	reader := strings.NewReader(sageflowString)

	// Set format to parse.
	viper.SetConfigType(format.String())
	viper.ReadConfig(reader)

	// Unmarshal string into object.
	if err := viper.Unmarshal(&config); err != nil {
		return SageflowConfig{}, err
	}

	return config, nil
}

// LoadSageflowConfig loads a sageflow config from file
func LoadSageflowConfig() (SageflowConfig, error) {
	// Load .env file.
	if err := godotenv.Load(); err != nil {
		return SageflowConfig{}, err
	}

	// Get config file path from env.
	path := os.Getenv("SAGEFLOW_CONFIG_FILE")
	if path == "" {
		return SageflowConfig{}, errors.New("SAGEFLOW_CONFIG_FILE env variable is missing or empty")
	}

	// Get file extension and use it to determine config format.
	format, err := ToConfigFormat(filepath.Ext(path)[1:])
	if err != nil {
		return SageflowConfig{}, err
	}

	// Read file.
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return SageflowConfig{}, err
	}

	return NewSageflowConfig(string(fileContent), format)
}
