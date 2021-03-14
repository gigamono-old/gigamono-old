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
// Sec: Secrets shouldn't be stored in this file.
type SageflowConfig struct {
	Version  uint       `json:"version"`
	Kind     ConfigKind `json:"kind"`
	Metadata struct {
		Authors []Author `json:"authors"`
	} `json:"metdata"`
	Execution struct {
		UseSubprocess bool `mapstructure:"use_subprocess" json:"use_subprocess"`
	} `json:"execution"`
	Services struct {
		TLS   struct{} `json:"tls"`
		Types struct {
			API struct {
				Port int `json:"port"`
			} `json:"api"`
			Engine struct {
				Port int `json:"port"`
			} `json:"engine"`
			Auth struct {
				Port int `json:"port"`
			} `json:"auth"`
		} `json:"types"`
	} `json:"services"`
	SecretsManager struct {
		kind string
	} `mapstructure:"secrets_manager" json:"secrets_manager"`
}

// NewSageflowConfig creates a SageflowConfig from string. Supports JSON, TOML and YAML string format.
func NewSageflowConfig(sageflowString string, format ConfigFormat) (SageflowConfig, error) {
	config := SageflowConfig{}
	reader := strings.NewReader(sageflowString)

	// Set format to parse.
	converter := viper.New()
	converter.SetConfigType(format.String())
	converter.ReadConfig(reader)

	// Unmarshal string into object.
	if err := converter.Unmarshal(&config); err != nil {
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
