package target

import (
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"go.szostok.io/magex/shx"
)

func SyncHomepage() {
	var (
		wd          = lo.Must(os.Getwd())
		homepageDir = filepath.Join(wd, "..", "landingpages")

		indexSrc = filepath.Join(homepageDir, "build", "index.html")
		indexDst = filepath.Join(wd, "docs", "mkdocs-theme")

		staticDirSrc = filepath.Join(homepageDir, "build", "static")
		staticDirDst = filepath.Join(wd, "docs", "static/")
	)

	lo.Must0(shx.MustCmdf("npm run build").In(homepageDir).Run())

	lo.Must0(shx.MustCmdf("cp %s %s", indexSrc, indexDst).Run())

	lo.Must0(shx.MustCmdf("rm -rf ./docs/static").Run())
	lo.Must0(shx.MustCmdf("cp -r %s %s", staticDirSrc, staticDirDst).Run())
}
