package extension

import (
	"os"
	"strings"

	"github.com/spf13/cobra"

	"go.szostok.io/version/printer"
	"go.szostok.io/version/upgrade"
)

var example = `
<cli> version
<cli> version -o=json
<cli> version -o=yaml
<cli> version -o=pretty
<cli> version -o=short
`

type CobraExtensionOption func(*CobraExtensionOptions)

type CobraExtensionOptions struct {
	PrinterOptions []printer.ContainerOption
}

func WithPrinterOptions(opts ...printer.ContainerOption) CobraExtensionOption {
	return func(options *CobraExtensionOptions) {
		options.PrinterOptions = opts
	}
}

func WithUpgradeNotice(owner, repo string, opts ...upgrade.Options) CobraExtensionOption {
	return func(options *CobraExtensionOptions) {
		options.PrinterOptions = append(options.PrinterOptions, printer.WithUpgradeNotice(owner, repo, opts...))
	}
}

// NewVersionCobraCmd returns a root cobra.Command for printing CLI version.
func NewVersionCobraCmd(opts ...CobraExtensionOption) *cobra.Command {
	var options CobraExtensionOptions
	for _, customize := range opts {
		customize(&options)
	}

	verPrinter := printer.New(options.PrinterOptions...)

	ver := &cobra.Command{
		Use:     "version",
		Short:   "Print the CLI version",
		Example: strings.ReplaceAll(example, "<cli>", os.Args[0]),
		RunE: func(cmd *cobra.Command, args []string) error {
			return verPrinter.Print(cmd.OutOrStdout())
		},
	}

	verPrinter.RegisterPFlags(ver.Flags())
	return ver
}
