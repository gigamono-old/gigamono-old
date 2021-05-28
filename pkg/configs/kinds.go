package configs

import (
	"errors"
	"strings"
)

// ConfigFormat is the different config format supported by Gigamono.
type ConfigFormat string

// ...
const (
	YAML ConfigFormat = "YAML"
	JSON ConfigFormat = "JSON"
	TOML ConfigFormat = "TOML"
)

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

// EnvironmentKind represents the kind of environment.
type EnvironmentKind string

// ...
const (
	Development EnvironmentKind = "Development"
	Production  EnvironmentKind = "Production"
)

// SecretsManagerKind represents the kind of secrets manager.
type SecretsManagerKind string

// ...
const (
	Env SecretsManagerKind = "Env"
)

// FilestoreManagerKind represents the kind of filestore manager.
type FilestoreManagerKind string

// ...
const (
	Local FilestoreManagerKind = "Local"
)

// ConfigKind represents the kind of config file.
type ConfigKind string

// ...
const (
	App            ConfigKind = "App"
	Workflow       ConfigKind = "Workflow"
	Gigamono       ConfigKind = "Gigamono"
	Appcredentials ConfigKind = "AppCredentials"
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

// StepKind is the type of task.
type StepKind string

// ...
const (
	StepKindTrigger StepKind = "Trigger"
	StepKindAction  StepKind = "Action"
)
