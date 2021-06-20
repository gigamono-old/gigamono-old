package configs

import (
	"encoding/json"
	"strings"

	"github.com/spf13/viper"
)

// IntegrationConfig holds the information about an app which can be used in tasks.
type IntegrationConfig struct {
	Version  uint       `json:"version"`
	Kind     ConfigKind `json:"kind"`
	Metadata struct {
		Name                string   `json:"name"`
		PublicID            *UUID    `mapstructure:"public_id" json:"public_id"`
		Version             string   `json:"version"`
		Description         string   `json:"description"`
		Category            string   `json:"category"`
		Builtin             bool     `json:"builtin"`
		Tags                []string `json:"tags"`
		AvatarURL           *string  `mapstructure:"avatar_url" json:"avatar_url"`
		HomepageURL         string   `mapstructure:"homepage_url" json:"homepage_url"`
		APIDocumentationURL string   `mapstructure:"api_documentation_url" json:"api_documentation_url"`
		ResourceNouns       []string `mapstructure:"resource_nouns" json:"resource_nouns"`
		Authors             []Author `json:"authors"`
	} `json:"metadata"`
	Auths struct {
		OAuth2s []OAuth2 `json:"oauth2s"`
		APIKeys []APIKey `json:"api_keys"`
	} `json:"auths"`
	Operations struct {
		Triggers map[string]Trigger `json:"triggers"`
		Actions  map[string]Action  `json:"actions"`
	} `json:"operations"`
}

// OAuth2 holds the necessary information for getting authorisation via OAuth2.
type OAuth2 struct {
	Label                      string   `json:"label"`
	Scopes                     []string `json:"scopes"`
	ShouldRefreshAutomatically bool     `mapstructure:"should_refresh_automatically" json:"should_refresh_automatically"`
	Inputs                     map[string]struct {
		Label            string    `json:"label"`
		Tip              string    `json:"tip"`
		IsRequired       bool      `mapstructure:"is_required" json:"is_required"`
		IsAdministrative bool      `mapstructure:"is_administrative" json:"is_administrative"`
		InputKind        InputKind `mapstructure:"input_kind" json:"input_kind"`
		DefaultValue     string    `mapstructure:"default_value" json:"default_value"`
		Dropdown         Dropdown  `json:"dropdown"`
	} `json:"inputs"`
	APIs struct {
		AuthorisationRequest Endpoint `mapstructure:"authorisation_request" json:"authorisation_request"`
		AccessTokenRequest   Endpoint `mapstructure:"access_token_request" json:"access_token_request"`
		RefreshTokenRequest  Endpoint `mapstructure:"refresh_token_request" json:"refresh_token_request"`
	} `json:"apis"`
}

// APIKey holds the necessary information for getting authorisation via api keys.
type APIKey struct {
	Inputs map[string]struct {
		Label            string    `json:"label"`
		Tip              string    `json:"tip"`
		IsRequired       bool      `mapstructure:"is_required" json:"is_required"`
		IsAdministrative bool      `mapstructure:"is_administrative" json:"is_administrative"`
		InputKind        InputKind `mapstructure:"input_kind" json:"input_kind"`
		DefaultValue     string    `mapstructure:"default_value" json:"default_value"`
		Dropdown         Dropdown  `json:"dropdown"`
	} `json:"inputs"`
	API Endpoint `json:"api"`
}

// Trigger specifies a trigger operation.
type Trigger struct {
	Label  string           `json:"label"`
	Tip    string           `json:"tip"`
	Inputs map[string]Input `json:"inputs"`
	APIs   struct {
		Polls     []Poll     `json:"polls"`
		RestHooks []RestHook `json:"resthook"`
	} `json:"apis"`
}

// Action specifies an action operation.
type Action struct {
	Label      string           `json:"label"`
	Tip        string           `json:"tip"`
	ActionKind ActionKind       `json:"action_kind"`
	Inputs     map[string]Input `json:"inputs"`
	API        Endpoint         `json:"api"`
}

// Input describes an input field.
type Input struct {
	Label            string    `json:"label"`
	Tip              string    `json:"tip"`
	IsRequired       bool      `mapstructure:"is_required" json:"is_required"`
	IsWriteOp        bool      `mapstructure:"is_write_op" json:"is_write_op"`
	IsIdentification bool      `mapstructure:"is_identification" json:"is_identification"`
	ResourceNoun     string    `json:"resource_noun"`
	InputKind        InputKind `mapstructure:"input_kind" json:"input_kind"`
	Dropdown         Dropdown  `json:"dropdown"`
}

// Poll describes how a trigger polls data.
type Poll struct {
	Endpoint `json:"endpoint"`
	AuthKind AuthKind `json:"auth_kind"`
}

// RestHook describes a webhook trigger API.
type RestHook struct {
	AuthKind   AuthKind `json:"auth_kind"`
	Operations struct {
		Subscribe   Endpoint `json:"subscribe"`
		Unsubscribe Endpoint `json:"unsubscribe"`
		List        Endpoint `json:"list"`
	} `json:"operations"`
}

// Endpoint specifies how a resource is resoved, fetched, updated, etc.
type Endpoint struct {
	Code     string `json:"code"`
	Language string `json:"language"`
	Form     struct {
		Method  string            `json:"method"`
		URL     string            `json:"url"`
		Headers map[string]string `json:"headers"`
		Params  map[string]string `json:"params"`
		Body    map[string]string `json:"body"`
	} `json:"form"`
}

// Dropdown specifies dropdown information.
type Dropdown struct {
	Kind           DropdownKind `json:"kind"`
	AllowsMultiple bool         `mapstructure:"allows_multiple" json:"allows_multiple"`
	AllowsCustom   bool         `mapstructure:"allows_custom" json:"allows_custom"`
	Options        []string     `json:"options"`
}

// NewIntegrationConfig creates an IntegrationConfig from string. Supports JSON, TOML and YAML string format.
func NewIntegrationConfig(integrationString string, format ConfigFormat) (IntegrationConfig, error) {
	// TODO: Sec: Validation
	integration := IntegrationConfig{}
	reader := strings.NewReader(integrationString)

	// Set format to parse.
	converter := viper.New()
	converter.SetConfigType(string(format))
	converter.ReadConfig(reader)

	// Unmarshal string into object.
	if err := converter.Unmarshal(&integration, getCustomDecoder()); err != nil {
		return IntegrationConfig{}, err
	}

	return integration, nil
}

// JSON converts config to json.
func (config *IntegrationConfig) JSON() (string, error) {
	// TODO: Sec: Validation
	bytes, err := json.Marshal(config)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
