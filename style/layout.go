package style

// Layout define the layout for printing the pretty format.
type Layout struct {
	GoTemplate string `json:"goTemplate,omitempty" yaml:"goTemplate"`
}

var (
	// KeyValueLayoutGoTpl prints all version data in a 'key  value' manner.
	KeyValueLayoutGoTpl = `
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

	// BoxLayoutGoTpl prints all version data in box.
	// https://knowyourmeme.com/memes/this-is-fine
	BoxLayoutGoTpl = `
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
