```go
import "go.szostok.io/version"
```

Go package to present your CLI version in **a classy way**. All magic included!

![](docs/assets/preview.gif)

Like the idea? Give a GitHub star ⭐!

## Quick Start

```bash
go get go.szostok.io/version
```

Visit [`version.szostok.io/quick-start`](https://version.szostok.io/quick-start) for the most popular way of the setup.

## Documentation

Visit [`version.szostok.io`](https://version.szostok.io) for complete documentation about setup and usage.

## Why?

If you create a new CLI, it's natural that you use a framework such as Cobra, `urfave/cli`, or similar. Each of your CLIs has also an option to show its version. But in this case, we repeat the same stuff: collecting and displaying related information.

This package aims to solve that problem. To register the version command, simply add:

```go
extension.NewVersionCobraCmd()
```

Go 1.18 simplified collecting version-related data, as commit, date, and other data are embedded. You can still override these fields with ldflags, e.g.:

```bash
-X go.szostok.io/version.version=1.42.0
```

You can gain more features, such as upgrade notice, just by adding:

```go
extension.WithUpgradeNotice("repo-owner", "repo-name")
```

## Functionality

- For Go 1.18+, detect `version`, `commit`, `commitDate`, and `dirtyBuild` automatically
  - Allow version data overriding via `-ldflags`
- Print the version in the YAML, JSON, short, and pretty formats
- Detect and display an upgrade notice if a newer version of your project has been released
- Automatically disable color output for non-tty output streams
  - Handle the version and upgrade notices separately
- Designed in a way that lets you use each component individually
- Everything can be enabled with a single line of code. For example, use `extension.NewVersionCobraCmd()` to enable the version command for Cobra
- Customize the output format and behaviour (e.g. timeouts, re-check intervals)
- Parse any dates and print them in the local date and time format
- All provided functionality is fully tested to ensure no regression
- Extend the version info with own fields just by assigning your Go struct
<!--- - Autodiscover installation method --->

## <img src="./docs/assets/bell-icon.png" /> Stay informed

Follow [@m_szostok](https://twitter.com/m_szostok) on Twitter to get the latest news. You can also subscribe for new [`version`](https://github.com/mszostok/version/releases) releases on GitHub, where you can find a detailed changelog for each of them.

For additional content, check [Mateusz Szostok's blog](https://szostok.io).
