package style

type (
	// Formatting holds formatting settings.
	Formatting struct {
		Header Header `json:"header" yaml:"header"`
		Key    Key    `json:"key"    yaml:"key"`
		Val    Val    `json:"val"    yaml:"val"`
		Date   Date   `json:"date"   yaml:"date"`
	}

	// Header holds template 'header' func related settings.
	Header struct {
		Prefix          string `json:"prefix"  yaml:"prefix"`
		FormatPrimitive `json:",inline" yaml:",inline"`
		Name            string `json:"name"    yaml:"name"`
	}

	// Key holds template 'key' func related settings.
	Key struct {
		FormatPrimitive `json:",inline"  yaml:",inline"`
	}

	// Val holds template 'val' func related settings.
	Val struct {
		FormatPrimitive `json:",inline"  yaml:",inline"`
	}

	// Date holds template 'date' func related settings.
	Date struct {
		EnableHumanizedSuffix bool `json:"enableHumanizedSuffix" yaml:"enableHumanizedSuffix"`
	}

	// FormatPrimitive holds general formatting options.
	FormatPrimitive struct {
		Color      string   `json:"color"       yaml:"color"`
		Background string   `json:"background"  yaml:"background"`
		Options    []string `json:"options"     yaml:"options"`
	}
)

var defaultFormatting = Formatting{
	Header: Header{
		Prefix: "▓▓▓ ",
		FormatPrimitive: FormatPrimitive{
			Color: "magenta",
		},
	},
	Key: Key{
		FormatPrimitive{
			Color:   "gray",
			Options: []string{"bold"},
		},
	},
	Val: Val{
		FormatPrimitive{
			Color: "white",
		},
	},
	Date: Date{
		EnableHumanizedSuffix: true,
	},
}
