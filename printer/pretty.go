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
	PrettyPostRenderFunc func(body string) (string, error)
)

var (
	// PrettyLayoutGoTpl prints all version data in a 'key  value' manner.
	PrettyLayoutGoTpl = `
{{ header .Meta.CLIName }}

  {{ key "Version" }}             {{ .Version                     | val }}
  {{ key "Git Commit" }}          {{ .GitCommit  | commit         | val }}
  {{ key "Build Date" }}          {{ .BuildDate  | fmtDate        | val }}
  {{ key "Commit Date" }}         {{ .CommitDate | fmtDate        | val }}
  {{ key "Dirty Build" }}         {{ .DirtyBuild | fmtBool        | val }}
  {{ key "Go version" }}          {{ .GoVersion  | trimPrefix "go"| val }}
  {{ key "Compiler" }}            {{ .Compiler                    | val }}
  {{ key "Platform" }}            {{ .Platform                    | val }}
`
)

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
	isSmartTerminal := term.IsSmart(w)
	out, err := p.execute(in, isSmartTerminal)
	if err != nil {
		return err
	}

	if p.postRenderFunc != nil {
		out, err = p.postRenderFunc(out)
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
