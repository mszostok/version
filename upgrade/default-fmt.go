package upgrade

import "strings"

var defaultLayoutGoTpl = removeStartingNewLines(`
  │ A new release is available: {{ .Version }} → {{ .NewVersion | Green }}
  │ {{ .ReleaseURL  | Underline | Blue }}
`)

func removeStartingNewLines(in string) string {
	in = strings.TrimPrefix(in, "\n")
	return in
}
