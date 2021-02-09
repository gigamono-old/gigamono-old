package configs

// SageflowConfig holds Sageflow configurations.
type SageflowConfig struct {
	Version  uint
	Kind     ConfigKind
	Metadata struct {
		Authors []Author
	}
	Execution struct {
		UseSubprocess bool `mapstructure:"use_subprocess"`
	}
	Server struct {
		API struct {
			Port int8
		}
		Engine struct {
			Port int8
		}
		Auth struct {
			Port int8
		}
	}
	Database struct {
		Resource struct {
			URI  string
			Type string
		}
		Auth struct {
			URI  string
			Type string
		}
	}
}
