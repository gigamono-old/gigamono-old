package configs

import "github.com/sageflow/sageflow/pkg/database/models/auth"

// AuthInfoConfig contains app-specific information for establishing auth.
type AuthInfoConfig struct {
	Model    *auth.AppCredentials
	Version  uint       `json:"version"`
	Kind     ConfigKind `json:"kind"`
	Metadata struct {
		AppName string   `mapstructure:"app_name" json:"app_name"`
		AppID   UUID     `mapstructure:"app_id" json:"app_id"`
		Authors []Author `json:"authors"`
	} `json:"metadata"`
	Auths struct {
		OAuth2s []struct {
			Fields map[string][]string `json:"fields"`
			Envs   struct {
				ClientID     string `mapstructure:"client_id" json:"client_id"`
				ClientSecret string `mapstructure:"client_secret" json:"client_secret"`
				RedirectURI  string `mapstructure:"redirect_uri" json:"redirect_uri"`
			} `json:"envs"`
		} `json:"oauth2s"`
		APIKeys []struct {
			Fields map[string][]string `json:"fields"`
			Envs   struct {
				APIKey string `mapstructure:"api_key" json:"api_key"`
			} `json:"envs"`
		} `json:"api_keys"`
	} `json:"auths"`
}
