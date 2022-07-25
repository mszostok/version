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

type Render struct {
	config Config
}

func NewRender() *Render {
	return &Render{
		config: DefaultConfig(),
	}
}

func (r *Render) Render(in any) (string, error) {
	tpl, err := template.New("pretty").
		Funcs(sprig.FuncMap()).
		Funcs(colorFuncMap).
		Funcs(template.FuncMap{
			"header":    r.header,
			"key":       r.key,
			"val":       r.val,
			"commit":    r.commit,
			"fmtDate":   r.fmtDate,
			"fmtBool":   r.fmtBool,
			"repeatMax": r.repeatMax,
		}).Parse(r.config.Layout.Raw)
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
		suffix = "(" + humanize.Time(t) + ")"
	}
	return fmt.Sprintf("%s%s", t.Local().Format(time.RFC822), suffix)
}
