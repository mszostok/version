package target

import (
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"
	"go.szostok.io/magex/shx"
)

func TakePreview() {
	buildPreviewCLI()
	takeShoot()
}

func takeShoot() {
	lo.Must0(
		shx.MustCmdf("./magefiles/hack/preview.sh").RunV(),
	)
}

func buildPreviewCLI() {
	var (
		buildDate  = time.Date(2022, time.April, 1, 12, 22, 14, 0, time.UTC).Format("2006-01-02T15:04:05Z0700")
		commitDate = time.Date(2022, time.March, 28, 15, 32, 14, 0, time.UTC).Format("2006-01-02T15:04:05Z0700")
	)

	ldflags := []string{
		"-X go.szostok.io/version.version=1.0.0",
		"-X go.szostok.io/version.dirtyBuild=falfse",
		fmt.Sprintf("-X 'go.szostok.io/version.buildDate=%s'", buildDate),
		fmt.Sprintf("-X go.szostok.io/version.commitDate=%s", commitDate),
	}

	lo.Must0(
		shx.MustCmdf(`go build -ldflags="%s" -o ./preview ./preview-stub`, strings.Join(ldflags, " ")).
			In("./magefiles").Run(),
	)
}
