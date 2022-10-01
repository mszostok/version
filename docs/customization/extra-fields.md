# Extra Info Fields

The `version` package supports most popular version fields natively.

??? example "Native fields"

     | Key           | Description                                                                                                     |
     |---------------|-----------------------------------------------------------------------------------------------------------------|
     | `.Version`    | Binary version value set via `-ldflags`, otherwise taken from `go install url/tool@version`.                    |
     | `.GitCommit`  | Git commit value set via `-ldfags`, otherwise taken from `debug.ReadBuildInfo()` - the `vcs.revision` tag.      |
     | `.BuildDate`  | Build date value set via `-ldflags`, otherwise empty.                                                           |
     | `.CommitDate` | Git commit date value set via `-ldfags`, otherwise taken from `debug.ReadBuildInfo()` from the `vcs.time` tag.  |
     | `.DirtyBuild` | Dirty build value, set via `-ldfags`, otherwise taken from `debug.ReadBuildInfo()` from the `vcs.modified` tag. |
     | `.GoVersion`  | Go version taken from `runtime.Version()`.                                                                      |
     | `.Compiler`   | Go compiler taken from `runtime.Compiler`.                                                                      |
     | `.Platform`   | Build platform, passed in the following format: `runtime.GOOS/runtime.GOARCH`.                                  |

     Check [build options](../../get-started/build-ldflags) to learn how to override them if you need.

However, each project may want to display more information such as documentation or changelog URLs, sometimes even domain related fields. You can provide them using your Go struct.

## Usage

!!! tip

    Want to try? See the [custom fields](/examples#custom-fields) example!

Steps:

1. Assign your custom struct to `Info.ExtraFields`.

   Go struct with nested fields are not properly supported in Pretty mode.

2. Use `json`, `yaml` and `pretty` field tags to define the field name for a given output format.
3. In the Pretty mode, fields are printed in the same order as defined in struct.

Example:

```go
--8<--
examples/custom-fields/main.go
--8<--
```
