package version

import (
	"fmt"
	"io"

	"github.com/mszostok/version/style"
)

var _ Printer = &Pretty{}

type PrettyRenderFunc func(in *Info) (string, error)

// Pretty prints human-readable version.
type Pretty struct {
	customRenderFn      PrettyRenderFunc
	defaultRender       *style.GoTemplateRender
	defaultRenderConfig *style.Config
}

func NewPrettyPrinter(opts ...PrettyPrinterOption) *Pretty {
	p := &Pretty{
		defaultRenderConfig: style.DefaultConfig(),
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
