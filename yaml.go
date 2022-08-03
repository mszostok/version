package version

import (
	"fmt"
	"io"
	"strings"

	"github.com/muesli/termenv"

	"go.szostok.io/version/style/termenvx"
	"go.szostok.io/version/term"
)

var _ Printer = &YAML{}

// YAML prints data in YAML format.
type YAML struct{}

// Print marshals input data to YAML format and writes it to a given writer.
func (p *YAML) Print(in *Info, w io.Writer) error {
	if in == nil {
		return nil
	}
	var buff strings.Builder

	profile := p.colorProfile(w)

	sep := p.string(profile, termenv.ANSIYellow)
	yamlLine := p.yamlLine(profile)

	buff.WriteString(sep("---\n"))
	buff.WriteString(yamlLine("version", in.Version, true))
	buff.WriteString(yamlLine("gitCommit", in.GitCommit, false))
	buff.WriteString(yamlLine("buildDate", in.BuildDate, false))
	buff.WriteString(yamlLine("commitDate", in.CommitDate, false))
	buff.WriteString(yamlLine("dirtyBuild", in.DirtyBuild, false))
	buff.WriteString(yamlLine("goVersion", in.GoVersion, true))
	buff.WriteString(yamlLine("compiler", in.Compiler, true))
	buff.WriteString(yamlLine("platform", in.Platform, true))

	_, err := fmt.Fprintln(w, buff.String())
	return err
}

func (p *YAML) yamlLine(profile termenv.Profile) func(k string, v interface{}, quote bool) string {
	key := p.string(profile, termenv.ANSIYellow)
	val := p.string(profile, termenv.ANSIWhite)

	return func(k string, v interface{}, quote bool) string {
		rv := fmt.Sprintf("%v", v)
		if quote {
			rv = fmt.Sprintf("%q", rv)
		}
		return key("%s: ", k) + val("%s\n", rv)
	}
}

func (*YAML) string(p termenv.Profile, color termenv.Color) func(format string, args ...interface{}) string {
	return func(format string, args ...interface{}) string {
		msg := fmt.Sprintf(format, args...)
		return termenv.
			String(msg).
			Foreground(p.Convert(color)).
			String()
	}
}

func (*YAML) colorProfile(w io.Writer) termenv.Profile {
	if term.IsSmart(w) {
		return termenvx.ColorProfile()
	}

	return termenv.Ascii
}
