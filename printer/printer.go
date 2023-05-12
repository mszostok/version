package printer

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/spf13/pflag"

	"go.szostok.io/version"
	"go.szostok.io/version/upgrade"
)

// Printer is an interface that knows how to print Info object.
type Printer interface {
	// Print receives Info, formats it and prints it to a writer.
	Print(in *version.Info, w io.Writer) error
}

// Container provides functionality to print version info in requested format.
// Can be configured with pflag.FlagSet.
type Container struct {
	output OutputFormat

	printers      map[OutputFormat]Printer
	postHookFunc  PostHookFunc
	upgradeNotice *upgrade.GitHubDetector
	excludedField version.Field
}

// New returns a new Container instance.
func New(options ...ContainerOption) *Container {
	var opts ContainerOptions
	for _, customize := range options {
		customize.ApplyToContainerOption(&opts)
	}

	return &Container{
		printers: map[OutputFormat]Printer{
			JSONFormat:   &JSON{},
			YAMLFormat:   &YAML{},
			ShortFormat:  &Short{},
			PrettyFormat: NewPretty(opts.PrettyOptions...),
		},
		output:        PrettyFormat,
		postHookFunc:  opts.PostHookFunc,
		upgradeNotice: opts.UpgradeNotice,
		excludedField: opts.ExcludedField,
	}
}

// RegisterPFlags registers `--output/-o` CLI flag.
func (r *Container) RegisterPFlags(flags *pflag.FlagSet) {
	flags.VarP(&r.output, "output", "o", fmt.Sprintf("Output format. One of: %s", r.availablePrinters()))
}

// OutputFormat returns default print format type.
func (r *Container) OutputFormat() OutputFormat {
	return r.output
}

// Print prints Info object in a requested format.
// It's just a syntax sugar for PrintInfo(w, version.GetInfo(excludedFields)).
func (r *Container) Print(w io.Writer) error {
	return r.PrintInfo(w, version.GetInfo(r.excludedField))
}

// PrintInfo prints a given Info object in a requested format.
func (r *Container) PrintInfo(w io.Writer, in *version.Info) error {
	printer, found := r.printers[r.output]
	if !found {
		return fmt.Errorf("printer %q is not available", r.output)
	}

	if err := printer.Print(in, w); err != nil {
		return err
	}

	if r.upgradeNotice != nil {
		if err := r.upgradeNotice.PrintIfFoundGreater(os.Stderr, in.Version); err != nil {
			return err
		}
	}

	if r.postHookFunc != nil { // TODO: should it be called irrespective of whether there was an error before?
		return r.postHookFunc()
	}
	return nil
}

func (r *Container) availablePrinters() string {
	var out []string
	for key := range r.printers {
		out = append(out, key.String())
	}

	// To generate doc automatically, it needs to be deterministic.
	sort.Strings(out)

	return strings.Join(out, " | ")
}
