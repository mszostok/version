package version

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/spf13/pflag"
)

// Printer is an interface that knows how to print objects.
type Printer interface {
	// Print receives an object, formats it and prints it to a writer.
	Print(in Info, w io.Writer) error
}

// VersionPrinter provides functionality to print version info in requested format.
// Can be configured with pflag.FlagSet.
type VersionPrinter struct {
	output OutputFormat

	printers map[OutputFormat]Printer
}

// NewPrinter returns a new VersionPrinter instance.
func NewPrinter(opts ...PrinterOption) *VersionPrinter {
	p := &VersionPrinter{
		printers: map[OutputFormat]Printer{
			JSONFormat:  &JSON{},
			YAMLFormat:  &YAML{},
			ShortFormat: &Short{},
		},
		output: PrettyFormat,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

// PrinterOption allows VersionPrinter instance customization.
type PrinterOption func(printer *VersionPrinter)

// WithDefaultOutputFormat sets a default format type.
func WithDefaultOutputFormat(format OutputFormat) PrinterOption {
	return func(r *VersionPrinter) {
		r.output = format
	}
}

// RegisterFlags registers VersionPrinter terminal flags.
func (r *VersionPrinter) RegisterFlags(flags *pflag.FlagSet) {
	flags.VarP(&r.output, "output", "o", fmt.Sprintf("Output format. One of: %s", r.availablePrinters()))
}

// OutputFormat returns default print format type.
func (r *VersionPrinter) OutputFormat() OutputFormat {
	return r.output
}

// Print prints received object in requested format.
func (r *VersionPrinter) Print(w io.Writer, in Info) error {
	printer, found := r.printers[r.output]
	if !found {
		return fmt.Errorf("printer %q is not available", r.output)
	}

	return printer.Print(in, w)
}

func (r *VersionPrinter) availablePrinters() string {
	var out []string
	for key := range r.printers {
		out = append(out, key.String())
	}

	// We generate doc automatically, so it needs to be deterministic
	sort.Strings(out)

	return strings.Join(out, " | ")
}
