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
		// you just need to add this, and you are done.
		extension.NewVersionCobraCmd(),
	)

	return cmd
}

func main() {
	if err := NewRoot().Execute(); err != nil {
		os.Exit(1)
	}
}
