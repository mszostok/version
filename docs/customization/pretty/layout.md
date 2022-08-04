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
{{ Header .Meta.CLIName }}

  {{ Key "Version" }}             {{ .Version                     | Val }}
  {{ Key "Git Commit" }}          {{ .GitCommit  | Commit         | Val }}
  {{ Key "Build Date" }}          {{ .BuildDate  | FmtDate        | Val }}
  {{ Key "Commit Date" }}         {{ .CommitDate | FmtDate        | Val }}
  {{ Key "Dirty Build" }}         {{ .DirtyBuild | FmtBool        | Val }}
  {{ Key "Go Version" }}          {{ .GoVersion  | trimPrefix "go"| Val }}
  {{ Key "Compiler" }}            {{ .Compiler                    | Val }}
  {{ Key "Platform" }}            {{ .Platform                    | Val }}
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
        {{ Header .Meta.CLIName }}

          {{ Key "Version" }}             {{ .Version                     | Val }}
          {{ Key "Git Commit" }}          {{ .GitCommit  | Commit         | Val }}
          {{ Key "Build Date" }}          {{ .BuildDate  | FmtDate        | Val }}
          {{ Key "Commit Date" }}         {{ .CommitDate | FmtDate        | Val }}
          {{ Key "Dirty Build" }}         {{ .DirtyBuild | FmtBool        | Val }}
          {{ Key "Go Version" }}          {{ .GoVersion  | trimPrefix "go"| Val }}
          {{ Key "Compiler" }}            {{ .Compiler                    | Val }}
          {{ Key "Platform" }}            {{ .Platform                    | Val }}
    ```
    <!-- YAMLLayout end -->

=== "JSON"

    !!! note ""

        You need to admit that it's not the best option for multiline strings 😬

    <!-- JSONLayout start -->
    ```json
    {
      "layout": {
        "goTemplate": "\n{{ header }}\n\n  {{ Key \"Version\" }}             {{ .Version                     | Val }}\n  {{ Key \"Git Commit\" }}          {{ .GitCommit  | Commit         | Val }}\n  {{ Key \"Build Date\" }}          {{ .BuildDate  | FmtDate        | Val }}\n  {{ Key \"Commit Date\" }}         {{ .CommitDate | FmtDate        | Val }}\n  {{ Key \"Dirty Build\" }}         {{ .DirtyBuild | FmtBool        | Val }}\n  {{ Key \"Go Version\" }}          {{ .GoVersion  | trimPrefix \"go\"| Val }}\n  {{ Key \"Compiler\" }}            {{ .Compiler                    | Val }}\n  {{ Key \"Platform\" }}            {{ .Platform                    | Val }}\n"
      }
    }
    ```
    <!-- JSONLayout end -->

