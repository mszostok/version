package version

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/dustin/go-humanize"
	"github.com/gookit/color"
)

var _ Printer = &Pretty{}

// Pretty prints human-readable version.
type Pretty struct{}

func (p *Pretty) Print(in Info, w io.Writer) error {
	var buff strings.Builder

	header := color.New(color.FgMagenta).Sprintf

	buff.WriteString("\n")
	buff.WriteString(header("▓▓▓ %s\n\n", in.name)) // TODO: extract as header...
	buff.WriteString(line("Version", in.Version))
	buff.WriteString(line("Git Commit", maxLen(in.GitCommit, 7)))
	buff.WriteString(line("Build Date", prettyTime(in.BuildDate)))
	buff.WriteString(line("Commit Date", prettyTime(in.CommitDate)))
	buff.WriteString(line("Dirty Build", formatBool(in.DirtyBuild)))
	buff.WriteString(line("Go Version", strings.TrimPrefix(in.GoVersion, "go")))
	buff.WriteString(line("Compiler", in.Compiler))
	buff.WriteString(line("Platform", in.Platform))

	_, err := fmt.Fprintln(w, buff.String())
	return err
}

func line(k string, v any) string {
	key := color.New(color.Bold, color.FgDarkGray).Sprintf
	val := color.White.Sprintf

	return key("  %-20s", k) + val("%v\n", v)
}

func formatBool(in bool) string {
	if in {
		return "yes"
	}
	return "no"
}

func prettyTime(in string) string {
	if in == "" {
		return ""
	}

	t, err := dateparse.ParseAny(in)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%s (%s)", t.Local().Format(time.RFC822), humanize.Time(t))
}

func maxLen(in string, max int) string {
	if len(in) <= max {
		return in
	}
	return in[:max]
}
