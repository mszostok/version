package printer

import (
	"fmt"
	"io"

	"go.szostok.io/version"
)

var _ Printer = &Short{}

// Short prints only the version param.
type Short struct{}

// Print writes a version number into a given writer.
func (p *Short) Print(in *version.Info, w io.Writer) error {
	if in == nil {
		return nil
	}
	_, err := fmt.Fprintln(w, in.Version)
	return err
}
