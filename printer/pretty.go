package printer

import (
	"fmt"
	"io"

	"go.szostok.io/version"
	"go.szostok.io/version/style"
	"go.szostok.io/version/term"
)

var _ Printer = &Pretty{}

type (
	// PrettyRenderFunc represents render function signature.
	PrettyRenderFunc func(in *version.Info, isSmartTerminal bool) (string, error)
	// PrettyPostRenderFunc represents post render function signature.
	PrettyPostRenderFunc func(body string, isSmartTerminal bool) (string, error)
)

// PrettyLayoutGoTpl prints all version data in a 'key  value' manner.
var PrettyLayoutGoTpl = `{{ AdjustKeyWidth .ExtraFields }}
{{ Header .Meta.CLIName }}

  {{ Key "Version"     }}    {{ .Version                     | Val }}
  {{ Key "Git Commit"  }}    {{ .GitCommit  | Commit         | Val }}
  {{ Key "Build Date"  }}    {{ .BuildDate  | FmtDate        | Val }}
  {{ Key "Commit Date" }}    {{ .CommitDate | FmtDate        | Val }}
  {{ Key "Dirty Build" }}    {{ .DirtyBuild | FmtBool        | Val }}
  {{ Key "Go version"  }}    {{ .GoVersion  | trimPrefix "go"| Val }}
  {{ Key "Compiler"    }}    {{ .Compiler                    | Val }}
  {{ Key "Platform"    }}    {{ .Platform                    | Val }}
  {{- range $item := (.ExtraFields | Extra) }}
  {{ $item.Key | Key   }}    {{ $item.Value | Val }}
  {{- end}}
`

// Pretty prints human-readable version printing.
type Pretty struct {
	customRenderFn PrettyRenderFunc
	postRenderFunc PrettyPostRenderFunc

	defaultRender *style.GoTemplateRender
}

// NewPretty returns a new Pretty instance.
func NewPretty(options ...PrettyOption) *Pretty {
	opts := PrettyOptions{
		RenderConfig: style.DefaultConfig(PrettyLayoutGoTpl),
	}
	for _, customize := range options {
		customize.ApplyToPrettyOption(&opts)
	}

	return &Pretty{
		customRenderFn: opts.CustomRenderFn,
		postRenderFunc: opts.PostRenderFunc,
		defaultRender:  style.NewGoTemplateRender(opts.RenderConfig),
	}
}

// Print prints a human-readable input version Info into a given writter.
func (p *Pretty) Print(in *version.Info, w io.Writer) error {
	if in == nil {
		return nil
	}

	isSmartTerminal := term.IsSmart(w)
	out, err := p.execute(in, isSmartTerminal)
	if err != nil {
		return err
	}

	if p.postRenderFunc != nil {
		out, err = p.postRenderFunc(out, isSmartTerminal)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprint(w, out)
	return err
}

func (p *Pretty) execute(in *version.Info, isSmartTerminal bool) (string, error) {
	if p.customRenderFn == nil {
		return p.defaultRender.Render(in, isSmartTerminal)
	}

	return p.customRenderFn(in, isSmartTerminal)
}

// PrettyDefaultRenderConfig returns the default render configuration when no customizations are provided.
func PrettyDefaultRenderConfig() *style.Config {
	return style.DefaultConfig(PrettyLayoutGoTpl)
}
