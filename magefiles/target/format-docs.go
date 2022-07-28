package target

import (
	"strings"

	"github.com/samber/lo"
	"go.szostok.io/magex/shx"
)

var excludedFromFormatting = []string{
	// TODO: add option to ignore section by mdox
	"docs/customization/build-ldflags/magefile.md",
	"docs/customization/usage/urfave-cli.md",
	"docs/examples.md",
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
