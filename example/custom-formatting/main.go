package main

import (
	"log"
	"os"

	"github.com/mszostok/version"
	"github.com/mszostok/version/style"
)

func main() {
	version.CollectFromBuildInfo()

	formatting := style.DefaultFormatting()
	formatting.Header = style.Header{
		Prefix: "💡 ",
		FormatPrimitive: style.FormatPrimitive{
			Color: "magenta",
			Options: []string{
				"underscore",
			},
		},
	}
	formatting.Key.Color = "yellow"
	formatting.Val.Color = ""
	printer := version.NewPrinter(version.WithPrettyFormatting(&formatting))
	if err := printer.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
