package main

import (
	"log"
	"os"

	"github.com/spf13/pflag"

	"go.szostok.io/version/printer"
)

func main() {
	p := printer.New()
	p.RegisterPFlags(pflag.CommandLine) // register `--output/-o` flag
	pflag.Parse()

	if err := p.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
