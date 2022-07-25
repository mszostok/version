package main

import (
	"log"
	"os"

	"github.com/spf13/pflag"

	"github.com/mszostok/version"
)

func main() {
	version.CollectFromBuildInfo()

	printer := version.NewPrinter()
	printer.RegisterPFlags(pflag.CommandLine) // optionally register `--output/-o` flag.
	pflag.Parse()

	if err := printer.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
