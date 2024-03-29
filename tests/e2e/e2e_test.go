//go:build e2e

package integration

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/araddon/dateparse"
	"github.com/mattn/go-runewidth"
	"github.com/mattn/go-shellwords"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestCases struct {
	cmd      string
	name     string
	dir      string
	bordered bool
}

var cases = []TestCases{
	// Cobra
	{
		name: "Should return cobra help menu",
		cmd:  ``,
		dir:  "cobra",
	},
	{
		name: "Should return cobra version in default Pretty format",
		cmd:  `version`,
		dir:  "cobra",
	},
	{
		name: "Should return cobra version in YAML format",
		cmd:  `version -oyaml`,
		dir:  "cobra",
	},
	{
		name: "Should return cobra version in JSON format",
		cmd:  `version -ojson`,
		dir:  "cobra",
	},
	{
		name: "Should return cobra version in short format",
		cmd:  `version -oshort`,
		dir:  "cobra",
	},
	{
		name: "Should return cobra version in short format with full flag name",
		cmd:  `version --output short`,
		dir:  "cobra",
	},
	{
		name: "Should return cobra version help",
		cmd:  `version --help`,
		dir:  "cobra",
	},
	{
		name: "Should return cobra version when alias ver is used",
		cmd:  `version ver`,
		dir:  "cobra",
	},

	// Cobra alias
	{
		name: "Should return cobra version help with v alias",
		cmd:  `version --help`,
		dir:  "cobra-alias",
	},
	{
		name: "Should return cobra version when alias v is used",
		cmd:  `version v`,
		dir:  "cobra-alias",
	},

	// Custom fields
	{
		name: "Should return version with custom fields in default Pretty format",
		cmd:  ``,
		dir:  "custom-fields",
	},
	{
		name: "Should return version with custom fields in JSON format",
		cmd:  `-ojson`,
		dir:  "custom-fields",
	},
	{
		name: "Should return version with custom fields in YAML format",
		cmd:  `-oyaml`,
		dir:  "custom-fields",
	},

	// Custom formatting
	{
		name: "Should return version with custom Pretty formatting",
		cmd:  ``,
		dir:  "custom-formatting",
	},

	// Custom layout
	{
		name:     "Should return version with custom Pretty layout",
		cmd:      ``,
		bordered: true,
		dir:      "custom-layout",
	},

	// Custom renderer
	{
		name:     "Should return version with custom Pretty renderer",
		cmd:      ``,
		bordered: true,
		dir:      "custom-renderer",
	},

	// Plain
	{
		name: "Should return plain version",
		cmd:  ``,
		dir:  "plain",
	},

	// Printer
	{
		name: "Should return default printer version",
		cmd:  ``,
		dir:  "printer",
	},

	// Printer post hook
	{
		name: "Should return version with executed post hook in default Pretty format",
		cmd:  ``,
		dir:  "printer-post-hook",
	},
	{
		name: "Should return version with executed post hook in JSON format",
		cmd:  `-ojson`,
		dir:  "printer-post-hook",
	},
	{
		name: "Should return version with executed post hook in YAML format",
		cmd:  `-oyaml`,
		dir:  "printer-post-hook",
	},

	// Upgrade notice cobra
	{
		name: "Should return version with upgrade notice in default Pretty format",
		cmd:  `version`,
		dir:  "upgrade-notice-cobra",
	},
	{
		name: "Should return version with upgrade notice in JSON format",
		cmd:  `version -ojson`,
		dir:  "upgrade-notice-cobra",
	},
	{
		name: "Should return version with upgrade notice in YAML format",
		cmd:  `version -oyaml`,
		dir:  "upgrade-notice-cobra",
	},
	{
		name: "Should return version with upgrade notice in short format",
		cmd:  `version -oshort`,
		dir:  "upgrade-notice-cobra",
	},

	// Upgrade notice custom
	{
		name: "Should return version with custom upgrade notice in default Pretty format",
		cmd:  ``,
		dir:  "upgrade-notice-custom",
	},
	{
		name: "Should return version with custom upgrade notice in JSON format",
		cmd:  `-ojson`,
		dir:  "upgrade-notice-custom",
	},
	{
		name: "Should return version with custom upgrade notice in YAML format",
		cmd:  `-oyaml`,
		dir:  "upgrade-notice-custom",
	},
	{
		name: "Should return version with custom upgrade notice in short format",
		cmd:  `-oshort`,
		dir:  "upgrade-notice-custom",
	},

	// Upgrade notice standalone
	{
		name: "Should return upgrade notice from standalone command",
		cmd:  ``,
		dir:  "upgrade-notice-standalone",
	},
}

// TestExamplesColorOutput tests examples usage with the colored output.
//
// This test is based on golden file. To update golden file, run:
//
//	UPDATE_GOLDEN=true TEST_NAME=TestExamplesColorOutput mage test:e2e
func TestExamplesColorOutput(t *testing.T) {
	if os.Getenv("CI") == "true" {
		t.Skip("Those tests are not stable on CI yet")
	}
	t.Parallel()

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			//t.Parallel() goexpect doesn't work in multi thread

			// given
			binaryPath := buildBinaryAllLDFlags(t, tc.dir)

			// when
			result, err := Exec(binaryPath, tc.cmd).
				AwaitColorResultAtMost(10 * time.Second)

			// then
			require.NoError(t, err)
			assert.Equal(t, 0, result.ExitCode)

			data := normalizeOutput(result.Stdout, tc.bordered)

			g := goldie.New(t, goldie.WithNameSuffix(".golden.txt"))

			g.Assert(t, t.Name(), []byte(data))
		})
	}
}

