package style

import (
	"bytes"
	"fmt"
	"os"
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
// The forOutput is needed to check if colors and formatting is supported.
func NewGoTemplateRender(cfg *Config) *GoTemplateRender {
	renderer := GoTemplateRender{
		config:       cfg,
		colorProfile: termenv.Ascii,
	}

	return &renderer
}

// Render renders input data based on configured given style.
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

func (r *GoTemplateRender) header() string {
	c := NewTermenvStyle(r.colorProfile, r.config.Formatting.Header.FormatPrimitive)

	name := r.config.Formatting.Header.Name
	if name == "" {
		// name = version.Get().Meta.CLIName TODO: after package split
		name = os.Args[0]
	}
	return c.Styled(fmt.Sprintf("%s%s", r.config.Formatting.Header.Prefix, name))
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

func (r *GoTemplateRender) fmtDate(in string) string {
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
