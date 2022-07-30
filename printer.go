package version

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/spf13/pflag"

	"github.com/mszostok/version/upgrade"
)

// Printer is an interface that knows how to print Info object.
type Printer interface {
	// Print receives Info, formats it and prints it to a writer.
	Print(in *Info, w io.Writer) error
}

// PrinterContainer provides functionality to print version info in requested format.
// Can be configured with pflag.FlagSet.
type PrinterContainer struct {
	output OutputFormat

	printers          map[OutputFormat]Printer
	name              string
	upgradeNotice     *upgrade.GitHubDetector
	upgradeNoticeChan chan bool
}

type PrinterContainerOptions struct {
	prettyOptions []PrettyPrinterOption
	upgradeNotice *upgrade.GitHubDetector
}

// NewPrinter returns a new PrinterContainer instance.
func NewPrinter(customize ...PrinterContainerOption) *PrinterContainer {
	var opts PrinterContainerOptions
	for _, opt := range customize {
		opt.ApplyToPrinterContainerOption(&opts)
	}

	p := &PrinterContainer{
		printers: map[OutputFormat]Printer{
			JSONFormat:   &JSON{},
			YAMLFormat:   &YAML{},
			ShortFormat:  &Short{},
			PrettyFormat: NewPrettyPrinter(opts.prettyOptions...),
		},
		name:              os.Args[0],
		output:            PrettyFormat,
		upgradeNotice:     opts.upgradeNotice,
		upgradeNoticeChan: make(chan bool, 1),
	}

	return p
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
	go r.checkReleasesIfNeeded()

	printer, found := r.printers[r.output]
	if !found {
		return fmt.Errorf("printer %q is not available", r.output)
	}

	if err := printer.Print(Get(r.name), w); err != nil {
		return err
	}

	if err := r.printUpgradeNoticeIfAvailable(); err != nil {
		return err
	}

	return nil
}

// PrintInfo prints received Info object in requested format.
func (r *PrinterContainer) PrintInfo(w io.Writer, in *Info) error {
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

func (r *PrinterContainer) printUpgradeNoticeIfAvailable() error {
	if r.upgradeNotice == nil {
		return nil
	}

	detectedNewVersion := <-r.upgradeNoticeChan
	if !detectedNewVersion {
		return nil
	}

	upgradeNotice, err := r.upgradeNotice.Render()
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(os.Stderr, "\n"+upgradeNotice) // TODO: customize os.Stderr/os.Stdout/file?
	return err
}

func (r *PrinterContainer) checkReleasesIfNeeded() {
	if r.upgradeNotice == nil {
		return
	}
	r.upgradeNoticeChan <- r.upgradeNotice.CheckForUpdate(version)
}
