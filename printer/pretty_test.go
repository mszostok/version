package printer_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"go.szostok.io/version"
	"go.szostok.io/version/printer"
	"go.szostok.io/version/style"
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

// TestPrettyPrinterOptions tests the custom Pretty options.
//
// It uses the golden files, to update them run:
//
//	go test -v -run=TestPrettyPrinterOptions ./printer/...  -update
func TestPrettyPrinterOptions(t *testing.T) {
	tests := []struct {
		name         string
		givenOptions []printer.PrettyOption
	}{
		{
			name: "Custom renderer",
			givenOptions: []printer.PrettyOption{
				printer.WithPrettyRenderer(func(in *version.Info, isSmartTerminal bool) (string, error) {
					return fmt.Sprintf("version: %s \nisSmart: %v", in.Version, isSmartTerminal), nil
				}),
			},
		},
		{
			name: "Custom formatting",
			givenOptions: []printer.PrettyOption{
				printer.WithPrettyFormatting(func() *style.Formatting {
					formatting := style.DefaultFormatting()
					formatting.Header = style.Header{
						Prefix: "💡 ",
					}
					formatting.Date.EnableHumanizedSuffix = false
					return &formatting
				}()),
			},
		},
		{
			name: "Custom layout",
			givenOptions: []printer.PrettyOption{
				printer.WithPrettyLayout(&style.Layout{
					GoTemplate: `
            | {{ Key "Version"     }}        {{ .Version                     | Val   }}
            | {{ Key "Git Commit"  }}        {{ .GitCommit  | Commit         | Val   }}
            | {{ Key "Build Date"  }}        {{ .BuildDate  | FmtDate        | Val   }}
            | {{ Key "Commit Date" }}        {{ .CommitDate | FmtDate        | Val   }}
            | {{ Key "Dirty Build" }}        {{ .DirtyBuild | FmtBool        | Val   }}
            | {{ Key "Go Version"  }}        {{ .GoVersion  | trimPrefix "go"| Val   }}
            | {{ Key "Compiler"    }}        {{ .Compiler                    | Val   }}
            | {{ Key "Platform"    }}        {{ .Platform                    | Val   }}`,
				}),
			},
		},
		{
			name: "Custom style",
			givenOptions: []printer.PrettyOption{
				printer.WithPrettyStyle(&style.Config{
					Layout: style.Layout{
						GoTemplate: `{{ Header .Meta.CLIName }}`,
					},
					Formatting: style.Formatting{
						Header: style.Header{
							Prefix: "💡 ",
						},
					},
				}),
			},
		},
		{
			name: "Custom post render hook",
			givenOptions: []printer.PrettyOption{
				printer.WithPrettyPostRenderHook(func(body string, isSmartTerminal bool) (string, error) {
					return fmt.Sprintf("%s\nCustom post render footer (isSmart: %v)", body, isSmartTerminal), nil
				}),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// given
			pretty := printer.NewPretty(tc.givenOptions...)
			var buff strings.Builder

			// when
			err := pretty.Print(&version.Info{
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
			}, &buff)

			// then
			require.NoError(t, err)
			assertGoldenFile(t, buff.String())
		})
	}
}
