package configs

import (
	"strings"

	"github.com/spf13/viper"
)

// WorkflowConfig represents a runnable workflow.
type WorkflowConfig struct {
	Version  uint
	Kind     ConfigKind
	Metadata struct {
		Name              string
		Description       string
		ExecutionContexts []ExecutionContext `mapstructure:"execution_contexts"`
		Authors           []Author
	}
	Tasks []Task
}

// Task is an executable step in a workflow.
type Task struct {
	Kind             TaskKind
	Name             string
	Index            uint
	Dependencies     []uint
	Position []uint
	ExecutionContext ExecutionContext `mapstructure:"execution_context"`
	AppName          string `mapstructure:"app_name"`
	AppID            UUID             `mapstructure:"app_id"`
	AccountID        UUID   `mapstructure:"account_id"`
	Fields           map[string][]string
}

// NewWorkflowConfig creates a WorkflowConfig from string. Supports JSON, TOML and YAML string format.
func NewWorkflowConfig(workflowString string, format ConfigFormat) (WorkflowConfig, error) {
	workflow := WorkflowConfig{}
	reader := strings.NewReader(workflowString)

	// Set viper to parse format.
	viper.SetConfigType(format.String())
	viper.ReadConfig(reader)

	// Convert format into Workflow object.
	if err := viper.Unmarshal(&workflow, getCustomDecoder()); err != nil {
		return WorkflowConfig{}, err
	}

	return workflow, nil
}
