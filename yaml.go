package version

import (
	"fmt"
	"io"
	"strings"

	"github.com/gookit/color"
)

var _ Printer = &YAML{}

// YAML prints data in YAML format.
type YAML struct{}

// Print marshals input data to YAML format and writes it to a given writer.
func (p *YAML) Print(in Info, w io.Writer) error {
	var buff strings.Builder

	sep := color.New(color.Yellow).Sprint
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

func yamlLine(k string, v any, quote bool) string {
	key := color.New(color.Yellow).Sprintf
	val := color.White.Sprintf

	rv := fmt.Sprintf("%v", v)
	if quote {
		rv = fmt.Sprintf("%q", rv)
	}
	return key("%s: ", k) + val("%s\n", rv)
}
