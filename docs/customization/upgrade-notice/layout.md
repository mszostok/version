# Layout

!!! note ""
    Layout focuses on structured arrangement of pretty version data.

To define the layout, use [Go templating](https://pkg.go.dev/html/template). You can use also the [version package's built-in functions](https://github.com/mszostok/version/blob/main/style/go-tpl-funcs.go). All helper functions defined by the [Sprig template library](https://masterminds.github.io/sprig/) are also available.

These are the fields that you can access in your Go template definition:

| Key            | Description                                              |
|----------------|----------------------------------------------------------|
| `.Version`     | Binary version                                           |
| `.NewVersion`  | New binary version taken from the latest GitHub release  |
| `.ReleaseURL`  | GitHub release URL                                       |

## Go

!!! tip

    Want to try? See the [custom layout](/examples/#custom-upgrade-notice) example!

Example usage:

```go
var forBoxLayoutGoTpl = heredoc.Doc(`
A new release is available: {{ .Version }} → {{ .NewVersion | Green }}
{{ .ReleaseURL  | Underline | Blue }}`)

func main() {
	upgradeOpts := []upgrade.Options{
		upgrade.WithLayout(&style.Layout{
			GoTemplate: forBoxLayoutGoTpl,
		}),
		upgrade.WithPostRenderHook(func(body string, isSmartTerminal bool) (string, error) {
			return body + "\n footer", nil
		}),
	}

	notice := upgrade.NewGitHubDetector("mszostok", "codeowners-validator", upgradeOpts...)
	err := notice.PrintIfFoundGreater(os.Stderr, "0.5.4")
	if err != nil {
		log.Fatal(err)
	}
}
```
