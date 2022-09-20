package termenvx

import (
	"os"
	"strings"

	"github.com/muesli/termenv"
)

func ColorProfile() termenv.Profile {
	var (
		term      = os.Getenv("TERM")
		colorTerm = os.Getenv("COLORTERM")
	)

	switch strings.ToLower(colorTerm) {
	case "24bit", "truecolor":
		if strings.HasPrefix(term, "screen") {
			// tmux supports TrueColor, screen only ANSI256
			if os.Getenv("TERM_PROGRAM") != "tmux" {
				return termenv.ANSI256
			}
		}
		return termenv.TrueColor
	case "yes", "true":
		return termenv.ANSI256
	}

	switch term {
	case "xterm-kitty":
		return termenv.TrueColor
	case "linux":
		return termenv.ANSI
	}

	if strings.Contains(term, "256color") {
		return termenv.ANSI256
	}
	if strings.Contains(term, "color") {
		return termenv.ANSI
	}
	if strings.Contains(term, "ansi") {
		return termenv.ANSI
	}

	return termenv.Ascii
}
