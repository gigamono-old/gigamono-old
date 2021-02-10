package configs

import (
	"fmt"
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
	switch strings.ToUpper(format) {
	case "YAML", "YML":
		return YAML, nil
	case "JSON":
		return JSON, nil
	case "TOML":
		return TOML, nil
	default:
		return "", errors.New(fmt.Sprint("Conversion from string", format, "to ConfigFormat not possible"))
	}
}

// ConfigKind represents the kind of config file.
type ConfigKind string

// ...
const (
	APP      ConfigKind = "APP"
	WORKFLOW ConfigKind = "WORKFLOW"
	SAGEFLOW ConfigKind = "SAGEFLOW"
	AUTHINFO ConfigKind = "AUTHINFO"
)

// ExecutionContext refers to how the engine is going to run each task.
type ExecutionContext string

// ...
const (
	PROTECTED ExecutionContext = "PROTECTED" // Sandboxed code execution
	BARE      ExecutionContext = "BARE"      // Non-sandboxed code execution
)

// AuthKind is the type of authorisation an App supports.
type AuthKind string

// ...
const (
	OAUTH2 AuthKind = "OAUTH2"
)

// InputKind is the type of user input.
type InputKind string

// ...
const (
	EMAIL  InputKind = "EMAIL"
	SELECT InputKind = "SELECT"
)

// DropdownKind is the type of dropdown.
type DropdownKind string

// ...
const (
	STATIC  DropdownKind = "STATIC"
	DYNAMIC DropdownKind = "DYNAMIC"
)

// ActionKind is the type of action.
type ActionKind string

// ...
const (
	ACTION ActionKind = "ACTION"
	SEARCH ActionKind = "SEARCH"
)

// TaskKind is the type of task.
type TaskKind string

// ...
const (
	TaskKindTRIGGER TaskKind = "TRIGGER"
	TaskKindACTION  TaskKind = "ACTION"
)
