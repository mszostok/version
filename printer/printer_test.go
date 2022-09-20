package printer_test

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.szostok.io/version"
	"go.szostok.io/version/printer"
)

// TestPrettyPrinter tests the Pretty format.
//
// It uses the golden files, to update them run:
//
//	go test -v -run=TestPrinter ./printer/...  -update
func TestPrinter(t *testing.T) {
	tests := []struct {
		name            string
		givenArgs       []string
		expOutputFormat printer.OutputFormat
	}{
		{
			name:            "Print in default format",
			expOutputFormat: printer.PrettyFormat,
		},
		{
			name:            "Print in Pretty format",
			givenArgs:       []string{"-opretty"},
			expOutputFormat: printer.PrettyFormat,
		},
		{
			name:            "Print in YAML format",
			givenArgs:       []string{"-oyaml"},
			expOutputFormat: printer.YAMLFormat,
		},
		{
			name:            "Print in JSON format",
			givenArgs:       []string{"-ojson"},
			expOutputFormat: printer.JSONFormat,
		},
		{
			name:            "Print in JSON format when long flag is used",
			givenArgs:       []string{"--output=json"},
			expOutputFormat: printer.JSONFormat,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given
			p := printer.New()
			var buff strings.Builder

			flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
			p.RegisterPFlags(flags)

			// when
			err := flags.Parse(tc.givenArgs)
			require.NoError(t, err)

			err = p.PrintInfo(&buff, &version.Info{
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
			})

			// then
			require.NoError(t, err)
			assert.Equal(t, tc.expOutputFormat, p.OutputFormat())
			assertGoldenFile(t, buff.String())
		})
	}
}

func TestPrinterUnknownFormatFlag(t *testing.T) {
	// given
	p := printer.New()

	flags := pflag.NewFlagSet("test", pflag.ContinueOnError)
	p.RegisterPFlags(flags)

	// when
	err := flags.Parse([]string{"--output=table"}) // unknown output format

	// then
	require.EqualError(t, err, `invalid argument "table" for "-o, --output" flag: invalid output format type`)
}

func TestPrinterUnknownFormat(t *testing.T) {
	// given
	p := printer.New()
	p.SetOutputFormat("table")

	// when
	err := p.PrintInfo(nil, nil)

	// then
	require.EqualError(t, err, `printer "table" is not available`)
}

// TestPrettyPrinter tests the Print method.
//
// It uses the golden files, to update them run:
//
//	go test -v -run=TestPrinterPrint ./printer/...  -update
func TestPrinterPrint(t *testing.T) {
	// given
	p := printer.New()
	var buff strings.Builder

	// when
	err := p.Print(&buff)

	// then
	require.NoError(t, err)
	normalized := normalizeOutput(buff.String())
	assertGoldenFile(t, normalized)
}

// normalizeOutput normalize dynamic fields such as platform and binary name.
func normalizeOutput(data string) string {
	data = strings.ReplaceAll(data, os.Args[0], "fixed-name")
	data = strings.ReplaceAll(data, fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH), "fixed-platform")
	return data
}
