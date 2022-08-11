package main

import (
	"log"
	"os"

	"go.szostok.io/version/printer"
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

	p := printer.New(printer.WithPrettyFormatting(&formatting))
	if err := p.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
