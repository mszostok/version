package version

import (
	"fmt"
	"io"
)

var _ Printer = &Short{}

// Short prints only the version param.
type Short struct{}

func (p *Short) Print(in *Info, w io.Writer) error {
	if in == nil {
		return nil
	}
	_, err := fmt.Fprintln(w, in.Version)
	return err
}
