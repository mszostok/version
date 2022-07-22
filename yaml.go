package version

import (
	"io"

	"gopkg.in/yaml.v3"
)

var _ Printer = &YAML{}

// YAML prints data in YAML format.
type YAML struct{}

// Print marshals input data to YAML format and writes it to a given writer.
func (p *YAML) Print(in Info, w io.Writer) error {
	out, err := yaml.Marshal(in)
	if err != nil {
		return err
	}

	_, err = w.Write(out)
	return err
}
