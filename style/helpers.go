package style

import (
	"strings"

	"github.com/gookit/color"
)

func colorSprintf(opts ...color.Color) func(in ...string) string {
	return func(in ...string) string {
		mg := color.New(opts...).Sprintf
		return mg(strings.Join(in, " "))
	}
}

func newGookitStyle(in PropertyFormat) color.Style {
	c := color.New()
	if in.Color != "" {
		c.Add(colorGookit(in.Color))
	}
	if in.Background != "" {
		c.Add(backgroundGookit(in.Background))
	}

	for _, opt := range in.Options {
		c.Add(color.AllOptions[opt])
	}

	return c
}

func colorGookit(in string) color.Color {
	if strings.HasPrefix(in, "#") {
		return color.HEX(in).Color()
	}
	if cs, found := color.FgColors[in]; found {
		return cs
	}
	if cs, found := color.ExFgColors[in]; found {
		return cs
	}

	return 0
}

func backgroundGookit(in string) color.Color {
	if strings.HasPrefix(in, "#") {
		return color.HEX(in, true).Color()
	}
	if cs, found := color.BgColors[in]; found {
		return cs
	}
	if cs, found := color.ExBgColors[in]; found {
		return cs
	}

	return 0
}
