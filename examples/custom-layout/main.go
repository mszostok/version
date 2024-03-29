package main

import (
	"log"
	"os"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/MakeNowJust/heredoc/v2"
	fcolor "github.com/fatih/color"
	"github.com/gookit/color"

	"github.com/muesli/reflow/indent"

	"go.szostok.io/version"
	"go.szostok.io/version/printer"
	"go.szostok.io/version/style"
)

func main() {
	styleFromEnv, err := printer.WithPrettyStyleFromEnv("CLI_STYLE")
	exitOnErr(err)

	p := printer.New(
		printer.WithPrettyPostRenderHook(SprintInBox),
		printer.WithPrettyLayout(&style.Layout{
			GoTemplate: layoutGoTpl,
		}),
		styleFromEnv,
	)

	err = p.Print(os.Stdout)
	exitOnErr(err)
}

var layoutGoTpl = heredoc.Doc(`
  {{ Key "Version"     }}        {{ .Version                     | Val   }}
  {{ Key "Git Commit"  }}        {{ .GitCommit  | Commit         | Val   }}
  {{ Key "Build Date"  }}        {{ .BuildDate  | FmtDate        | Val   }}
  {{ Key "Commit Date" }}        {{ .CommitDate | FmtDate        | Val   }}
  {{ Key "Dirty Build" }}        {{ .DirtyBuild | FmtBool        | Val   }}
  {{ Key "Go Version"  }}        {{ .GoVersion  | trimPrefix "go"| Val   }}
  {{ Key "Compiler"    }}        {{ .Compiler                    | Val   }}
  {{ Key "Platform"    }}        {{ .Platform                    | Val   }}`)

func SprintInBox(body string, isSmartTerminal bool) (string, error) {
	cfg := box.Config{Px: 2, Py: 1, Type: "Round", Color: "Yellow", ContentAlign: "Left", TitlePos: "Top"}
	boxed := box.New(cfg)

	body = boxed.String(fcolor.MagentaString("▓▓▓ %s", version.Get().Meta.CLIName), body)
	body = indent.String(body, 2)
	if !isSmartTerminal {
		body = color.ClearCode(body)
	}
	return "\n" + body, nil
}

func exitOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
