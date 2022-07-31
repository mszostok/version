package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/mszostok/version"
	"github.com/mszostok/version/upgrade"
)

func NewVersionWithCheck() *cobra.Command {
	verCmd := version.NewCobraCmd()

	verCmd.AddCommand(&cobra.Command{
		Use:   "check",
		Short: "Check GitHub for a new release",
		RunE: func(cmd *cobra.Command, args []string) error {
			ghUpgrade := upgrade.NewGitHubDetector(
				"mszostok", "codeowners-validator",
				upgrade.WithMinElapseTimeForRecheck(time.Second),
			)

			checkInput := upgrade.LookForLatestReleaseInput{
				CurrentVersion: version.Get().Version,
			}
			out, err := ghUpgrade.LookForLatestRelease(checkInput)
			if err != nil {
				return err
			}

			if !out.Found {
				return nil
			}

			body, err := ghUpgrade.Render(out.ReleaseInfo)
			if err != nil {
				return err
			}
			_, err = fmt.Fprint(cmd.ErrOrStderr(), "\n"+body)
			return err
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
