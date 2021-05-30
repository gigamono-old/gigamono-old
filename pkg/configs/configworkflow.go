package configs

import (
	"encoding/json"
	"strings"

	"github.com/spf13/viper"
)

// WorkflowConfig represents a runnable workflow.
type WorkflowConfig struct {
	Version  uint       `json:"version"`
	Kind     ConfigKind `json:"kind"`
	Metadata struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Authors     []Author `json:"authors"`
	} `json:"metadata"`
	Integration []struct {
		Name    string `json:"name"`
		ID      UUID   `json:"id"`
		FileURL UUID   `mapstructure:"file_url" json:"file_url"`
	} `json:"integration"`
	Steps []Step `json:"steps"`
}

// Step is an executable step in a workflow.
type Step struct {
	Index            uint                `json:"index"`
	IntegrationIndex *uint               `mapstructure:"integration_index" json:"integration_index"`
	OperationKey     string              `mapstructure:"operation_key" json:"operation_key"`
	Dependencies     []uint              `json:"dependencies"`
	Fields           map[string][]string `json:"fields"`
}

// NewWorkflowConfig creates a WorkflowConfig from string. Supports JSON, TOML and YAML string format.
func NewWorkflowConfig(workflowString string, format ConfigFormat) (WorkflowConfig, error) {
	// TODO: Sec: Validation
	workflow := WorkflowConfig{}
	reader := strings.NewReader(workflowString)

	// Set format to parse.
	converter := viper.New()
	converter.SetConfigType(string(format))
	converter.ReadConfig(reader)

	// Unmarshal string into object.
	if err := converter.Unmarshal(&workflow, getCustomDecoder()); err != nil {
		return WorkflowConfig{}, err
	}

	return workflow, nil
}

// JSON converts config to json.
func (config *WorkflowConfig) JSON() (string, error) {
	// TODO: Sec: Validation
	bytes, err := json.Marshal(config)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
