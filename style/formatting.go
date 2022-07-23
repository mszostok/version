package style

type Formatting struct {
	Header Header `json:"header" yaml:"header"`
	Date   Date   `json:"date" yaml:"date"`
}

type Date struct {
	EnableHumanizedSuffix bool `json:"enableHumanizedSuffix" yaml:"enableHumanizedSuffix"`
}

type Header struct {
	Prefix string `json:"prefix" yaml:"prefix"`
	Color  string `json:"color" yaml:"color"`
	Bold   bool   `json:"bold" yaml:"bold"`
	Name   string `json:"name" yaml:"name"`
}

var DefaultFormatting = Formatting{
	Header: Header{
		Prefix: "▓▓▓ ",
		Color:  "magenta",
		Bold:   false,
	},
	Date: Date{
		EnableHumanizedSuffix: true,
	},
}
