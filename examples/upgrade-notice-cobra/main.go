package main

import (
	"os"

	"github.com/spf13/cobra"

	"go.szostok.io/version/extension"
)

// NewRoot returns a root cobra.Command for the whole CLI.
func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "An example CLI built with github.com/spf13/cobra",
	}

	cmd.AddCommand(
		// 1. Register 'version' command
		extension.NewVersionCobraCmd(
			// 2. Explict turn on upgrade notice
			extension.WithUpgradeNotice("mszostok", "codeowners-validator"),
		),
	)

	return cmd
}

func main() {
	if err := NewRoot().Execute(); err != nil {
		os.Exit(1)
	}
}
