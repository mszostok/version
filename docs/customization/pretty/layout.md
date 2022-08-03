# Layout

!!! note ""
    Layout focuses on structured arrangement of a pretty version's data.

To define the layout use the [Go templating](https://pkg.go.dev/html/template). You can use also [version's pkg built-in functions](https://github.com/mszostok/version/blob/main/style/go-tpl-funcs.go) that respect the [formatting settings](./format.md). Additionally, all helper functions defined by the [Sprig template library](https://masterminds.github.io/sprig/) are also available.

These fields can be access in your Go template definition:

| Key           | Description                                                                                                  |
|---------------|--------------------------------------------------------------------------------------------------------------|
| `.Version`    | Binary version value set via `-ldflags`, otherwise taken from `go install url/tool@version`.                 |
| `.GitCommit`  | Git commit value set via `-ldfags`, otherwise taken from `debug.ReadBuildInfo()` - the `vcs.revision` tag.   |
| `.BuildDate`  | Build date value set via `-ldflags`, otherwise empty.                                                        |
| `.CommitDate` | Git commit date value set via `-ldfags`, otherwise taken from `debug.ReadBuildInfo()` - the `vcs.time` tag.  |
| `.DirtyBuild` | Dirty build value, set via `-ldfags`, otherwise taken from `debug.ReadBuildInfo()` - the `vcs.modified` tag. |
| `.GoVersion`  | Go version taken from `runtime.Version()`.                                                                   |
| `.Compiler`   | Go compiler taken from `runtime.Compiler`.                                                                   |
| `.Platform`   | Platform build, in format of `runtime.GOOS/runtime.GOARCH`.                                                  |

## Go

!!! tip

    Want to try? See the [custom layout](/examples#custom-layout) example!

Example usage:

```go
var CustomLayoutGoTpl = `
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

func main() {
	// ...
	layout := style.Layout{
		GoTemplate: CustomLayoutGoTpl,
	}
	version.NewPrinter(version.WithPrettyLayout(layout))
}
```


## Config file

!!! coming-soon "Coming soon"

    See the [mszostok/version#13](https://github.com/mszostok/version/issues/13) issue for a reference. If you want to see it, please add 👍 under the issue.

The config file can be loaded by:

- enabling loading style from environment variable via `version.WithPrettyStyleFromEnv("ENV_NAME_FOR_FILE_PATH")`,
- or using `version.WithPrettyStyleFile` function directly.

=== "YAML"

    <!-- YAMLLayout start -->
    ```yaml
    layout:
      goTemplate: |2
        {{ header }}
    
          {{ key "Version" }}             {{ .Version                     | val }}
          {{ key "Git Commit" }}          {{ .GitCommit  | commit         | val }}
          {{ key "Build Date" }}          {{ .BuildDate  | fmtDate        | val }}
          {{ key "Commit Date" }}         {{ .CommitDate | fmtDate        | val }}
          {{ key "Dirty Build" }}         {{ .DirtyBuild | fmtBool        | val }}
          {{ key "Go Version" }}          {{ .GoVersion  | trimPrefix "go"| val }}
          {{ key "Compiler" }}            {{ .Compiler                    | val }}
          {{ key "Platform" }}            {{ .Platform                    | val }}
    ```
    <!-- YAMLLayout end -->

=== "JSON"

    !!! note ""

        You need to admit that it's not the best option for multiline strings 😬

    <!-- JSONLayout start -->
    ```json
    {
      "layout": {
        "goTemplate": "\n{{ header }}\n\n  {{ key \"Version\" }}             {{ .Version                     | val }}\n  {{ key \"Git Commit\" }}          {{ .GitCommit  | commit         | val }}\n  {{ key \"Build Date\" }}          {{ .BuildDate  | fmtDate        | val }}\n  {{ key \"Commit Date\" }}         {{ .CommitDate | fmtDate        | val }}\n  {{ key \"Dirty Build\" }}         {{ .DirtyBuild | fmtBool        | val }}\n  {{ key \"Go Version\" }}          {{ .GoVersion  | trimPrefix \"go\"| val }}\n  {{ key \"Compiler\" }}            {{ .Compiler                    | val }}\n  {{ key \"Platform\" }}            {{ .Platform                    | val }}\n"
      }
    }
    ```
    <!-- JSONLayout end -->

