package configs

// SageflowConfig holds Sageflow configurations.
type SageflowConfig struct {
	Execution struct {
		UseSubprocess bool `mapstructure:"use_subprocess"`
	}
}
