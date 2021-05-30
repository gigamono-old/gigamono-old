package configs

// IntegrationConfig holds the information about an app which can be used in tasks.
type IntegrationConfig struct {
	Version  uint       `json:"version"`
	Kind     ConfigKind `json:"kind"`
	Metadata struct {
		Name          string   `json:"name"`
		PublicID      UUID     `mapstructure:"public_id" json:"public_id"`
		Version       string   `json:"version"`
		Description   string   `json:"description"`
		Category      string   `json:"category"`
		Tags          []string `json:"tags"`
		AvatarURL     string   `mapstructure:"avatar_url" json:"avatar_url"`
		HomepageURL   string   `json:"homepage_url"`
		ResourceNouns []string `json:"resource_nouns"`
		Authors       []Author `json:"authors"`
	} `json:"metadata"`
	Auths struct {
		OAuth2s []OAuth2 `json:"oauth2s"`
		APIKeys []APIKey `json:"api_keys"`
	} `json:"auths"`
	Operations struct {
		Triggers []Trigger `json:"triggers"`
		Actions  []Action  `json:"actions"`
	} `json:"operations"`
}

// OAuth2 holds the necessary information for getting authorisation via OAuth2.
type OAuth2 struct {
	Label                      string   `json:"label"`
	Scopes                     []string `json:"scopes"`
	ShouldRefreshAutomatically bool     `mapstructure:"should_refresh_automatically" json:"should_refresh_automatically"`
	Fields                     []struct {
		Label            string    `json:"label"`
		Key              string    `json:"key"`
		Tip              string    `json:"tip"`
		IsRequired       bool      `mapstructure:"is_required" json:"is_required"`
		IsAdministrative bool      `mapstructure:"is_administrative" json:"is_administrative"`
		InputKind        InputKind `mapstructure:"input_kind" json:"input_kind"`
		DefaultValue     string    `mapstructure:"default_value" json:"default_value"`
		Dropdown         Dropdown  `json:"dropdown"`
	} `json:"fields"`
	APIs struct {
		AuthorisationRequest Endpoint `mapstructure:"authorisation_request" json:"authorisation_request"`
		AccessTokenRequest   Endpoint `mapstructure:"access_token_request" json:"access_token_request"`
		RefreshTokenRequest  Endpoint `mapstructure:"refresh_token_request" json:"refresh_token_request"`
	} `json:"apis"`
}

// APIKey holds the necessary information for getting authorisation via api keys.
type APIKey struct {
	Fields []struct {
		Label            string    `json:"label"`
		Key              string    `json:"key"`
		Tip              string    `json:"tip"`
		IsRequired       bool      `mapstructure:"is_required" json:"is_required"`
		IsAdministrative bool      `mapstructure:"is_administrative" json:"is_administrative"`
		InputKind        InputKind `mapstructure:"input_kind" json:"input_kind"`
		DefaultValue     string    `mapstructure:"default_value" json:"default_value"`
		Dropdown         Dropdown  `json:"dropdown"`
	} `json:"fields"`
	API Endpoint `json:"api"`
}

// Trigger specifies a trigger operation.
type Trigger struct {
	Key    string  `json:"key"`
	Label  string  `json:"label"`
	Tip    string  `json:"tip"`
	Fields []Field `json:"field"`
	APIs   struct {
		Polls     []Poll     `json:"polls"`
		RestHooks []RestHook `json:"resthook"`
	} `json:"apis"`
}

// Action specifies an action operation.
type Action struct {
	Key        string     `json:"key"`
	Label      string     `json:"label"`
	Tip        string     `json:"tip"`
	ActionKind ActionKind `json:"action_kind"`
	Fields     []Field    `json:"field"`
	API        Endpoint   `json:"api"`
}

// Field describes an input field.
type Field struct {
	Key              string    `json:"key"`
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
