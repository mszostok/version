package termenvx

import (
	"github.com/muesli/termenv"
)

func ColorTermenv(p termenv.Profile, in string) termenv.Color {
	color, found := ColorsMapping[in]
	if found {
		return p.Convert(color)
	}
	return p.Color(in)
}

func ColorOptionsTermenv(style termenv.Style, opt string) termenv.Style {
	switch opt {
	case "Bold":
		return style.Bold()
	case "Faint":
		return style.Faint()
	case "Italic":
		return style.Italic()
	case "Underline":
		return style.Underline()
	case "Overline":
		return style.Overline()
	case "Blink":
		return style.Blink()
	case "Reverse":
		return style.Reverse()
	case "CrossOut":
		return style.CrossOut()
	}
	return style
}
