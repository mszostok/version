package termenvx

import "github.com/muesli/termenv"

var ColorsMapping = map[string]termenv.Color{
	// Foreground colors
	"Black":   termenv.ANSIBlack,
	"Red":     termenv.ANSIRed,
	"Green":   termenv.ANSIGreen,
	"Yellow":  termenv.ANSIYellow,
	"Blue":    termenv.ANSIBlue,
	"Magenta": termenv.ANSIMagenta,
	"Cyan":    termenv.ANSICyan,
	"White":   termenv.ANSIWhite,
	"Gray":    termenv.RGBColor("#93a1a1"),
}
