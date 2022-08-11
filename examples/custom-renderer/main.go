package main

import (
	"log"
	"os"

	"github.com/Delta456/box-cli-maker/v2"

	"go.szostok.io/version"
	"go.szostok.io/version/printer"
	"go.szostok.io/version/style"
)

func main() {
	p := printer.New(printer.WithPrettyRenderer(prettyRender))
	if err := p.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func prettyRender(in *version.Info, isSmartTerminal bool) (string, error) {
	renderBody := style.NewGoTemplateRender(style.DefaultConfig(printer.PrettyLayoutGoTpl))
	body, err := renderBody.Render(in, isSmartTerminal)
	if err != nil {
		return "", err
	}

	Box := box.New(box.Config{Px: 0, Py: 0, Type: "Round", Color: "Yellow", ContentAlign: "Left"})
	Box.TitlePos = "Top"

	return Box.String("Box", body), nil
}
