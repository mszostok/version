package printer

import (
	"fmt"
	"io"

	"github.com/hokaccha/go-prettyjson"

	"go.szostok.io/version"
)

var _ Printer = &JSON{}

// JSON prints data in JSON format.
type JSON struct{}

// Print marshals input data to JSON format and writes it to a given writer.
func (p *JSON) Print(in *version.Info, w io.Writer) error {
	out, err := prettyjson.Marshal(in)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(w, string(out))
	return err
}
