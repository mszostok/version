package version

import (
	"fmt"
	"io"

	"github.com/mszostok/version/style"
)

var _ Printer = &Pretty{}

type (
	// PrettyRenderFunc represents render function signature.
	PrettyRenderFunc func(in *Info) (string, error)
	// PrettyPostRenderFunc represents post render function signature.
	PrettyPostRenderFunc func(body string) (string, error)
)

var (
	// PrettyKVLayoutGoTpl prints all version data in a 'key  value' manner.
	PrettyKVLayoutGoTpl = `
{{ header }}

  {{ key "Version" }}             {{ .Version                     | val }}
  {{ key "Git Commit" }}          {{ .GitCommit  | commit         | val }}
  {{ key "Build Date" }}          {{ .BuildDate  | fmtDate        | val }}
  {{ key "Commit Date" }}         {{ .CommitDate | fmtDate        | val }}
  {{ key "Dirty Build" }}         {{ .DirtyBuild | fmtBool        | val }}
  {{ key "Go Version" }}          {{ .GoVersion  | trimPrefix "go"| val }}
  {{ key "Compiler" }}            {{ .Compiler                    | val }}
  {{ key "Platform" }}            {{ .Platform                    | val }}
`
)

// Pretty prints human-readable version.
type Pretty struct {
	customRenderFn      PrettyRenderFunc
	defaultRender       *style.GoTemplateRender
	defaultRenderConfig *style.Config
	postRenderFunc      PrettyPostRenderFunc
}

func NewPrettyPrinter(opts ...PrettyPrinterOption) *Pretty {
	p := &Pretty{
		defaultRenderConfig: style.DefaultConfig(PrettyKVLayoutGoTpl),
	}

	for _, opt := range opts {
		opt.ApplyPrettyPrinterOption(p)
	}

	p.defaultRender = style.NewGoTemplateRender(p.defaultRenderConfig)

	return p
}

func (p *Pretty) Print(in *Info, w io.Writer) error {
	out, err := p.execute(in)
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

func (p *Pretty) execute(in *Info) (string, error) {
	if p.customRenderFn == nil {
		return p.defaultRender.Render(in)
	}

	return p.customRenderFn(in)
}
