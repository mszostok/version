package printer

import (
	"fmt"
)

// OutputFormat is a type for capturing supported output formats.
// Implements pflag.Value interface.
type OutputFormat string

// ErrInvalidOutputFormatType is returned when an unsupported format type is used.
var ErrInvalidOutputFormatType = fmt.Errorf("invalid output format type")

const (
	// PrettyFormat represents human-readable format.
	PrettyFormat OutputFormat = "pretty"
	// JSONFormat represents JSON data format.
	JSONFormat OutputFormat = "json"
	// YAMLFormat represents YAML data format.
	YAMLFormat OutputFormat = "yaml"
	// ShortFormat represents short (version only) format.
	ShortFormat OutputFormat = "short"
)

// IsValid returns true if OutputFormat is valid.
func (o OutputFormat) IsValid() bool {
	switch o {
	case PrettyFormat, JSONFormat, YAMLFormat, ShortFormat:
		return true
	}
	return false
}

// String returns the string representation of the Format. Required by pflag.Value interface.
func (o OutputFormat) String() string {
	return string(o)
}

// Set format type to a given input. Required by pflag.Value interface.
func (o *OutputFormat) Set(in string) error {
	*o = OutputFormat(in)
	if !o.IsValid() {
		return ErrInvalidOutputFormatType
	}
	return nil
}

// Type returns data type. Required by pflag.Value interface.
func (o *OutputFormat) Type() string {
	return "string"
}
