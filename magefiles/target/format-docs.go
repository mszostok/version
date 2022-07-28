package target

import (
	"strings"

	"github.com/samber/lo"
	"go.szostok.io/magex/shx"
)

func FmtDocs(onlyCheck bool) error {
	mdFiles := lo.Must(shx.FindFiles(".", shx.FindFilesOpts{
		Ext: []string{".md"},
		// TODO: add option to ignore section by mdox
		IgnorePrefix: []string{"docs/examples.md", "docs/customization/usage/urfave-cli.md"},
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
