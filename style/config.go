package style

type Config struct {
	Formatting Formatting `json:"formatting" yaml:"formatting"`
	Layout     Layout     `json:"layout" yaml:"layout"`
}

func DefaultConfig() Config {
	return Config{
		Formatting: DefaultFormatting,
		Layout: Layout{
			Raw: DefaultLayoutTpl,
		},
	}
}
