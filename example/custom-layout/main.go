package main

import (
	"log"
	"os"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/fatih/color"
	"github.com/muesli/reflow/indent"

	"go.szostok.io/version"
	"go.szostok.io/version/style"
)

func main() {
	version.CollectFromBuildInfo()

	opts := []version.PrinterContainerOption{
		version.WithPrettyPostRenderHook(SprintInBox),
		version.WithPrettyLayout(style.Layout{
			GoTemplate: BoxLayoutGoTpl,
		}),
	}
	printer := version.NewPrinter(opts...)
	if err := printer.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

// BoxLayoutGoTpl prints all version data in box.
// https://knowyourmeme.com/memes/this-is-fine
var BoxLayoutGoTpl = heredoc.Doc(`
  {{ key "Version" }}             {{ .Version                     | val   }}
  {{ key "Git Commit" }}          {{ .GitCommit  | commit         | val   }}
  {{ key "Build Date" }}          {{ .BuildDate  | fmtDate        | val   }}
  {{ key "Commit Date" }}         {{ .CommitDate | fmtDate        | val   }}
  {{ key "Dirty Build" }}         {{ .DirtyBuild | fmtBool        | val   }}
  {{ key "Go Version" }}          {{ .GoVersion  | trimPrefix "go"| val   }}
  {{ key "Compiler" }}            {{ .Compiler                    | val   }}
  {{ key "Platform" }}            {{ .Platform                    | val   }}`)

func SprintInBox(body string) (string, error) {
	cfg := box.Config{Px: 2, Py: 1, Type: "Round", Color: "Yellow", ContentAlign: "Left", TitlePos: "Top"}
	boxed := box.New(cfg)

	body = boxed.String(color.MagentaString("▓▓▓ %s", version.Get().Meta.CLIName), body)
	body = indent.String(body, 2)
	return "\n" + body, nil
}
