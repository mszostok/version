package target

import (
	"strings"

	"github.com/samber/lo"
	"go.szostok.io/magex/shx"
)

var excludedFromFormatting = []string{
	// TODO: add option to ignore section by mdox
	"docs/get-started/build-ldflags/magefile.md",
	"docs/get-started/usage/urfave-cli.md",
	"docs/examples.md",
	"docs/customization/pretty/layout.md",
	"docs/customization/pretty/format.md",
}

func FmtDocs(onlyCheck bool) error {
	mdFiles := lo.Must(shx.FindFiles(".", shx.FindFilesOpts{
		Ext:          []string{".md"},
		IgnorePrefix: excludedFromFormatting,
	}))

	return shx.MustCmdf(`./bin/mdox fmt --soft-wraps %s %s`,
		WithOptArg("--check", onlyCheck),
		strings.Join(mdFiles, " "),
	).Run()
}

func WithOptArg(key string, shouldAdd bool) string {
	if shouldAdd {
		return key
	}
	return ""
}
