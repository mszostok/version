package version

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// TODO: func options to inform if server version
// TODO: to add template?

var example = `
<cli> version
<cli> version -o=json
<cli> version -o=yaml
<cli> version -o=pretty
<cli> version -o=short
`

// NewCobraCmd returns a root cobra.Command for printing CLI version.
func NewCobraCmd() *cobra.Command {
	printer := NewPrinter()
	CollectFromBuildInfo()

	ver := &cobra.Command{
		Use:     "version",
		Short:   "Print the CLI version",
		Example: strings.ReplaceAll(example, "<cli>", os.Args[0]),
		RunE: func(cmd *cobra.Command, args []string) error {
			return printer.Print(cmd.OutOrStdout())
		},
	}

	printer.RegisterPFlags(ver.Flags())
	return ver
}
