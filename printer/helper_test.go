package printer_test

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var update = flag.Bool("update", false, "Update golden test file fixture")

type CustomFields struct {
	DocumentationURL string `json:"documentationURL" yaml:"documentationURL" pretty:"Documentation URL"`
	RepoURL          string `json:"repositoryURL"    yaml:"repoURL"          pretty:"Repository URL"`
	IsFun            bool   `json:"isFun"            yaml:"isFun"            pretty:"Is Fun"`
	Counter          int    `json:"counter"          yaml:"counter"          pretty:"Counter"`
}

// trimLineSpace trim all leading and trailing white space in each line.
func trimLineSpace(in string) string {
	var buff strings.Builder
	for _, line := range strings.Split(in, "\n") {
		buff.WriteString(strings.TrimSpace(line) + "\n")
	}
	return buff.String()
}

func assertGoldenFile(t *testing.T, actualData string) {
	t.Helper()

	path := filepath.Join("testdata", t.Name()+".golden.txt")
	if *update {
		updateGolden(t, path, actualData)
		return
	}
	golden, err := os.ReadFile(path)
	require.NoError(t, err)

	assert.Equal(t, string(golden), actualData)
}

func updateGolden(t *testing.T, goldenFile, actualData string) {
	goldenFileDir := filepath.Dir(goldenFile)
	err := ensureDir(goldenFileDir)
	require.NoError(t, err)

	err = os.WriteFile(goldenFile, []byte(actualData), 0o600)
	require.NoError(t, err)
}

func ensureDir(loc string) error {
	_, err := os.Stat(loc)
	if err != nil && os.IsNotExist(err) {
		return os.MkdirAll(loc, 0o755)
	}
	return err
}
