package main

import (
	"log"
	"os"

	"github.com/spf13/pflag"

	"go.szostok.io/version/printer"
)

func main() {
	verPrinter := printer.New()
	verPrinter.RegisterPFlags(pflag.CommandLine) // register `--output/-o` flag
	pflag.Parse()

	if err := verPrinter.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
