package version

import (
	"fmt"
	"io"
)

var _ Printer = &Short{}

// Short prints short version.
type Short struct{}

// Print marshals input data to JSON format and writes it to a given writer.
func (p *Short) Print(in Info, w io.Writer) error {
	_, err := fmt.Fprintln(w, in.Version)
	return err
}
