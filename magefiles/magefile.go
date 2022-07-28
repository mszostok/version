package main

import (
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/samber/lo"
	"go.szostok.io/magex/deps"
	"go.szostok.io/magex/shx"
)

const (
	GolangciLintVersion = "1.47.2"
	bin                 = "bin"
)

var (
	//Default = Build

	Aliases = map[string]interface{}{
		"l": Lint,
	}
)

// "Go" Targets

// Lint Runs linters on the codebase
func Lint() error {
	lo.Must0(deps.EnsureGolangciLint(bin, GolangciLintVersion))
	return shx.MustCmdf(`./bin/golangci-lint run --fix ./...`).Run()
}

// "Docs" Targets

type Docs mg.Namespace

// Fmt Formats markdown documentation
func (d Docs) Fmt() error {
	return d.fmt(false)
}

// Check Checks formatting and links in *.md files
func (d Docs) Check() error {
	return d.fmt(true)
}

func (d Docs) fmt(onlyCheck bool) error {
	lo.Must0(deps.EnsureMdox(bin, ""))

	mdFiles := lo.Must(shx.FindFiles(".", shx.FindFilesOpts{
		Ext: []string{".md"},
	}))

	return shx.MustCmdf(`./bin/mdox fmt --soft-wraps %s %s`,
		WithOptArg("--check", onlyCheck),
		strings.Join(mdFiles, " "),
	).Run()
}

func WithOptArg(key string, shouldAdd bool) string {
	if shouldAdd {
		return key
	}
	return ""
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

// "Other" Targets

//// Generate Generates files
//func Generate() {
//	mg.Deps(Docs.Generate, generators.GitHubWorkflows)
//}
