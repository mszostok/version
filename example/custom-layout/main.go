package main

import (
	"log"
	"os"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/fatih/color"
	"github.com/muesli/reflow/indent"

	"go.szostok.io/version"
	"go.szostok.io/version/printer"
	"go.szostok.io/version/style"
)

func main() {
	opts := []printer.ContainerOption{
		printer.WithPrettyPostRenderHook(SprintInBox),
		printer.WithPrettyLayout(style.Layout{
			GoTemplate: layoutGoTpl,
		}),
	}
	p := printer.New(opts...)
	if err := p.Print(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

var layoutGoTpl = heredoc.Doc(`
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
