package printer

import (
	fmt "fmt"
	"io"

	"github.com/goccy/go-yaml/lexer"
	"github.com/goccy/go-yaml/printer"
	"github.com/muesli/termenv"
	"gopkg.in/yaml.v3"

	"go.szostok.io/version"
	"go.szostok.io/version/style/termenvx"
	"go.szostok.io/version/term"
)

var _ Printer = &YAML{}

// YAML prints data in YAML format.
type YAML struct{}

// Print marshals input data to YAML format and writes it to a given writer.
// Prints colored output only if a given writer supports that.
func (y *YAML) Print(in *version.Info, w io.Writer) error {
	out, err := yaml.Marshal(in)
	if err != nil {
		return fmt.Errorf("while marshaling: %w", err)
	}
	cp := y.colorProfile(w)

	data := y.colorizedYAML(cp, string(out))

	_, err = fmt.Fprintln(w, data)
	return err
}

func (*YAML) colorProfile(w io.Writer) termenv.Profile {
	if term.IsSmart(w) {
		return termenvx.ColorProfile()
	}

	return termenv.Ascii
}

func (*YAML) escape(p termenv.Profile) string {
	if p.Convert(termenv.ANSIWhite).Sequence(false) == "" {
		return ""
	}
	return fmt.Sprintf("%sm", termenv.CSI+termenv.ResetSeq)
}

func (*YAML) color(p termenv.Profile, color termenv.Color) string {
	seq := p.Convert(color).Sequence(false)
	if seq == "" {
		return ""
	}
	return fmt.Sprintf("%s%sm", termenv.CSI, seq)
}

func (y *YAML) colorizedYAML(cp termenv.Profile, in string) string {
	suffix := y.escape(cp)

	tokens := lexer.Tokenize(in)
	var p printer.Printer
	p.Bool = func() *printer.Property {
		return &printer.Property{
			Prefix: y.color(cp, termenv.ANSIYellow),
			Suffix: suffix,
		}
	}
	p.Number = func() *printer.Property {
		return &printer.Property{
			Prefix: y.color(cp, termenv.ANSIMagenta),
			Suffix: suffix,
		}
	}
	p.MapKey = func() *printer.Property {
		return &printer.Property{
			Prefix: y.color(cp, termenv.ANSIWhite),
			Suffix: suffix,
		}
	}
	p.Anchor = func() *printer.Property {
		return &printer.Property{
			Prefix: y.color(cp, termenv.ANSIBrightYellow),
			Suffix: suffix,
		}
	}
	p.Alias = func() *printer.Property {
		return &printer.Property{
			Prefix: y.color(cp, termenv.ANSIBrightYellow),
			Suffix: suffix,
		}
	}
	p.String = func() *printer.Property {
		return &printer.Property{
			Prefix: y.color(cp, termenv.ANSIBrightGreen),
			Suffix: suffix,
		}
	}

	return p.PrintTokens(tokens)
}
