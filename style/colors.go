package style

import (
	"html/template"
	"strings"

	"github.com/gookit/color"
)

func colorFuncMap() template.FuncMap {
	return template.FuncMap{
		// Foreground colors
		"black":        colorSprintf(color.FgBlack),
		"red":          colorSprintf(color.FgRed),
		"green":        colorSprintf(color.FgGreen),
		"yellow":       colorSprintf(color.FgYellow),
		"blue":         colorSprintf(color.FgBlue),
		"magenta":      colorSprintf(color.FgMagenta),
		"cyan":         colorSprintf(color.FgCyan),
		"white":        colorSprintf(color.FgWhite),
		"lightRed":     colorSprintf(color.FgLightRed),
		"lightGreen":   colorSprintf(color.FgLightGreen),
		"lightYellow":  colorSprintf(color.FgLightYellow),
		"lightBlue":    colorSprintf(color.FgLightBlue),
		"lightMagenta": colorSprintf(color.FgLightMagenta),
		"lightCyan":    colorSprintf(color.FgLightCyan),
		"lightWhite":   colorSprintf(color.FgLightWhite),
		"gray":         colorSprintf(color.FgGray),

		// Option settings
		"opBold":          colorSprintf(color.OpBold),
		"opFuzzy":         colorSprintf(color.OpFuzzy),
		"opItalic":        colorSprintf(color.OpItalic),
		"opUnderscore":    colorSprintf(color.OpUnderscore),
		"opReverse":       colorSprintf(color.OpReverse),
		"opConcealed":     colorSprintf(color.OpConcealed),
		"opStrikethrough": colorSprintf(color.OpStrikethrough),

		// Background colors
		"bgBlack":        colorSprintf(color.BgBlack),
		"bgRed":          colorSprintf(color.BgRed),
		"bgGreen":        colorSprintf(color.BgGreen),
		"bgYellow":       colorSprintf(color.BgYellow),
		"bgBlue":         colorSprintf(color.BgBlue),
		"bgMagenta":      colorSprintf(color.BgMagenta),
		"bgCyan":         colorSprintf(color.BgCyan),
		"bgWhite":        colorSprintf(color.BgWhite),
		"bgLightRed":     colorSprintf(color.BgLightRed),
		"bgLightGreen":   colorSprintf(color.BgLightGreen),
		"bgLightYellow":  colorSprintf(color.BgLightYellow),
		"bgLightBlue":    colorSprintf(color.BgLightBlue),
		"bgLightMagenta": colorSprintf(color.BgLightMagenta),
		"bgLightCyan":    colorSprintf(color.BgLightCyan),
		"bgLightWhite":   colorSprintf(color.BgLightWhite),
		"bgGray":         colorSprintf(color.BgGray),
	}
}

func colorSprintf(opts ...color.Color) func(in ...string) string {
	return func(in ...string) string {
		mg := color.New(opts...).Sprintf
		return mg(strings.Join(in, " "))
	}
}
