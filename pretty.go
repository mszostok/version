package version

import (
	"fmt"
	"io"

	"github.com/mszostok/version/style"
)

var _ Printer = &Pretty{}

type prettyRenderer interface {
	Render(in any) (string, error)
}

// Pretty prints human-readable version.
type Pretty struct {
	render prettyRenderer
}

func NewPrettyPrinter() *Pretty {
	return &Pretty{
		render: style.NewRender(),
	}
}

func (p *Pretty) Print(in Info, w io.Writer) error {
	out, err := p.render.Render(in)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(w, out)
	return err
}
