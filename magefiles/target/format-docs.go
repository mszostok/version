package target

import (
	"go.szostok.io/magex/shx"
)

func EnsurePrettier(bin string) error {
	return shx.MustCmdf(`npm install -D --prefix %s prettier`, bin).Run()
}

func FmtDocs(onlyCheck bool) error {
	return shx.MustCmdf(`./bin/node_modules/.bin/prettier --write "docs/**/*.md" %s`,
		WithOptArg("--check", onlyCheck),
	).RunV()
}

func WithOptArg(key string, shouldAdd bool) string {
	if shouldAdd {
		return key
	}
	return ""
}
