package main

import (
	"log"
	"os"

	"go.szostok.io/version"
	"go.szostok.io/version/style"
)

func main() {
	formatting := style.DefaultFormatting()
	formatting.Header = style.Header{
		Prefix: "💡 ",
		FormatPrimitive: style.FormatPrimitive{
			Color: "Magenta",
			Options: []string{
				"Underline",
			},
		},
	}
	formatting.Key.Color = "Yellow"
	formatting.Val.Color = ""
	printer := version.NewPrinter(version.WithPrettyFormatting(&formatting))
	if err := printer.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
