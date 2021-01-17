package configs

import (
	"reflect"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// WorkflowConfig represents a runnable workflow.
type WorkflowConfig struct {
	Version  uint
	Metadata struct {
		Name              string
		ExecutionContexts []ExecutionContext `mapstructure:"execution_contexts"`
	}
	Tasks []Task
}

// Task is an executable step in a workflow.
type Task struct {
	Kind             TaskKind
	Name             string
	Index            uint
	Dependencies     []uint
	ExecutionContext ExecutionContext `mapstructure:"execution_context"`
	AppID            UUID             `mapstructure:"app_id"`
	Fields           map[string][]string
	AppName          string `mapstructure:"app_name"`
	AccountID        UUID   `mapstructure:"account_id"`
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

func getCustomDecoder() viper.DecoderConfigOption {
	return viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
			switch t {
			case reflect.TypeOf(UUID{}):
				parsedID, err := uuid.FromString(data.(string))
				if err != nil {
					return UUID{}, err
				}
				id := UUID(parsedID)
				return id, nil
			}
			return data, nil
		},
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	))
}
