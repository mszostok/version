package main

import (
	"github.com/magefile/mage/mg"
	"github.com/samber/lo"
	"go.szostok.io/magex/deps"
	"go.szostok.io/magex/shx"

	"tools/target"
)

const (
	GolangciLintVersion = "1.47.2"
	MdoxVersion         = "" // latest
	MuffetVersion       = "2.5.0"
	bin                 = "bin"
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
	mg.Deps(mg.F(deps.EnsureMdox, bin, MdoxVersion))

	return target.FmtDocs(false)
}

// Check Checks formatting and links in *.md files
func (d Docs) Check() error {
	mg.Deps(mg.F(deps.EnsureMdox, bin, MdoxVersion))

	return target.FmtDocs(true)
}

// CheckDeadLinks Detects dead links in documentation.
func (d Docs) CheckDeadLinks() error {
	mg.Deps(mg.F(deps.EnsureMuffet, bin, MuffetVersion))
	lo.Must0(shx.MustCmdf(" pip install -r requirements.txt").RunS())
	return target.CheckDeadLinks()
}

// "Test" Targets

type Test mg.Namespace

// Unit Executes Go unit tests.
func (Test) Unit() error {
	return shx.MustCmdf(`go test -coverprofile=coverage.out ./...`).Run()
}

// Coverage Generates file with unit test coverage data and open it in browser
func (t Test) Coverage() error {
	mg.Deps(t.Unit)
	return shx.MustCmdf(`go tool cover -html=coverage.out`).Run()
}

type Gen mg.Namespace

func (Gen) All() {
	mg.Deps(Gen.PrettyExamples)
}

func (Gen) PrettyExamples() {
	target.EmbedDefaultPrettyFormatting()
	target.EmbedDefaultPrettyLayout()
}
