package term

import (
	"io"
	"os"
	"runtime"

	"github.com/mattn/go-isatty"
)

// IsSmart returns true if the writer w is a terminal AND
// we think that the terminal is smart enough to use VT escape codes etc.
func IsSmart(w io.Writer) bool {
	type fileDescriptor interface {
		Fd() uintptr
	}
	fd, ok := w.(fileDescriptor)
	if !ok {
		return false
	}
	if !isatty.IsTerminal(fd.Fd()) {
		return false
	}

	// Explicitly dumb terminals are not smart
	// https://en.wikipedia.org/wiki/Computer_terminal#Dumb_terminals
	term := os.Getenv("TERM")
	if term == "dumb" {
		return false
	}

	// On Windows WT_SESSION is set by the modern terminal component.
	// Older terminals have poor support for UTF-8, VT escape codes, etc.
	if runtime.GOOS == "windows" && os.Getenv("WT_SESSION") == "" {
		return false
	}

	// OK, we'll assume it's smart now, given no evidence otherwise.
	return true
}
