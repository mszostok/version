package main

import (
	"log"
	"os"

	"go.szostok.io/version"
	"go.szostok.io/version/printer"
)

type Custom struct {
	// In the pretty mode, fields are printed in the same order as defined in struct.
	BuiltBy string `json:"builtBy" yaml:"builtBy" pretty:"Built By"`
	RepoURL string `json:"repoURL" yaml:"repoURL" pretty:"Repository URL"`
	DocsURL string `json:"docsURL" yaml:"docsURL" pretty:"Documentation URL"`
}

func main() {
	custom := Custom{
		RepoURL: "https://github.com/mszostok/version",
		DocsURL: "https://szostok.io/projects/version",
		BuiltBy: "GoReleaser",
	}

	info := version.Get()
	info.ExtraFields = custom

	p := printer.New()
	if err := p.PrintInfo(os.Stdout, info); err != nil {
		log.Fatal(err)
	}
}
