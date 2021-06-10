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
		Name                 string `json:"name"`
		ID                   UUID   `json:"id"`
		SpecificationFileURL UUID   `mapstructure:"specification_file_url" json:"specification_file_url"`
	} `json:"integration"`
	Steps struct {
		Max   uint64              `json:"max"`
		Items map[uint64]StepItem `json:"items"`
	} `json:"steps"`
}

// StepItem is a step in a workflow.
type StepItem struct {
	IntegrationIndex uint              `mapstructure:"integration_index" json:"integration_index"`
	OperationKey     string            `mapstructure:"operation_key" json:"operation_key"`
	Dependencies     []uint            `json:"dependencies"`
	Inputs           map[string]string `json:"inputs"`
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
