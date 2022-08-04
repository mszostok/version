package main

import (
	"log"
	"os"

	"github.com/Delta456/box-cli-maker/v2"

	"go.szostok.io/version"
	"go.szostok.io/version/style"
)

func main() {
	printer := version.NewPrinter(version.WithPrettyRenderer(prettyRender))
	if err := printer.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func prettyRender(in *version.Info, isSmartTerminal bool) (string, error) {
	renderBody := style.NewGoTemplateRender(style.DefaultConfig(version.PrettyKVLayoutGoTpl))
	body, err := renderBody.Render(in, isSmartTerminal)
	if err != nil {
		return "", err
	}

	Box := box.New(box.Config{Px: 0, Py: 0, Type: "Round", Color: "Yellow", ContentAlign: "Left"})
	Box.TitlePos = "Top"

	return Box.String("Box", body), nil
}
