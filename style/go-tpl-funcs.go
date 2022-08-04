package style

import (
	"strings"
	"text/template"

	"github.com/muesli/termenv"

	"go.szostok.io/version/style/termenvx"
)

func (r *GoTemplateRender) styleFuncMap() template.FuncMap {
	return template.FuncMap{
		"header":           r.header,
		"key":              r.key,
		"val":              r.val,
		"fmtDate":          r.fmtDate,
		"fmtDateHumanized": r.fmtDateHumanized,
		"fmtBool":          r.fmtBool,
	}
}

func (r *GoTemplateRender) generalHelpersFuncMap() template.FuncMap {
	return template.FuncMap{
		"commit": r.commit,
	}
}

func (r *GoTemplateRender) termenvColorFuncMap() template.FuncMap {
	return template.FuncMap{
		// Foreground colors
		"Black":   r.colorSprintf(termenvx.ColorsMapping["Black"]),
		"Red":     r.colorSprintf(termenvx.ColorsMapping["Red"]),
		"Green":   r.colorSprintf(termenvx.ColorsMapping["Green"]),
		"Yellow":  r.colorSprintf(termenvx.ColorsMapping["Yellow"]),
		"Blue":    r.colorSprintf(termenvx.ColorsMapping["Blue"]),
		"Magenta": r.colorSprintf(termenvx.ColorsMapping["Magenta"]),
		"Cyan":    r.colorSprintf(termenvx.ColorsMapping["Cyan"]),
		"White":   r.colorSprintf(termenvx.ColorsMapping["White"]),
		"Gray":    r.colorSprintf(termenvx.ColorsMapping["Gray"]),

		// Background colors
		"BgBlack":   r.bgColorSprintf(termenvx.ColorsMapping["Black"]),
		"BgRed":     r.bgColorSprintf(termenvx.ColorsMapping["Red"]),
		"BgGreen":   r.bgColorSprintf(termenvx.ColorsMapping["Green"]),
		"BgYellow":  r.bgColorSprintf(termenvx.ColorsMapping["Yellow"]),
		"BgBlue":    r.bgColorSprintf(termenvx.ColorsMapping["Blue"]),
		"BgMagenta": r.bgColorSprintf(termenvx.ColorsMapping["Magenta"]),
		"BgCyan":    r.bgColorSprintf(termenvx.ColorsMapping["Cyan"]),
		"BgWhite":   r.bgColorSprintf(termenvx.ColorsMapping["White"]),
		"BgGray":    r.bgColorSprintf(termenvx.ColorsMapping["Gray"]),
	}
}

func (r *GoTemplateRender) colorSprintf(c termenv.Color) func(in ...string) string {
	return func(in ...string) string {
		return termenv.
			String().
			Foreground(r.colorProfile.Convert(c)).
			Styled(strings.Join(in, " "))
	}
}

func (r *GoTemplateRender) bgColorSprintf(c termenv.Color) func(in ...string) string {
	return func(in ...string) string {
		return termenv.
			String().
			Background(r.colorProfile.Convert(c)).
			Styled(strings.Join(in, " "))
	}
}

func (r *GoTemplateRender) termenvHelpersFuncMap() template.FuncMap {
	return termenvx.TemplateFuncs(r.colorProfile)
}
