package configs

import (
	"github.com/sageflow/sagedb/pkg/models"
)

// AppConfig holds the information about an app which can be used in tasks.
type AppConfig struct {
	Model    *models.App
	Version  uint
	Kind     string
	Metadata struct {
		Name     string
		PublicID UUID `mapstructure:"public_id"`
		Version  string
	}
	Auths struct {
		OAuth2s []OAuth2
		APIKeys []APIKey
	}
	Operations struct {
		Triggers []Trigger
		Actions  []Action
	}
}

// OAuth2 holds the necessary information for getting authorisation via OAuth2.
type OAuth2 struct {
	Scopes                     []string
	ShouldRefreshAutomatically bool `mapstructure:"should_refresh_automatically"`
	Fields                     []struct {
		Key              string
		IsRequired       bool      `mapstructure:"is_required"`
		IsAdministrative bool      `mapstructure:"is_administrative"`
		InputKind        InputKind `mapstructure:"input_kind"`
		DefaultValue     string    `mapstructure:"default_value"`
		Dropdown         Dropdown
	}
	APIs struct {
		AuthorisationRequest Endpoint `mapstructure:"authorisation_request"`
		AccessTokenRequest   Endpoint `mapstructure:"access_token_request"`
		RefreshTokenRequest  Endpoint `mapstructure:"refresh_token_request"`
	}
}

// APIKey holds the necessary information for getting authorisation via api keys.
type APIKey struct {
	Fields []struct {
		Key              string
		IsRequired       bool      `mapstructure:"is_required"`
		IsAdministrative bool      `mapstructure:"is_administrative"`
		InputKind        InputKind `mapstructure:"input_kind"`
		DefaultValue     string    `mapstructure:"default_value"`
		Dropdown         Dropdown
	}
	API Endpoint
}

// Trigger specifies a trigger operation.
type Trigger struct {
	Key    string
	Fields []struct {
		Key              string
		IsRequired       bool      `mapstructure:"is_required"`
		IsWriteOp        bool      `mapstructure:"is_write_op"`
		IsIdentification bool      `mapstructure:"is_identification"`
		InputKind        InputKind `mapstructure:"input_kind"`
		Dropdown         Dropdown
	}
	APIs struct {
		Polls     []Poll
		RestHooks []RestHook
	}
}

// Action specifies an action operation.
type Action struct {
	Key        string
	ActionKind ActionKind
	Fields     []struct {
		Key              string
		IsRequired       bool      `mapstructure:"is_required"`
		IsWriteOp        bool      `mapstructure:"is_write_op"`
		IsIdentification bool      `mapstructure:"is_identification"`
		InputKind        InputKind `mapstructure:"input_kind"`
		Dropdown         Dropdown
	}
	API Endpoint
}

// Poll describes how a trigger polls data.
type Poll struct {
	Endpoint
	AuthKind AuthKind
}

// RestHook describes a webhook trigger API.
type RestHook struct {
	AuthKind   AuthKind
	Operations struct {
		Subscribe   Endpoint
		Unsubscribe Endpoint
		List        Endpoint
	}
}

// Endpoint specifies how a resource is resoved, fetched, updated, etc.
type Endpoint struct {
	Code     string
	Language string
	Form     struct {
		Method  string
		URL     string
		Headers map[string]string
		Params  map[string]string
		Body    map[string]string
	}
}

// Dropdown specifies dropdown information.
type Dropdown struct {
	Kind           DropdownKind
	AllowsMultiple bool `mapstructure:"allows_multiple"`
	AllowsCustom   bool `mapstructure:"allows_custom"`
	Options        []string
}