// TestExamplesNoColorOutput tests examples usage with the non-colored output.
//
// This test is based on golden file. To update golden file, run:
//
//	UPDATE_GOLDEN=true TEST_NAME=TestExamplesNoColorOutput mage test:e2e
func TestExamplesNoColorOutput(t *testing.T) {
	t.Parallel()

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if runtime.GOOS != "windows" {
				t.Parallel()
			}
			// given
			binaryPath := buildBinaryAllLDFlags(t, tc.dir)

			// when
			result, err := Exec(binaryPath, tc.cmd).
				AwaitResultAtMost(10 * time.Second)

			// then
			require.NoError(t, err)
			assert.Equal(t, 0, result.ExitCode)

			data := result.Stdout + result.Stderr
			data = normalizeOutput(data, tc.bordered)

			g := goldie.New(t, goldie.WithNameSuffix(".golden.txt"))

			g.Assert(t, t.Name(), []byte(data))
		})
	}
}

// TestExamplesRecheckInterval tests that the recheck interval is respected.
//
// This test is based on golden file. To update golden file, run:
//
//	UPDATE_GOLDEN=true TEST_NAME=TestExamplesRecheckInterval mage test:e2e
func TestExamplesRecheckInterval(t *testing.T) {
	t.Parallel()

	// given
	binaryPath := buildBinaryAllLDFlags(t, "upgrade-notice-sub-cmd")

	// when
	result, err := Exec(binaryPath, "version check").
		AwaitResultAtMost(10 * time.Second)

	// then
	// Should return upgrade notice for the first time
	require.NoError(t, err, result.Stderr)
	assert.Equal(t, 0, result.ExitCode)

	data := result.Stdout + result.Stderr
	g := goldie.New(t, goldie.WithNameSuffix(".golden.txt"))
	g.Assert(t, t.Name(), []byte(data))

	// when
	result, err = Exec(binaryPath, "version check").
		AwaitResultAtMost(10 * time.Second)

	// then
	require.NoError(t, err, result.Stderr)
	assert.Equal(t, 0, result.ExitCode)
	data = result.Stdout + result.Stderr
	g = goldie.New(t, goldie.WithNameSuffix("recheck.golden.txt"))
	g.Assert(t, t.Name(), []byte(data))
}

func normalizeOutput(data string, bordered bool) string {
	platform := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

	normalizedPlatform := "normalized"
	if bordered { // if bordered we need to be careful to do not mess up the padding
		padding := runewidth.StringWidth(platform) - runewidth.StringWidth(normalizedPlatform)
		normalizedPlatform += strings.Repeat(" ", padding)

	}
	return strings.ReplaceAll(data, platform, normalizedPlatform)
}

var prettyResolvedFieldsFormat = heredoc.Doc(`

💡 %s

  Version              (devel)
  Git Commit           %s
  Build Date           N/A
  Commit Date          %s
  Dirty Build          %s
  Go version           1.19.4
  Compiler             gc
  Platform             %s
`)

// TestExamplesColorOutput tests that version can resolve the info fields automatically.
func TestResolvesDefaultFields(t *testing.T) {
	t.Parallel()

	// given
	var bin = "auto-resolved-fields"
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}

	var (
		dir        = filepath.Join(exampleDir, "custom-formatting")
		binaryPath = filepath.Join(dir, bin)
	)

	args, err := shellwords.Parse(fmt.Sprintf(`build -o %s . `, bin))
	require.NoError(t, err)
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	require.NoError(t, cmd.Run())

	// when
	result, err := Exec(binaryPath, "").
		AwaitResultAtMost(10 * time.Second)

	commit, commitDate, dirtyBuild := getGitDetails(t)
	expOutput := fmt.Sprintf(prettyResolvedFieldsFormat, binaryPath, commit, commitDate, dirtyBuild, fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))

	// then
	require.NoError(t, err)
	assert.Equal(t, 0, result.ExitCode)
	assert.Equal(t, expOutput, result.Stdout)
	assert.Empty(t, result.Stderr)
}

func getGitDetails(t *testing.T) (string, string, string) {
	t.Helper()
	var (
		commit = runGitCmd(t, `rev-parse HEAD`)
		date   = fmtDate(runGitCmd(t, `--no-pager log -1 --format="%cD"`))
		dirty  = fmtBool(runGitCmd(t, `status --short`) != "")
	)

	return fmt.Sprintf("%.7s", commit), date, dirty
}
func fmtBool(in bool) string {
	if in {
		return "yes"
	}
	return "no"
}

func runGitCmd(t *testing.T, rawArgs string) string {
	t.Helper()

	args, err := shellwords.Parse(rawArgs)
	require.NoError(t, err)

	out, err := exec.Command("git", args...).CombinedOutput()
	require.NoError(t, err)
	return strings.TrimSpace(string(out))
}

func fmtDate(in string) string {
	normalized, _ := dateparse.ParseAny(in)
	return normalized.Local().Format(time.RFC822)
}
