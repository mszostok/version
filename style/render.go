package style

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"text/template"
	"time"

	"github.com/Masterminds/sprig/v3"
	"github.com/araddon/dateparse"
	"github.com/dustin/go-humanize"
	"github.com/mattn/go-runewidth"
	"github.com/muesli/termenv"
	"gopkg.in/yaml.v3"

	"go.szostok.io/version/style/termenvx"
)

// GoTemplateRender provides a functionality to render input data based on a given style.
type GoTemplateRender struct {
	config       *Config
	colorProfile termenv.Profile
	width        int
}

// NewGoTemplateRender returns new GoTemplateRender instance.
func NewGoTemplateRender(cfg *Config) *GoTemplateRender {
	renderer := GoTemplateRender{
		config:       cfg,
		colorProfile: termenv.Ascii,
		width:        15,
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

func (r *GoTemplateRender) adjustKeyWidth(in any) string {
	if in == nil {
		return ""
	}

	// could be any underlying type
	val := reflect.ValueOf(in)
	// if it's a pointer, resolve its value
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	// should double-check we now have a struct (could still be anything)
	if val.Kind() != reflect.Struct {
		return ""
	}

	// check for the longest name
	structType := val.Type()
	for i := 0; i < structType.NumField(); i++ {
		var (
			field = structType.Field(i)
			tag   = field.Tag
			name  = field.Name
		)
		if val, ok := tag.Lookup("pretty"); ok {
			name = val
		}

		width := runewidth.StringWidth(name)
		if r.width < width {
			r.width = width
		}
	}

	return ""
}

func (r *GoTemplateRender) header(in string) string {
	c := NewTermenvStyle(r.colorProfile, r.config.Formatting.Header.FormatPrimitive)
	return c.Styled(fmt.Sprintf("%s%s", r.config.Formatting.Header.Prefix, in))
}

func (r *GoTemplateRender) key(in string) string {
	ff := r.config.Formatting.Key.FormatPrimitive

	c := NewTermenvStyle(r.colorProfile, ff)
	initial := runewidth.StringWidth(in)
	newV := c.Styled(in)
	varr := runewidth.StringWidth(newV) - initial + 2
	return fmt.Sprintf("%-*s", r.width+varr, newV)
}

type KV struct {
	Key   string
	Value any
}

func (r *GoTemplateRender) extra(in any) []KV {
	var out []KV
	if in == nil {
		return out
	}

	// could be any underlying type
	val := reflect.ValueOf(in)
	// if its a pointer, resolve its value
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	// should double-check we now have a struct (could still be anything)
	if val.Kind() != reflect.Struct {
		return out
	}

	// now we grab our values
	structType := val.Type()
	for i := 0; i < structType.NumField(); i++ {
		var (
			field = structType.Field(i)
			tag   = field.Tag
			name  = field.Name
		)

		if val, ok := tag.Lookup("pretty"); ok {
			name = val
		}
		out = append(out, KV{
			Key:   name,
			Value: val.Field(i).Interface(),
		})
	}

	return out
}

func (*GoTemplateRender) commit(in string) string {
	return strings.TrimSpace(fmt.Sprintf("%.7s", in))
}

func (r *GoTemplateRender) val(in any) string {
	c := NewTermenvStyle(r.colorProfile, r.config.Formatting.Val.FormatPrimitive)

	switch str := in.(type) {
	case string:
		return c.Styled(str)
	default:
		out, err := yaml.Marshal(in)
		if err != nil {
			return err.Error()
		}
		return c.Styled(string(out))
	}
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
