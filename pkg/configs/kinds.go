package configs

import (
	"errors"
	"strings"
)

// ConfigFormat is the different config format supported by Sageflow.
type ConfigFormat string

// ...
const (
	YAML ConfigFormat = "YAML"
	JSON ConfigFormat = "JSON"
	TOML ConfigFormat = "TOML"
)

func (format *ConfigFormat) String() string {
	return string(*format)
}

// ToConfigFormat convert a string to ConfigFormat.
func ToConfigFormat(format string) (ConfigFormat, error) {
	switch strings.ToLower(format) {
	case "yaml", "yml":
		return YAML, nil
	case "json":
		return JSON, nil
	case "toml":
		return TOML, nil
	default:
		return "", errors.New("Conversion from string `" + format + "` to ConfigFormat not possible")
	}
}

// ConfigKind represents the kind of config file.
type ConfigKind string

// ...
const (
	App            ConfigKind = "App"
	Workflow       ConfigKind = "Workflow"
	Sageflow       ConfigKind = "Sageflow"
	Appcredentials ConfigKind = "AppCredentials"
)

// ExecutionContext refers to how the engine is going to run each task.
type ExecutionContext string

// ...
const (
	Protected ExecutionContext = "Protected" // Sandboxed code execution
	Bare      ExecutionContext = "Bare"      // Non-sandboxed code execution
)

// AuthKind is the type of authorisation an App supports.
type AuthKind string

// ...
const (
	AuthKindOAuth2 AuthKind = "OAuth2"
)

// InputKind is the type of user input.
type InputKind string

// ...
const (
	Email  InputKind = "Email"
	Select InputKind = "Select"
)

// DropdownKind is the type of dropdown.
type DropdownKind string

// ...
const (
	Static  DropdownKind = "Static"
	Dynamic DropdownKind = "Dynamic"
)

// ActionKind is the type of action.
type ActionKind string

// ...
const (
	ActionKindAction ActionKind = "Action"
	ActionKindSearch ActionKind = "Search"
)

// TaskKind is the type of task.
type TaskKind string

// ...
const (
	TaskKindTrigger TaskKind = "Trigger"
	TaskKindAction  TaskKind = "Action"
)
