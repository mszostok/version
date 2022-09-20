package printer_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"go.szostok.io/version"
	"go.szostok.io/version/printer"
)

// TestPrettyPrinter tests the Pretty format.
//
// It uses the golden files, to update them run:
//
//	go test -v -run=TestPrettyPrinter ./printer/...  -update
func TestPrettyPrinter(t *testing.T) {
	tests := []struct {
		name      string
		givenInfo *version.Info
	}{
		{
			name: "Built-in Info fields",
			givenInfo: &version.Info{
				Version:    "0.6.1",
				GitCommit:  "324d022c190ce49e0440e6bdac6383e4874c7c70",
				BuildDate:  "2022-04-01T12:22:14Z",
				CommitDate: "2022-03-28T15:32:14Z",
				DirtyBuild: false,
				GoVersion:  "go1.19.1",
				Compiler:   "gc",
				Platform:   "darwin/amd64",
				Meta: version.Meta{
					CLIName: "testing",
				},
			},
		},
		{
			name: "Custom Info fields",
			givenInfo: &version.Info{
				Version:    "0.6.1",
				GitCommit:  "324d022c190ce49e0440e6bdac6383e4874c7c70",
				BuildDate:  "2022-04-01T12:22:14Z",
				CommitDate: "2022-03-28T15:32:14Z",
				DirtyBuild: false,
				GoVersion:  "go1.19.1",
				Compiler:   "gc",
				Platform:   "darwin/amd64",
				Meta: version.Meta{
					CLIName: "testing",
				},
				ExtraFields: CustomFields{
					DocumentationURL: "https://example.com/docs",
					RepoURL:          "https://example.com/repo",
					IsFun:            true,
					Counter:          42,
				},
			},
		},
		{
			name:      "Nil Info",
			givenInfo: nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given
			pretty := printer.NewPretty()
			var buff strings.Builder

			// when
			err := pretty.Print(tc.givenInfo, &buff)

			// then
			require.NoError(t, err)
			assertGoldenFile(t, buff.String())
		})
	}
}
