package style

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/araddon/dateparse"
	"github.com/dustin/go-humanize"
	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
)

// Render provides a functionality to render input data based on a given style.
type Render struct {
	config Config
}

// NewRender returns new Render instance.
func NewRender() *Render {
	return &Render{
		config: DefaultConfig(),
	}
}

// Render renders input data based on configured given style.
func (r *Render) Render(in any) (string, error) {
	tpl, err := template.New("pretty").
		Funcs(sprig.FuncMap()).
		Funcs(colorFuncMap).
		Funcs(r.styleFuncMap()).
		Funcs(r.generalHelpersFuncMap()).
		Parse(r.config.Layout.GoTemplate)
	if err != nil {
		return "", err
	}

	var buff bytes.Buffer
	if err := tpl.Execute(&buff, in); err != nil {
		return "", err
	}

	return buff.String(), nil
}

func (r *Render) header() string {
	c := newGookitStyle(r.config.Formatting.Header.PropertyFormat)
	name := r.config.Formatting.Header.Name
	if name == "" {
		name = os.Args[0]
	}
	return c.Sprintf("%s%s", r.config.Formatting.Header.Prefix, name)
}

func (r *Render) key(in string) string {
	c := newGookitStyle(r.config.Formatting.Key.PropertyFormat)

	return c.Sprint(in)
}

func (*Render) commit(in string) string {
	return strings.TrimSpace(fmt.Sprintf("%.7s", in))
}

func (r *Render) val(in string) string {
	c := newGookitStyle(r.config.Formatting.Val.PropertyFormat)
	return c.Sprintf("%-*s", 37, in)
}

func (*Render) repeatMax(max int, sing, in string) string {
	max -= runewidth.StringWidth(color.ClearCode(in))
	return fmt.Sprintf("%s%s", in, strings.Repeat(sing, max))
}

func (*Render) fmtBool(in bool) string {
	if in {
		return "yes"
	}
	return "no"
}

func (r *Render) fmtDate(in string) string {
	if in == "" {
		return ""
	}

	t, err := dateparse.ParseAny(in)
	if err != nil {
		return ""
	}

	suffix := ""
	if r.config.Formatting.Date.EnableHumanizedSuffix {
		suffix = " (" + humanize.Time(t) + ")"
	}
	return fmt.Sprintf("%s%s", t.Local().Format(time.RFC822), suffix)
}
