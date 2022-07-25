package style

type Layout struct {
	Raw string `json:"raw,omitempty" yaml:"raw"`
}

var (
	// DefaultLayoutTpl the default layout that prints all version data.
	DefaultLayoutTpl = `
{{ header }}

  {{ key "Version" }}             {{ .Version                     | val }}
  {{ key "Git Commit" }}          {{ .GitCommit  | commit         | val }}
  {{ key "Build Date" }}          {{ .BuildDate  | fmtDate        | val }}
  {{ key "Commit Date" }}         {{ .CommitDate | fmtDate        | val }}
  {{ key "Dirty Build" }}         {{ .DirtyBuild | fmtBool        | val }}
  {{ key "Go Version" }}          {{ .GoVersion  | trimPrefix "go"| val }}
  {{ key "Compiler" }}            {{ .Compiler                    | val }}
  {{ key "Platform" }}            {{ .Platform                    | val }}
`

	// BoxLayoutTpl https://knowyourmeme.com/memes/this-is-fine
	BoxLayoutTpl = `
╭───{{ repeatMax 57 "─" header }}{{/* ─────────────────────────────────── */}}╮
│                                  {{ repeatMax 25 " " ""                  }} │
│  {{ key "Version" }}             {{ .Version                     | val   }} │
│  {{ key "Git Commit" }}          {{ .GitCommit  | commit         | val   }} │
│  {{ key "Build Date" }}          {{ .BuildDate  | fmtDate        | val   }} │
│  {{ key "Commit Date" }}         {{ .CommitDate | fmtDate        | val   }} │
│  {{ key "Dirty Build" }}         {{ .DirtyBuild | fmtBool        | val   }} │
│  {{ key "Go Version" }}          {{ .GoVersion  | trimPrefix "go"| val   }} │
│  {{ key "Compiler" }}            {{ .Compiler                    | val   }} │
│  {{ key "Platform" }}            {{ .Platform                    | val   }} │
╰───{{ repeatMax 57 "─" ""}}{{/* ──────────────────────────────────────── */}}╯
`
)
