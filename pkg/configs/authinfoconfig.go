package configs

// AuthInfoConfig contains app-specific information for establishing auth.
type AuthInfoConfig struct {
	ID       UUID
	Version  uint
	Kind     string
	Metadata struct {
		AppName string `mapstructure:"app_name"`
		AppID   UUID   `mapstructure:"app_id"`
	}
	Auths struct {
		OAuth2s []struct {
			Fields map[string][]string
			Envs   struct {
				ClientID     string `mapstructure:"client_id"`
				ClientSecret string `mapstructure:"client_secret"`
				RedirectURI  string `mapstructure:"redirect_uri"`
			}
		}
		APIKeys []struct {
			Fields map[string][]string
			Envs   struct {
				APIKey string `mapstructure:"api_key"`
			}
		}
	}
}
