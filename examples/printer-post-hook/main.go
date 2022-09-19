package main

import (
	"log"
	"os"

	"github.com/spf13/pflag"

	"go.szostok.io/version"
	"go.szostok.io/version/printer"
	"go.szostok.io/version/upgrade"
)

func main() {
	ghUpgradeNotice := upgrade.NewGitHubDetector(
		"mszostok", "codeowners-validator",
	)

	p := printer.New(
		printer.WithPostHook(func() error {
			return ghUpgradeNotice.PrintIfFoundGreater(os.Stderr, version.Get().Version)
		}),
	)
	p.RegisterPFlags(pflag.CommandLine) // register `--output/-o` flag
	pflag.Parse()

	if err := p.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
