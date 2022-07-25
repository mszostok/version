package style

// Config holds style configuration.
type Config struct {
	Formatting Formatting `json:"formatting" yaml:"formatting"`
	Layout     Layout     `json:"layout" yaml:"layout"`
}

// DefaultConfig returns default style config.
func DefaultConfig() Config {
	return Config{
		Formatting: defaultFormatting,
		Layout: Layout{
			Raw: KeyValueLayoutTpl,
		},
	}
}
