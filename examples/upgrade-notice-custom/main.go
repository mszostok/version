package main

import (
	"log"
	"os"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/muesli/reflow/indent"
	"github.com/spf13/pflag"

	"go.szostok.io/version/printer"
	"go.szostok.io/version/style"
	"go.szostok.io/version/upgrade"
)

func main() {
	upgradeOpts := []upgrade.Options{
		upgrade.WithLayout(&style.Layout{
			GoTemplate: forBoxLayoutGoTpl,
		}),
		upgrade.WithPostRenderHook(SprintInBox),
	}

	p := printer.New(
		printer.WithUpgradeNotice("mszostok", "codeowners-validator", upgradeOpts...),
	)
	p.RegisterPFlags(pflag.CommandLine) // register `--output/-o` flag
	pflag.Parse()

	if err := p.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

var forBoxLayoutGoTpl = heredoc.Doc(`
A new release is available: {{ .Version | Red }} → {{ .NewVersion | Green }}
{{ .ReleaseURL  | Underline | Blue }}`)

func SprintInBox(body string) (string, error) {
	cfg := box.Config{Px: 1, Py: 0, Type: "Round", Color: "Yellow", ContentAlign: "Left"}
	boxed := box.New(cfg)

	body = boxed.String("", body)
	body = indent.String(body, 2)
	return body, nil
}
