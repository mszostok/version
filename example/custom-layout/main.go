package main

import (
	"log"
	"os"

	"github.com/mszostok/version"
	"github.com/mszostok/version/style"
)

func main() {
	version.CollectFromBuildInfo()

	printer := version.NewPrinter(version.WithPrettyLayout(style.Layout{
		GoTemplate: style.BoxLayoutGoTpl,
	}))
	if err := printer.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
