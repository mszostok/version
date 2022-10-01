# Layout

!!! note ""

    Layout focuses on structured arrangement of pretty version data.

To define the layout, use [Go templating](https://pkg.go.dev/html/template). You can also use the [`version` package's built-in functions](https://github.com/mszostok/version/blob/main/style/go-tpl-funcs.go) that respect the [formatting settings](./format.md). All helper functions defined by the [Sprig template library](https://masterminds.github.io/sprig/) are also available.

These are the fields that you can access in your Go template definition:

| Key           | Description                                                                                                     |
| ------------- | --------------------------------------------------------------------------------------------------------------- |
| `.Version`    | Binary version value set via `-ldflags`, otherwise taken from `go install url/tool@version`.                    |
| `.GitCommit`  | Git commit value set via `-ldfags`, otherwise taken from `debug.ReadBuildInfo()` - the `vcs.revision` tag.      |
| `.BuildDate`  | Build date value set via `-ldflags`, otherwise empty.                                                           |
| `.CommitDate` | Git commit date value set via `-ldfags`, otherwise taken from `debug.ReadBuildInfo()` from the `vcs.time` tag.  |
| `.DirtyBuild` | Dirty build value, set via `-ldfags`, otherwise taken from `debug.ReadBuildInfo()` from the `vcs.modified` tag. |
| `.GoVersion`  | Go version taken from `runtime.Version()`.                                                                      |
| `.Compiler`   | Go compiler taken from `runtime.Compiler`.                                                                      |
| `.Platform`   | Build platform, passed in the following format: `runtime.GOOS/runtime.GOARCH`.                                  |

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
	printer.New(printer.WithPrettyLayout(&layout))
}
```

## Config file

!!! coming-soon "Coming soon"

    See the [mszostok/version#13](https://github.com/mszostok/version/issues/13) issue for reference. If you'd like to see it included in a future release, add 👍 under the issue.

To load the config file, you can:

- Enable loading your custom style from an environment variable via `printer.WithPrettyStyleFromEnv("ENV_NAME_FOR_FILE_PATH")`
- Use `printer.WithPrettyStyleFile` function directly

=== "YAML"

    <!-- YAMLLayout start -->
    ```yaml
    layout:
      goTemplate: |
        {{ AdjustKeyWidth .ExtraFields }}
        {{ Header .Meta.CLIName }}

          {{ Key "Version"     }}    {{ .Version                     | Val }}
          {{ Key "Git Commit"  }}    {{ .GitCommit  | Commit         | Val }}
          {{ Key "Build Date"  }}    {{ .BuildDate  | FmtDate        | Val }}
          {{ Key "Commit Date" }}    {{ .CommitDate | FmtDate        | Val }}
          {{ Key "Dirty Build" }}    {{ .DirtyBuild | FmtBool        | Val }}
          {{ Key "Go version"  }}    {{ .GoVersion  | trimPrefix "go"| Val }}
          {{ Key "Compiler"    }}    {{ .Compiler                    | Val }}
          {{ Key "Platform"    }}    {{ .Platform                    | Val }}
          {{- range $item := (.ExtraFields | Extra) }}
          {{ $item.Key | Key   }}    {{ $item.Value | Val }}
          {{- end}}
    ```
    <!-- YAMLLayout end -->

=== "JSON"

    !!! note ""

        You need to admit that it's not the best option for multiline strings 😬

    <!-- JSONLayout start -->
    ```json
    {
      "layout": {
        "goTemplate": "{{ AdjustKeyWidth .ExtraFields }}\n{{ Header .Meta.CLIName }}\n\n  {{ Key \"Version\"     }}    {{ .Version                     | Val }}\n  {{ Key \"Git Commit\"  }}    {{ .GitCommit  | Commit         | Val }}\n  {{ Key \"Build Date\"  }}    {{ .BuildDate  | FmtDate        | Val }}\n  {{ Key \"Commit Date\" }}    {{ .CommitDate | FmtDate        | Val }}\n  {{ Key \"Dirty Build\" }}    {{ .DirtyBuild | FmtBool        | Val }}\n  {{ Key \"Go version\"  }}    {{ .GoVersion  | trimPrefix \"go\"| Val }}\n  {{ Key \"Compiler\"    }}    {{ .Compiler                    | Val }}\n  {{ Key \"Platform\"    }}    {{ .Platform                    | Val }}\n  {{- range $item := (.ExtraFields | Extra) }}\n  {{ $item.Key | Key   }}    {{ $item.Value | Val }}\n  {{- end}}\n"
      }
    }
    ```
    <!-- JSONLayout end -->
