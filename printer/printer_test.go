package printer_test

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"go.szostok.io/version"
	"go.szostok.io/version/printer"
	"go.szostok.io/version/upgrade"
)

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

func TestPrinterPostHookOption(t *testing.T) {
	// given
	var postHookExecuted bool
	p := printer.New(printer.WithPostHook(func() error {
		postHookExecuted = true
		return nil
	}))

	// when
	err := p.PrintInfo(io.Discard, &version.Info{Version: "0.6.1"})

	// then
	require.NoError(t, err)
	assert.True(t, postHookExecuted)
}

func TestPrinterPostHookOptionError(t *testing.T) {
	// given
	customErr := errors.New("custom error")
	p := printer.New(printer.WithPostHook(func() error {
		return customErr
	}))

	// when
	err := p.PrintInfo(io.Discard, &version.Info{Version: "0.6.1"})

	// then
	require.ErrorIs(t, err, customErr)
}

func TestPrinterUpgradeNoticeOption(t *testing.T) {
	// given
	p := printer.New(printer.WithUpgradeNotice(
		"owner", "repo",
		upgrade.WithRenderer(func(in *upgrade.Info, isSmartTerminal bool) (string, error) {
			return fmt.Sprintf("version: %s \nisSmart: %v", in.Version, isSmartTerminal), nil
		})),
	)

	// when
	// I don't want to run a really upgrade checker, instead we check if it's set properly.
	out, err := p.GetUpgradeNotice().Render(&upgrade.Info{Version: "0.6.1"}, false)

	// then
	require.NoError(t, err)
	assert.Equal(t, "version: 0.6.1 \nisSmart: false", out)
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

// TestPrinterPrint tests the Print method.
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

// It uses the golden files, to update them run:
//
//	go test -v -run=TestPrinterStyleFileOptions ./printer/...  -update
func TestPrinterStyleFileOptions(t *testing.T) {
	tests := []struct {
		testName string
		fileName string
	}{
		{
			testName: "Print custom layout",
			fileName: "testdata/TestPrinterStyleFileOptions/customStyle",
		},
		{
			testName: "Print default layout",
			fileName: "testdata/TestPrinterStyleFileOptions/invalidStyle",
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			yamlStyle, err := printer.WithPrettyStyleFile(tc.fileName + ".yaml")
			require.NoError(t, err)

			jsonStyle, err := printer.WithPrettyStyleFile(tc.fileName + ".json")
			require.NoError(t, err)

			t.Setenv("config-file", tc.fileName+".yaml")

			envStyle, err := printer.WithPrettyStyleFromEnv("config-file")
			require.NoError(t, err)

			validateStyle(yamlStyle, t)
			validateStyle(jsonStyle, t)
			validateStyle(envStyle, t)
		})
	}
}

// It uses the golden files, to update them run:
//
//	go test -v -run=TestPrinterStyleFromEnvOptionsUseDefaults ./printer/...  -update
func TestPrinterStyleFromEnvOptionsUseDefaults(t *testing.T) {
	emptyStyle, err := printer.WithPrettyStyleFromEnv("empty-variable")
	require.NoError(t, err)

	validateStyle(emptyStyle, t)
}

func validateStyle(customStyle *printer.CustomPrettyStyle, t *testing.T) {
	t.Helper()

	p := printer.New(customStyle)
	var buff strings.Builder

	err := p.Print(&buff)

	require.NoError(t, err)

	stripped := strings.TrimSpace(buff.String())
	normalized := normalizeOutput(stripped)
	assertGoldenFile(t, normalized)
}

// normalizeOutput normalize dynamic fields such as platform and binary name.
func normalizeOutput(data string) string {
	data = strings.ReplaceAll(data, os.Args[0], "fixed-name")
	data = strings.ReplaceAll(data, fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH), "fixed-platform")
	return data
}
