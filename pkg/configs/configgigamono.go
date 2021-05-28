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
	Environment EnvironmentKind `json:"environment"`
	Services    struct {
		API struct {
			Ports Ports `json:"ports"`
		} `json:"api"`
		Auth struct {
			Ports Ports `json:"ports"`
		} `json:"auth"`
		WorkflowEngine struct {
			MainServer struct {
				Ports Ports `json:"ports"`
			} `mapstructure:"main_server" json:"main_server"`
			WebhookService struct {
				Ports Ports `json:"ports"`
			} `mapstructure:"webhook_service" json:"webhook_service"`
			RunnableSupervisor struct {
				Ports Ports `json:"ports"`
			} `mapstructure:"runnable_supervisor" json:"runnable_supervisor"`
		} `mapstructure:"workflow_engine" json:"workflow_engine"`
		DocumentEngine struct {
			Ports Ports `json:"ports"`
		} `mapstructure:"document_engine" json:"document_engine"`
	} `json:"services"`
	Secrets   SecretsManagerKind `json:"secrets"`
	Filestore struct {
		Workflow   FilestoreInfo `json:"workflow"`
		Serverless FilestoreInfo `json:"serverless"`
	} `json:"filestore"`
}

// Ports represents public and private ports.
type Ports struct {
	Public  uint `json:"public_port"`
	Private uint `json:"private_port"`
}

// FilestoreInfo represents information for managing certain type of file.
type FilestoreInfo struct {
	Kind FilestoreManagerKind `json:"kind"`
	Path string               `json:"path"`
}

// NewGigamonoConfig creates a GigamonoConfig from string. Supports JSON, TOML and YAML string format.
func NewGigamonoConfig(gigamonoString string, format ConfigFormat) (GigamonoConfig, error) {
	// TODO: Sec: Validation
	config := GigamonoConfig{}
	reader := strings.NewReader(gigamonoString)

	// Set format to parse.
	converter := viper.New()
	converter.SetConfigType(string(format))
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
