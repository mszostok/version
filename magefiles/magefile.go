package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/samber/lo"
	"go.szostok.io/magex/deps"
	"go.szostok.io/magex/shx"

	"tools/target"
)

const (
	GolangciLintVersion = "1.49.0"
	MdoxVersion         = "" // latest
	MuffetVersion       = "2.5.0"
	bin                 = "bin"

	TestsDir                   = "./tests"
	TestUpdateGoldenEnvVarName = "UPDATE_GOLDEN"
	TestNameEnvVarName         = "TEST_NAME"
)

var (
	Default = Docs.Fmt

	Aliases = map[string]interface{}{
		"l":   Lint,
		"fmt": Docs.Fmt,
	}
)

// "Go" Targets

// Lint Runs linters on the codebase
func Lint() error {
	lo.Must0(deps.EnsureGolangciLint(bin, GolangciLintVersion))
	return shx.MustCmdf(`./bin/golangci-lint run --fix ./...`).RunV()
}

// "Docs" Targets

type Docs mg.Namespace

// Fmt Formats markdown documentation
func (d Docs) Fmt() error {
	mg.Deps(mg.F(target.EnsurePrettier, bin))

	return target.FmtDocs(false)
}

// Check Checks formatting and links in *.md files
func (d Docs) Check() error {
	mg.Deps(mg.F(target.EnsurePrettier, bin))

	return target.FmtDocs(true)
}

// CheckDeadLinks Detects dead links in documentation.
func (d Docs) CheckDeadLinks(warmup time.Duration) error {
	mg.Deps(mg.F(deps.EnsureMuffet, bin, MuffetVersion))
	lo.Must0(shx.MustCmdf("pip install -r requirements.txt").RunS())
	return target.CheckDeadLinks(warmup)
}

// "Test" Targets

type Test mg.Namespace

// Unit Executes Go unit tests.
func (Test) Unit() error {
	return shx.MustCmdf(`go test -v -count 1 -coverprofile=coverage.out ./...`).RunV()
}

// E2e Executes E2E tests.
func (Test) E2e() error {
	var (
		shouldUpdate = shx.GetEnvVal(TestUpdateGoldenEnvVarName, "false")
		testName     = shx.GetEnvVal(TestNameEnvVarName, "")
	)

	if shouldUpdate == "true" {
		path := filepath.Join(TestsDir, "e2e", "testdata", testName)
		lo.Must0(os.RemoveAll(path))
		lo.Must0(os.MkdirAll(path, 0o775))
	}
	return shx.MustCmdf(`go test -v -tags=e2e ./e2e/...`).
		In(TestsDir).
		WithArg("-update", shouldUpdate).
		WithArg("-run", testName).
		RunV()
}

// Coverage Generates file with unit test coverage data and open it in browser
func (t Test) Coverage() error {
	mg.Deps(t.Unit)
	return shx.MustCmdf(`go tool cover -html=coverage.out`).Run()
}

type Gen mg.Namespace

func (Gen) All() {
	mg.Deps(Gen.PrettyExamples)
	mg.Deps(Gen.Homepage)
}

func (Gen) PrettyExamples() {
	target.EmbedDefaultPrettyFormatting()
	target.EmbedDefaultPrettyLayout()
}

func (Gen) Homepage() {
	target.SyncHomepage()
}

func (Gen) Preview() {
	target.TakePreview()
}
