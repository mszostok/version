package version

import (
	"fmt"
	"io"

	"github.com/mszostok/version/style"
)

var _ Printer = &Pretty{}

type PrettyRenderFunc func(in *Info) (string, error)

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

	// PrettyBoxLayoutGoTpl prints all version data in box.
	// https://knowyourmeme.com/memes/this-is-fine
	PrettyBoxLayoutGoTpl = `
╭───{{ repeatMax 57 "─" header }}{{/* ─────────────────────────────────── */}}╮
│                                  {{ repeatMax 25 " " ""                  }} │
│  {{ key "Version" }}             {{ .Version                     | val   }} │
│  {{ key "Git Commit" }}          {{ .GitCommit  | commit         | val   }} │
│  {{ key "Build Date" }}          {{ .BuildDate  | fmtDate        | val   }} │
│  {{ key "Commit Date" }}         {{ .CommitDate | fmtDate        | val   }} │
│  {{ key "Dirty Build" }}         {{ .DirtyBuild | fmtBool        | val   }} │
│  {{ key "Go Version" }}          {{ .GoVersion  | trimPrefix "go"| val   }} │
│  {{ key "Compiler" }}            {{ .Compiler                    | val   }} │
│  {{ key "Platform" }}            {{ .Platform                    | val   }} │
╰───{{ repeatMax 57 "─" ""}}{{/* ──────────────────────────────────────── */}}╯
`
)

// Pretty prints human-readable version.
type Pretty struct {
	customRenderFn      PrettyRenderFunc
	defaultRender       *style.GoTemplateRender
	defaultRenderConfig *style.Config
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

	_, err = fmt.Fprint(w, out)
	return err
}

func (p *Pretty) execute(in *Info) (string, error) {
	if p.customRenderFn == nil {
		return p.defaultRender.Render(in)
	}

	return p.customRenderFn(in)
}
