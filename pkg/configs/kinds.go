package configs

// ConfigFormat is the different config format supported by Sageflow.
type ConfigFormat string

// The different types of config format supported.
const (
	YAML  ConfigFormat = "YAML"
	JSON  ConfigFormat = "JSON"
	TOML  ConfigFormat = "TOML"
)

func (format *ConfigFormat) String() string {
	return string(*format)
}

// ExecutionContext refers to how the engine is going to run each task.
type ExecutionContext string

// The different types of execution contexts.
const (
	PROTECTED ExecutionContext = "PROTECTED" // Sandboxed code execution
	BARE      ExecutionContext = "BARE"      // Non-sandboxed code execution
)

// AuthKind is the type of authorisation an App supports.
type AuthKind string

// The different types of authorization.
const (
	OAUTH2 AuthKind = "OAUTH2"
)

// InputKind is the type of user input.
type InputKind string

// The different types of user input.
const (
	EMAIL  InputKind = "EMAIL"
	SELECT InputKind = "SELECT"
)

// DropdownKind is the type of dropdown.
type DropdownKind string

// The different types of dropdown.
const (
	STATIC  DropdownKind = "STATIC"
	DYNAMIC DropdownKind = "DYNAMIC"
)

// ActionKind is the type of action.
type ActionKind string

// The different types of action.
const (
	ACTION ActionKind = "ACTION"
	SEARCH ActionKind = "SEARCH"
)

// TaskKind is the type of task.
type TaskKind string

// The different types of task.
const (
	TaskKindTRIGGER TaskKind = "TRIGGER"
	TaskKindACTION  TaskKind = "ACTION"
)
