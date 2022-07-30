package upgrade

import (
	"github.com/Delta456/box-cli-maker/v2"
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/muesli/reflow/indent"
)

var forBoxLayoutGoTpl = heredoc.Doc(`
		A new release is available: {{ .Version }} → {{ .NewVersion | green }}
		{{ .ReleaseURL  | underscore | blue }}`)

type BoxColor string

const (
	BoxBlack     BoxColor = "Black"
	BoxBlue               = "Blue"
	BoxRed                = "Red"
	BoxGreen              = "Green"
	BoxYellow             = "Yellow"
	BoxCyan               = "Cyan"
	BoxMagenta            = "Magenta"
	BoxWhite              = "White"
	BoxHiBlack            = "HiBlack"
	BoxHiBlue             = "HiBlue"
	BoxHiRed              = "HiRed"
	BoxHiGreen            = "HiGreen"
	BoxHiYellow           = "HiYellow"
	BoxHiCyan             = "HiCyan"
	BoxHiMagenta          = "HiMagenta"
	BoxHiWhite            = "HiWhite"
)

func (b BoxColor) Color() string {
	switch b {
	case BoxBlack, BoxBlue, BoxRed, BoxGreen, BoxYellow, BoxCyan, BoxMagenta, BoxWhite, BoxHiBlack, BoxHiBlue, BoxHiRed, BoxHiGreen, BoxHiYellow, BoxHiCyan, BoxHiMagenta, BoxHiWhite:
		return string(b)
	default:
		return BoxWhite
	}
}

func SprintInBox(body string, color BoxColor) string {
	cfg := box.Config{Px: 1, Py: 0, Type: "Round", Color: color.Color(), ContentAlign: "Left"}
	boxed := box.New(cfg)

	body = boxed.String("", body)
	body = indent.String(body, 2)
	return body
}
