package style

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/araddon/dateparse"
	"github.com/dustin/go-humanize"
	"github.com/muesli/termenv"

	"go.szostok.io/version/style/termenvx"
)

// GoTemplateRender provides a functionality to render input data based on a given style.
type GoTemplateRender struct {
	config       *Config
	colorProfile termenv.Profile
}

// NewGoTemplateRender returns new GoTemplateRender instance.
func NewGoTemplateRender(cfg *Config) *GoTemplateRender {
	renderer := GoTemplateRender{
		config:       cfg,
		colorProfile: termenv.Ascii,
	}

	return &renderer
}

// Render renders input data based on configured given style.
// If isSmartTerminal is set to 'true', colors and formatting are used.
func (r *GoTemplateRender) Render(in interface{}, isSmartTerminal bool) (string, error) {
	if isSmartTerminal {
		r.colorProfile = termenvx.ColorProfile()
	}
	tpl, err := template.New("tpl").
		Funcs(sprig.TxtFuncMap()).
		Funcs(r.styleFuncMap()).
		Funcs(r.generalHelpersFuncMap()).
		Funcs(r.termenvColorFuncMap()).
		Funcs(r.termenvHelpersFuncMap()).
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

func (r *GoTemplateRender) header(in string) string {
	c := NewTermenvStyle(r.colorProfile, r.config.Formatting.Header.FormatPrimitive)
	return c.Styled(fmt.Sprintf("%s%s", r.config.Formatting.Header.Prefix, in))
}

func (r *GoTemplateRender) key(in string) string {
	c := NewTermenvStyle(r.colorProfile, r.config.Formatting.Key.FormatPrimitive)
	return c.Styled(in)
}

func (*GoTemplateRender) commit(in string) string {
	return strings.TrimSpace(fmt.Sprintf("%.7s", in))
}

func (r *GoTemplateRender) val(in string) string {
	c := NewTermenvStyle(r.colorProfile, r.config.Formatting.Val.FormatPrimitive)
	return c.Styled(in)
}

func (*GoTemplateRender) fmtBool(in bool) string {
	if in {
		return "yes"
	}
	return "no"
}

func (r *GoTemplateRender) fmtDate(in interface{}) string {
	var normalized time.Time
	switch date := in.(type) {
	case time.Time:
		normalized = date
	case string:
		t, err := dateparse.ParseAny(date)
		if err != nil {
			return ""
		}
		normalized = t
	}
	suffix := ""
	if r.config.Formatting.Date.EnableHumanizedSuffix {
		suffix = " (" + humanize.Time(normalized) + ")"
	}
	return fmt.Sprintf("%s%s", normalized.Local().Format(time.RFC822), suffix)
}

func (r *GoTemplateRender) fmtDateHumanized(in interface{}) string {
	var normalized time.Time
	switch date := in.(type) {
	case time.Time:
		normalized = date
	case string:
		t, err := dateparse.ParseAny(date)
		if err != nil {
			return ""
		}
		normalized = t
	}
	return humanize.Time(normalized)
}
