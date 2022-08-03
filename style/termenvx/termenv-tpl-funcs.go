package termenvx

import (
	"text/template"

	"github.com/muesli/termenv"
)

// TemplateFuncs contains a few useful template helpers.
// Copied from: https://github.com/muesli/termenv/blob/v0.12.0/templatehelper.go
// Applied adjustments from: https://github.com/muesli/termenv/pull/14
func TemplateFuncs(p termenv.Profile) template.FuncMap {
	if p == termenv.Ascii {
		return noopTemplateFuncs
	}
	return template.FuncMap{
		"Color": func(values ...interface{}) string {
			s := termenv.String(values[len(values)-1].(string))
			switch len(values) {
			case 2:
				s = s.Foreground(p.Color(values[0].(string)))
			case 3:
				s = s.
					Foreground(p.Color(values[0].(string))).
					Background(p.Color(values[1].(string)))
			}

			return s.String()
		},
		"Foreground": func(values ...interface{}) string {
			s := termenv.String(values[len(values)-1].(string))
			if len(values) == 2 {
				s = s.Foreground(p.Color(values[0].(string)))
			}

			return s.String()
		},
		"Background": func(values ...interface{}) string {
			s := termenv.String(values[len(values)-1].(string))
			if len(values) == 2 {
				s = s.Background(p.Color(values[0].(string)))
			}

			return s.String()
		},
		"Bold":      styleFunc(termenv.Style.Bold),
		"Faint":     styleFunc(termenv.Style.Faint),
		"Italic":    styleFunc(termenv.Style.Italic),
		"Underline": styleFunc(termenv.Style.Underline),
		"Overline":  styleFunc(termenv.Style.Overline),
		"Blink":     styleFunc(termenv.Style.Blink),
		"Reverse":   styleFunc(termenv.Style.Reverse),
		"CrossOut":  styleFunc(termenv.Style.CrossOut),
	}
}

func styleFunc(f func(termenv.Style) termenv.Style) func(...interface{}) string {
	return func(values ...interface{}) string {
		s := termenv.String(values[0].(string))
		return f(s).String()
	}
}

var noopTemplateFuncs = template.FuncMap{
	"Color":      noColorFunc,
	"Foreground": noColorFunc,
	"Background": noColorFunc,
	"Bold":       noStyleFunc,
	"Faint":      noStyleFunc,
	"Italic":     noStyleFunc,
	"Underline":  noStyleFunc,
	"Overline":   noStyleFunc,
	"Blink":      noStyleFunc,
	"Reverse":    noStyleFunc,
	"CrossOut":   noStyleFunc,
}

func noColorFunc(in string) string {
	return in
}

func noStyleFunc(in string) string {
	return in
}
