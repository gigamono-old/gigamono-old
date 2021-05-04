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
		Name              string             `json:"name"`
		Description       string             `json:"description"`
		Authors           []Author           `json:"authors"`
	} `json:"metadata"`
	Steps []Step `json:"steps"`
}

// Step is an executable step in a workflow.
type Step struct {
	Kind             StepKind            `json:"kind"`
	Name             string              `json:"name"`
	Index            uint                `json:"index"`
	Dependencies     []uint              `json:"dependencies"`
	Position         []uint              `json:"position"`
	AppName          string              `mapstructure:"app_name" json:"app_name"`
	AppID            UUID                `mapstructure:"app_id" json:"app_id"`
	AccountID        UUID                `mapstructure:"account_id" json:"account_id"`
	Fields           map[string][]string `json:"fields"`
}

// NewWorkflowConfig creates a WorkflowConfig from string. Supports JSON, TOML and YAML string format.
func NewWorkflowConfig(workflowString string, format ConfigFormat) (WorkflowConfig, error) {
	// TODO: Sec: Validation
	workflow := WorkflowConfig{}
	reader := strings.NewReader(workflowString)

	// Set format to parse.
	converter := viper.New()
	converter.SetConfigType(format.String())
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
