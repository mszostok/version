package printer_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.szostok.io/version"
	"go.szostok.io/version/printer"
)

func TestJSONPrinter(t *testing.T) {
	tests := []struct {
		name      string
		givenInfo *version.Info
		expOutput string
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
			expOutput: `{
			 "buildDate": "2022-04-01T12:22:14Z",
			 "commitDate": "2022-03-28T15:32:14Z",
			 "compiler": "gc",
			 "gitCommit": "324d022c190ce49e0440e6bdac6383e4874c7c70",
			 "goVersion": "go1.19.1",
			 "platform": "darwin/amd64",
			 "version": "0.6.1"
			}`,
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
			expOutput: `{
			 "buildDate": "2022-04-01T12:22:14Z",
			 "commitDate": "2022-03-28T15:32:14Z",
			 "compiler": "gc",
			 "gitCommit": "324d022c190ce49e0440e6bdac6383e4874c7c70",
			 "goVersion": "go1.19.1",
			 "platform": "darwin/amd64",
			 "version": "0.6.1",
			 "counter": 42,
			 "isFun": true,
			 "documentationURL": "https://example.com/docs",
			 "repositoryURL": "https://example.com/repo"
			}`,
		},
		{
			name:      "Nil Info",
			givenInfo: nil,
			expOutput: "{}",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given
			json := &printer.JSON{}
			var buff strings.Builder

			// when
			err := json.Print(tc.givenInfo, &buff)

			// then
			require.NoError(t, err)
			assert.JSONEq(t, tc.expOutput, buff.String())
		})
	}
}
