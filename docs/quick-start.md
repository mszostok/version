# Quick Start

The quick start describes the most popular way of creating CLIs in Go. It uses [Cobra](https://cobra.dev/) and [GoReleaser](https://goreleaser.com/).

## Register `version` command

```go
package main

import (
	"os"

	"github.com/mszostok/version"
	"github.com/spf13/cobra"
)

// NewRoot returns a root cobra.Command for the whole CLI.
func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cobra",
		Short: "An example CLI built with github.com/spf13/cobra",
	}

	cmd.AddCommand(
		// you just need to add this, and you are done.
		version.NewCobraCmd(),
	)

	return cmd
}

func main() {
	if err := NewRoot().Execute(); err != nil {
		os.Exit(1)
	}
}
```

In that way you get a fully working `<cli> version` command.

![](assets/examples/screen-cobra-version.png)
![](assets/examples/screen-cobra-version_-h.png)

## GoReleaser versioning info with `-ldflags`

```yaml
# .goreleaser.yaml
builds:
  -
    # .. your settings ..

    ldflags:
      - -s -w
      - -X go.szostok.io/version.version={{.Version}}
      - -X go.szostok.io/version.buildDate={{.Date}}
```

### Summary

As you saw, in a few seconds, you got a powerful `version` command! However, this only scratches the surfaces of possible configuration options.

See the customization documentation for more guidelines on how to meet what you need. For example:
- [usage examples](/get-started/usage)

- [build options](/get-started/build-ldflags)
- and [customization](/customization/)
