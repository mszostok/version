package style

import (
	"github.com/muesli/termenv"

	"go.szostok.io/version/style/termenvx"
)

func NewTermenvStyle(p termenv.Profile, in FormatPrimitive) termenv.Style {
	printer := termenv.String()
	if p == termenv.Ascii {
		return printer // no settings
	}

	if in.Color != "" {
		printer = printer.Foreground(termenvx.ColorTermenv(p, in.Color))
	}
	if in.Background != "" {
		printer = printer.Background(termenvx.ColorTermenv(p, in.Background))
	}

	for _, opt := range in.Options {
		printer = termenvx.ColorOptionsTermenv(printer, opt)
	}

	return printer
}
