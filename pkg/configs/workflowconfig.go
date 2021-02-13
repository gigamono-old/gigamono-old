package configs

import (
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
		ExecutionContexts []ExecutionContext `mapstructure:"execution_contexts" json:"execution_contexts"`
		Authors           []Author           `json:"authors"`
	} `json:"metadata"`
	Tasks []Task `json:"tasks"`
}

// Task is an executable step in a workflow.
type Task struct {
	Kind             TaskKind            `json:"kind"`
	Name             string              `json:"name"`
	Index            uint                `json:"index"`
	Dependencies     []uint              `json:"dependencies"`
	Position         []uint              `json:"position"`
	AppName          string              `mapstructure:"app_name" json:"app_name"`
	ExecutionContext ExecutionContext    `mapstructure:"execution_context" json:"execution_context"`
	AppID            UUID                `mapstructure:"app_id" json:"app_id"`
	AccountID        UUID                `mapstructure:"account_id" json:"account_id"`
	Fields           map[string][]string `json:"fields"`
}

// NewWorkflowConfig creates a WorkflowConfig from string. Supports JSON, TOML and YAML string format.
func NewWorkflowConfig(workflowString string, format ConfigFormat) (WorkflowConfig, error) {
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
