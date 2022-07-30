package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/mszostok/version"
	"github.com/mszostok/version/upgrade"
)

// NewRoot returns a root cobra.Command for the whole CLI.
func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "An example CLI built with github.com/spf13/cobra",
	}

	cmd.AddCommand(
		// 1. Register 'version' command
		version.NewCobraCmd(
			// 2. Explict turn on upgrade notice
			version.WithUpgradeNotice("mszostok", "codeowners-validator", upgrade.WithBoxed(upgrade.BoxYellow)),
		),
	)

	return cmd
}

func main() {
	if err := NewRoot().Execute(); err != nil {
		os.Exit(1)
	}
}
