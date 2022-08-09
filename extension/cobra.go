package extension

import (
	"os"
	"strings"

	"github.com/spf13/cobra"

	"go.szostok.io/version/printer"
)

var example = `
<cli> version
<cli> version -o=json
<cli> version -o=yaml
<cli> version -o=pretty
<cli> version -o=short
`

// NewVersionCobraCmd returns a root cobra.Command for printing CLI version.
func NewVersionCobraCmd(opts ...CobraOption) *cobra.Command {
	var options CobraOptions
	for _, customize := range opts {
		customize.ApplyToCobraOption(&options)
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
