# Custom Renderer

If the custom formatting and [layout](./layout.md) don't meet you needs, you can simply specify your own rendering function:

```go linenums="1" hl_lines="19-22"
// NewRoot returns a root cobra.Command for the whole CLI.
func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "example",
		Short: "An example CLI built with github.com/spf13/cobra",
	}

	opts := []upgrade.Options{
		upgrade.WithRenderer(func(in *upgrade.Info, isSmartTerminal bool) (string, error) {
			return fmt.Sprintf(`
      Version             %q
      New Version         %q
   `, in.Version, in.NewVersion), nil
		}),
	}

	cmd.AddCommand(
		// 1. Register the 'version' command
		extension.NewVersionCobraCmd(
			// 2. Explicitly enable the upgrade notice
			extension.WithUpgradeNotice("mszostok", "codeowners-validator", opts...)),
	)

	return cmd
}
```
