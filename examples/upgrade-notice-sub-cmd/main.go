package main

import (
	"os"
	"time"

	"github.com/spf13/cobra"

	"go.szostok.io/version"
	"go.szostok.io/version/extension"
	"go.szostok.io/version/style"
	"go.szostok.io/version/upgrade"
)

var layoutGoTpl = `
  │ A new release is available: {{ .Version | Red }} → {{ .NewVersion | Green }}
  │ {{ .ReleaseURL  | Underline | Blue }}
  │
  │ {{ "Resolved from cache:" | Italic }} {{ .IsFromCache | FmtBool | Italic }}
`

func NewVersionWithCheck() *cobra.Command {
	verCmd := extension.NewVersionCobraCmd()

	verCmd.AddCommand(&cobra.Command{
		Use:   "check",
		Short: "Check GitHub for a new release",
		RunE: func(cmd *cobra.Command, args []string) error {
			ghUpgrade := upgrade.NewGitHubDetector(
				"mszostok", "codeowners-validator",
				upgrade.WithMinElapseTimeForRecheck(30*time.Second),
				upgrade.WithLayout(&style.Layout{
					// Learn more at https://version.szostok.io/customization/upgrade-notice/layout/
					GoTemplate: layoutGoTpl,
				}),
			)
			return ghUpgrade.PrintIfFoundGreater(cmd.ErrOrStderr(), version.Get().Version)
		},
	})

	return verCmd
}

// NewRoot returns a root cobra.Command for the whole CLI.
func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "An example CLI built with github.com/spf13/cobra",
	}

	cmd.AddCommand(
		NewVersionWithCheck(),
	)

	return cmd
}

func main() {
	if err := NewRoot().Execute(); err != nil {
		os.Exit(1)
	}
}
