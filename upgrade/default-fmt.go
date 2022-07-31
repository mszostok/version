package upgrade

import "strings"

var defaultLayoutGoTpl = removeStartingNewLines(`
  │ A new release is available: {{ .Version }} → {{ .NewVersion | green }}
  │ {{ .ReleaseURL  | underscore | blue }}
`)

func removeStartingNewLines(in string) string {
	in = strings.TrimPrefix(in, "\n")
	return in
}
