```go
import "go.szostok.io/version"
```

Importable Go package for printing the CLI version. All magic included!

![](docs/assets/preview.png)

## Features

<img src="./docs/assets/pretty.png" width="77%" align="right"/>

### `pretty` format

Pretty format a.k.a human-readable.
<br /><br /> <br /><br /> <br /><br />
<br /><br /> <br /><br /> <br /><br />

<img src="./docs/assets/json.png" width="65%" align="left"/>

### `json` format

JSON format that can be useful for CI examples, e.g.
```
<cli> version -ojson | jq .gitCommit
```

<br /><br />
<br /><br />
<br /><br />

<img src="./docs/assets/json.png" width="65%" align="right"/>

### `yaml` format

YAML format that can be useful for CI examples, e.g. `<cli> version -oyaml | yq .gitCommit`

<br /><br />
<br /><br />
<br /><br />

<img src="./docs/assets/json.png" width="65%" align="left"/>

### `short` format

JSON format that can be useful for CI examples, e.g. `<cli> version -ojson | jq .gitCommit`

<br /><br />
<br /><br />
<br /><br />

## Usage

### Custom

```go
func foo() {
  version.CollectFromBuildInfo()

  printer := version.NewPrinter()
  printer.RegisterFlags(ver.Flags())

  printer.Print(os.Stdout, version.Get("name"))
}
```
### Cobra

```go
package cmd

import (
	"log"

	"github.com/mszostok/version"
	"github.com/spf13/cobra"
)

// NewRoot returns a root cobra.Command for the whole CLI.
func NewRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gimme",
		Short: "Insights about a Git(Hub) repository.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				log.Fatalln(err)
			}
		},
	}

	rootCmd.AddCommand(
		// Just register and you are done!
		version.NewCobraCmd("gimme"),
	)

	return rootCmd
}
```

In that way you get a fully working `<cli> version` command.

![](docs/assets/help.png)
