package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"go.szostok.io/version"
	"go.szostok.io/version/term"
	"go.szostok.io/version/upgrade"
)

const recheckInterval = 30 * time.Second

func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "An example CLI built with github.com/spf13/cobra",
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			gh := upgrade.NewGitHubDetector(
				"mszostok", "codeowners-validator",
				upgrade.WithMinElapseTimeForRecheck(recheckInterval),
			)

			rel, err := gh.LookForGreaterRelease(upgrade.LookForGreaterReleaseInput{
				CurrentVersion: version.Get().Version,
			})
			if err != nil {
				return err
			}

			if !rel.Found {
				// no new version available
				return nil
			}

			if rel.ReleaseInfo.IsFromCache {
				// The time for re-checking for a new release has not elapsed yet,
				// so the cached version is returned.
				//
				// NOTE: This check is run for all CLI commands. You can display it, but it can be quite spammy.
				// For example, the GitHub CLI `gh` displays an upgrade notice only once per 24h,
				// so it acts more like a "reminder".
				// However, IT'S ALWAYS YOUR CALL :)
				// You can even display information if it was from cache or not, see:
				//   https://github.com/mszostok/version/blob/main/examples/upgrade-notice-sub-cmd/main.go
				return nil
			}

			// NOTE:
			// 1. Print the upgrade notice on a standard error channel (stderr).
			//    As a result, output processing for a given command works properly even if the upgrade notice is displayed.
			//
			// 2. Use 'term.IsSmart' so that the renderer can disable colored output for non-tty output streams.
			out, err := gh.Render(rel.ReleaseInfo, term.IsSmart(cmd.OutOrStderr()))
			if err != nil {
				return err
			}

			_, err = fmt.Fprint(cmd.OutOrStderr(), out)
			return err
		},
	}

	return cmd
}

func main() {
	cmd := NewRoot()

	// Register example commands, the upgrade notice should be displayed only once per 30s.
	// It doesn't matter which command you will run.
	cmd.AddCommand(
		NewHelloWorld(),
		NewHakunaMatata(),
	)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func NewHelloWorld() *cobra.Command {
	return &cobra.Command{
		Use:   "hello [name]",
		Short: "Greets to the given arg",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello %s\n", args[0])
		},
	}
}

func NewHakunaMatata() *cobra.Command {
	return &cobra.Command{
		Use:   "hakuna",
		Short: "Says hakuna matata",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hakuna Matata!")
		},
	}
}
