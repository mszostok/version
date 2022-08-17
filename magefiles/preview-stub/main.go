package main

import (
	"log"
	"os"
	"time"

	"go.szostok.io/version"
	"go.szostok.io/version/printer"
	"go.szostok.io/version/style"
	"go.szostok.io/version/upgrade"
)

var upgradeTpl = `
  │ A new release is available: {{ .Version }} → {{ "1.4.2" | Green }}
  │ {{ "https://github.com/mszostok/gimme/releases/tag/v1.4.2"  | Underline | Blue }}
`

func main() {
	p := printer.New(
		printer.WithUpgradeNotice("mszostok", "codeowners-validator",
			upgrade.WithMinElapseTimeForRecheck(time.Second),
			upgrade.WithIsVersionGreater(func(string, string) bool {
				return true
			}),
			upgrade.WithLayout(&style.Layout{
				GoTemplate: upgradeTpl,
			})),
	)

	info := version.Get()
	info.Meta.CLIName = "gimme"
	if err := p.PrintInfo(os.Stdout, info); err != nil {
		log.Fatal(err)
	}
}
