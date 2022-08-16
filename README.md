```go
import "go.szostok.io/version"
```

Importable Go package to present your CLI version in a classy way. All magic included!

![](docs/assets/preview.gif)

Like the idea? Give a GitHub star ⭐!

## Quick Start

```bash
go get go.szostok.io/version
```

Visit [`version.szostok.io/quick-start`](https://version.szostok.io/quick-start) for the most popular way of the setup.

## Documentation

<!--- Curious why? See the [blogpost about displaying the CLI version](). --->

Visit [`version.szostok.io`](https://version.szostok.io) for complete documentation about setup and usage.

## Functionality

- For Go 1.18+, detect `version`, `commit`, `commitDate`, and `dirtyBuild` automatically
  - Allow version data overriding via `-ldflags`
- Print the version in the YAML, JSON, short, and pretty formats
- Detect and display an upgrade notice if a newer version has been released
- Automatically disable color output for non-tty output streams
  - Handle the version and upgrade notices separately
- Customize the output format and behaviour (e.g. timeouts, re-check intervals)
- Parse any dates and print them in the local date and time format
<!--- - Autodiscover installation method --->
<!--- - All provided functionality is fully tested to ensure no regression --->

## <img src="./docs/assets/bell-icon.png" /> Stay informed

Follow [@m_szostok](https://twitter.com/m_szostok) on Twitter to get the latest news. You can also subscribe for new [`version`](https://github.com/mszostok/version/releases) releases on GitHub, where you can find a detailed changelog for each of them.

For additional content, check [Mateusz Szostok's blog](https://szostok.io).
