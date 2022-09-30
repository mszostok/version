package main

import (
	"log"
	"os"
	"time"

	"go.szostok.io/version"
	"go.szostok.io/version/style"
	"go.szostok.io/version/upgrade"
)

var layoutGoTpl = `
  │ A new release is available: {{ .Version | Red }} → {{ .NewVersion | Green }}
  │ {{ .ReleaseURL  | Underline | Blue }}
  │
  │ {{ "Resolved from cache:" | Italic }} {{ .IsFromCache | FmtBool | Italic }}
`

func main() {
	upgradeOpts := []upgrade.Options{
		upgrade.WithLayout(&style.Layout{
			// Learn more at https://version.szostok.io/customization/upgrade-notice/layout/
			GoTemplate: layoutGoTpl,
		}),
		upgrade.WithMinElapseTimeForRecheck(time.Second * 30),
	}

	ghUpgrade := upgrade.NewGitHubDetector(
		"mszostok", "codeowners-validator",
		upgradeOpts...,
	)

	// it's good to print on stderr as user can still grep the output without dealing with adhoc upgrade message
	err := ghUpgrade.PrintIfFoundGreater(os.Stderr, version.Get().Version)
	if err != nil {
		log.Fatal(err)
	}
}
