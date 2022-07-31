package style

// Config holds style configuration.
type Config struct {
	Formatting Formatting `json:"formatting,omitempty" yaml:"formatting,omitempty"`
	Layout     Layout     `json:"layout,omitempty" yaml:"layout,omitempty"`
}

// DefaultConfig returns default style config.
func DefaultConfig(layoutGoTpl string) *Config {
	return &Config{
		Formatting: defaultFormatting,
		Layout: Layout{
			GoTemplate: layoutGoTpl,
		},
	}
}

func DefaultFormatting() Formatting {
	return defaultFormatting
}
