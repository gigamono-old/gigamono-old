package configs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// GigamonoConfig holds Gigamono configurations.
// Sec: Secrets shouldn't be stored in this file.
type GigamonoConfig struct {
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
				PublicPort  uint `mapstructure:"public_port" json:"public_port"`
				PrivatePort uint `mapstructure:"private_port" json:"private_port"`
			} `json:"api"`
			Auth struct {
				PublicPort  uint `mapstructure:"public_port" json:"public_port"`
				PrivatePort uint `mapstructure:"private_port" json:"private_port"`
			} `json:"auth"`
			WorkflowEngine struct {
				PublicPorts struct {
					MainServer         uint `mapstructure:"main_server" json:"main_server"`
					WebhookService     uint `mapstructure:"webhook_service" json:"webhook_service"`
					RunnableSupervisor uint `mapstructure:"runnable_supervisor" json:"runnable_supervisor"`
				} `mapstructure:"public_ports" json:"public_ports"`
				PrivatePorts struct {
					MainServer         uint `mapstructure:"main_server" json:"main_server"`
					WebhookService     uint `mapstructure:"webhook_service" json:"webhook_service"`
					RunnableSupervisor uint `mapstructure:"runnable_supervisor" json:"runnable_supervisor"`
				} `mapstructure:"private_ports" json:"private_ports"`
			} `mapstructure:"workflow_engine" json:"workflow_engine"`
			DocumentEngine struct {
				PublicPort  uint `mapstructure:"public_port" json:"public_port"`
				PrivatePort uint `mapstructure:"private_port" json:"private_port"`
			} `mapstructure:"document_engine" json:"document_engine"`
		} `json:"types"`
	} `json:"services"`
	SecretsManager struct {
		kind string
	} `mapstructure:"secrets_manager" json:"secrets_manager"`
}

// NewGigamonoConfig creates a GigamonoConfig from string. Supports JSON, TOML and YAML string format.
func NewGigamonoConfig(gigamonoString string, format ConfigFormat) (GigamonoConfig, error) {
	// TODO: Sec: Validation
	config := GigamonoConfig{}
	reader := strings.NewReader(gigamonoString)

	// Set format to parse.
	converter := viper.New()
	converter.SetConfigType(format.String())
	converter.ReadConfig(reader)

	// Unmarshal string into object.
	if err := converter.Unmarshal(&config); err != nil {
		return GigamonoConfig{}, err
	}

	return config, nil
}

// LoadGigamonoConfig loads a gigamono config from file.
func LoadGigamonoConfig() (GigamonoConfig, error) {
	// Load .env file.
	if err := godotenv.Load(); err != nil {
		return GigamonoConfig{}, err
	}

	// Get config file path from env.
	path := os.Getenv("GIGAMONO_CONFIG_FILE")
	if path == "" {
		return GigamonoConfig{}, errors.New("GIGAMONO_CONFIG_FILE env variable is missing or empty")
	}

	// Get file extension and use it to determine config format.
	format, err := ToConfigFormat(filepath.Ext(path)[1:])
	if err != nil {
		return GigamonoConfig{}, err
	}

	// Read file.
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return GigamonoConfig{}, err
	}

	return NewGigamonoConfig(string(fileContent), format)
}

// JSON converts config to json.
func (config *GigamonoConfig) JSON() (string, error) {
	// TODO: Sec: Validation
	bytes, err := json.Marshal(config)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
