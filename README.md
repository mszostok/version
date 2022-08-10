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

Visit the [version.szostok.io/quick-start](https://version.szostok.io/quick-start) for the most popular way of the setup.

## Documentation

<!--- Curious why? See the [blogpost about displaying the CLI version](). --->

Visit the [version.szostok.io](https://version.szostok.io) for complete documentation about setup and usage.

## Functionality

- For Go 1.18+ detect `version`, `commit`, `commitDate`, `dirtyBuild` automatically
  - Allow to override the data via `-ldflags`
- Print version in YAML, JSON, short and pretty formats
- Detect and display an upgrade notice if a newer version was released
- Automatically disable color output for non-tty output streams
  - Handle version and upgrade notice separately
- Highly customizable output format
- Parse any dates and print it in the local time
<!--- - Autodiscover installation method --->
<!--- - All provided functionality is fully tested to ensure no regression --->

<br /><br />

<img src="./docs/assets/pretty-gif.gif" width="74%" align="right"/>

### `pretty` format

It comes with a built-in style. However, you can easily create your own. You can customize formatting or layout or provide own renderer functions.
<br /><br /> <br /><br /> <br /><br />

<img src="./docs/assets/json.png" width="65%" align="left"/>

### `json` format

JSON format can be useful for scripting purposes, e.g.

```
<cli> version -ojson | jq .gitCommit
```

<br /><br />
<br /><br />

<img src="./docs/assets/yaml.png" width="65%" align="right"/>

### `yaml` format

YAML format can be useful for scripting purposes, e.g.

```
<cli> version -oyaml | yq .gitCommit
```

<br /><br />
<br /><br />

<img src="./docs/assets/short.png" width="45%" align="left"/>

### `short` format

Short format can be useful for CI purposes to easily validate that the correct CLI version is used.

<br /><br />

### Cobra

There is a built-in support for https://github.com/spf13/cobra.

### Release upgrade notice

Checks for new GitHub releases. If a newer version was found, displays upgrade notice.
You can customize all aspect of upgrade check!

### Data Autodiscovery

For Go 1.18+ detects <code>version, commit, commitDate, dirtyBuild</code> automatically.

### Multiple Output Formats

Print version in YAML, JSON, short and pretty formats

### Full Customization

By default, you can enable everything with a single command such as `extension.NewVersionCobraCmd()`. But we give you an option to customize almost all aspect of the default look and feel.
For that you can use options such as `printer.WithFormatting()`, `printer.WithPostRenderHook()`, etc. To learn more about possible customization navigate to [version.szostok.io/customization](https://szostok.io/projects/version/customization/).

### New Release Discovery

Detects and display an upgrade notice if a newer version was released

### Well documented

Rich examples, documented all aspect of possible customization.

### Full modularity

It's design in a way that you can use each component individually. For example, use notice upgrade component, or only the version collector logic, or only printer with own info object.


## <img src="./docs/assets/bell-icon.png" /> Staying Informed

Follow the [@mszostok](https://twitter.com/m_szostok) on Twitter to get the latest news. You can also subscribe for new [`version`](https://github.com/mszostok/version/releases) releases on GitHub. I post there a detailed changelog for every release.

For more additional content check the [Mateusz's Szostok blog](https://szostok.io).
