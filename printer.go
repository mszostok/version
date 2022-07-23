package version

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/spf13/pflag"
)

// Printer is an interface that knows how to print Info object.
type Printer interface {
	// Print receives Info, formats it and prints it to a writer.
	Print(in Info, w io.Writer) error
}

// PrinterContainer provides functionality to print version info in requested format.
// Can be configured with pflag.FlagSet.
type PrinterContainer struct {
	output OutputFormat

	printers map[OutputFormat]Printer
	name     string
}

// NewPrinter returns a new PrinterContainer instance.
func NewPrinter(opts ...PrinterOption) *PrinterContainer {
	p := &PrinterContainer{
		printers: map[OutputFormat]Printer{
			JSONFormat:   &JSON{},
			YAMLFormat:   &YAML{},
			ShortFormat:  &Short{},
			PrettyFormat: &Pretty{},
		},
		name:   os.Args[0],
		output: PrettyFormat,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

// PrinterOption allows PrinterContainer instance customization.
type PrinterOption func(printer *PrinterContainer)

// WithDefaultOutputFormat sets a default format type.
func WithDefaultOutputFormat(format OutputFormat) PrinterOption {
	return func(r *PrinterContainer) {
		r.output = format
	}
}

// WithCLIName sets a custom CLI name.
func WithCLIName(name string) PrinterOption {
	return func(r *PrinterContainer) {
		r.name = name
	}
}

// RegisterPFlags registers PrinterContainer terminal flags.
func (r *PrinterContainer) RegisterPFlags(flags *pflag.FlagSet) {
	flags.VarP(&r.output, "output", "o", fmt.Sprintf("Output format. One of: %s", r.availablePrinters()))
}

// OutputFormat returns default print format type.
func (r *PrinterContainer) OutputFormat() OutputFormat {
	return r.output
}

// Print prints Info object in requested format.
func (r *PrinterContainer) Print(w io.Writer) error {
	printer, found := r.printers[r.output]
	if !found {
		return fmt.Errorf("printer %q is not available", r.output)
	}

	return printer.Print(Get(r.name), w)
}

// PrintInfo prints received Info object in requested format.
func (r *PrinterContainer) PrintInfo(w io.Writer, in Info) error {
	printer, found := r.printers[r.output]
	if !found {
		return fmt.Errorf("printer %q is not available", r.output)
	}

	return printer.Print(in, w)
}

func (r *PrinterContainer) availablePrinters() string {
	var out []string
	for key := range r.printers {
		out = append(out, key.String())
	}

	// We generate doc automatically, so it needs to be deterministic
	sort.Strings(out)

	return strings.Join(out, " | ")
}
