package extension

import (
	"strings"

	"github.com/spf13/cobra"

	"go.szostok.io/version"
	"go.szostok.io/version/printer"
)

var example = `
<cli> version
<cli> version -o=json
<cli> version -o=yaml
<cli> version -o=short
`

// NewVersionCobraCmd returns a root cobra.Command for printing CLI version.
func NewVersionCobraCmd(opts ...CobraOption) *cobra.Command {
	options := CobraOptions{
		Aliases: []string{"ver"},
	}

	for _, customize := range opts {
		customize.ApplyToCobraOption(&options)
	}

	verPrinter := printer.New(options.PrinterOptions...)

	ver := &cobra.Command{
		Use:     "version",
		Short:   "Print the CLI version",
		Example: strings.ReplaceAll(example, "<cli>", version.Get().Meta.CLIName),
		Aliases: options.Aliases,
		RunE: func(cmd *cobra.Command, args []string) error {
			if options.PreHook != nil {
				if err := options.PreHook(cmd.Context()); err != nil {
					return err
				}
			}

			if err := verPrinter.Print(cmd.OutOrStdout()); err != nil {
				return err
			}

			if options.PostHook != nil {
				return options.PostHook(cmd.Context())
			}

			return nil
		},
	}

	verPrinter.RegisterPFlags(ver.Flags())
	return ver
}
