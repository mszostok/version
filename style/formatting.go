package style

type (
	Formatting struct {
		Header Header `json:"header" yaml:"header"`
		Key    Key    `json:"key"    yaml:"key"`
		Val    Val    `json:"val"    yaml:"val"`
		Date   Date   `json:"date"   yaml:"date"`
	}
	Header struct {
		Prefix         string `json:"prefix"  yaml:"prefix"`
		PropertyFormat `json:",inline" yaml:",inline"`
		Name           string `json:"name"    yaml:"name"`
	}

	Key struct {
		PropertyFormat `json:",inline"  yaml:",inline"`
	}

	Val struct {
		PropertyFormat `json:",inline"  yaml:",inline"`
	}

	Date struct {
		EnableHumanizedSuffix bool `json:"enableHumanizedSuffix" yaml:"enableHumanizedSuffix"`
	}

	PropertyFormat struct {
		Color   string   `json:"color"   yaml:"color"`
		Options []string `json:"options" yaml:"options"`
	}
)

var DefaultFormatting = Formatting{
	Header: Header{
		Prefix: "▓▓▓ ",
		PropertyFormat: PropertyFormat{
			Color: "magenta",
		},
	},
	Key: Key{
		PropertyFormat{
			Color:   "gray",
			Options: []string{"bold"},
		},
	},
	Val: Val{
		PropertyFormat{
			Color: "white",
		},
	},
	Date: Date{
		EnableHumanizedSuffix: true,
	},
}
