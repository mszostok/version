package main

import (
	"os"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/muesli/reflow/indent"
	"github.com/spf13/cobra"

	"github.com/mszostok/version"
	"github.com/mszostok/version/style"
	"github.com/mszostok/version/upgrade"
)

// NewRoot returns a root cobra.Command for the whole CLI.
func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "An example CLI built with github.com/spf13/cobra",
	}

	opts := []upgrade.Options{
		upgrade.WithLayout(&style.Layout{
			GoTemplate: forBoxLayoutGoTpl,
		}),
		upgrade.WithPostRenderHook(func(body string) (string, error) {
			return body + "\ncustom footer", nil
		}),
	}

	format := style.DefaultFormatting()
	format.Header.Color = "yellow"
	cmd.AddCommand(
		// 1. Register 'version' command
		version.NewCobraCmd(
			version.WithPrettyFormatting(&format),
			// 2. Explict turn on upgrade notice
			version.WithUpgradeNotice("mszostok", "codeowners-validator", opts...)),
	)

	return cmd
}

func main() {
	if err := NewRoot().Execute(); err != nil {
		os.Exit(1)
	}
}

var forBoxLayoutGoTpl = heredoc.Doc(`
A new release is available: {{ .Version }} → {{ .NewVersion | green }}
{{ .ReleaseURL  | underscore | blue }}`)

func SprintInBox(body string) (string, error) {
	cfg := box.Config{Px: 1, Py: 0, Type: "Round", Color: "Yellow", ContentAlign: "Left"}
	boxed := box.New(cfg)

	body = boxed.String("", body)
	body = indent.String(body, 2)
	return body, nil
}
